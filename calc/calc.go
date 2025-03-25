package calc

import "github.com/ignite-laboratories/core"

// Difference returns a-b.
func Difference[TValue core.Numeric](a TValue, b TValue) TValue {
	return a - b
}

// Delta returns b-a.
func Delta[TValue core.Numeric](a TValue, b TValue) TValue {
	return b - a
}
