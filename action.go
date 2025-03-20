package core

// Action functions are provided temporal context when activated.
type Action func(ctx Context)

// NewActionPotential creates a new Action that only activates if the provided potential returns true.
func NewActionPotential(action Action, potential func(ctx Context) bool) Action {
	return func(ctx Context) {
		if potential(ctx) {
			action(ctx)
		}
	}
}
