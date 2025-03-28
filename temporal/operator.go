package temporal

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/when"
)

// Operation represents the state of A and B that generated the resulting Value.
type Operation[TValue core.Numeric] struct {
	Value TValue
	A     std.Data[TValue]
	B     std.Data[TValue]
}

func Operator[TValue core.Numeric](engine *core.Engine, potential core.Potential, muted bool, operator core.Operate[TValue], a *Dimension[TValue, any], b *Dimension[TValue, any]) *Dimension[Operation[TValue], any] {
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
		data := std.Data[Operation[TValue]]{
			Context: ctx,
			Point:   operation,
		}
		d.update(data)
	}, potential, muted)
	return &d
}
