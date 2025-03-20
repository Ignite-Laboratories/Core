// Package when provides a way of creating action potentials.
package when

import "github.com/ignite-laboratories/core"

// Always provides a potential that always fires the action.
func Always(action core.Action) core.Action {
	return action
}

// Never provides a potential that never fires.
func Never() core.Action {
	return func(ctx core.Context) {}
}

// Downbeats provides a potential that fires the action when the beat is 0.
func Downbeats(action core.Action) core.Action {
	return func(ctx core.Context) {
		if ctx.Beat == 0 {
			action(ctx)
		}
	}
}

// Even provides a potential that fires the action when the beat is even.
func Even(action core.Action) core.Action {
	return func(ctx core.Context) {
		if ctx.Beat%2 == 0 {
			action(ctx)
		}
	}
}

// Odd provides a potential that returns true when the beat is odd.
func Odd(action core.Action) core.Action {
	return func(ctx core.Context) {
		if ctx.Beat%2 != 0 {
			action(ctx)
		}
	}
}

// Modulo provides the following potential: "beat % value == 0".
func Modulo(value int, action core.Action) core.Action {
	return func(ctx core.Context) {
		if ctx.Beat%value == 0 {
			action(ctx)
		}
	}
}

// On provides the following potential: "beat == value".
func On(beat int, action core.Action) core.Action {
	return func(ctx core.Context) {
		if ctx.Beat == beat {
			action(ctx)
		}
	}
}
