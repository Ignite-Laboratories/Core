package std

import (
	"fmt"
	"golang.org/x/exp/rand"
)

// RGBA is a structure for holding color values.
type RGBA struct {
	// R is the red channel.
	R float32

	// G is the green channel.
	G float32

	// B is the blue channel.
	B float32

	// A is the alpha channel.
	A float32
}

// RandomRGB generates a random set of RGB values with an alpha of 1.0.
func RandomRGB() RGBA {
	return RGBA{
		R: rand.Float32(),
		G: rand.Float32(),
		B: rand.Float32(),
		A: 1.0,
	}
}

// RandomRGBA generates a random set of RGBA values.
func RandomRGBA() RGBA {
	return RGBA{
		R: rand.Float32(),
		G: rand.Float32(),
		B: rand.Float32(),
		A: rand.Float32(),
	}
}

// SplitRGBA quickly provides the R G B A values as individual variables.
func (c RGBA) SplitRGBA() (r, g, b, a float32) {
	return c.R, c.G, c.B, c.A
}

// SplitRGB quickly provides the R G B values as individual variables.
func (c RGBA) SplitRGB() (r, g, b float32) {
	return c.R, c.G, c.B
}

// Comparator returns if the two RGBA values are equal in values.
func Comparator(a RGBA, b RGBA) bool {
	return a.R == b.R && a.G == b.G && a.B == b.B && a.A == b.A
}

func (c RGBA) String() string {
	return fmt.Sprintf("(%f, %f, %f, %f)", c.R, c.G, c.B, c.A)
}
