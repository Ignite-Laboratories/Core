package main

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/potential"
	"time"
)

func main() {
	// Loop every 16 beats as fast as calculable
	core.Self.Loop(core.When(potential.Modulo(16), calculation))

	// Block every 16 beats by 1 second
	core.Self.Block(core.When(potential.Modulo(16), regulation))

	// Stimulate every 16 beats
	core.Self.Stimulate(core.When(potential.Modulo(16), stimulation))

	// Make it so
	_ = core.Self.Ignite()
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
	time.Sleep(3 * time.Second)
}
