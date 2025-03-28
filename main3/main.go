package main

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/host/mouse"
	"math"
)

func init() {
	//mouse.Coordinates.Unmute()
	//temporal.Analyzer[std.XY[int], any, any](core.Impulse, when.EighthSpeed(&mouse.SampleRate), false, Print, mouse.Coordinates)
	mouse.Reaction(std.HardRef(2048.0).Ref, Velocity)
	mouse.Reaction(&mouse.SampleRate, Feedback)
}

func main() {
	core.Impulse.Spark()
}

func Velocity(ctx core.Context, old *std.Data[std.XY[int]], current *std.Data[std.XY[int]]) {
	if old == nil {
		return
	}
	delta := current.Point.X - old.Point.X
	deltaAbs := math.Abs(float64(delta))
	if deltaAbs > 100 {
		fmt.Println("Slow down!")
	}
}

func Feedback(ctx core.Context, old *std.Data[std.XY[int]], current *std.Data[std.XY[int]]) {
	if current.Point.X > 1024 {
		mouse.SampleRate = 2048.0
	} else {
		mouse.SampleRate = 2.0
	}
	fmt.Println(current.Point)
}

func Print(ctx core.Context, cache *any, data []std.Data[std.XY[int]]) any {
	points := make([]std.XY[int], len(data))
	for i, v := range data {
		points[i] = v.Point
	}
	fmt.Printf("[%d] %v\n", ctx.Beat, points)
	return nil
}
