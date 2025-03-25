// Package condition provides a set of helper functions for creating conditional potentials
package condition

import (
	"github.com/ignite-laboratories/core"
	"time"
)

// Frequency provides a potential that activates at the specified Hertz.
func Frequency(hertz *float64) core.Potential {
	d := core.HertzToDuration(*hertz)
	return Duration(&d)
}

// Duration provides the following potential:
//
//	time.Now().Sub(ctx.LastActivation.Inception) > duration
func Duration(duration *time.Duration) core.Potential {
	return func(ctx core.Context) bool {
		return time.Now().Sub(ctx.LastActivation.Inception) > *duration
	}
}

// Pace provides a potential that counts to the provided value before returning true.
//
// NOTE: This is a impulse slowing operation!
func Pace(value *uint64) core.Potential {
	return func(ctx core.Context) bool {
		for i := uint64(0); i < *value; i++ {
		}
		return true
	}
}

// Always provides a potential that always fires.
func Always(ctx core.Context) bool {
	return true
}

// Never provides a potential that never fires.
func Never() bool {
	return false
}
