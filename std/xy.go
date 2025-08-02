package std

import (
	"fmt"
	"github.com/ignite-laboratories/core/math"
)

// XY is a general structure for holding generic (x,y) coordinate values.
type XY[T math.Numeric] struct {
	X T
	Y T
}

// RandomXY returns a pseudo-random XY[T] of the provided type using math.RandomNumber[T].
//
// If requesting a floating point type, the resulting number will be bounded
// in the fully closed interval [0.0, 1.0]
//
// If requesting an integer type, the resulting number will be bounded
// in the fully closed interval [0, n] - where n is the maximum value of
// the provided type.
func RandomXY[T math.Numeric]() XY[T] {
	return XY[T]{
		X: math.RandomNumber[T](),
		Y: math.RandomNumber[T](),
	}
}

// RandomXYUpTo returns a pseudo-random XY[T] of the provided type bounded in the closed interval [0, max].
func RandomXYUpTo[T math.Numeric](xUpper T, yUpper T) XY[T] {
	return XY[T]{
		X: math.RandomNumberRange[T](math.Tuple[T]{B: xUpper}),
		Y: math.RandomNumberRange[T](math.Tuple[T]{B: yUpper}),
	}
}

// RandomXYRange returns a pseudo-random XY[T] of the provided type bounded in the closed interval [min, max].
func RandomXYRange[T math.Numeric](xRange math.Tuple[T], yRange math.Tuple[T]) XY[T] {
	return XY[T]{
		X: math.RandomNumberRange[T](xRange),
		Y: math.RandomNumberRange[T](yRange),
	}
}

// NormalizeXY32 returns an XY[float32] ranging from 0.0-1.0.
func NormalizeXY32[T math.Integer](source XY[T]) XY[float32] {
	return XY[float32]{
		X: math.NormalizeToFloat32(source.X),
		Y: math.NormalizeToFloat32(source.Y),
	}
}

// NormalizeXY64 returns an XYZ[float64] ranging from 0.0-1.0.
func NormalizeXY64[T math.Integer](source XY[T]) XY[float64] {
	return XY[float64]{
		X: math.NormalizeToFloat64(source.X),
		Y: math.NormalizeToFloat64(source.Y),
	}
}

// ScaleToTypeXY32 returns a scaled value of the provided type in the range [0, T.MaxValue].
//
// NOTE: This will panic if the provided value is greater than the maximum value of the provided type.
func ScaleToTypeXY32[TOut math.Integer](source XY[float32]) XY[TOut] {
	return XY[TOut]{
		X: math.ScaleFloat32ToType[TOut](source.X),
		Y: math.ScaleFloat32ToType[TOut](source.Y),
	}
}

// ScaleToTypeXY64 returns a scaled value of the provided type in the range [0, T.MaxValue].
//
// NOTE: This will panic if the provided value is greater than the maximum value of the provided type.
func ScaleToTypeXY64[TOut math.Integer](source XY[float64]) XY[TOut] {
	return XY[TOut]{
		X: math.ScaleFloat64ToType[TOut](source.X),
		Y: math.ScaleFloat64ToType[TOut](source.Y),
	}
}

// XYComparator returns if the two XY values are equal in values.
func XYComparator[T math.Numeric](a XY[T], b XY[T]) bool {
	return a.X == b.X && a.Y == b.Y
}

func (c XY[T]) String() string {
	return fmt.Sprintf("(%v, %v)", c.X, c.Y)
}
