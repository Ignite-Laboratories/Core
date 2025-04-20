package std

import (
	"fmt"
	"github.com/ignite-laboratories/core"
)

// XY is a general structure for holding generic (x,y) coordinate values.
type XY[T core.Numeric] struct {
	X T
	Y T
}

// RandomXY generates a random set of XY values of the provided type.
func RandomXY[T core.Numeric]() XY[T] {
	return XY[T]{
		X: core.RandomNumber[T](),
		Y: core.RandomNumber[T](),
	}
}

// XYComparator returns if the two XY values are equal in values.
func XYComparator[T core.Numeric](a XY[T], b XY[T]) bool {
	return a.X == b.X && a.Y == b.Y
}

func (c XY[T]) String() string {
	return fmt.Sprintf("(%d, %d)", c.X, c.Y)
}
