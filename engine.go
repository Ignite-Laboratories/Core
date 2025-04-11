package core

import (
	"math"
	"sync"
	"time"
)

// Engine is a neural impulse driver.
type Engine struct {
	Entity

	// Active indicates if the engine is currently firing activations or not.
	Active bool

	// Ignition provides the first impulse moment of this engine.
	Ignition time.Time

	// Last provides statistics regarding the last impulse.
	Last Runtime

	// Resistance indicates how much to resist the next impulse, with 0 (default) providing no resistance.
	Resistance int

	// Beat provides the current count of impulses fired while performing asynchronous activations.
	//
	// It will loop to 0 whenever all activations are finished.
	Beat int

	// MaxFrequency is the maximum frequency of impulse.
	MaxFrequency float64

	// OnStop is called whenever the engine stops, if it's provided.
	OnStop func()

	stopPotential Potential
	neurons       map[uint64]*Neuron
	mutex         sync.Mutex
}

// NewEngine creates and configures a new neural impulse engine instance.
func NewEngine() *Engine {
	e := Engine{}
	e.ID = NextID()
	e.MaxFrequency = math.MaxFloat64

	// Make the neural map
	e.neurons = make(map[uint64]*Neuron)

	// Set up impulse regulation
	regulator := func(ctx Context) {
		for i := 0; i < e.Resistance; i++ {
		}
	}
	e.Block(regulator, func(ctx Context) bool { return true }, false)

	return &e
}

// addNeuron provides a thread-safe way of adding neurons to the internal map.
func (e *Engine) addNeuron(n *Neuron) {
	defer e.mutex.Unlock()
	e.mutex.Lock()
	e.neurons[n.ID] = n
}

// Stop causes the impulse engine to cease firing neural activations.
func (e *Engine) Stop() {
	Verbosef(ModuleName, "[%d] stopping %v\n", e.ID, e.Name)
	e.Active = false
	if e.OnStop != nil {
		e.OnStop()
	}
}

// StopWhen causes the impulse engine to cease firing neural activations when the provided potential returns true.
func (e *Engine) StopWhen(potential Potential) {
	e.stopPotential = potential
}

// MuteByID suppresses the identified neuron until Unmute is called.
func (e *Engine) MuteByID(id uint64) {
	defer e.mutex.Unlock()
	e.mutex.Lock()
	for _, n := range e.neurons {
		if n.ID == id {
			n.Muted = true
			return
		}
	}
}

// UnmuteByID un-suppresses the identified neuron.
func (e *Engine) UnmuteByID(id uint64) {
	defer e.mutex.Unlock()
	e.mutex.Lock()
	for _, n := range e.neurons {
		if n.ID == id {
			n.Muted = false
			return
		}
	}
}

// remove deletes the identified neuron from the internal neural map.
func (e *Engine) remove(id uint64) {
	defer e.mutex.Unlock()
	e.mutex.Lock()
	delete(e.neurons, id)
}

// Range provides a slice of the current neural activations.
func (e *Engine) Range() []*Neuron {
	defer e.mutex.Unlock()
	e.mutex.Lock()
	out := make([]*Neuron, len(e.neurons))
	for _, n := range e.neurons {
		out = append(out, n)
	}
	return out
}

// Block activates the provided action on every impulse in a blocking fashion, if the potential returns true.
//
// If 'muted' is true, the neuron is lies dormant until un-muted.
func (e *Engine) Block(action Action, potential Potential, muted bool) *Neuron {
	var n Neuron
	n.ID = NextID()
	n.engine = e
	n.Action = func(ctx Context) {
		n.executing = true
		action(ctx)
		n.ActivationCount++
		n.executing = false
	}
	n.Potential = potential
	n.Muted = muted
	e.addNeuron(&n)
	return &n
}

// Stimulate activates the provided action on every impulse in an asynchronous fashion, if the potential returns true.
//
// If 'muted' is true, the neuron is lies dormant until un-muted.
func (e *Engine) Stimulate(action Action, potential Potential, muted bool) *Neuron {
	// NOTE: The trick here is that it never sets 'Executing' =)
	var n Neuron
	n.ID = NextID()
	n.engine = e
	n.Action = func(ctx Context) {
		go action(ctx)
		n.ActivationCount++
	}
	n.Potential = potential
	n.Muted = muted
	e.addNeuron(&n)
	return &n
}

// Loop activates the provided action in an asynchronous fashion cyclically, if the potential returns true.
//
// NOTE: This fires a new goroutine for every activation
//
// If 'muted' is true, the neuron is lies dormant until un-muted.
func (e *Engine) Loop(action Action, potential Potential, muted bool) *Neuron {
	var n Neuron
	n.ID = NextID()
	n.engine = e
	n.Action = func(ctx Context) {
		n.executing = true
		go func() {
			action(ctx)
			n.ActivationCount++
			n.executing = false
		}()
	}
	n.Potential = potential
	n.Muted = muted
	e.addNeuron(&n)
	return &n
}

// Trigger fires the provided action one time, if the potential returns true.
//
// If 'async' is true, the action is called asynchronously - otherwise, it blocks the firing impulse.
func (e *Engine) Trigger(action Action, potential Potential, async bool) {
	defer e.mutex.Unlock()
	e.mutex.Lock()

	// Grab 'now' ASAP!
	now := time.Now()

	lastImpulse := e.Last

	// Create n temporal context
	var ctx Context
	ctx.ID = NextID()
	ctx.Moment = now
	ctx.Period = now.Sub(lastImpulse.Inception)
	ctx.Beat = e.Beat
	ctx.LastImpulse = e.Last

	// Build the neuron
	var n Neuron
	n.ID = NextID()
	n.engine = e
	n.Action = func(ctx Context) {
		if async {
			go action(ctx)
		} else {
			action(ctx)
		}
		n.ActivationCount++
		n.Destroy()
	}
	n.Potential = potential
	var wg sync.WaitGroup
	e.fire(now, ctx, &n, &wg)
}

// Spark begins driving impulses.
func (e *Engine) Spark() {
	if e.Active {
		return
	}
	e.Active = true

	Verbosef(ModuleName, "[%d] sparking %v\n", e.ID, e.Name)

	// Set up a wait group for blocking operations
	var wg sync.WaitGroup

	// On the first impulse time is oriented to the system's inception moment
	lastFinishTime := Inception
	lastNow := Inception

	// Loop =)
	for Alive && e.Active {
		// Grab 'now' ASAP!
		now := time.Now()
		period := now.Sub(lastNow)

		// Don't fire faster than the maximum operating frequency
		if period < HertzToDuration(e.MaxFrequency) {
			continue
		}

		// Get the current impulse wave of neurons
		e.mutex.Lock()
		neurons := make([]*Neuron, 0, len(e.neurons))
		var hasExecution bool
		for _, n := range e.neurons {
			neurons = append(neurons, n)
			if n.executing {
				hasExecution = true
			}
		}
		e.mutex.Unlock()

		// If none have execution, loop the Beat back to 0
		if !hasExecution {
			e.Beat = 0
		}

		// Calculate the impulse stats
		e.Last.Inception = lastNow
		e.Last.Start = lastNow
		e.Last.End = lastFinishTime
		e.Last.Duration = e.Last.End.Sub(e.Last.Start)
		e.Last.RefractoryPeriod = now.Sub(lastFinishTime)

		// Build a temporal context
		var ctx Context
		ctx.ID = NextID()
		ctx.Moment = now
		ctx.Period = period
		ctx.Beat = e.Beat
		ctx.LastImpulse = e.Last

		// Check if the engine's stopping potential has been set...
		if e.stopPotential != nil && e.stopPotential(ctx) {
			// ...If so, end execution
			e.Stop()
			break
		}

		// Launch the wave of neurons
		for _, neuron := range neurons {
			// Grab this neuron's start ASAP!
			start := time.Now()

			// Don't re-activate anything that's still executing or muted
			if neuron.executing || neuron.Muted {
				continue
			}

			// Handle the rest asynchronously...
			wg.Add(1)
			go e.fire(start, ctx, neuron, &wg)
		}
		wg.Wait()
		finishTime := time.Now()

		// Save off the incrementer variables
		lastNow = now
		lastFinishTime = finishTime
		e.Beat++
	}
	// Give everyone a brief moment to cleanup
	time.Sleep(time.Millisecond * 250)
}

// fire is what activates each neuron.
func (e *Engine) fire(start time.Time, ctx Context, neuron *Neuron, wg *sync.WaitGroup) {
	defer wg.Done()
	defer func() {
		// Check if the neuron had a failure of some kind
		if r := recover(); r != nil {
			// Mark it as not executing and log the issue
			neuron.executing = false
			neuron.Last.End = time.Now()
			Printf(ModuleName, "[%d] Neural panic ", neuron.ID)
		}
	}()

	// Test the potential first
	ctx.LastActivation = neuron.Last
	if !neuron.Potential(ctx) {
		return
	}

	// Calculate the refractory period
	neuron.Last.RefractoryPeriod = start.Sub(neuron.Last.End)

	// Save off the runtime info
	ctx.LastActivation = neuron.Last

	// Fire the neuron
	neuron.Action(ctx)
	end := time.Now()

	// Update the runtime info
	neuron.Last.Inception = ctx.Moment
	neuron.Last.Start = start
	neuron.Last.End = end
}
