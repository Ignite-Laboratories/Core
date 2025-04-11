package main

import (
	"fmt"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/temporal"
	"github.com/ignite-laboratories/core/when"
	"github.com/ignite-laboratories/host/hydra"
	"github.com/ignite-laboratories/host/mouse"
	"time"
)

func main() {
	temporal.Loop(core.Impulse, when.Frequency(std.HardRef(4.0).Ref), false, sample)
	hydra.CreateFullscreenWindow(core.Impulse, "glitter", Render, when.Frequency(std.HardRef(60.0).Ref), false)
	core.Impulse.Spark()
	time.Sleep(time.Second)
}

func sample(ctx core.Context) {
	fmt.Println(mouse.Sample())
}

func Render(ctx core.Context) {
	gl.ClearColor(1.0, 0.0, 0.0, 1.0) // RGB color
	gl.Clear(gl.COLOR_BUFFER_BIT)
}
