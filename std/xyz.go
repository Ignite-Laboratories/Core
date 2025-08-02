package std

import (
	"fmt"
	"github.com/ignite-laboratories/core/math"
)

// XYZ is a general structure for holding generic (x,y,z) coordinate values.
type XYZ[T math.Numeric] struct {
	X T
	Y T
	Z T
}

// RandomXYZ returns a pseudo-random XYZ[T] of the provided type using math.RandomNumber[T].
//
// If requesting a floating point type, the resulting number will be bounded
// in the fully closed interval [0.0, 1.0]
//
// If requesting an integer type, the resulting number will be bounded
// in the fully closed interval [0, n] - where n is the maximum value of
// the provided type.
func RandomXYZ[T math.Numeric]() XYZ[T] {
	return XYZ[T]{
		X: math.RandomNumber[T](),
		Y: math.RandomNumber[T](),
		Z: math.RandomNumber[T](),
	}
}

// RandomXYZUpTo returns a pseudo-random XYZ[T] of the provided type bounded in the closed interval [0, max].
func RandomXYZUpTo[T math.Numeric](xUpper T, yUpper T, zUpper T) XYZ[T] {
	return XYZ[T]{
		X: math.RandomNumberRange[T](math.Tuple[T]{B: xUpper}),
		Y: math.RandomNumberRange[T](math.Tuple[T]{B: yUpper}),
		Z: math.RandomNumberRange[T](math.Tuple[T]{B: zUpper}),
	}
}

// RandomXYZRange returns a pseudo-random XYZ[T] of the provided type bounded in the closed interval [min, max].
func RandomXYZRange[T math.Numeric](xRange math.Tuple[T], yRange math.Tuple[T], zRange math.Tuple[T]) XYZ[T] {
	return XYZ[T]{
		X: math.RandomNumberRange[T](xRange),
		Y: math.RandomNumberRange[T](yRange),
		Z: math.RandomNumberRange[T](zRange),
	}
}

// NormalizeXYZ32 returns an XYZ[float32] ranging from 0.0-1.0.
func NormalizeXYZ32[T math.Integer](source XYZ[T]) XYZ[float32] {
	return XYZ[float32]{
		X: math.NormalizeToFloat32(source.X),
		Y: math.NormalizeToFloat32(source.Y),
		Z: math.NormalizeToFloat32(source.Z),
	}
}

// NormalizeXYZ64 returns an XYZ[float64] ranging from 0.0-1.0.
func NormalizeXYZ64[T math.Integer](source XYZ[T]) XYZ[float64] {
	return XYZ[float64]{
		X: math.NormalizeToFloat64(source.X),
		Y: math.NormalizeToFloat64(source.Y),
		Z: math.NormalizeToFloat64(source.Z),
	}
}

// ScaleToTypeXYZ32 returns a scaled value of the provided type in the range [0, T.MaxValue].
//
// NOTE: This will panic if the provided value is greater than the maximum value of the provided type.
func ScaleToTypeXYZ32[TOut math.Integer](source XYZ[float32]) XYZ[TOut] {
	return XYZ[TOut]{
		X: math.ScaleFloat32ToType[TOut](source.X),
		Y: math.ScaleFloat32ToType[TOut](source.Y),
		Z: math.ScaleFloat32ToType[TOut](source.Z),
	}
}

// ScaleToTypeXYZ64 returns a scaled value of the provided type in the range [0, T.MaxValue].
//
// NOTE: This will panic if the provided value is greater than the maximum value of the provided type.
func ScaleToTypeXYZ64[TOut math.Integer](source XYZ[float64]) XYZ[TOut] {
	return XYZ[TOut]{
		X: math.ScaleFloat64ToType[TOut](source.X),
		Y: math.ScaleFloat64ToType[TOut](source.Y),
		Z: math.ScaleFloat64ToType[TOut](source.Z),
	}
}

// XYZComparator returns if the two XYZ values are equal in values.
func XYZComparator[T math.Numeric](a XYZ[T], b XYZ[T]) bool {
	return a.X == b.X && a.Y == b.Y && a.Z == b.Z
}

func (c XYZ[T]) String() string {
	return fmt.Sprintf("(%v, %v, %v)", c.X, c.Y, c.Z)
}

/**
Swizzling
*/

func (c XYZ[T]) XX() (T, T) { return c.X, c.X }
func (c XYZ[T]) XY() (T, T) { return c.X, c.Y }
func (c XYZ[T]) XZ() (T, T) { return c.X, c.Z }
func (c XYZ[T]) YX() (T, T) { return c.Y, c.X }
func (c XYZ[T]) YY() (T, T) { return c.Y, c.Y }
func (c XYZ[T]) YZ() (T, T) { return c.Y, c.Z }
func (c XYZ[T]) ZX() (T, T) { return c.Z, c.X }
func (c XYZ[T]) ZY() (T, T) { return c.Z, c.Y }
func (c XYZ[T]) ZZ() (T, T) { return c.Z, c.Z }

func (c XYZ[T]) XXX() (T, T, T) { return c.X, c.X, c.X }
func (c XYZ[T]) XXY() (T, T, T) { return c.X, c.X, c.Y }
func (c XYZ[T]) XXZ() (T, T, T) { return c.X, c.X, c.Z }
func (c XYZ[T]) XYX() (T, T, T) { return c.X, c.Y, c.X }
func (c XYZ[T]) XYY() (T, T, T) { return c.X, c.Y, c.Y }
func (c XYZ[T]) XYZ() (T, T, T) { return c.X, c.Y, c.Z }
func (c XYZ[T]) XZX() (T, T, T) { return c.X, c.Z, c.X }
func (c XYZ[T]) XZY() (T, T, T) { return c.X, c.Z, c.Y }
func (c XYZ[T]) XZZ() (T, T, T) { return c.X, c.Z, c.Z }
func (c XYZ[T]) YXX() (T, T, T) { return c.Y, c.X, c.X }
func (c XYZ[T]) YXY() (T, T, T) { return c.Y, c.X, c.Y }
func (c XYZ[T]) YXZ() (T, T, T) { return c.Y, c.X, c.Z }
func (c XYZ[T]) YYX() (T, T, T) { return c.Y, c.Y, c.X }
func (c XYZ[T]) YYY() (T, T, T) { return c.Y, c.Y, c.Y }
func (c XYZ[T]) YYZ() (T, T, T) { return c.Y, c.Y, c.Z }
func (c XYZ[T]) YZX() (T, T, T) { return c.Y, c.Z, c.X }
func (c XYZ[T]) YZY() (T, T, T) { return c.Y, c.Z, c.Y }
func (c XYZ[T]) YZZ() (T, T, T) { return c.Y, c.Z, c.Z }
func (c XYZ[T]) ZXX() (T, T, T) { return c.Z, c.X, c.X }
func (c XYZ[T]) ZXY() (T, T, T) { return c.Z, c.X, c.Y }
func (c XYZ[T]) ZXZ() (T, T, T) { return c.Z, c.X, c.Z }
func (c XYZ[T]) ZYX() (T, T, T) { return c.Z, c.Y, c.X }
func (c XYZ[T]) ZYY() (T, T, T) { return c.Z, c.Y, c.Y }
func (c XYZ[T]) ZYZ() (T, T, T) { return c.Z, c.Y, c.Z }
func (c XYZ[T]) ZZX() (T, T, T) { return c.Z, c.Z, c.X }
func (c XYZ[T]) ZZY() (T, T, T) { return c.Z, c.Z, c.Y }
func (c XYZ[T]) ZZZ() (T, T, T) { return c.Z, c.Z, c.Z }
