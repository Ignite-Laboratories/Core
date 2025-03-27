package temporal

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/when"
)

// Operation represents the state of A and B that generated the resulting Value.
type Operation[TValue core.Numeric] struct {
	Value TValue
	A     Data[TValue]
	B     Data[TValue]
}

func NewOperation[TValue core.Numeric](engine *core.Engine, potential core.Potential, muted bool, operator core.Operate[TValue], a *Dimension[TValue, any], b *Dimension[TValue, any]) *Dimension[Operation[TValue], any] {
	d := Dimension[Operation[TValue], any]{}
	d.ID = core.NextID()
	d.Window = core.DefaultWindow
	d.Trimmer = engine.Loop(d.Trim, when.Always, false)
	d.Stimulator = engine.Stimulate(func(ctx core.Context) {
		operation := Operation[TValue]{
			A: a.Current,
			B: b.Current,
		}
		operation.Value = operator(operation.A.Point, operation.B.Point)
		data := Data[Operation[TValue]]{
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
