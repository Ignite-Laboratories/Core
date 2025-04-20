package core

// Tuple represents up to 8 generic values.
type Tuple[T any] struct {
	A T
	B T
	C T
	D T
	E T
	F T
	G T
	H T
}

// NumberRange creates a Tuple with A representing the start and B representing the stop.
func NumberRange[T Numeric](start T, stop T) Tuple[T] {
	return Tuple[T]{
		A: start,
		B: stop,
	}
}
