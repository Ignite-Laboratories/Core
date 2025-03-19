package core

import (
	"time"
)

// Context represents contextually relevant temporal information.
type Context struct {
	Entity

	// Beat is the engine's current impulse count.
	//
	// 	NOTE: This will loop over when all activations finish at the same time.
	Beat int

	// Moment is the moment in time this impulse represents.
	Moment time.Time

	// Delta is the amount of time that has passed since the last impulse's Moment.
	Delta time.Duration

	// ImpulseStats provides information about the last impulse of the engine.
	ImpulseStats runtimeStats

	// ActivationStats provides information about the last activation.
	ActivationStats runtimeStats
}
