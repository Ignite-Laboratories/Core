package temporal

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/when"
)

// Blend represents the state of A and B that generated the resulting blended Value.
type Blend[TValue core.Numeric] struct {
	Value TValue
	A     Data[TValue]
	B     Data[TValue]
}

// NewBlend creates a dimension that blends the point value of two input dimensions for every impulse that the potential returns true.
//
// NOTE: The potential function gates the creation of timeline indexes!
// This can adjust the "resolution" of output data =)
//
// Muted indicates if the stimulator of this dimension should be created muted.
func NewBlend[TValue core.Numeric](engine *core.Engine, potential core.Potential, muted bool, blend core.Operate[TValue], a *Dimension[TValue, any], b *Dimension[TValue, any]) *Dimension[Blend[TValue], any] {
	d := Dimension[Blend[TValue], any]{}
	d.ID = core.NextID()
	d.Window = core.DefaultWindow
	d.Trimmer = engine.Loop(d.Trim, when.Always, false)
	d.Stimulator = engine.Stimulate(func(ctx core.Context) {
		mux := Blend[TValue]{
			A: a.Current,
			B: b.Current,
		}
		mux.Value = blend(mux.A.Point, mux.B.Point)
		data := Data[Blend[TValue]]{
			Context: ctx,
			Point:   mux,
		}
		d.Mutex.Lock()
		d.Timeline = append(d.Timeline, data)
		d.Current = data
		d.Mutex.Unlock()
	}, potential, muted)
	return &d
}
