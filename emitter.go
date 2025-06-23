package core

// Emitter is an interface that can be called to emit a single value to a buffer.
type Emitter[T any] interface {
	// Emit places the provided value into the buffer.
	Emit(value T)
}
