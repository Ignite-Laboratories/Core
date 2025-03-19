package core

// Action functions are provided temporal context when invoked.
type Action func(ctx Context)

// Potential functions are provided temporal context when invoked in order to make decisions.
type Potential func(ctx Context) bool

// NewActionPotential creates a new Action that only activates if the provided Potential returns true.
func NewActionPotential(action Action, potential Potential) Action {
	return func(ctx Context) {
		if potential(ctx) {
			action(ctx)
		}
	}
}
