package core

import (
	"golang.org/x/exp/constraints"
)

// Numeric represents any integer or floating-point type.
type Numeric interface {
	constraints.Integer | constraints.Float
}

// Integer represents any integer type.
type Integer interface {
	constraints.Integer
}

// Float represents any floating-point type.
type Float interface {
	constraints.Float
}
