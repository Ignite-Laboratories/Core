package std

import (
	"github.com/ignite-laboratories/core"
)

// Data is a contextual point value in time.
type Data[T any] struct {
	core.Context

	// Point is the recorded value of this contextual moment.
	Point T
}
