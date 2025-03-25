package core

import (
	"sync"
	"time"
)

// Dimension is a way of observing a target value across time, limited to a window of observance.
type Dimension[TValue any, TCache any] struct {
	Entity

	// Current is the currently held value of this dimension.
	Current Data[TValue]

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

// Trim removes anything on the timeline that is older than the dimension's window of observance.
func (o *Dimension[TValue, TCache]) Trim(ctx Context) {
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
