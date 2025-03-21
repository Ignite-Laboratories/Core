package core

// Loopable represents a looping system.
type Loopable interface {
	GetActive() bool
	Activate()
	ActivateSynchronously()
	GetActivation() *Activation
	Pace(ctx Context) bool
	Loop(ctx Context)
}
