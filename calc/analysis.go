package calc

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/condition"
)

// NewAnalysis creates a new dimension that records the result of the provided integral function cyclically.
// The integral function is always called with the exact timeline data since the last analysis started.
//
// NOTE: The potential function gates analysis.
// This can adjust "reactivity" to input data =)
//
// Muted indicates if the stimulator of this dimension should be created muted.
func NewAnalysis[TSource any, TValue any, TCache any](engine *core.Engine, potential core.Potential, muted bool, integrate core.Integral[core.Data[TSource], TValue, TCache], target *core.Dimension[TSource, TCache]) *core.Dimension[TValue, TCache] {
	d := core.Dimension[TValue, TCache]{}
	d.ID = core.NextID()
	d.Trimmer = engine.Loop(d.Trim, condition.Always, false)

	d.Stimulator = engine.Loop(func(ctx core.Context) {
		// Get target timeline data
		target.Mutex.Lock()
		last := target.Current
		var data []core.Data[TSource]
		copy(data, target.Timeline)
		target.Mutex.Unlock()

		// Trim any indices that were handled by the last analysis
		var trimCount int
		for i, v := range data {
			if v.Moment.After(last.Moment) {
				trimCount = i - 1
				break
			}
		}
		if trimCount < 0 {
			trimCount = 0
		}
		data = data[trimCount:]

		// Perform integration
		out := core.Data[TValue]{
			Context: ctx,
			Point:   integrate(ctx, d.Cache, data),
		}

		// Record the result
		d.Mutex.Lock()
		d.Timeline = append(d.Timeline, out)
		d.Current = out
		d.Mutex.Unlock()
	}, potential, muted)
	return &d
}
