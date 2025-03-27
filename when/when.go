// Package when provides a set of helper functions for creating conditional potentials
package when

import (
	"github.com/ignite-laboratories/core"
	"time"
)

// Frequency provides a potential that activates at the specified Hertz.
func Frequency(hertz *float64) core.Potential {
	d := core.HertzToDuration(*hertz)
	return Duration(&d)
}

// ResonantWith provides a potential that activates at a sympathetic frequency to the provided source.
//
// A sympathetic frequency is a multiple of the source - thus, source / sympathetic.
//
// If you would like something to resonate at half the rate of the source frequency, provide a sympathetic value of 2.0
func ResonantWith(source *float64, sympathetic *float64) core.Potential {
	resonance := *source / *sympathetic
	return Frequency(&resonance)
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
