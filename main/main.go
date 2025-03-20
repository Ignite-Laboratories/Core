package main

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/potential"
	"time"
)

// In this example -
//	- The loop should activate every third 16th beat
// 	- The block should activate every 16th beat
//	- The stimulation should activate every 8th beat

func main() {
	// Loop every 16th beat as fast as calculable
	core.Impulse.Loop(core.When(potential.Modulo(16), calculation))

	// Block every 16 beats by 1 second
	core.Impulse.Block(core.When(potential.Modulo(16), regulation))

	// Stimulate every 8 beats
	core.Impulse.Stimulate(core.When(potential.Modulo(8), stimulation))

	// Make it so
	_ = core.Impulse.Spark()
}

func stimulation(ctx core.Context) {
	println("Stimulating beat ", ctx.Beat)
}

func regulation(ctx core.Context) {
	println("Regulating beat ", ctx.Beat)
	time.Sleep(1 * time.Second)
}

func calculation(ctx core.Context) {
	println("Calculating on beat ", ctx.Beat)
	time.Sleep(2500 * time.Millisecond)
}
