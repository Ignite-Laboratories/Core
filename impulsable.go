package core

// Impulsable represents the lifecycle of an impulsable structure.
type Impulsable interface {
	Initialize()
	Impulse(Context)
	Cleanup()
	Lock()
	Unlock()
}
