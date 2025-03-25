// Package condition provides a set of helper functions for creating conditional potentials
package condition

import (
	"github.com/ignite-laboratories/core"
	"math"
	"time"
)

// Frequency provides a potential that activates at the provided frequency, in Hertz.
func Frequency(hertz *float64) core.Potential {
	hz := *hertz
	if hz <= 0 {
		// No division by zero
		hz = math.SmallestNonzeroFloat64
	}
	secondsPerCycle := 1 / hz
	nanosecondsPerCycle := secondsPerCycle * 1e9
	duration := time.Duration(nanosecondsPerCycle)
	return Duration(&duration)
}

// Duration provides the following potential: "time.Now().Sub(ctx.LastActivation.Inception) > duration"
func Duration(duration *time.Duration) core.Potential {
	return func(ctx core.Context) bool {
		return time.Now().Sub(ctx.LastActivation.Inception) > *duration
	}
}

// Pace provides a potential that counts to the provided value before returning true.
//
// NOTE: This is a impulse blocking operation!
//
// This is a rudimentary way of slowing an activation off the impulse moment.
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
