package std

type reference[T any] struct {
	Ref *T
}

// HardRef provides a way to create inline pointer references.
func HardRef[T any](val T) *reference[T] {
	return &reference[T]{Ref: &val}
}
