package temporal

// Recorder is a type of core.Emitter that records values onto a dimension.
type Recorder[TValue any] struct {
	*Dimension[TValue, any]
}

// NewRecorder creates a recorder that emits typed values onto the provided dimension.
func NewRecorder[TValue any](dimension *Dimension[TValue, any]) *Recorder[TValue] {
	r := Recorder[TValue]{}
	r.Dimension = dimension
	return &r
}

// Emit places the provided value on the dimension.
func (r *Recorder[TValue]) Emit(value interface{}) {
	r.Write(value)
}
