// Package when provides helpers for creating potential functions.
package when

import (
	"github.com/ignite-laboratories/core"
	"time"
)

// Always provides a potential that always fires the action.
func Always(ctx core.Context) bool {
	return true
}

// Never provides a potential that never fires.
func Never() bool {
	return false
}

// Downbeats provides a potential that fires the action when the beat is 0.
func Downbeats(ctx core.Context) bool {
	return ctx.Beat == 0
}

// Even provides a potential that fires the action when the beat is even.
func Even(ctx core.Context) bool {
	return ctx.Beat%2 == 0
}

// Odd provides a potential that returns true when the beat is odd.
func Odd(ctx core.Context) bool {
	return ctx.Beat%2 != 0
}

// Modulo provides the following potential: "beat % value == 0".
func Modulo(value int) core.Potential {
	return func(ctx core.Context) bool {
		return ctx.Beat%value == 0
	}
}

// On provides the following potential: "beat == value".
func On(beat int) core.Potential {
	return func(ctx core.Context) bool {
		return ctx.Beat == beat
	}
}

type _after struct {
}

// After provides potentials relative to something external to the context meeting a condition.
var After _after

// Period provides a potential that checks if the amount of time since
// the last activation's -inception- exceeds 'duration' before re-activation.
func (a _after) Period(duration time.Duration) core.Potential {
	return func(ctx core.Context) bool {
		return time.Now().Sub(ctx.LastActivation.Inception) > duration
	}
}

// RefractionPeriod provides a potential that checks if the amount of time since
// the last activation's -end- exceeds 'duration' before re-activation.
func (a _after) RefractionPeriod(duration time.Duration) core.Potential {
	return func(ctx core.Context) bool {
		return time.Now().Sub(ctx.LastActivation.End) > duration
	}
}

// Count provides a potential that counts to the provided value before activation.
func (a _after) Count(value uint64) core.Potential {
	return func(ctx core.Context) bool {
		for i := uint64(0); i < value; i++ {
		}
		return true
	}
}
