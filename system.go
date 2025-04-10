package core

// System represents an impulse-able structure.
type System struct {
	Entity
	*Neuron

	// Alive indicates if the system is still alive.
	Alive bool

	// Dying indicates if the alive will shortly be set to false.
	Dying bool
}

func CreateSystem(engine *Engine, action Action, potential Potential, muted bool) *System {
	sys := &System{}
	sys.ID = NextID()
	sys.Alive = true
	sys.Neuron = engine.Loop(func(ctx Context) {
		if sys.Dying {
			sys.Neuron.Destroy()
			sys.Alive = false
			return
		}
		action(ctx)
	}, potential, muted)
	return sys
}

// Stop will signal the system to die, then block until it finishes cleaning itself up.
func (s *System) Stop() {
	s.Dying = true
	for s.Alive {
		// Hold until finished
	}
}
