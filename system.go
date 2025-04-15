package core

import (
	"time"
)

// System represents an impulsable structure.
type System struct {
	NamedEntity
	*Neuron

	// Alive indicates if the system is still alive.
	Alive bool

	// Stopping indicates if Alive will shortly be set to false.
	Stopping bool

	// Cleanup is called (if provided) after Stop() finishes setting Alive to false.
	Cleanup func()
}

// CreateSystem creates a new structure which fires the provided action whenever the potential returns true.
func CreateSystem(engine *Engine, action Action, potential Potential, muted bool) *System {
	sys := &System{}
	sys.NamedEntity = NewNamedEntity()
	sys.Alive = true
	sys.Neuron = engine.Loop(func(ctx Context) {
		if sys.Stopping {
			sys.Neuron.Destroy()
			sys.Alive = false
			return
		}
		action(ctx)
	}, potential, muted)
	return sys
}

// Stop will signal the system to stop and then block until it completes.
func (sys *System) Stop() {
	sys.Stopping = true
	for Alive && sys.Alive {
		// Hold until finished
		time.Sleep(time.Millisecond)
	}
	if sys.Cleanup != nil {
		sys.Cleanup()
	}
}
