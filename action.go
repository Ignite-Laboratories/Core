package core

// Action functions are provided temporal context when activated.
type Action func(ctx Context)

// When creates a new Action that only activates if the provided potential returns true.
func When(potential func(ctx Context) bool, action Action) Action {
	return func(ctx Context) {
		if potential(ctx) {
			action(ctx)
		}
	}
}
