// Package when provides helpers for creating potential functions.
package when

import (
	"github.com/ignite-laboratories/core"
)

// Always provides a potential that always fires.
func Always(ctx core.Context) bool {
	return true
}

// Never provides a potential that never fires.
func Never() bool {
	return false
}

// Downbeats provides a potential that fires when the beat is 0.
func Downbeats(ctx core.Context) bool {
	return ctx.Beat == 0
}

// Even provides a potential that fires when the beat is even.
func Even(ctx core.Context) bool {
	return ctx.Beat%2 == 0
}

// Odd provides a potential that returns true when the beat is odd.
func Odd(ctx core.Context) bool {
	return ctx.Beat%2 != 0
}

// Modulo provides the following potential: "beat % value == 0"
func Modulo(value int) core.Potential {
	return func(ctx core.Context) bool {
		return ctx.Beat%value == 0
	}
}

// Over provides the following potential: "beat > value"
func Over(value int) core.Potential {
	return func(ctx core.Context) bool {
		return ctx.Beat > value
	}
}

// On provides the following potential: "beat == value"
func On(beat int) core.Potential {
	return func(ctx core.Context) bool {
		return ctx.Beat == beat
	}
}

// Under provides the following potential: "beat < value"
func Under(value int) core.Potential {
	return func(ctx core.Context) bool {
		return ctx.Beat < value
	}
}
