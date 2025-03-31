package core

// Action functions are provided temporal context when invoked.
type Action func(ctx Context)

// Potential functions are provided temporal context when invoked in order to make decisions.
type Potential func(ctx Context) bool

// alwaysFire provides a potential that always returns true.
func alwaysFire(ctx Context) bool {
	// This is here because `core` cannot cyclically reference `when`
	return true
}
