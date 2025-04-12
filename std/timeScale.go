package std

import (
	"github.com/ignite-laboratories/core"
	"time"
)

// TimeScale represents a pairing of duration and an abstract "height."
type TimeScale[T core.Numeric] struct {
	Duration time.Duration
	Height   T
}
