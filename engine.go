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
	activations map[uint64]*Activation // This is instantiated on init()
	mutex       sync.Mutex
}

func (e *engine) addActivation(a *Activation) {
	e.mutex.Lock()
	e.activations[a.ID] = a
	e.mutex.Unlock()
}

// Stop causes the impulse engine to cease firing neural activations.
func (e *engine) Stop() {
	e.Active = false
}

// Block activates the provided Action on every impulse in a blocking fashion.
func (e *engine) Block(action Action) {
	e.addActivation(newBlockingActivation(action))
}

// Stimulate activates the provided Action on every impulse in an asynchronous fashion.
func (e *engine) Stimulate(action Action) {
	e.addActivation(newImpulsiveActivation(action))
}

// Loop activates the provided Action in an asynchronous fashion cyclically.
func (e *engine) Loop(action Action) {
	e.addActivation(newLoopingActivation(action))
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
		e.mutex.Lock()
		activations := make([]*Activation, len(e.activations))
		var hasExecution bool
		for i, a := range e.activations {
			activations[i] = a
			if a.Executing {
				hasExecution = true
			}
		}
		// If none have execution, loop the beat back to 0
		if !hasExecution {
			e.beat = 0
		}
		e.mutex.Unlock()

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
		ctx.Beat = e.beat
		ctx.LastImpulse = e.LastImpulse

		// Launch the wave of activations
		e.impulse(ctx, activations, &wg)
		wg.Wait()
		finishTime := time.Now()

		// Save off the incrementer variables
		lastNow = now
		lastFinishTime = finishTime
		e.beat++
	}
	return nil
}

func (e *engine) impulse(ctx Context, activations []*Activation, wg *sync.WaitGroup) {
	// Launch all kernels...
	for _, a := range activations {
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
}
