package calc

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/condition"
)

// NewCalculation creates a dimension that performs a calculation for every impulse that the potential returns true.
//
// NOTE: The potential function gates the creation of timeline indexes!
// This can adjust the "resolution" of output data =)
//
// Muted indicates if the stimulator of this dimension should be created muted.
func NewCalculation[TValue any](engine *core.Engine, potential core.Potential, muted bool, calculate core.CalculatePoint[TValue]) *core.Dimension[TValue, any] {
	d := core.Dimension[TValue, any]{}
	d.ID = core.NextID()
	d.Trimmer = engine.Loop(d.Trim, condition.Always, false)
	d.Stimulator = engine.Stimulate(func(ctx core.Context) {
		value := calculate(ctx)
		data := core.Data[TValue]{
			Context: ctx,
			Point:   value,
		}
		d.Mutex.Lock()
		d.Timeline = append(d.Timeline, data)
		d.Current = data
		d.Mutex.Unlock()
	}, potential, muted)
	return &d
}
