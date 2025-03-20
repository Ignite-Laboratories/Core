package core

import "time"

// runtime provides information about how long an activation took to execute relative to its Inception.
type runtime struct {
	// RefractoryPeriod is the amount of time between the End of the last impulse and Inception.
	RefractoryPeriod time.Duration

	// Inception is the moment the impulse started.
	Inception time.Time

	// Start is the moment execution began.
	Start time.Time

	// End is the moment execution completed.
	End time.Time
}
