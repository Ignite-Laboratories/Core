package main

import (
	"fmt"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/when"
	"github.com/ignite-laboratories/host/hydra"
	"github.com/ignite-laboratories/host/mouse"
	"math"
)

func init() {
	mouse.Reaction(core.Impulse, &mouse.SampleRate, Velocity)
}

func main() {
	hydra.CreateFullscreenWindow(core.Impulse, "glitter", Render, when.Frequency(std.HardRef(60.0).Ref), false)
	core.Impulse.Spark()
}

func Render(ctx core.Context) {
	gl.ClearColor(0.0, 0.0, 0.0, 0.0) // RGB color
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

func Velocity(ctx core.Context, old std.Data[std.MouseState], current std.Data[std.MouseState]) {
	delta := current.Point.Position.X - old.Point.Position.X
	deltaAbs := math.Abs(float64(delta))
	if deltaAbs > 100 {
		fmt.Println("Slow down!")
	}
}
