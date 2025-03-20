package core

// Activation is control surface for a logical unit of execution.
type Activation struct {
	// ID is the unique identifier for this activation.
	ID uint64

	// executing indicates if the activation is currently activating.
	executing bool

	// Muted indicates if the activation has been explicitly told to temporarily stop activating.
	Muted bool

	// Potential is what this activation could do.
	Potential Action

	// Last provides temporal runtime information for the last activation.
	Last runtime
}
