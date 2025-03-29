package main

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/temporal"
	"github.com/ignite-laboratories/core/when"
	"github.com/ignite-laboratories/host/mouse"
)

func main() {
	temporal.Calculation(core.Impulse, when.Frequency(&mouse.SampleRate), false, CalcCoords)
	core.Impulse.Spark()
}

func CalcCoords(ctx core.Context) std.XY[int] {
	coords := mouse.SampleCoordinates()
	if coords.X < 1024 {
		fmt.Println(coords)
	}
	return *coords
}
