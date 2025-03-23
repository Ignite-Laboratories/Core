package core

// Activation is a logical unit of execution.
type Activation struct {
	Entity

	// executing indicates if the activation is currently active.
	executing bool

	// Muted indicates if the activation has been explicitly suppressed from activation.
	Muted bool

	// Action is what this activation does.
	Action Action

	// Potential must return true when called for activation to occur.
	Potential Potential

	// Last provides temporal runtime information for the last activation.
	Last runtime
}

// Mute suppresses activation.
func (a *Activation) Mute() *Activation {
	a.Muted = true
	return a
}

// Unmute un-suppresses activation.
func (a *Activation) Unmute() *Activation {
	a.Muted = false
	return a
}
