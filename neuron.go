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

	// Potential must return true when called for activation to occur.
	Potential Potential

	// Last provides temporal runtime information for the last activation.
	Last runtime
}
