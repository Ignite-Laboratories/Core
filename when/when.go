package when

import (
	"github.com/ignite-laboratories/core"
	"time"
)

// Frequency provides a potential that activates at the specified frequency (Hertz).
func Frequency(hertz *float64) core.Potential {
	return func(ctx core.Context) bool {
		d := core.HertzToDuration(*hertz)
		return time.Now().Sub(ctx.LastActivation.Inception) > d
	}
}

// Resonant provides a potential that activates at a sympathetic frequency (Hertz) to the provided source.
//
//	Resonance = Source / Subdivision
func Resonant(source *float64, subdivision *float64) core.Potential {
	return func(ctx core.Context) bool {
		resonance := *source / *subdivision
		d := core.HertzToDuration(resonance)
		return time.Now().Sub(ctx.LastActivation.Inception) > d
	}
}

// HalfSpeed provides a potential that activates at half the rate of the source frequency (Hertz).
func HalfSpeed(hertz *float64) core.Potential {
	half := 2.0
	return Resonant(hertz, &half)
}

// QuarterSpeed provides a potential that activates at half the rate of the source frequency (Hertz).
func QuarterSpeed(hertz *float64) core.Potential {
	half := 4.0
	return Resonant(hertz, &half)
}

// EighthSpeed provides a potential that activates at half the rate of the source frequency (Hertz).
func EighthSpeed(hertz *float64) core.Potential {
	half := 8.0
	return Resonant(hertz, &half)
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
