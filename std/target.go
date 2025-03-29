package std

// TargetFunc functions return a pointer to a value.
type TargetFunc[TValue any] func() *TValue

func Target[TValue any](val *TValue) TargetFunc[TValue] {
	return func() *TValue {
		return val
	}
}
