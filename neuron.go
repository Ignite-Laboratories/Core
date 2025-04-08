package core

// Neuron is a logical unit of execution.
type Neuron struct {
	Entity

	// executing indicates if the neuron is currently active.
	executing bool

	// Muted indicates if the neuron has been explicitly suppressed from activation.
	Muted bool

	// Action is what this neuron does.
	Action Action

	// Destroyed indicates if this neuron has been destroyed and can be used to make cleanup decisions.
	Destroyed bool

	// Potential must return true when called for activation to occur.
	Potential Potential

	// Last provides temporal runtime information for the last activation.
	Last Runtime

	// ActivationCount provides the number of times this neuron has been activated.
	ActivationCount uint64

	engine *Engine
}

// Trigger fires the provided action one time, if the potential returns true.
//
// If 'async' is true, the action is called asynchronously - otherwise, it blocks the firing impulse.
func (n *Neuron) Trigger(async bool) {
	n.engine.Trigger(n.Action, n.Potential, async)
}

// Destroy removes this neuron from the engine entirely.
func (n *Neuron) Destroy() {
	n.Destroyed = true
	n.engine.remove(n.ID)
}
