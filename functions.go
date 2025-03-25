package core

// Action functions are provided temporal context when invoked.
type Action func(ctx Context)

// Potential functions are provided temporal context when invoked in order to make decisions.
type Potential func(ctx Context) bool

// CalculatePoint functions calculate a contextual value.
type CalculatePoint[T any] func(Context) T

// Integral functions take in a set of contextual values and calculate a result.
type Integral[TIn any, TOut any] func(Context, []TIn) TOut

// alwaysFire provides a potential that always fires.
func alwaysFire(ctx Context) bool {
	// This is here because 'core' cannot cyclically reference 'when'
	return true
}

// neverFire provides a potential that never fires.
func neverFire(ctx Context) bool {
	// This is here because 'core' cannot cyclically reference 'when'
	return false
}
