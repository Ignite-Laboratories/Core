package core

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// Engine is a neural impulse driver.
type Engine struct {
	// Active indicates if the neural impulse engine is currently firing activations or not.
	Active bool

	// Ignition provides the first impulse moment of this neural impulse engine.
	Ignition time.Time

	// LastImpulse provides statistics regarding the last neural impulse.
	LastImpulse runtime

	Beat        int
	activations map[uint64]*Activation // This is instantiated on init()
	mutex       sync.Mutex
}

func (e *Engine) addActivation(a *Activation) {
	defer e.mutex.Unlock()
	e.mutex.Lock()
	e.activations[a.ID] = a
}

// Initialize must be called before the first Spark.
func (e *Engine) Initialize() {
	e.activations = make(map[uint64]*Activation)
}

// Stop causes the impulse engine to cease firing neural activations.
func (e *Engine) Stop() {
	e.Active = false
}

// Mute stops the provided activation ID from activating until Unmute is called.
func (e *Engine) Mute(id uint64) {
	defer e.mutex.Unlock()
	e.mutex.Lock()
	for _, a := range e.activations {
		if a.ID == id {
			a.Muted = true
			return
		}
	}
}

// Unmute lets the provided activation ID begin activating again.
func (e *Engine) Unmute(id uint64) {
	defer e.mutex.Unlock()
	e.mutex.Lock()
	for _, a := range e.activations {
		if a.ID == id {
			a.Muted = false
			return
		}
	}
}

// ClearActivation removes the activation ID from neural activity.
func (e *Engine) ClearActivation(id uint64) {
	defer e.mutex.Unlock()
	e.mutex.Lock()
	delete(e.activations, id)
}

// Range provides a slice of the current neural activations.
func (e *Engine) Range() []*Activation {
	defer e.mutex.Unlock()
	e.mutex.Lock()
	out := make([]*Activation, len(e.activations))
	for _, a := range e.activations {
		out = append(out, a)
	}
	return out
}

// Once invokes the provided Action once in an asynchronous fashion, if the
// provided Potential returns true.
func (e *Engine) Once(action Action, potential Potential) {
	defer e.mutex.Unlock()
	e.mutex.Lock()

	// Grab 'now' ASAP!
	now := time.Now()

	lastImpulse := e.LastImpulse

	// Create a temporal context
	var ctx Context
	ctx.ID = NextID()
	ctx.Moment = now
	ctx.Delta = now.Sub(lastImpulse.Inception)
	ctx.Beat = e.Beat
	ctx.LastImpulse = e.LastImpulse

	// Build the activation
	var a Activation
	a.ID = NextID()
	a.Action = func(ctx Context) {
		go action(ctx)
	}
	a.Potential = potential

	// Impulse the activation
	var wg sync.WaitGroup
	e.impulse(ctx, &a, &wg)
}

// Block activates the provided Action on every impulse in a blocking fashion, if the provided
// Potential returns true.
//
// It returns an Activation control surface.
func (e *Engine) Block(action Action, potential Potential) *Activation {
	var a Activation
	a.ID = NextID()
	a.Action = func(ctx Context) {
		a.executing = true
		action(ctx)
		a.executing = false
	}
	a.Potential = potential
	e.addActivation(&a)
	return &a
}

// Stimulate activates the provided Action on every impulse in an asynchronous fashion, if the
// provided Potential returns true.
//
// It returns an Activation control surface.
func (e *Engine) Stimulate(action Action, potential Potential) *Activation {
	// NOTE: The trick here is that it never sets 'Executing' =)
	var a Activation
	a.ID = NextID()
	a.Action = func(ctx Context) {
		go action(ctx)
	}
	a.Potential = potential
	e.addActivation(&a)
	return &a
}

// Loop activates the provided Action in an asynchronous fashion cyclically, if the provided
// Potential returns true.
//
// It returns an Activation control surface.
func (e *Engine) Loop(action Action, potential Potential) *Activation {
	var a Activation
	a.ID = NextID()
	a.Action = func(ctx Context) {
		a.executing = true
		go func() {
			action(ctx)
			a.executing = false
		}()
	}
	a.Potential = potential
	e.addActivation(&a)
	return &a
}

// Spark begins neural activation.
func (e *Engine) Spark() error {
	if e.Active {
		return fmt.Errorf("this neural impulse engine is already active")
	}
	e.Active = true

	// Set up a wait group for blocking operations
	var wg sync.WaitGroup

	// On the first impulse time is oriented to the system's inception moment
	lastFinishTime := Inception
	lastNow := Inception

	// Loop =)
	for Alive && e.Active {
		// Grab 'now' ASAP!
		now := time.Now()

		// Get the current impulse wave of activations
		e.mutex.Lock() // Lock synchronized data
		activations := make([]*Activation, 0, len(e.activations))
		var hasExecution bool
		for _, a := range e.activations {
			activations = append(activations, a)
			if a.executing {
				hasExecution = true
			}
		}
		e.mutex.Unlock() // Unlock

		// If none have execution, loop the Beat back to 0
		if !hasExecution {
			e.Beat = 0
		}

		// Calculate the impulse stats
		e.LastImpulse.Inception = lastNow
		e.LastImpulse.Start = lastNow
		e.LastImpulse.End = lastFinishTime
		e.LastImpulse.RefractoryPeriod = now.Sub(lastFinishTime)

		// Build a temporal context
		var ctx Context
		ctx.ID = NextID()
		ctx.Moment = now
		ctx.Delta = now.Sub(lastNow)
		ctx.Beat = e.Beat
		ctx.LastImpulse = e.LastImpulse

		// Launch the wave of activations
		for _, a := range activations {
			e.impulse(ctx, a, &wg)
		}
		wg.Wait()
		finishTime := time.Now()

		// Save off the incrementer variables
		lastNow = now
		lastFinishTime = finishTime
		e.Beat++
	}
	return nil
}

func (e *Engine) impulse(ctx Context, activation *Activation, wg *sync.WaitGroup) {
	// Grab this activation's start ASAP!
	start := time.Now()

	// Don't re-activate anything that's still executing or muted
	if activation.executing || activation.Muted {
		return
	}

	// Handle the rest asynchronously...
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer func() {
			// The activation had a failure of some kind
			if r := recover(); r != nil {
				// Mark it as not executing and log the issue
				activation.executing = false
				activation.Last.End = time.Now()
				log.Printf("[%d] Activation panic ", activation.ID)
			}
		}()

		// Test the potential first
		ctx.LastActivation = activation.Last
		if !activation.Potential(ctx) {
			return
		}

		// Calculate the refractory period
		activation.Last.RefractoryPeriod = start.Sub(activation.Last.End)

		// Save off the runtime info
		ctx.LastActivation = activation.Last

		// Fire the activation
		activation.Action(ctx)
		end := time.Now()

		// Update the runtime info
		activation.Last.Inception = ctx.Moment
		activation.Last.Start = start
		activation.Last.End = end
	}()
}
