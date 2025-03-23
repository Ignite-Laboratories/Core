package when

import (
	"github.com/ignite-laboratories/core"
	"time"
)

type _after struct {
}

// After provides time-oriented potentials.
var After _after

// Period provides a potential that checks if the amount of time since
// the last activation's -inception- exceeds 'duration' before re-activation.
func (a _after) Period(duration *time.Duration) core.Potential {
	return func(ctx core.Context) bool {
		return time.Now().Sub(ctx.LastActivation.Inception) > *duration
	}
}

// RefractionPeriod provides a potential that checks if the amount of time since
// the last activation's -end- exceeds 'duration' before re-activation.
func (a _after) RefractionPeriod(duration *time.Duration) core.Potential {
	return func(ctx core.Context) bool {
		return time.Now().Sub(ctx.LastActivation.End) > *duration
	}
}

// Count provides a potential that counts to the provided value before activation.
func (a _after) Count(value *uint64) core.Potential {
	return func(ctx core.Context) bool {
		for i := uint64(0); i < *value; i++ {
		}
		return true
	}
}
