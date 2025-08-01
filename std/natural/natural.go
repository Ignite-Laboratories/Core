package natural

import (
	"github.com/ignite-laboratories/core/std"
)

// From takes a Measurement of the provided unsigned integer value as a Natural number.
func From(value uint) std.Natural {
	return std.Natural{
		Measurement: std.FromBytes(internal.Measure(value)[0]...),
	}
}

// FromString creates a new Natural measurement that represents the provided base-encoded string.
//
// NOTE: The input string must be encoded as expected by Real.SetBase()
func FromString(base byte, value string) std.Natural {
	// TODO: Implement this
	panic("unsupported")
}
