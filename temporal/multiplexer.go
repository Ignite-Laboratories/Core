package temporal

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/when"
)

// Multiplexer creates a dimension that's a blend of the point value of many input dimensions for every impulse that the potential returns true.
//
// NOTE: The potential function gates the creation of timeline indexes!
// This can adjust the "resolution" of output data =)
//
// Muted indicates if the stimulator of this dimension should be created muted.
func Multiplexer[TValue core.Numeric](engine *core.Engine, potential core.Potential, muted bool, blend core.Blend[TValue], dimensions ...*Dimension[any, any]) *Dimension[TValue, any] {
	d := Dimension[TValue, any]{}
	d.ID = core.NextID()
	d.Window = core.DefaultWindow
	d.Trimmer = engine.Loop(d.Trim, when.Always, false)
	d.Stimulator = engine.Stimulate(func(ctx core.Context) {
		values := make([]any, len(dimensions))
		for i, otherD := range dimensions {
			values[i] = otherD.Current
		}
		data := std.Data[TValue]{
			Context: ctx,
			Point:   blend(values),
		}
		d.update(data)
	}, potential, muted)
	return &d
}
