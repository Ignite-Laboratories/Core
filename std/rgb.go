package std

import (
	"fmt"
	"github.com/ignite-laboratories/core"
)

// RGB is a structure for holding red, green, and blue color values.
type RGB[T core.Numeric] struct {
	// R is the red channel.
	R T

	// G is the green channel.
	G T

	// B is the blue channel.
	B T
}

// RGBFromHex converts the provided RGB hex values into a std.RGB[byte].
func RGBFromHex(value uint32) RGB[byte] {
	return RGB[byte]{
		R: byte((value >> 16) & 0xFF),
		G: byte((value >> 8) & 0xFF),
		B: byte(value & 0xFF),
	}
}

// RandomRGB returns a pseudo-random RGB[T] of the provided type using core.RandomNumber[T].
//
// If requesting a floating point type, the resulting number will be bounded
// in the fully closed interval [0.0, 1.0]
//
// If requesting an integer type, the resulting number will be bounded
// in the fully closed interval [0, n] - where n is the maximum value of
// the provided type.
func RandomRGB[T core.Numeric]() RGB[T] {
	return RGB[T]{
		R: core.RandomNumber[T](),
		G: core.RandomNumber[T](),
		B: core.RandomNumber[T](),
	}
}

// RandomRGBUpTo returns a pseudo-random RGB[T] of the provided type bounded in the closed interval [0, max].
func RandomRGBUpTo[T core.Numeric](rUpper T, gUpper T, bUpper T) RGB[T] {
	return RGB[T]{
		R: core.RandomNumberRange[T](core.NumericRange[T]{Stop: rUpper}),
		G: core.RandomNumberRange[T](core.NumericRange[T]{Stop: gUpper}),
		B: core.RandomNumberRange[T](core.NumericRange[T]{Stop: bUpper}),
	}
}

// RandomRGBRange returns a pseudo-random RGB[T] of the provided type bounded in the closed interval [min, max].
func RandomRGBRange[T core.Numeric](rRange core.NumericRange[T], gRange core.NumericRange[T], bRange core.NumericRange[T]) RGB[T] {
	return RGB[T]{
		R: core.RandomNumberRange[T](rRange),
		G: core.RandomNumberRange[T](gRange),
		B: core.RandomNumberRange[T](bRange),
	}
}

// NormalizeRGB32 returns an RGB[float32] ranging from 0.0-1.0.
func NormalizeRGB32[T core.Integer](source RGB[T]) RGB[float32] {
	return RGB[float32]{
		R: core.NormalizeToFloat32(source.R),
		G: core.NormalizeToFloat32(source.G),
		B: core.NormalizeToFloat32(source.B),
	}
}

// NormalizeRGB64 returns an RGB[float64] ranging from 0.0-1.0.
func NormalizeRGB64[T core.Integer](source RGB[T]) RGB[float64] {
	return RGB[float64]{
		R: core.NormalizeToFloat64(source.R),
		G: core.NormalizeToFloat64(source.G),
		B: core.NormalizeToFloat64(source.B),
	}
}

// ScaleToTypeRGB32 returns a scaled value of the provided type in the range [0, T.MaxValue].
//
// NOTE: This will panic if the provided value is greater than the maximum value of the provided type.
func ScaleToTypeRGB32[TOut core.Integer](source RGB[float32]) RGB[TOut] {
	return RGB[TOut]{
		R: core.ScaleFloat32ToType[TOut](source.R),
		G: core.ScaleFloat32ToType[TOut](source.G),
		B: core.ScaleFloat32ToType[TOut](source.B),
	}
}

// ScaleToTypeRGB64 returns a scaled value of the provided type in the range [0, T.MaxValue].
//
// NOTE: This will panic if the provided value is greater than the maximum value of the provided type.
func ScaleToTypeRGB64[TOut core.Integer](source RGB[float64]) RGB[TOut] {
	return RGB[TOut]{
		R: core.ScaleFloat64ToType[TOut](source.R),
		G: core.ScaleFloat64ToType[TOut](source.G),
		B: core.ScaleFloat64ToType[TOut](source.B),
	}
}

// RGBComparator returns if the two RGB values are equal in values.
func RGBComparator[T core.Numeric](a RGB[T], b RGB[T]) bool {
	return a.R == b.R && a.G == b.G && a.B == b.B
}

func (c RGB[T]) String() string {
	return fmt.Sprintf("(%v, %v, %v)", c.R, c.G, c.B)
}

/**
Swizzling
*/

func (c RGB[T]) RR() (T, T) { return c.R, c.R }
func (c RGB[T]) RG() (T, T) { return c.R, c.G }
func (c RGB[T]) RB() (T, T) { return c.R, c.B }
func (c RGB[T]) GR() (T, T) { return c.G, c.R }
func (c RGB[T]) GG() (T, T) { return c.G, c.G }
func (c RGB[T]) GB() (T, T) { return c.G, c.B }
func (c RGB[T]) BR() (T, T) { return c.B, c.R }
func (c RGB[T]) BG() (T, T) { return c.B, c.G }
func (c RGB[T]) BB() (T, T) { return c.B, c.B }

func (c RGB[T]) RRR() (T, T, T) { return c.R, c.R, c.R }
func (c RGB[T]) RRG() (T, T, T) { return c.R, c.R, c.G }
func (c RGB[T]) RRB() (T, T, T) { return c.R, c.R, c.B }
func (c RGB[T]) RGR() (T, T, T) { return c.R, c.G, c.R }
func (c RGB[T]) RGG() (T, T, T) { return c.R, c.G, c.G }
func (c RGB[T]) RGB() (T, T, T) { return c.R, c.G, c.B }
func (c RGB[T]) RBR() (T, T, T) { return c.R, c.B, c.R }
func (c RGB[T]) RBG() (T, T, T) { return c.R, c.B, c.G }
func (c RGB[T]) RBB() (T, T, T) { return c.R, c.B, c.B }
func (c RGB[T]) GRR() (T, T, T) { return c.G, c.R, c.R }
func (c RGB[T]) GRG() (T, T, T) { return c.G, c.R, c.G }
func (c RGB[T]) GRB() (T, T, T) { return c.G, c.R, c.B }
func (c RGB[T]) GGR() (T, T, T) { return c.G, c.G, c.R }
func (c RGB[T]) GGG() (T, T, T) { return c.G, c.G, c.G }
func (c RGB[T]) GGB() (T, T, T) { return c.G, c.G, c.B }
func (c RGB[T]) GBR() (T, T, T) { return c.G, c.B, c.R }
func (c RGB[T]) GBG() (T, T, T) { return c.G, c.B, c.G }
func (c RGB[T]) GBB() (T, T, T) { return c.G, c.B, c.B }
func (c RGB[T]) BRR() (T, T, T) { return c.B, c.R, c.R }
func (c RGB[T]) BRG() (T, T, T) { return c.B, c.R, c.G }
func (c RGB[T]) BRB() (T, T, T) { return c.B, c.R, c.B }
func (c RGB[T]) BGR() (T, T, T) { return c.B, c.G, c.R }
func (c RGB[T]) BGG() (T, T, T) { return c.B, c.G, c.G }
func (c RGB[T]) BGB() (T, T, T) { return c.B, c.G, c.B }
func (c RGB[T]) BBR() (T, T, T) { return c.B, c.B, c.R }
func (c RGB[T]) BBG() (T, T, T) { return c.B, c.B, c.G }
func (c RGB[T]) BBB() (T, T, T) { return c.B, c.B, c.B }
