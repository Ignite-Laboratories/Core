package calc

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/condition"
)

// Operation represents the state of A and B that generated the resulting Value.
type Operation[TA core.Numeric, TB core.Numeric, TValue core.Numeric] struct {
	Value TValue
	A     core.Data[TA]
	B     core.Data[TB]
}

func NewOperation[TA core.Numeric, TB core.Numeric, TValue core.Numeric](engine *core.Engine, potential core.Potential, muted bool, operator core.Operate[TA, TB, TValue], a *core.Dimension[TA, any], b *core.Dimension[TB, any]) *core.Dimension[Operation[TA, TB, TValue], any] {
	d := core.Dimension[Operation[TA, TB, TValue], any]{}
	d.ID = core.NextID()
	d.Trimmer = engine.Loop(d.Trim, condition.Always, false)
	d.Stimulator = engine.Stimulate(func(ctx core.Context) {
		operation := Operation[TA, TB, TValue]{
			A: a.Current,
			B: b.Current,
		}
		operation.Value = operator(operation.A.Point, operation.B.Point)
		data := core.Data[Operation[TA, TB, TValue]]{
			Context: ctx,
			Point:   operation,
		}
		d.Mutex.Lock()
		d.Timeline = append(d.Timeline, data)
		d.Current = data
		d.Mutex.Unlock()
	}, potential, muted)
	return &d
}
