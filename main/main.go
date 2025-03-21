package main

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/when"
	"time"
)

var threshold = 1024

func main() {
	ls := NewLoopingSystem(time.Millisecond * 500)
	ls.Activate()

	ms := NewMutingSystem(ls, time.Second*3)
	ms.Activate()

	core.Impulse.Loop(Calculate, when.Always)
	core.Impulse.Block(Regulate, when.Always)
	core.Impulse.Stimulate(PrintBeat, when.Always)

	core.Impulse.Spark()
}

func Calculate(ctx core.Context) {
	for core.Impulse.Beat < threshold {
	}
}

func Regulate(ctx core.Context) {
	time.Sleep(time.Millisecond * 10)
}

func PrintBeat(ctx core.Context) {
	fmt.Printf("Beat %d\n", ctx.Beat)
}
