package core

import "time"

// runtime provides information about how long an activation took to execute relative to its Inception.
type runtime struct {
	// RefractoryPeriod is the amount of time between the End of the last activation and Inception.
	RefractoryPeriod time.Duration

	// Inception is the moment the impulse started.
	Inception time.Time

	// Start is the moment of activation.
	Start time.Time

	// End is the moment activation completed.
	End time.Time
}
