// Package after provides helpers for creating delayed potentials
package after

import (
	"github.com/ignite-laboratories/core"
	"time"
)

// Period provides the following potential: "time.Now().Sub(ctx.LastActivation.Inception) > duration"
func Period(duration *time.Duration) core.Potential {
	return func(ctx core.Context) bool {
		return time.Now().Sub(ctx.LastActivation.Inception) > *duration
	}
}

// RefractionPeriod provides the following potential: "time.Now().Sub(ctx.LastActivation.End) > duration"
func RefractionPeriod(duration *time.Duration) core.Potential {
	return func(ctx core.Context) bool {
		return time.Now().Sub(ctx.LastActivation.End) > *duration
	}
}

// Pacing provides a potential that counts to the provided value before returning true.
//
// This is a rudimentary way of slowing an activation off the impulse moment.
func Pacing(value *uint64) core.Potential {
	return func(ctx core.Context) bool {
		for i := uint64(0); i < *value; i++ {
		}
		return true
	}
}
