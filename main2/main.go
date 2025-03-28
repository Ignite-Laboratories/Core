package main

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/host/mouse"
)

func init() {
	//temporal.Analyzer[std.XY[int], any, any](core.Impulse, when.EighthSpeed(&mouse.SampleRate), false, Print, mouse.Coordinates)
	mouse.Coordinates.OnChange = Feedback
	mouse.Coordinates.Unmute()
}

func main() {
	core.Impulse.Spark()
}

func Feedback(oldVal std.Data[std.XY[int]], newVal std.Data[std.XY[int]]) {
	if newVal.Point.X > 1024 {
		mouse.SampleRate = 2048.0
	} else {
		mouse.SampleRate = 2.0
	}
	fmt.Println(newVal.Point)
}

func Print(ctx core.Context, cache *any, data []std.Data[std.XY[int]]) any {
	points := make([]std.XY[int], len(data))
	for i, v := range data {
		points[i] = v.Point
	}
	fmt.Printf("[%d] %v\n", ctx.Beat, points)
	return nil
}
