package core

import "time"

// runtimeStats provides information about how long something took to execute relative to its impulse moment.
type runtimeStats struct {
	// Inception is the moment the impulse started.
	Inception time.Time

	// Start is the moment execution began.
	Start time.Time

	// End is the moment execution completed.
	End time.Time

	// RefractoryPeriod is the amount of time between the end of the last impulse and inception of the current.
	RefractoryPeriod time.Duration
}
