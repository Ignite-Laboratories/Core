package temporal

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/when"
)

// Observer creates a dimension that records the target value across time, if the provided potential returns true.
//
// NOTE: The potential function gates the creation of timeline indexes!
// This can adjust the "resolution" of output data =)
//
// Muted indicates if the stimulator of this dimension should be created muted.
func Observer[TValue any](engine *core.Engine, potential core.Potential, muted bool, target *TValue) *Dimension[TValue, any] {
	d := Dimension[TValue, any]{}
	d.ID = core.NextID()
	d.Window = core.DefaultWindow
	d.Trimmer = engine.Loop(d.Trim, when.Always, false)
	d.Stimulator = engine.Stimulate(func(ctx core.Context) {
		data := std.Data[TValue]{
			Context: ctx,
			Point:   *target,
		}
		d.update(data)
	}, potential, muted)
	return &d
}
