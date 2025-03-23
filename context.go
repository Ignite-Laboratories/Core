package core

import (
	"time"
)

// Context represents contextually relevant temporal information.
type Context struct {
	Entity

	// Beat is the engine's current impulse count.
	Beat int

	// Moment is the moment in time this impulse represents.
	Moment time.Time

	// Period is the amount of time that has passed since the last impulse's Moment.
	Period time.Duration

	// LastImpulse provides runtime information about the last impulse of the engine.
	LastImpulse runtime

	// LastActivation provides runtime information about the last activation.
	LastActivation runtime
}
