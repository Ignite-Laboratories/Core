package std

// Target functions return a pointer to a value.
type Target[TValue any] func() *TValue
