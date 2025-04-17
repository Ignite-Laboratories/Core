package std

import (
	"errors"
	"fmt"
	"golang.org/x/exp/rand"
	"strconv"
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

// RGBFromHex converts the provided RGB hex values into std.RGBA float32 form.
//
// Each channel must be provided a two-digit string input.
func RGBFromHex(r string, g string, b string) (RGBA, error) {
	rf, rErr := hexToFloat(r)
	if rErr != nil {
		return RGBA{}, rErr
	}
	gf, gErr := hexToFloat(g)
	if gErr != nil {
		return RGBA{}, gErr
	}
	bf, bErr := hexToFloat(b)
	if bErr != nil {
		return RGBA{}, bErr
	}
	return RGBA{
		R: rf,
		G: gf,
		B: bf,
	}, nil
}

// RGBAFromHex converts the provided RGBA hex values into std.RGBA float32 form.
//
// Each channel must be provided a two-digit string input.
func RGBAFromHex(r string, g string, b string, a string) (RGBA, error) {
	rf, rErr := hexToFloat(r)
	if rErr != nil {
		return RGBA{}, rErr
	}
	gf, gErr := hexToFloat(g)
	if gErr != nil {
		return RGBA{}, gErr
	}
	bf, bErr := hexToFloat(b)
	if bErr != nil {
		return RGBA{}, bErr
	}
	af, aErr := hexToFloat(a)
	if aErr != nil {
		return RGBA{}, aErr
	}
	return RGBA{
		R: rf,
		G: gf,
		B: bf,
		A: af,
	}, nil
}

func hexToFloat(hex string) (float32, error) {
	if len(hex) != 2 {
		return 0, errors.New("input must be a 2-character hexadecimal string")
	}
	// Parse the hexadecimal string to an integer
	intValue, err := strconv.ParseInt(hex, 16, 0)
	if err != nil {
		return 0, err
	}
	// Normalize the integer to 0-1 range as a float64
	return float32(intValue) / 255.0, nil
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

// SplitRGBAWithLocation quickly provides the R G B A values as individual variables prefixed by their GL location.
//
// This is purely a convenience method for inline GL operations.
func (c RGBA) SplitRGBAWithLocation(l int32) (location int32, r, g, b, a float32) {
	return l, c.R, c.G, c.B, c.A
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
