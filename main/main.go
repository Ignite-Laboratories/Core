package main

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/when"
	"time"
)

var muter = NewMuter(pacer, time.Second*3)
var pacer = NewPacer(time.Millisecond * 500)

func main() {
	muter.Activate(true)
	pacer.Activate(true)
	pacer.Mute()
	
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
