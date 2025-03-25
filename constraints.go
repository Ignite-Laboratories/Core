// Package constraints provides a singular point for referencing 'Numeric' types of both Integer and Float.
package core

import (
	"golang.org/x/exp/constraints"
)

// Numeric represents any integer or floating-point type.
type Numeric interface {
	constraints.Integer | constraints.Float
}
