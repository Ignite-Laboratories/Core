package core

// Activate provides a set of commonly used Potential functions.
var Activate _activate

// _activate is a placeholder struct for Activate.
type _activate struct{}

// Always provides a Potential that always returns true.
func (w _activate) Always() Potential {
	return func(ctx Context) bool {
		return true
	}
}

// Never provides a Potential that always returns false.
func (w _activate) Never() Potential {
	return func(ctx Context) bool {
		return false
	}
}

// Downbeats provides a Potential that returns true when the beat is 0.
func (w _activate) Downbeats() Potential {
	return func(ctx Context) bool {
		return ctx.Beat == 0
	}
}

// EvenBeats provides a Potential that returns true when the beat is even.
func (w _activate) EvenBeats() Potential {
	return func(ctx Context) bool {
		return ctx.Beat%2 == 0
	}
}

// OddBeats provides a Potential that returns true when the beat is odd.
func (w _activate) OddBeats() Potential {
	return func(ctx Context) bool {
		return ctx.Beat%2 != 0
	}
}

// Modulo provides the following Potential: "beat % value == 0".
func (w _activate) Modulo(value int) Potential {
	return func(ctx Context) bool {
		return ctx.Beat%value == 0
	}
}

// On provides the following Potential: "beat == value".
func (w _activate) On(beat int) Potential {
	return func(ctx Context) bool {
		return ctx.Beat == beat
	}
}
