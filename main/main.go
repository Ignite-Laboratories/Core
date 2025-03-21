package main

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/when"
	"time"
)

var mutable *core.Activation

func main() {
	var muteDuration = time.Second * 3
	core.Impulse.Loop(toggleMute, when.After.Period(muteDuration))

	var loopDuration = time.Millisecond * 500
	mutable = core.Impulse.Loop(loop, when.After.Period(loopDuration))
	mutable.Mute()

	core.Impulse.Spark()
}

func toggleMute(ctx core.Context) {
	mutable.Muted = !mutable.Muted
	if mutable.Muted {
		fmt.Printf("Muted on %d\n", ctx.Beat)
	} else {
		fmt.Printf("Un-muted on %d\n", ctx.Beat)
	}
}

func loop(ctx core.Context) {
	fmt.Printf("Stimulated on %d\n", ctx.Beat)
}
