package core

// Action functions are provided temporal context when invoked.
type Action func(ctx Context)

// Potential functions are provided temporal context when invoked in order to make decisions.
type Potential func(ctx Context) bool

// CalculatePoint functions calculate a contextual value.
type CalculatePoint[T any] func(Context) T

// Integral functions take in a set of contextual values and calculate a result.
//
// They are also provided with a cache pointer that can hold values between activations.
type Integral[TIn any, TOut any, TCache any] func(Context, *TCache, []TIn) TOut

// alwaysFire provides a potential that always returns true.
func alwaysFire(ctx Context) bool {
	// This is here because `core` cannot cyclically reference `condition`
	return true
}
