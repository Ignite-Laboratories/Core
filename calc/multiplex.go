package calc

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/condition"
)

// NewMultiplex creates a dimension that's a blend of the point value of many input dimensions for every impulse that the potential returns true.
//
// NOTE: The potential function gates the creation of timeline indexes!
// This can adjust the "resolution" of output data =)
//
// Muted indicates if the stimulator of this dimension should be created muted.
func NewMultiplex[TValue core.Numeric](engine *core.Engine, potential core.Potential, muted bool, blend core.Blend[TValue], dimensions ...*core.Dimension[any, any]) *core.Dimension[TValue, any] {
	d := core.Dimension[TValue, any]{}
	d.ID = core.NextID()
	d.Trimmer = engine.Loop(d.Trim, condition.Always, false)
	d.Stimulator = engine.Stimulate(func(ctx core.Context) {
		values := make([]any, len(dimensions))
		for i, otherD := range dimensions {
			values[i] = otherD.Current
		}
		data := core.Data[TValue]{
			Context: ctx,
			Point:   blend(values),
		}
		d.Mutex.Lock()
		d.Timeline = append(d.Timeline, data)
		d.Current = data
		d.Mutex.Unlock()
	}, potential, muted)
	return &d
}
