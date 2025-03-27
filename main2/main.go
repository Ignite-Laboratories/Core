package main

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/temporal"
	"github.com/ignite-laboratories/core/when"
	"github.com/ignite-laboratories/host"
	"time"
)

type Coordinates struct {
	X int
	Y int
}

var sampleFreq = 16.0
var resonance = 4.0

var mouser = temporal.NewCalculation[Coordinates](core.Impulse, when.Frequency(&sampleFreq), false, GetCoordinates)
var analyzer = temporal.NewAnalysis[Coordinates, any, time.Time](core.Impulse, when.ResonantWith(&sampleFreq, &resonance), false, PrintCoordinates, mouser)

func main() {
	core.Impulse.Spark()
}

var last = core.Inception

func PrintCoordinates(ctx core.Context, cache *time.Time, data []temporal.Data[Coordinates]) any {
	points := make([]Coordinates, len(data))
	for i, v := range data {
		points[i] = v.Point
	}
	fmt.Println(points)
	time.Sleep(100 * time.Millisecond)
	return nil
}

func GetCoordinates(ctx core.Context) Coordinates {
	x, y, _ := host.Mouse.GetCoordinates()
	return Coordinates{x, y}
}
