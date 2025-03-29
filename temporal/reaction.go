package temporal

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/when"
)

// Reaction creates a dimension that calls the reaction function if the provided potential returns true.
//
// NOTE: The potential function gates the creation of timeline indexes!
// This can adjust the "resolution" of output data =)
//
// Muted indicates if the stimulator of this dimension should be created muted.
//
// Looping indicates if the stimulator of this dimension should activate impulsively, or as a loop.
func Reaction[TValue any](engine *core.Engine, potential core.Potential, muted bool, target std.TargetFunc[TValue], change Change[TValue]) *Dimension[TValue, any] {
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
		var old *std.Data[TValue]
		if d.Current != nil {
			old = d.Current
		}
		d.Timeline = append(d.Timeline, data)
		d.Current = &data
		change(ctx, old, d.Current)
		d.Mutex.Unlock()
	}
	d.Stimulator = engine.Stimulate(f, potential, muted)
	return &d
}
