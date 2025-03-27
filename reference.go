package core

type reference[T any] struct {
	Val *T
}

// QuickRef provides a way to easily create inline pointer references.
func QuickRef[T any](val T) *reference[T] {
	return &reference[T]{Val: &val}
}
