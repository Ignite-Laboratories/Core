package std

import (
	"fmt"
	"golang.org/x/exp/rand"
)

// RGBA is a general structure for holding color values.
type RGBA struct {
	R float32
	G float32
	B float32
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

func Comparator(a RGBA, b RGBA) bool {
	return a.R == b.R && a.G == b.G && a.B == b.B && a.A == b.A
}

func (c RGBA) String() string {
	return fmt.Sprintf("(%f, %f, %f, %f)", c.R, c.G, c.B, c.A)
}
