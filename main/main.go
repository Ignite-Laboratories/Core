package main

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/when"
	"time"
)

// Setup Systems
var muter = NewMutingSystem(&basic.System, time.Second*3)
var basic = NewBasicSystem(time.Millisecond * 500)

// initialize
func init() {
	basic.Mute()
}

// Run
func main() {
	core.Impulse.Loop(TrimResistance, when.Always)
	core.Impulse.Stimulate(PrintBeat, when.Always)

	core.Impulse.Resistance = 1024000000
	core.Impulse.Spark()
}

/**
Loops
*/

func TrimResistance(ctx core.Context) {
	for core.Impulse.Beat < 22 {
	}
	core.Impulse.Resistance /= 2
}

func PrintBeat(ctx core.Context) {
	fmt.Printf("Beat %d\n", ctx.Beat)
}
