package core

import (
	"golang.org/x/exp/constraints"
)

// Numeric represents any integer or floating-point type.
type Numeric interface {
	constraints.Integer | constraints.Float
}
