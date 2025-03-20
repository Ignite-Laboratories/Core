package potential

import "github.com/ignite-laboratories/core"

// Always provides a potential that always returns true.
func Always(ctx core.Context) bool {
	return true
}

// Never provides a potential that always returns false.
func Never(ctx core.Context) bool {
	return false
}

// Downbeats provides a potential that returns true when the beat is 0.
func Downbeats(ctx core.Context) bool {
	return ctx.Beat == 0
}

// EvenBeats provides a potential that returns true when the beat is even.
func EvenBeats(ctx core.Context) bool {
	return ctx.Beat%2 == 0
}

// OddBeats provides a potential that returns true when the beat is odd.
func OddBeats(ctx core.Context) bool {
	return ctx.Beat%2 != 0
}

// Modulo provides the following potential: "beat % value == 0".
func Modulo(value int) func(ctx core.Context) bool {
	return func(ctx core.Context) bool {
		return ctx.Beat%value == 0
	}
}

// On provides the following potential: "beat == value".
func On(beat int) func(ctx core.Context) bool {
	return func(ctx core.Context) bool {
		return ctx.Beat == beat
	}
}
