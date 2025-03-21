package core

// Activation is control surface for a logical unit of execution.
type Activation struct {
	// ID is the unique identifier for this activation.
	ID uint64

	// executing indicates if the activation is currently activating.
	executing bool

	// Muted indicates if the activation has been explicitly told to temporarily stop activating.
	Muted bool

	// Action is what this activation does.
	Action Action

	// Potential must return true for this activation to activate.
	Potential Potential

	// Last provides temporal runtime information for the last activation.
	Last runtime
}

// Mute suppresses the activation from activation.
func (a *Activation) Mute() *Activation {
	a.Muted = true
	return a
}

// Unmute un-suppresses the activation from activation.
func (a *Activation) Unmute() *Activation {
	a.Muted = false
	return a
}
