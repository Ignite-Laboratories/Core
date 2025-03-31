package std

import (
	"fmt"
	"github.com/ignite-laboratories/core"
)

// XY is a general structure for holding (x,y) coordinate values.
type XY[T core.Numeric] struct {
	X T
	Y T
	Z T
	W T
}

func XYComparator[T core.Numeric](a XY[T], b XY[T]) bool {
	return a.X == b.X && a.Y == b.Y
}

func (c XY[T]) String() string {
	return fmt.Sprintf("(%d, %d)", c.X, c.Y)
}

// XYZ is a general structure for holding (x,y,z) coordinate values.
type XYZ[T core.Numeric] struct {
	X T
	Y T
	Z T
}

func XYZComparator[T core.Numeric](a XYZ[T], b XYZ[T]) bool {
	return a.X == b.X && a.Y == b.Y && a.Z == b.Z
}

func (c XYZ[T]) String() string {
	return fmt.Sprintf("(%d, %d, %d)", c.X, c.Y, c.Z)
}

// XYZW is a general structure for holding (x,y,z,w) coordinate values.
type XYZW[T core.Numeric] struct {
	X T
	Y T
	Z T
	W T
}

func XYZWComparator[T core.Numeric](a XYZW[T], b XYZW[T]) bool {
	return a.X == b.X && a.Y == b.Y && a.Z == b.Z && a.W == b.W
}

func (c XYZW[T]) String() string {
	return fmt.Sprintf("(%d, %d, %d, %d)", c.X, c.Y, c.Z, c.W)
}
