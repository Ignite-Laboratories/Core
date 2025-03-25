package core

import (
	"sync"
	"time"
)

// Dimension is a way of observing a target value across time, limited to a window of observance.
type Dimension[TValue any, TCache any] struct {
	Entity

	// Value is the current value of this dimension.
	Value Data[TValue]

	// Cache is a place where a looping stimulator can save information for the next activation of the loop.
	Cache *TCache

	// Timeline is the historical values of this dimension.
	Timeline []Data[TValue]

	// Window is the duration to hold onto recorded values for.
	Window time.Duration

	// Mutex should be locked for any operations that need a momentary snapshot of the timeline.
	Mutex sync.Mutex

	// Stimulator is the activation that drives the function that populates this timeline.
	Stimulator *Activation

	// Trimmer is the activation that trims the timeline of entries beyond the window of observance.
	Trimmer *Activation
}

// trim removes anything on the timeline that is older than the dimension's window of observance.
func (o *Dimension[TValue, TCache]) trim(ctx Context) {
	o.Mutex.Lock()
	var trimCount int
	for i, v := range o.Timeline {
		if time.Now().Sub(v.Moment) < o.Window {
			trimCount = i
			break
		}
	}
	o.Timeline = o.Timeline[trimCount:]
	o.Mutex.Unlock()

	for o.Stimulator.Muted {
		// If the stimulator is muted, don't bother looping until it un-mutes
	}
}

// Mute suppresses the stimulator of this dimension.
func (d *Dimension[TValue, TCache]) Mute() {
	d.Stimulator.Muted = true
}

// Unmute un-suppresses the stimulator of this dimension.
func (d *Dimension[TValue, TCache]) Unmute() {
	d.Stimulator.Muted = false
}

// NewObservation creates a dimension that records the target value across time, if the provided potential returns true.
//
// NOTE: The potential function gates the creation of timeline indexes!
// This can adjust the "resolution" of output data =)
//
// Muted indicates if the stimulator of this dimension should be created muted.
func NewObservation[TValue any](engine *Engine, target *TValue, potential Potential, muted bool) *Dimension[TValue, any] {
	d := Dimension[TValue, any]{}
	d.ID = NextID()
	d.Trimmer = engine.Loop(d.trim, alwaysFire, false)
	d.Stimulator = engine.Stimulate(func(ctx Context) {
		d.Mutex.Lock()
		data := Data[TValue]{
			Context: ctx,
			Value:   *target,
		}
		d.Timeline = append(d.Timeline, data)
		d.Value = data
		d.Mutex.Unlock()
	}, potential, muted)
	return &d
}

// NewCalculation creates a dimension that performs a calculation for every impulse that the potential returns true.
//
// NOTE: The potential function gates the creation of timeline indexes!
// This can adjust the "resolution" of output data =)
//
// Muted indicates if the stimulator of this dimension should be created muted.
func NewCalculation[TValue any](engine *Engine, calculate CalculatePoint[TValue], potential Potential, muted bool) *Dimension[TValue, any] {
	d := Dimension[TValue, any]{}
	d.ID = NextID()
	d.Trimmer = engine.Loop(d.trim, alwaysFire, false)
	d.Stimulator = engine.Stimulate(func(ctx Context) {
		d.Mutex.Lock()
		value := calculate(ctx)
		data := Data[TValue]{
			Context: ctx,
			Value:   value,
		}
		d.Timeline = append(d.Timeline, data)
		d.Value = data
		d.Mutex.Unlock()
	}, potential, muted)
	return &d
}

// NewAnalysis creates a new dimension that records the result of the provided integral function cyclically.
// The integral function is always called with the exact timeline data since the last analysis started.
//
// NOTE: The potential function gates analysis.
// This can adjust "reactivity" to input data =)
//
// Muted indicates if the stimulator of this dimension should be created muted.
func NewAnalysis[TIn any, TOut any, TCache any](engine *Engine, target *Dimension[TIn, TCache], integrate Integral[Data[TIn], TOut, TCache], potential Potential, muted bool) *Dimension[TOut, TCache] {
	d := Dimension[TOut, TCache]{}
	d.ID = NextID()
	d.Trimmer = engine.Loop(d.trim, alwaysFire, false)

	d.Stimulator = engine.Loop(func(ctx Context) {
		// Get target timeline data
		target.Mutex.Lock()
		last := target.Value
		var data []Data[TIn]
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
		out := Data[TOut]{
			Context: ctx,
			Value:   integrate(ctx, d.Cache, data),
		}

		// Record the result
		d.Mutex.Lock()
		d.Timeline = append(d.Timeline, out)
		d.Value = out
		d.Mutex.Unlock()
	}, potential, muted)
	return &d
}
