package core

// Data is a contextual point value in time.
type Data[T any] struct {
	Context

	// Value is the recorded value of this contextual moment.
	Value T
}
