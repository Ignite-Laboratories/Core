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

// Blend functions take in many dissimilar-typed values and output a result.
type Blend[TOut any] func(...any) TOut

// Operate functions take in two numeric values and output a result.
type Operate[TValue Numeric] func(TValue, TValue) TValue
