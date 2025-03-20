package main

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/when"
	"time"
)

func main() {
	// Block every 16th beat by 1 second
	core.Impulse.Block(when.Modulo(16, regulate))

	// Loop every 16th beat as fast as calculable
	// The loop takes 2.5 seconds, so it'll activate with every third blocking impulse.
	core.Impulse.Loop(when.Modulo(16, pulse))

	// Stimulate every 8th beat
	core.Impulse.Stimulate(when.Modulo(8, stimulate))

	// Make it so
	_ = core.Impulse.Spark()
}

func regulate(ctx core.Context) {
	println("Regulating beat ", ctx.Beat)
	time.Sleep(1 * time.Second)
}

func pulse(ctx core.Context) {
	println("Pulsing on beat ", ctx.Beat)
	time.Sleep(2500 * time.Millisecond)
}

func stimulate(ctx core.Context) {
	println("Stimulating beat ", ctx.Beat)
}
