package core

import (
	"time"
)

// Context represents contextually relevant temporal information.
type Context struct {
	Entity

	// Beat is the engine's current impulse count.
	//
	// 	NOTE:
	//	This will loop over when all activations finish at the same time.
	//  For impulsive activations they never flag as 'Executing', so they naturally don't increment the Beat.
	Beat int

	// Moment is the moment in time this impulse represents.
	Moment time.Time

	// Delta is the amount of time that has passed since the last impulse's Moment.
	Delta time.Duration

	// LastImpulse provides runtime information about the last impulse of the engine.
	LastImpulse runtime

	// LastActivation provides runtime information about the last activation.
	LastActivation runtime
}
