package main

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/when"
	"time"
)

func main() {
	ls := NewLoopingSystem(time.Millisecond * 500)
	ls.Activate()

	ms := NewMutingSystem(ls, time.Second*3)
	ms.Activate()

	core.Impulse.Loop(TrimResistance, when.Always)
	core.Impulse.Stimulate(PrintBeat, when.Always)

	core.Impulse.Resistance = 1024000000
	core.Impulse.Spark()
}

func TrimResistance(ctx core.Context) {
	for core.Impulse.Beat < 22 {
	}
	core.Impulse.Resistance /= 2
}

func PrintBeat(ctx core.Context) {
	fmt.Printf("Beat %d\n", ctx.Beat)
}
