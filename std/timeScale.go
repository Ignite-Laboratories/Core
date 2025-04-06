package std

import (
	"github.com/ignite-laboratories/core"
	"time"
)

// TimeScale represents a pairing of duration and abstract "height," as determined by the provided numeric type.
type TimeScale[T core.Numeric] struct {
	Duration time.Duration
	Height   T
}
