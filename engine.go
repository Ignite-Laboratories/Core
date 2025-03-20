package core

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// engine is a neural impulse driver.
type engine struct {
	// Active indicates if the neural impulse engine is currently firing activations or not.
	Active bool

	// Ignition provides the first impulse moment of this neural impulse engine.
	Ignition time.Time

	// LastImpulse provides statistics regarding the last neural impulse.
	LastImpulse runtime

	beat        int
	activations map[uint64]*activation // This is instantiated on init()
	mutex       sync.Mutex
}

func (e *engine) addActivation(a *activation) {
	defer e.mutex.Unlock()
	e.mutex.Lock()
	e.activations[a.ID] = a
}

// Stop causes the impulse engine to cease firing neural activations.
func (e *engine) Stop() {
	e.Active = false
}

// Mute stops the provided activation ID from activating until Unmute is called.
func (e *engine) Mute(id uint64) {
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
func (e *engine) Unmute(id uint64) {
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
func (e *engine) ClearActivation(id uint64) {
	defer e.mutex.Unlock()
	e.mutex.Lock()
	delete(e.activations, a.ID)
}

// Range provides a slice of the current neural activations.
func (e *engine) Range() []*activation {
	defer e.mutex.Unlock()
	e.mutex.Lock()
	out := make([]*activation, len(e.activations))
	for _, a := range e.activations {
		out = append(out, a)
	}
	return out
}

// Stimulate activates the provided Action once in an asynchronous fashion.
func (e *engine) Stimulate(action Action) {
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
	ctx.Beat = e.beat
	ctx.LastImpulse = e.LastImpulse

	// Build the activation
	var a activation
	a.ID = NextID()
	a.Potential = func(ctx Context) {
		go action(ctx)
	}

	// Impulse the activation
	wg := &sync.WaitGroup
	e.impulse(ctx, &a, wg)
}

// Block activates the provided Action on every impulse in a blocking fashion.
//
// It returns the activation's ID.
func (e *engine) Block(action Action) *activation {
	var a activation
	a.ID = NextID()
	a.Potential = func(ctx Context) {
		a.executing = true
		action(ctx)
		a.executing = false
	}
	e.addActivation(&a)
	return &a
}

// Pulse activates the provided Action on every impulse in an asynchronous fashion.
//
// It returns the activation's ID.
func (e *engine) Pulse(action Action) *activation {
	// NOTE: The trick here is that it never sets 'Executing' =)
	var a activation
	a.ID = NextID()
	a.Potential = func(ctx Context) {
		go action(ctx)
	}
	e.addActivation(&a)
	return &a
}

// Loop activates the provided Action in an asynchronous fashion cyclically.
//
// It returns the activation's ID.
func (e *engine) Loop(action Action) *activation {
	var a activation
	a.ID = NextID()
	a.Potential = func(ctx Context) {
		a.executing = true
		go func() {
			action(ctx)
			a.executing = false
		}()
	}
	e.addActivation(&a)
	return &a
}

// Spark begins neural activation.
func (e *engine) Spark() error {
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
		activations := make([]*activation, len(e.activations))
		var hasExecution bool
		for i, a := range e.activations {
			activations[i] = a
			if a.executing {
				hasExecution = true
			}
		}
		// If none have execution, loop the beat back to 0
		if !hasExecution {
			e.beat = 0
		}

		// Calculate the impulse stats
		e.LastImpulse.Inception = lastNow
		e.LastImpulse.Start = lastNow
		e.LastImpulse.End = lastFinishTime
		e.LastImpulse.RefractoryPeriod = now.Sub(lastFinishTime)
		e.mutex.Unlock()

		// Build a temporal context
		var ctx Context
		ctx.ID = NextID()
		ctx.Moment = now
		ctx.Delta = now.Sub(lastNow)
		ctx.Beat = e.beat
		ctx.LastImpulse = e.LastImpulse

		// Launch the wave of activations
		for _, a := range activations {
			e.impulse(ctx, activations, &wg)
		}
		wg.Wait()
		finishTime := time.Now()

		// Save off the incrementer variables
		lastNow = now
		lastFinishTime = finishTime
		e.beat++
	}
	return nil
}

func (e *engine) impulse(ctx Context, activation *activation, wg *sync.WaitGroup) {
	// Grab this activation's start ASAP!
	start := time.Now()

	// Don't re-activate anything that's still executing
	if a.Executing {
		continue
	}

	// Handle the rest asynchronously...
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer func() {
			// The activation had a failure of some kind
			if r := recover(); r != nil {
				// Mark it as not executing and log the issue
				a.Executing = false
				a.Last.End = time.Now()
				log.Printf("[%d] activation panic ", a.ID)
			}
		}()

		// Calculate the refractory period
		a.Last.RefractoryPeriod = start.Sub(a.Last.End)

		// Save off the runtime info
		ctx.LastActivation = a.Last

		// Fire the activation
		a.Potential(ctx)
		end := time.Now()

		// Update the runtime info
		a.Last.Inception = ctx.Moment
		a.Last.Start = start
		a.Last.End = end
		a.Last.RefractoryPeriod = 0
	}()
}
