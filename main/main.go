package main

import (
	"github.com/ignite-laboratories/core"
	"time"
)

func main() {
	loopingSystem := NewLoopingSystem(time.Millisecond * 500)
	loopingSystem.Activate()

	mutingSystem := NewMutingSystem(loopingSystem, time.Second*3)
	mutingSystem.Activate()

	core.Impulse.Spark()
}
