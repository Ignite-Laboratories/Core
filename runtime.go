package core

import "time"

// runtime provides information about how long something took to run relative to its inception.
type runtime struct {
	// RefractoryPeriod is the amount of time between the end of the last activation and inception.
	RefractoryPeriod time.Duration

	// Inception is the moment the impulse started.
	Inception time.Time

	// Start is the moment of activation.
	Start time.Time

	// End is the moment activation completed.
	End time.Time
}
