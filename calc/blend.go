package calc

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/condition"
)

// Blend represents the state of A and B that generated the resulting muxed Value.
type Blend[TA core.Numeric, TB core.Numeric, TOut core.Numeric] struct {
	Value TOut
	A     core.Data[TA]
	B     core.Data[TB]
}

// NewBlend creates a dimension that blends the point value of two input dimensions for every impulse that the potential returns true.
//
// NOTE: The potential function gates the creation of timeline indexes!
// This can adjust the "resolution" of output data =)
//
// Muted indicates if the stimulator of this dimension should be created muted.
func NewBlend[TA core.Numeric, TB core.Numeric, TOut core.Numeric](engine *core.Engine, potential core.Potential, muted bool, blend core.Operate[TA, TB, TOut], a *core.Dimension[TA, any], b *core.Dimension[TB, any]) *core.Dimension[Blend[TA, TB, TOut], any] {
	d := core.Dimension[Blend[TA, TB, TOut], any]{}
	d.ID = core.NextID()
	d.Trimmer = engine.Loop(d.Trim, condition.Always, false)
	d.Stimulator = engine.Stimulate(func(ctx core.Context) {
		mux := Blend[TA, TB, TOut]{
			A: a.Current,
			B: b.Current,
		}
		mux.Value = blend(mux.A.Point, mux.B.Point)
		data := core.Data[Blend[TA, TB, TOut]]{
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
