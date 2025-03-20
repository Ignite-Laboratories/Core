package main

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/when"
)

var mutable *core.Activation

func main() {
	mutable = core.Impulse.Loop(loop)
	core.Impulse.Stimulate(when.Modulo(16, toggle))
	_ = core.Impulse.Spark()
}

func toggle(ctx core.Context) {
	mutable.Muted = !mutable.Muted
}

func loop(ctx core.Context) {
	core.Impulse.Once(when.Even(stimulation))
}

func stimulation(ctx core.Context) {
	fmt.Println("Chain stimulated")
}
