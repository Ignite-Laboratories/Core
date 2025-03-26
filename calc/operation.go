package calc

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/condition"
)

// Operation represents the state of A and B that generated the resulting Value.
type Operation[TValue core.Numeric] struct {
	Value TValue
	A     core.Data[TValue]
	B     core.Data[TValue]
}

func NewOperation[TValue core.Numeric](engine *core.Engine, potential core.Potential, muted bool, operator core.Operate[TValue], a *core.Dimension[TValue, any], b *core.Dimension[TValue, any]) *core.Dimension[Operation[TValue], any] {
	d := core.Dimension[Operation[TValue], any]{}
	d.ID = core.NextID()
	d.Window = core.DefaultWindow
	d.Trimmer = engine.Loop(d.Trim, condition.Always, false)
	d.Stimulator = engine.Stimulate(func(ctx core.Context) {
		operation := Operation[TValue]{
			A: a.Current,
			B: b.Current,
		}
		operation.Value = operator(operation.A.Point, operation.B.Point)
		data := core.Data[Operation[TValue]]{
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
