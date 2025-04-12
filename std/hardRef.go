package std

// HardRef provides a way to create inline pointer references.
func HardRef[T any](val T) *T {
	return &val
}
