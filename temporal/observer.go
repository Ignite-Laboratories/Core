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
//
// Looping indicates if the stimulator of this dimension should activate impulsively, or as a loop.
func Observer[TValue any](engine *core.Engine, potential core.Potential, muted bool, looping bool, target std.Target[TValue]) *Dimension[TValue, any] {
	d := Dimension[TValue, any]{}
	d.ID = core.NextID()
	d.Window = core.DefaultWindow
	d.Trimmer = engine.Loop(d.Trim, when.Always, false)
	f := func(ctx core.Context) {
		data := std.Data[TValue]{
			Context: ctx,
			Point:   *target(),
		}
		d.Mutex.Lock()
		d.Timeline = append(d.Timeline, data)
		d.Current = &data
		d.Mutex.Unlock()
	}
	if looping {
		d.Stimulator = engine.Loop(f, potential, muted)
	} else {
		d.Stimulator = engine.Stimulate(f, potential, muted)
	}
	return &d
}
