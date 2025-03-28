package temporal

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/when"
)

// Analyzer creates a new dimension that records the result of the provided integral function cyclically.
// The integral function is always called with the exact timeline data since the last analysis started.
//
// NOTE: The potential function gates analysis.
// This can adjust "reactivity" to input data =)
//
// Muted indicates if the stimulator of this dimension should be created muted.
func Analyzer[TSource any, TValue any, TCache any](engine *core.Engine, potential core.Potential, muted bool, integrate core.Integral[std.Data[TSource], TValue, TCache], target *Dimension[TSource, any]) *Dimension[TValue, TCache] {
	d := Dimension[TValue, TCache]{}
	d.ID = core.NextID()
	d.Window = core.DefaultWindow
	d.Trimmer = engine.Loop(d.Trim, when.Always, false)
	d.lastCycle = core.Inception

	d.Stimulator = engine.Loop(func(ctx core.Context) {
		// Get target timeline data
		target.Mutex.Lock()
		data := make([]std.Data[TSource], len(target.Timeline))
		copy(data, target.Timeline)
		target.Mutex.Unlock()

		// Trim any indices that were handled by the last analysis
		var trimCount int
		for _, v := range data {
			if v.Moment.After(d.lastCycle) {
				break
			}
			trimCount++
		}
		data = data[trimCount:]

		// Save off the last moment for the next cycle
		lastCycle := d.lastCycle
		if len(data) > 0 {
			lastCycle = data[len(data)-1].Moment
		}

		// Perform integration
		point := integrate(ctx, d.Cache, data)
		out := std.Data[TValue]{
			Context: ctx,
			Point:   point,
		}

		// Record the result
		d.lastCycle = lastCycle
		d.update(out)
	}, potential, muted)
	return &d
}
