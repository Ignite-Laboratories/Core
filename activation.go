package core

// Activation is a logical unit of execution.
type Activation struct {
	Entity

	// Executing indicates if the activation is currently executing.
	Executing bool

	// Action is what to execute.
	Action Action

	// Last provides temporal runtime information for the last completed activation.
	Last runtime
}

// newBlockingActivation creates a new activation that blocks the impulse.
func newBlockingActivation(function Action) *Activation {
	var a Activation
	a.ID = NextID()
	a.Action = func(ctx Context) {
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
	a.Action = func(ctx Context) {
		go a.Action(ctx)
	}
	return &a
}

// newLoopingActivation creates a new activation that fires in an asynchronous loop.
func newLoopingActivation(function Action) *Activation {
	var a Activation
	a.ID = NextID()
	a.Action = func(ctx Context) {
		a.Executing = true
		go func() {
			function(ctx)
			a.Executing = false
		}()
	}
	return &a
}
