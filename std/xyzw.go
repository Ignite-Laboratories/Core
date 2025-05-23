package std

import (
	"fmt"
	"github.com/ignite-laboratories/core"
)

// XYZW is a general structure for holding generic (x,y,z,w) coordinate values.
type XYZW[T core.Numeric] struct {
	X T
	Y T
	Z T
	W T
}

// RandomXYZW returns a pseudo-random XYZW[T] of the provided type using core.RandomNumber[T].
//
// If requesting a floating point type, the resulting number will be bounded
// in the fully closed interval [0.0, 1.0]
//
// If requesting an integer type, the resulting number will be bounded
// in the fully closed interval [0, n] - where n is the maximum value of
// the provided type.
func RandomXYZW[T core.Numeric]() XYZW[T] {
	return XYZW[T]{
		X: core.RandomNumber[T](),
		Y: core.RandomNumber[T](),
		Z: core.RandomNumber[T](),
		W: core.RandomNumber[T](),
	}
}

// RandomXYZWUpTo returns a pseudo-random XYZW[T] of the provided type bounded in the closed interval [0, max].
func RandomXYZWUpTo[T core.Numeric](xUpper T, yUpper T, zUpper T, wUpper T) XYZW[T] {
	return XYZW[T]{
		X: core.RandomNumberRange[T](core.Tuple[T]{B: xUpper}),
		Y: core.RandomNumberRange[T](core.Tuple[T]{B: yUpper}),
		Z: core.RandomNumberRange[T](core.Tuple[T]{B: zUpper}),
		W: core.RandomNumberRange[T](core.Tuple[T]{B: wUpper}),
	}
}

// RandomXYZWRange returns a pseudo-random XYZW[T] of the provided type bounded in the closed interval [min, max].
func RandomXYZWRange[T core.Numeric](xRange core.Tuple[T], yRange core.Tuple[T], zRange core.Tuple[T], wRange core.Tuple[T]) XYZW[T] {
	return XYZW[T]{
		X: core.RandomNumberRange[T](xRange),
		Y: core.RandomNumberRange[T](yRange),
		Z: core.RandomNumberRange[T](zRange),
		W: core.RandomNumberRange[T](wRange),
	}
}

// NormalizeXYZW32 returns an XYZW[float32] ranging from 0.0-1.0.
func NormalizeXYZW32[T core.Integer](source XYZW[T]) XYZW[float32] {
	return XYZW[float32]{
		X: core.NormalizeToFloat32(source.X),
		Y: core.NormalizeToFloat32(source.Y),
		Z: core.NormalizeToFloat32(source.Z),
		W: core.NormalizeToFloat32(source.W),
	}
}

// NormalizeXYZW64 returns an XYZW[float64] ranging from 0.0-1.0.
func NormalizeXYZW64[T core.Integer](source XYZW[T]) XYZW[float64] {
	return XYZW[float64]{
		X: core.NormalizeToFloat64(source.X),
		Y: core.NormalizeToFloat64(source.Y),
		Z: core.NormalizeToFloat64(source.Z),
		W: core.NormalizeToFloat64(source.W),
	}
}

// ScaleToTypeXYZW32 returns a scaled value of the provided type in the range [0, T.MaxValue].
//
// NOTE: This will panic if the provided value is greater than the maximum value of the provided type.
func ScaleToTypeXYZW32[TOut core.Integer](source XYZW[float32]) XYZW[TOut] {
	return XYZW[TOut]{
		X: core.ScaleFloat32ToType[TOut](source.X),
		Y: core.ScaleFloat32ToType[TOut](source.Y),
		Z: core.ScaleFloat32ToType[TOut](source.Z),
		W: core.ScaleFloat32ToType[TOut](source.W),
	}
}

// ScaleToTypeXYZW64 returns a scaled value of the provided type in the range [0, T.MaxValue].
//
// NOTE: This will panic if the provided value is greater than the maximum value of the provided type.
func ScaleToTypeXYZW64[TOut core.Integer](source XYZW[float64]) XYZW[TOut] {
	return XYZW[TOut]{
		X: core.ScaleFloat64ToType[TOut](source.X),
		Y: core.ScaleFloat64ToType[TOut](source.Y),
		Z: core.ScaleFloat64ToType[TOut](source.Z),
		W: core.ScaleFloat64ToType[TOut](source.W),
	}
}

// XYZWComparator returns if the two XYZW values are equal in values.
func XYZWComparator[T core.Numeric](a XYZW[T], b XYZW[T]) bool {
	return a.X == b.X && a.Y == b.Y && a.Z == b.Z && a.W == b.W
}

func (c XYZW[T]) String() string {
	return fmt.Sprintf("(%v, %v, %v, %v)", c.X, c.Y, c.Z, c.W)
}

/**
Swizzling
*/

func (c XYZW[T]) XX() (T, T) { return c.X, c.X }
func (c XYZW[T]) XY() (T, T) { return c.X, c.Y }
func (c XYZW[T]) XZ() (T, T) { return c.X, c.Z }
func (c XYZW[T]) XW() (T, T) { return c.X, c.W }
func (c XYZW[T]) YX() (T, T) { return c.Y, c.X }
func (c XYZW[T]) YY() (T, T) { return c.Y, c.Y }
func (c XYZW[T]) YZ() (T, T) { return c.Y, c.Z }
func (c XYZW[T]) YW() (T, T) { return c.Y, c.W }
func (c XYZW[T]) ZX() (T, T) { return c.Z, c.X }
func (c XYZW[T]) ZY() (T, T) { return c.Z, c.Y }
func (c XYZW[T]) ZZ() (T, T) { return c.Z, c.Z }
func (c XYZW[T]) ZW() (T, T) { return c.Z, c.W }
func (c XYZW[T]) WX() (T, T) { return c.W, c.X }
func (c XYZW[T]) WY() (T, T) { return c.W, c.Y }
func (c XYZW[T]) WZ() (T, T) { return c.W, c.Z }
func (c XYZW[T]) WW() (T, T) { return c.W, c.W }

func (c XYZW[T]) XXX() (T, T, T) { return c.X, c.X, c.X }
func (c XYZW[T]) XXY() (T, T, T) { return c.X, c.X, c.Y }
func (c XYZW[T]) XXZ() (T, T, T) { return c.X, c.X, c.Z }
func (c XYZW[T]) XXW() (T, T, T) { return c.X, c.X, c.W }
func (c XYZW[T]) XYX() (T, T, T) { return c.X, c.Y, c.X }
func (c XYZW[T]) XYY() (T, T, T) { return c.X, c.Y, c.Y }
func (c XYZW[T]) XYZ() (T, T, T) { return c.X, c.Y, c.Z }
func (c XYZW[T]) XYW() (T, T, T) { return c.X, c.Y, c.W }
func (c XYZW[T]) XZX() (T, T, T) { return c.X, c.Z, c.X }
func (c XYZW[T]) XZY() (T, T, T) { return c.X, c.Z, c.Y }
func (c XYZW[T]) XZZ() (T, T, T) { return c.X, c.Z, c.Z }
func (c XYZW[T]) XZW() (T, T, T) { return c.X, c.Z, c.W }
func (c XYZW[T]) XWX() (T, T, T) { return c.X, c.W, c.X }
func (c XYZW[T]) XWY() (T, T, T) { return c.X, c.W, c.Y }
func (c XYZW[T]) XWZ() (T, T, T) { return c.X, c.W, c.Z }
func (c XYZW[T]) XWW() (T, T, T) { return c.X, c.W, c.W }
func (c XYZW[T]) YXX() (T, T, T) { return c.Y, c.X, c.X }
func (c XYZW[T]) YXY() (T, T, T) { return c.Y, c.X, c.Y }
func (c XYZW[T]) YXZ() (T, T, T) { return c.Y, c.X, c.Z }
func (c XYZW[T]) YXW() (T, T, T) { return c.Y, c.X, c.W }
func (c XYZW[T]) YYX() (T, T, T) { return c.Y, c.Y, c.X }
func (c XYZW[T]) YYY() (T, T, T) { return c.Y, c.Y, c.Y }
func (c XYZW[T]) YYZ() (T, T, T) { return c.Y, c.Y, c.Z }
func (c XYZW[T]) YYW() (T, T, T) { return c.Y, c.Y, c.W }
func (c XYZW[T]) YZX() (T, T, T) { return c.Y, c.Z, c.X }
func (c XYZW[T]) YZY() (T, T, T) { return c.Y, c.Z, c.Y }
func (c XYZW[T]) YZZ() (T, T, T) { return c.Y, c.Z, c.Z }
func (c XYZW[T]) YZW() (T, T, T) { return c.Y, c.Z, c.W }
func (c XYZW[T]) YWX() (T, T, T) { return c.Y, c.W, c.X }
func (c XYZW[T]) YWY() (T, T, T) { return c.Y, c.W, c.Y }
func (c XYZW[T]) YWZ() (T, T, T) { return c.Y, c.W, c.Z }
func (c XYZW[T]) YWW() (T, T, T) { return c.Y, c.W, c.W }
func (c XYZW[T]) ZXX() (T, T, T) { return c.Z, c.X, c.X }
func (c XYZW[T]) ZXY() (T, T, T) { return c.Z, c.X, c.Y }
func (c XYZW[T]) ZXZ() (T, T, T) { return c.Z, c.X, c.Z }
func (c XYZW[T]) ZXW() (T, T, T) { return c.Z, c.X, c.W }
func (c XYZW[T]) ZYX() (T, T, T) { return c.Z, c.Y, c.X }
func (c XYZW[T]) ZYY() (T, T, T) { return c.Z, c.Y, c.Y }
func (c XYZW[T]) ZYZ() (T, T, T) { return c.Z, c.Y, c.Z }
func (c XYZW[T]) ZYW() (T, T, T) { return c.Z, c.Y, c.W }
func (c XYZW[T]) ZZX() (T, T, T) { return c.Z, c.Z, c.X }
func (c XYZW[T]) ZZY() (T, T, T) { return c.Z, c.Z, c.Y }
func (c XYZW[T]) ZZZ() (T, T, T) { return c.Z, c.Z, c.Z }
func (c XYZW[T]) ZZW() (T, T, T) { return c.Z, c.Z, c.W }
func (c XYZW[T]) ZWX() (T, T, T) { return c.Z, c.W, c.X }
func (c XYZW[T]) ZWY() (T, T, T) { return c.Z, c.W, c.Y }
func (c XYZW[T]) ZWZ() (T, T, T) { return c.Z, c.W, c.Z }
func (c XYZW[T]) ZWW() (T, T, T) { return c.Z, c.W, c.W }
func (c XYZW[T]) WXX() (T, T, T) { return c.W, c.X, c.X }
func (c XYZW[T]) WXY() (T, T, T) { return c.W, c.X, c.Y }
func (c XYZW[T]) WXZ() (T, T, T) { return c.W, c.X, c.Z }
func (c XYZW[T]) WXW() (T, T, T) { return c.W, c.X, c.W }
func (c XYZW[T]) WYX() (T, T, T) { return c.W, c.Y, c.X }
func (c XYZW[T]) WYY() (T, T, T) { return c.W, c.Y, c.Y }
func (c XYZW[T]) WYZ() (T, T, T) { return c.W, c.Y, c.Z }
func (c XYZW[T]) WYW() (T, T, T) { return c.W, c.Y, c.W }
func (c XYZW[T]) WZX() (T, T, T) { return c.W, c.Z, c.X }
func (c XYZW[T]) WZY() (T, T, T) { return c.W, c.Z, c.Y }
func (c XYZW[T]) WZZ() (T, T, T) { return c.W, c.Z, c.Z }
func (c XYZW[T]) WZW() (T, T, T) { return c.W, c.Z, c.W }
func (c XYZW[T]) WWX() (T, T, T) { return c.W, c.W, c.X }
func (c XYZW[T]) WWY() (T, T, T) { return c.W, c.W, c.Y }
func (c XYZW[T]) WWZ() (T, T, T) { return c.W, c.W, c.Z }
func (c XYZW[T]) WWW() (T, T, T) { return c.W, c.W, c.W }

func (c XYZW[T]) XXXX() (T, T, T, T) { return c.X, c.X, c.X, c.X }
func (c XYZW[T]) XXXY() (T, T, T, T) { return c.X, c.X, c.X, c.Y }
func (c XYZW[T]) XXXZ() (T, T, T, T) { return c.X, c.X, c.X, c.Z }
func (c XYZW[T]) XXXW() (T, T, T, T) { return c.X, c.X, c.X, c.W }
func (c XYZW[T]) XXYX() (T, T, T, T) { return c.X, c.X, c.Y, c.X }
func (c XYZW[T]) XXYY() (T, T, T, T) { return c.X, c.X, c.Y, c.Y }
func (c XYZW[T]) XXYZ() (T, T, T, T) { return c.X, c.X, c.Y, c.Z }
func (c XYZW[T]) XXYW() (T, T, T, T) { return c.X, c.X, c.Y, c.W }
func (c XYZW[T]) XXZX() (T, T, T, T) { return c.X, c.X, c.Z, c.X }
func (c XYZW[T]) XXZY() (T, T, T, T) { return c.X, c.X, c.Z, c.Y }
func (c XYZW[T]) XXZZ() (T, T, T, T) { return c.X, c.X, c.Z, c.Z }
func (c XYZW[T]) XXZW() (T, T, T, T) { return c.X, c.X, c.Z, c.W }
func (c XYZW[T]) XXWX() (T, T, T, T) { return c.X, c.X, c.W, c.X }
func (c XYZW[T]) XXWY() (T, T, T, T) { return c.X, c.X, c.W, c.Y }
func (c XYZW[T]) XXWZ() (T, T, T, T) { return c.X, c.X, c.W, c.Z }
func (c XYZW[T]) XXWW() (T, T, T, T) { return c.X, c.X, c.W, c.W }
func (c XYZW[T]) XYXX() (T, T, T, T) { return c.X, c.Y, c.X, c.X }
func (c XYZW[T]) XYXY() (T, T, T, T) { return c.X, c.Y, c.X, c.Y }
func (c XYZW[T]) XYXZ() (T, T, T, T) { return c.X, c.Y, c.X, c.Z }
func (c XYZW[T]) XYXW() (T, T, T, T) { return c.X, c.Y, c.X, c.W }
func (c XYZW[T]) XYYX() (T, T, T, T) { return c.X, c.Y, c.Y, c.X }
func (c XYZW[T]) XYYY() (T, T, T, T) { return c.X, c.Y, c.Y, c.Y }
func (c XYZW[T]) XYYZ() (T, T, T, T) { return c.X, c.Y, c.Y, c.Z }
func (c XYZW[T]) XYYW() (T, T, T, T) { return c.X, c.Y, c.Y, c.W }
func (c XYZW[T]) XYZX() (T, T, T, T) { return c.X, c.Y, c.Z, c.X }
func (c XYZW[T]) XYZY() (T, T, T, T) { return c.X, c.Y, c.Z, c.Y }
func (c XYZW[T]) XYZZ() (T, T, T, T) { return c.X, c.Y, c.Z, c.Z }
func (c XYZW[T]) XYZW() (T, T, T, T) { return c.X, c.Y, c.Z, c.W }
func (c XYZW[T]) XYWX() (T, T, T, T) { return c.X, c.Y, c.W, c.X }
func (c XYZW[T]) XYWY() (T, T, T, T) { return c.X, c.Y, c.W, c.Y }
func (c XYZW[T]) XYWZ() (T, T, T, T) { return c.X, c.Y, c.W, c.Z }
func (c XYZW[T]) XYWW() (T, T, T, T) { return c.X, c.Y, c.W, c.W }
func (c XYZW[T]) XZXX() (T, T, T, T) { return c.X, c.Z, c.X, c.X }
func (c XYZW[T]) XZXY() (T, T, T, T) { return c.X, c.Z, c.X, c.Y }
func (c XYZW[T]) XZXZ() (T, T, T, T) { return c.X, c.Z, c.X, c.Z }
func (c XYZW[T]) XZXW() (T, T, T, T) { return c.X, c.Z, c.X, c.W }
func (c XYZW[T]) XZYX() (T, T, T, T) { return c.X, c.Z, c.Y, c.X }
func (c XYZW[T]) XZYY() (T, T, T, T) { return c.X, c.Z, c.Y, c.Y }
func (c XYZW[T]) XZYZ() (T, T, T, T) { return c.X, c.Z, c.Y, c.Z }
func (c XYZW[T]) XZYW() (T, T, T, T) { return c.X, c.Z, c.Y, c.W }
func (c XYZW[T]) XZZX() (T, T, T, T) { return c.X, c.Z, c.Z, c.X }
func (c XYZW[T]) XZZY() (T, T, T, T) { return c.X, c.Z, c.Z, c.Y }
func (c XYZW[T]) XZZZ() (T, T, T, T) { return c.X, c.Z, c.Z, c.Z }
func (c XYZW[T]) XZZW() (T, T, T, T) { return c.X, c.Z, c.Z, c.W }
func (c XYZW[T]) XZWX() (T, T, T, T) { return c.X, c.Z, c.W, c.X }
func (c XYZW[T]) XZWY() (T, T, T, T) { return c.X, c.Z, c.W, c.Y }
func (c XYZW[T]) XZWZ() (T, T, T, T) { return c.X, c.Z, c.W, c.Z }
func (c XYZW[T]) XZWW() (T, T, T, T) { return c.X, c.Z, c.W, c.W }
func (c XYZW[T]) XWXX() (T, T, T, T) { return c.X, c.W, c.X, c.X }
func (c XYZW[T]) XWXY() (T, T, T, T) { return c.X, c.W, c.X, c.Y }
func (c XYZW[T]) XWXZ() (T, T, T, T) { return c.X, c.W, c.X, c.Z }
func (c XYZW[T]) XWXW() (T, T, T, T) { return c.X, c.W, c.X, c.W }
func (c XYZW[T]) XWYX() (T, T, T, T) { return c.X, c.W, c.Y, c.X }
func (c XYZW[T]) XWYY() (T, T, T, T) { return c.X, c.W, c.Y, c.Y }
func (c XYZW[T]) XWYZ() (T, T, T, T) { return c.X, c.W, c.Y, c.Z }
func (c XYZW[T]) XWYW() (T, T, T, T) { return c.X, c.W, c.Y, c.W }
func (c XYZW[T]) XWZX() (T, T, T, T) { return c.X, c.W, c.Z, c.X }
func (c XYZW[T]) XWZY() (T, T, T, T) { return c.X, c.W, c.Z, c.Y }
func (c XYZW[T]) XWZZ() (T, T, T, T) { return c.X, c.W, c.Z, c.Z }
func (c XYZW[T]) XWZW() (T, T, T, T) { return c.X, c.W, c.Z, c.W }
func (c XYZW[T]) XWWX() (T, T, T, T) { return c.X, c.W, c.W, c.X }
func (c XYZW[T]) XWWY() (T, T, T, T) { return c.X, c.W, c.W, c.Y }
func (c XYZW[T]) XWWZ() (T, T, T, T) { return c.X, c.W, c.W, c.Z }
func (c XYZW[T]) XWWW() (T, T, T, T) { return c.X, c.W, c.W, c.W }
func (c XYZW[T]) YXXX() (T, T, T, T) { return c.Y, c.X, c.X, c.X }
func (c XYZW[T]) YXXY() (T, T, T, T) { return c.Y, c.X, c.X, c.Y }
func (c XYZW[T]) YXXZ() (T, T, T, T) { return c.Y, c.X, c.X, c.Z }
func (c XYZW[T]) YXXW() (T, T, T, T) { return c.Y, c.X, c.X, c.W }
func (c XYZW[T]) YXYX() (T, T, T, T) { return c.Y, c.X, c.Y, c.X }
func (c XYZW[T]) YXYY() (T, T, T, T) { return c.Y, c.X, c.Y, c.Y }
func (c XYZW[T]) YXYZ() (T, T, T, T) { return c.Y, c.X, c.Y, c.Z }
func (c XYZW[T]) YXYW() (T, T, T, T) { return c.Y, c.X, c.Y, c.W }
func (c XYZW[T]) YXZX() (T, T, T, T) { return c.Y, c.X, c.Z, c.X }
func (c XYZW[T]) YXZY() (T, T, T, T) { return c.Y, c.X, c.Z, c.Y }
func (c XYZW[T]) YXZZ() (T, T, T, T) { return c.Y, c.X, c.Z, c.Z }
func (c XYZW[T]) YXZW() (T, T, T, T) { return c.Y, c.X, c.Z, c.W }
func (c XYZW[T]) YXWX() (T, T, T, T) { return c.Y, c.X, c.W, c.X }
func (c XYZW[T]) YXWY() (T, T, T, T) { return c.Y, c.X, c.W, c.Y }
func (c XYZW[T]) YXWZ() (T, T, T, T) { return c.Y, c.X, c.W, c.Z }
func (c XYZW[T]) YXWW() (T, T, T, T) { return c.Y, c.X, c.W, c.W }
func (c XYZW[T]) YYXX() (T, T, T, T) { return c.Y, c.Y, c.X, c.X }
func (c XYZW[T]) YYXY() (T, T, T, T) { return c.Y, c.Y, c.X, c.Y }
func (c XYZW[T]) YYXZ() (T, T, T, T) { return c.Y, c.Y, c.X, c.Z }
func (c XYZW[T]) YYXW() (T, T, T, T) { return c.Y, c.Y, c.X, c.W }
func (c XYZW[T]) YYYX() (T, T, T, T) { return c.Y, c.Y, c.Y, c.X }
func (c XYZW[T]) YYYY() (T, T, T, T) { return c.Y, c.Y, c.Y, c.Y }
func (c XYZW[T]) YYYZ() (T, T, T, T) { return c.Y, c.Y, c.Y, c.Z }
func (c XYZW[T]) YYYW() (T, T, T, T) { return c.Y, c.Y, c.Y, c.W }
func (c XYZW[T]) YYZX() (T, T, T, T) { return c.Y, c.Y, c.Z, c.X }
func (c XYZW[T]) YYZY() (T, T, T, T) { return c.Y, c.Y, c.Z, c.Y }
func (c XYZW[T]) YYZZ() (T, T, T, T) { return c.Y, c.Y, c.Z, c.Z }
func (c XYZW[T]) YYZW() (T, T, T, T) { return c.Y, c.Y, c.Z, c.W }
func (c XYZW[T]) YYWX() (T, T, T, T) { return c.Y, c.Y, c.W, c.X }
func (c XYZW[T]) YYWY() (T, T, T, T) { return c.Y, c.Y, c.W, c.Y }
func (c XYZW[T]) YYWZ() (T, T, T, T) { return c.Y, c.Y, c.W, c.Z }
func (c XYZW[T]) YYWW() (T, T, T, T) { return c.Y, c.Y, c.W, c.W }
func (c XYZW[T]) YZXX() (T, T, T, T) { return c.Y, c.Z, c.X, c.X }
func (c XYZW[T]) YZXY() (T, T, T, T) { return c.Y, c.Z, c.X, c.Y }
func (c XYZW[T]) YZXZ() (T, T, T, T) { return c.Y, c.Z, c.X, c.Z }
func (c XYZW[T]) YZXW() (T, T, T, T) { return c.Y, c.Z, c.X, c.W }
func (c XYZW[T]) YZYX() (T, T, T, T) { return c.Y, c.Z, c.Y, c.X }
func (c XYZW[T]) YZYY() (T, T, T, T) { return c.Y, c.Z, c.Y, c.Y }
func (c XYZW[T]) YZYZ() (T, T, T, T) { return c.Y, c.Z, c.Y, c.Z }
func (c XYZW[T]) YZYW() (T, T, T, T) { return c.Y, c.Z, c.Y, c.W }
func (c XYZW[T]) YZZX() (T, T, T, T) { return c.Y, c.Z, c.Z, c.X }
func (c XYZW[T]) YZZY() (T, T, T, T) { return c.Y, c.Z, c.Z, c.Y }
func (c XYZW[T]) YZZZ() (T, T, T, T) { return c.Y, c.Z, c.Z, c.Z }
func (c XYZW[T]) YZZW() (T, T, T, T) { return c.Y, c.Z, c.Z, c.W }
func (c XYZW[T]) YZWX() (T, T, T, T) { return c.Y, c.Z, c.W, c.X }
func (c XYZW[T]) YZWY() (T, T, T, T) { return c.Y, c.Z, c.W, c.Y }
func (c XYZW[T]) YZWZ() (T, T, T, T) { return c.Y, c.Z, c.W, c.Z }
func (c XYZW[T]) YZWW() (T, T, T, T) { return c.Y, c.Z, c.W, c.W }
func (c XYZW[T]) YWXX() (T, T, T, T) { return c.Y, c.W, c.X, c.X }
func (c XYZW[T]) YWXY() (T, T, T, T) { return c.Y, c.W, c.X, c.Y }
func (c XYZW[T]) YWXZ() (T, T, T, T) { return c.Y, c.W, c.X, c.Z }
func (c XYZW[T]) YWXW() (T, T, T, T) { return c.Y, c.W, c.X, c.W }
func (c XYZW[T]) YWYX() (T, T, T, T) { return c.Y, c.W, c.Y, c.X }
func (c XYZW[T]) YWYY() (T, T, T, T) { return c.Y, c.W, c.Y, c.Y }
func (c XYZW[T]) YWYZ() (T, T, T, T) { return c.Y, c.W, c.Y, c.Z }
func (c XYZW[T]) YWYW() (T, T, T, T) { return c.Y, c.W, c.Y, c.W }
func (c XYZW[T]) YWZX() (T, T, T, T) { return c.Y, c.W, c.Z, c.X }
func (c XYZW[T]) YWZY() (T, T, T, T) { return c.Y, c.W, c.Z, c.Y }
func (c XYZW[T]) YWZZ() (T, T, T, T) { return c.Y, c.W, c.Z, c.Z }
func (c XYZW[T]) YWZW() (T, T, T, T) { return c.Y, c.W, c.Z, c.W }
func (c XYZW[T]) YWWX() (T, T, T, T) { return c.Y, c.W, c.W, c.X }
func (c XYZW[T]) YWWY() (T, T, T, T) { return c.Y, c.W, c.W, c.Y }
func (c XYZW[T]) YWWZ() (T, T, T, T) { return c.Y, c.W, c.W, c.Z }
func (c XYZW[T]) YWWW() (T, T, T, T) { return c.Y, c.W, c.W, c.W }
func (c XYZW[T]) ZXXX() (T, T, T, T) { return c.Z, c.X, c.X, c.X }
func (c XYZW[T]) ZXXY() (T, T, T, T) { return c.Z, c.X, c.X, c.Y }
func (c XYZW[T]) ZXXZ() (T, T, T, T) { return c.Z, c.X, c.X, c.Z }
func (c XYZW[T]) ZXXW() (T, T, T, T) { return c.Z, c.X, c.X, c.W }
func (c XYZW[T]) ZXYX() (T, T, T, T) { return c.Z, c.X, c.Y, c.X }
func (c XYZW[T]) ZXYY() (T, T, T, T) { return c.Z, c.X, c.Y, c.Y }
func (c XYZW[T]) ZXYZ() (T, T, T, T) { return c.Z, c.X, c.Y, c.Z }
func (c XYZW[T]) ZXYW() (T, T, T, T) { return c.Z, c.X, c.Y, c.W }
func (c XYZW[T]) ZXZX() (T, T, T, T) { return c.Z, c.X, c.Z, c.X }
func (c XYZW[T]) ZXZY() (T, T, T, T) { return c.Z, c.X, c.Z, c.Y }
func (c XYZW[T]) ZXZZ() (T, T, T, T) { return c.Z, c.X, c.Z, c.Z }
func (c XYZW[T]) ZXZW() (T, T, T, T) { return c.Z, c.X, c.Z, c.W }
func (c XYZW[T]) ZXWX() (T, T, T, T) { return c.Z, c.X, c.W, c.X }
func (c XYZW[T]) ZXWY() (T, T, T, T) { return c.Z, c.X, c.W, c.Y }
func (c XYZW[T]) ZXWZ() (T, T, T, T) { return c.Z, c.X, c.W, c.Z }
func (c XYZW[T]) ZXWW() (T, T, T, T) { return c.Z, c.X, c.W, c.W }
func (c XYZW[T]) ZYXX() (T, T, T, T) { return c.Z, c.Y, c.X, c.X }
func (c XYZW[T]) ZYXY() (T, T, T, T) { return c.Z, c.Y, c.X, c.Y }
func (c XYZW[T]) ZYXZ() (T, T, T, T) { return c.Z, c.Y, c.X, c.Z }
func (c XYZW[T]) ZYXW() (T, T, T, T) { return c.Z, c.Y, c.X, c.W }
func (c XYZW[T]) ZYYX() (T, T, T, T) { return c.Z, c.Y, c.Y, c.X }
func (c XYZW[T]) ZYYY() (T, T, T, T) { return c.Z, c.Y, c.Y, c.Y }
func (c XYZW[T]) ZYYZ() (T, T, T, T) { return c.Z, c.Y, c.Y, c.Z }
func (c XYZW[T]) ZYYW() (T, T, T, T) { return c.Z, c.Y, c.Y, c.W }
func (c XYZW[T]) ZYZX() (T, T, T, T) { return c.Z, c.Y, c.Z, c.X }
func (c XYZW[T]) ZYZY() (T, T, T, T) { return c.Z, c.Y, c.Z, c.Y }
func (c XYZW[T]) ZYZZ() (T, T, T, T) { return c.Z, c.Y, c.Z, c.Z }
func (c XYZW[T]) ZYZW() (T, T, T, T) { return c.Z, c.Y, c.Z, c.W }
func (c XYZW[T]) ZYWX() (T, T, T, T) { return c.Z, c.Y, c.W, c.X }
func (c XYZW[T]) ZYWY() (T, T, T, T) { return c.Z, c.Y, c.W, c.Y }
func (c XYZW[T]) ZYWZ() (T, T, T, T) { return c.Z, c.Y, c.W, c.Z }
func (c XYZW[T]) ZYWW() (T, T, T, T) { return c.Z, c.Y, c.W, c.W }
func (c XYZW[T]) ZZXX() (T, T, T, T) { return c.Z, c.Z, c.X, c.X }
func (c XYZW[T]) ZZXY() (T, T, T, T) { return c.Z, c.Z, c.X, c.Y }
func (c XYZW[T]) ZZXZ() (T, T, T, T) { return c.Z, c.Z, c.X, c.Z }
func (c XYZW[T]) ZZXW() (T, T, T, T) { return c.Z, c.Z, c.X, c.W }
func (c XYZW[T]) ZZYX() (T, T, T, T) { return c.Z, c.Z, c.Y, c.X }
func (c XYZW[T]) ZZYY() (T, T, T, T) { return c.Z, c.Z, c.Y, c.Y }
func (c XYZW[T]) ZZYZ() (T, T, T, T) { return c.Z, c.Z, c.Y, c.Z }
func (c XYZW[T]) ZZYW() (T, T, T, T) { return c.Z, c.Z, c.Y, c.W }
func (c XYZW[T]) ZZZX() (T, T, T, T) { return c.Z, c.Z, c.Z, c.X }
func (c XYZW[T]) ZZZY() (T, T, T, T) { return c.Z, c.Z, c.Z, c.Y }
func (c XYZW[T]) ZZZZ() (T, T, T, T) { return c.Z, c.Z, c.Z, c.Z }
func (c XYZW[T]) ZZZW() (T, T, T, T) { return c.Z, c.Z, c.Z, c.W }
func (c XYZW[T]) ZZWX() (T, T, T, T) { return c.Z, c.Z, c.W, c.X }
func (c XYZW[T]) ZZWY() (T, T, T, T) { return c.Z, c.Z, c.W, c.Y }
func (c XYZW[T]) ZZWZ() (T, T, T, T) { return c.Z, c.Z, c.W, c.Z }
func (c XYZW[T]) ZZWW() (T, T, T, T) { return c.Z, c.Z, c.W, c.W }
func (c XYZW[T]) ZWXX() (T, T, T, T) { return c.Z, c.W, c.X, c.X }
func (c XYZW[T]) ZWXY() (T, T, T, T) { return c.Z, c.W, c.X, c.Y }
func (c XYZW[T]) ZWXZ() (T, T, T, T) { return c.Z, c.W, c.X, c.Z }
func (c XYZW[T]) ZWXW() (T, T, T, T) { return c.Z, c.W, c.X, c.W }
func (c XYZW[T]) ZWYX() (T, T, T, T) { return c.Z, c.W, c.Y, c.X }
func (c XYZW[T]) ZWYY() (T, T, T, T) { return c.Z, c.W, c.Y, c.Y }
func (c XYZW[T]) ZWYZ() (T, T, T, T) { return c.Z, c.W, c.Y, c.Z }
func (c XYZW[T]) ZWYW() (T, T, T, T) { return c.Z, c.W, c.Y, c.W }
func (c XYZW[T]) ZWZX() (T, T, T, T) { return c.Z, c.W, c.Z, c.X }
func (c XYZW[T]) ZWZY() (T, T, T, T) { return c.Z, c.W, c.Z, c.Y }
func (c XYZW[T]) ZWZZ() (T, T, T, T) { return c.Z, c.W, c.Z, c.Z }
func (c XYZW[T]) ZWZW() (T, T, T, T) { return c.Z, c.W, c.Z, c.W }
func (c XYZW[T]) ZWWX() (T, T, T, T) { return c.Z, c.W, c.W, c.X }
func (c XYZW[T]) ZWWY() (T, T, T, T) { return c.Z, c.W, c.W, c.Y }
func (c XYZW[T]) ZWWZ() (T, T, T, T) { return c.Z, c.W, c.W, c.Z }
func (c XYZW[T]) ZWWW() (T, T, T, T) { return c.Z, c.W, c.W, c.W }
func (c XYZW[T]) WXXX() (T, T, T, T) { return c.W, c.X, c.X, c.X }
func (c XYZW[T]) WXXY() (T, T, T, T) { return c.W, c.X, c.X, c.Y }
func (c XYZW[T]) WXXZ() (T, T, T, T) { return c.W, c.X, c.X, c.Z }
func (c XYZW[T]) WXXW() (T, T, T, T) { return c.W, c.X, c.X, c.W }
func (c XYZW[T]) WXYX() (T, T, T, T) { return c.W, c.X, c.Y, c.X }
func (c XYZW[T]) WXYY() (T, T, T, T) { return c.W, c.X, c.Y, c.Y }
func (c XYZW[T]) WXYZ() (T, T, T, T) { return c.W, c.X, c.Y, c.Z }
func (c XYZW[T]) WXYW() (T, T, T, T) { return c.W, c.X, c.Y, c.W }
func (c XYZW[T]) WXZX() (T, T, T, T) { return c.W, c.X, c.Z, c.X }
func (c XYZW[T]) WXZY() (T, T, T, T) { return c.W, c.X, c.Z, c.Y }
func (c XYZW[T]) WXZZ() (T, T, T, T) { return c.W, c.X, c.Z, c.Z }
func (c XYZW[T]) WXZW() (T, T, T, T) { return c.W, c.X, c.Z, c.W }
func (c XYZW[T]) WXWX() (T, T, T, T) { return c.W, c.X, c.W, c.X }
func (c XYZW[T]) WXWY() (T, T, T, T) { return c.W, c.X, c.W, c.Y }
func (c XYZW[T]) WXWZ() (T, T, T, T) { return c.W, c.X, c.W, c.Z }
func (c XYZW[T]) WXWW() (T, T, T, T) { return c.W, c.X, c.W, c.W }
func (c XYZW[T]) WYXX() (T, T, T, T) { return c.W, c.Y, c.X, c.X }
func (c XYZW[T]) WYXY() (T, T, T, T) { return c.W, c.Y, c.X, c.Y }
func (c XYZW[T]) WYXZ() (T, T, T, T) { return c.W, c.Y, c.X, c.Z }
func (c XYZW[T]) WYXW() (T, T, T, T) { return c.W, c.Y, c.X, c.W }
func (c XYZW[T]) WYYX() (T, T, T, T) { return c.W, c.Y, c.Y, c.X }
func (c XYZW[T]) WYYY() (T, T, T, T) { return c.W, c.Y, c.Y, c.Y }
func (c XYZW[T]) WYYZ() (T, T, T, T) { return c.W, c.Y, c.Y, c.Z }
func (c XYZW[T]) WYYW() (T, T, T, T) { return c.W, c.Y, c.Y, c.W }
func (c XYZW[T]) WYZX() (T, T, T, T) { return c.W, c.Y, c.Z, c.X }
func (c XYZW[T]) WYZY() (T, T, T, T) { return c.W, c.Y, c.Z, c.Y }
func (c XYZW[T]) WYZZ() (T, T, T, T) { return c.W, c.Y, c.Z, c.Z }
func (c XYZW[T]) WYZW() (T, T, T, T) { return c.W, c.Y, c.Z, c.W }
func (c XYZW[T]) WYWX() (T, T, T, T) { return c.W, c.Y, c.W, c.X }
func (c XYZW[T]) WYWY() (T, T, T, T) { return c.W, c.Y, c.W, c.Y }
func (c XYZW[T]) WYWZ() (T, T, T, T) { return c.W, c.Y, c.W, c.Z }
func (c XYZW[T]) WYWW() (T, T, T, T) { return c.W, c.Y, c.W, c.W }
func (c XYZW[T]) WZXX() (T, T, T, T) { return c.W, c.Z, c.X, c.X }
func (c XYZW[T]) WZXY() (T, T, T, T) { return c.W, c.Z, c.X, c.Y }
func (c XYZW[T]) WZXZ() (T, T, T, T) { return c.W, c.Z, c.X, c.Z }
func (c XYZW[T]) WZXW() (T, T, T, T) { return c.W, c.Z, c.X, c.W }
func (c XYZW[T]) WZYX() (T, T, T, T) { return c.W, c.Z, c.Y, c.X }
func (c XYZW[T]) WZYY() (T, T, T, T) { return c.W, c.Z, c.Y, c.Y }
func (c XYZW[T]) WZYZ() (T, T, T, T) { return c.W, c.Z, c.Y, c.Z }
func (c XYZW[T]) WZYW() (T, T, T, T) { return c.W, c.Z, c.Y, c.W }
func (c XYZW[T]) WZZX() (T, T, T, T) { return c.W, c.Z, c.Z, c.X }
func (c XYZW[T]) WZZY() (T, T, T, T) { return c.W, c.Z, c.Z, c.Y }
func (c XYZW[T]) WZZZ() (T, T, T, T) { return c.W, c.Z, c.Z, c.Z }
func (c XYZW[T]) WZZW() (T, T, T, T) { return c.W, c.Z, c.Z, c.W }
func (c XYZW[T]) WZWX() (T, T, T, T) { return c.W, c.Z, c.W, c.X }
func (c XYZW[T]) WZWY() (T, T, T, T) { return c.W, c.Z, c.W, c.Y }
func (c XYZW[T]) WZWZ() (T, T, T, T) { return c.W, c.Z, c.W, c.Z }
func (c XYZW[T]) WZWW() (T, T, T, T) { return c.W, c.Z, c.W, c.W }
func (c XYZW[T]) WWXX() (T, T, T, T) { return c.W, c.W, c.X, c.X }
func (c XYZW[T]) WWXY() (T, T, T, T) { return c.W, c.W, c.X, c.Y }
func (c XYZW[T]) WWXZ() (T, T, T, T) { return c.W, c.W, c.X, c.Z }
func (c XYZW[T]) WWXW() (T, T, T, T) { return c.W, c.W, c.X, c.W }
func (c XYZW[T]) WWYX() (T, T, T, T) { return c.W, c.W, c.Y, c.X }
func (c XYZW[T]) WWYY() (T, T, T, T) { return c.W, c.W, c.Y, c.Y }
func (c XYZW[T]) WWYZ() (T, T, T, T) { return c.W, c.W, c.Y, c.Z }
func (c XYZW[T]) WWYW() (T, T, T, T) { return c.W, c.W, c.Y, c.W }
func (c XYZW[T]) WWZX() (T, T, T, T) { return c.W, c.W, c.Z, c.X }
func (c XYZW[T]) WWZY() (T, T, T, T) { return c.W, c.W, c.Z, c.Y }
func (c XYZW[T]) WWZZ() (T, T, T, T) { return c.W, c.W, c.Z, c.Z }
func (c XYZW[T]) WWZW() (T, T, T, T) { return c.W, c.W, c.Z, c.W }
func (c XYZW[T]) WWWX() (T, T, T, T) { return c.W, c.W, c.W, c.X }
func (c XYZW[T]) WWWY() (T, T, T, T) { return c.W, c.W, c.W, c.Y }
func (c XYZW[T]) WWWZ() (T, T, T, T) { return c.W, c.W, c.W, c.Z }
func (c XYZW[T]) WWWW() (T, T, T, T) { return c.W, c.W, c.W, c.W }
