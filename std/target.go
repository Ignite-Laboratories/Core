package std

import "github.com/ignite-laboratories/core"

// TargetFunc functions return a pointer to a value.
type TargetFunc[TValue any] func() *TValue

// Target returns a function that retrieves a reference to the target on demand.
func Target[TValue any](val *TValue) TargetFunc[TValue] {
	return func() *TValue {
		return val
	}
}

// BooleanTarget functions provide a potential from a boolean reference.
func BooleanTarget(value *bool) core.Potential {
	return func(ctx core.Context) bool {
		return *value
	}
}
