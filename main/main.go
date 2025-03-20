package main

import (
	"github.com/ignite-laboratories/core"
)

func main() {
	// Make it so
	loopingActivation := core.Impulse.Loop(loop)
	loopingActivation.Potential()
	_ = core.Impulse.Spark()
}

func loop(ctx core.Context) {

}
