package temporal

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
)

// Difference returns a-b.
func Difference[TValue core.Numeric](a TValue, b TValue) TValue {
	return a - b
}

// Delta returns b-a.
func Delta[TValue core.Numeric](a TValue, b TValue) TValue {
	return b - a
}

// Change functions are called when a dimension's current point value changes.
type Change[TValue any] func(ctx core.Context, old *std.Data[TValue], current *std.Data[TValue])
