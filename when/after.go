package when

import (
	"github.com/ignite-laboratories/core"
	"time"
)

type _after struct {
}

// After provides time-oriented potentials.
var After _after

// Period provides the following potential: "time.Now().Sub(ctx.LastActivation.Inception) > duration"
func (a _after) Period(duration *time.Duration) core.Potential {
	return func(ctx core.Context) bool {
		return time.Now().Sub(ctx.LastActivation.Inception) > *duration
	}
}

// RefractionPeriod provides the following potential: "time.Now().Sub(ctx.LastActivation.End) > duration"
func (a _after) RefractionPeriod(duration *time.Duration) core.Potential {
	return func(ctx core.Context) bool {
		return time.Now().Sub(ctx.LastActivation.End) > *duration
	}
}

// Count provides a potential that counts to the provided value before returning true.
//
// This is a rudimentary way of slowing an activation off the impulse moment.
func (a _after) Count(value *uint64) core.Potential {
	return func(ctx core.Context) bool {
		for i := uint64(0); i < *value; i++ {
		}
		return true
	}
}
