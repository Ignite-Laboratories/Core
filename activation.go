package core

// Activation is a logical unit of execution.
type Activation struct {
	// ID is the unique identifier for this Activation.
	ID uint64

	// Executing indicates if the activation is currently executing.
	Executing bool

	// Potential is what this Activation could do.
	Potential Action

	// Last provides temporal runtime information for the last completed activation.
	Last runtime
}

// newBlockingActivation creates a new activation that blocks the impulse.
func newBlockingActivation(function Action) *Activation {
	var a Activation
	a.ID = NextID()
	a.Potential = func(ctx Context) {
		a.Executing = true
		function(ctx)
		a.Executing = false
	}
	return &a
}

// newImpulsiveActivation creates a new activation that fires asynchronously on every impulse.
func newImpulsiveActivation(function Action) *Activation {
	var a Activation
	a.ID = NextID()
	a.Potential = func(ctx Context) {
		go a.Potential(ctx)
	}
	return &a
}

// newLoopingActivation creates a new activation that fires in an asynchronous loop.
func newLoopingActivation(function Action) *Activation {
	var a Activation
	a.ID = NextID()
	a.Potential = func(ctx Context) {
		a.Executing = true
		go func() {
			function(ctx)
			a.Executing = false
		}()
	}
	return &a
}
