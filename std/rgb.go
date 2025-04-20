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

func (c RGB[T]) String() string {
	return fmt.Sprintf("(r:%d, g:%d, b:%d)", c.R, c.G, c.B)
}

// RGBFromHex converts the provided RGB hex values into a std.RGB[byte].
func RGBFromHex(value uint32) RGB[byte] {
	return RGB[byte]{
		R: byte((value >> 16) & 0xFF),
		G: byte((value >> 8) & 0xFF),
		B: byte(value & 0xFF),
	}
}

// RandomRGB generates a random set of RGB values of the provided type.
func RandomRGB[T core.Numeric]() RGBA[T] {
	return RGB[T]{
		R: core.RandomNumber[T](),
		G: core.RandomNumber[T](),
		B: core.RandomNumber[T](),
	}
}

// Normalize32 returns an RGB[float32] ranging from 0.0-1.0.
func (c RGB[byte]) Normalize32() RGB[float32] {
	return RGB[float32]{
		R: float32(c.R) / 255.0,
		G: float32(c.G) / 255.0,
		B: float32(c.B) / 255.0,
	}
}

// Normalize64 returns an RGB[float64] ranging from 0.0-1.0.
func (c RGB[byte]) Normalize64() RGB[float64] {
	return RGB[float64]{
		R: float64(c.R) / 255.0,
		G: float64(c.G) / 255.0,
		B: float64(c.B) / 255.0,
	}
}

// Denormalize64 returns an RGB[byte] ranging from 0-255.
func (c RGB[float64]) Denormalize64() RGB[byte] {
	return RGB[byte]{
		R: byte(c.R * 255.0),
		G: byte(c.G * 255.0),
		B: byte(c.B * 255.0),
	}
}

// Denormalize32 returns an RGB[byte] ranging from 0-255.
func (c RGB[float32]) Denormalize32() RGB[byte] {
	return RGB[byte]{
		R: byte(c.R * 255.0),
		G: byte(c.G * 255.0),
		B: byte(c.B * 255.0),
	}
}

// RGBComparator returns if the two RGB values are equal in values.
func RGBComparator[T core.Numeric](a RGB[T], b RGB[T]) bool {
	return a.R == b.R && a.G == b.G && a.B == b.B
}

// 2-component swizzles
func (c RGB[T]) RR() (T, T) { return c.R, c.R }
func (c RGB[T]) RG() (T, T) { return c.R, c.G }
func (c RGB[T]) RB() (T, T) { return c.R, c.B }
func (c RGB[T]) GR() (T, T) { return c.G, c.R }
func (c RGB[T]) GG() (T, T) { return c.G, c.G }
func (c RGB[T]) GB() (T, T) { return c.G, c.B }
func (c RGB[T]) BR() (T, T) { return c.B, c.R }
func (c RGB[T]) BG() (T, T) { return c.B, c.G }
func (c RGB[T]) BB() (T, T) { return c.B, c.B }

// 3-component swizzles
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
