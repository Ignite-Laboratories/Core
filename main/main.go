package main

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/when"
	"time"
)

var stim = core.Impulse.Loop(Stimulate, when.Frequency(std.HardRef(16.0).Ref), true)

// Make it so
func init() {
	go core.Impulse.Spark()
	core.Impulse.MaxFrequency = 1
}

func main() {
	// Mute/Unmute the stimulation every three seconds
	core.Impulse.Loop(Toggle, when.Frequency(std.HardRef(0.5).Ref), false)

	// Trim down the resistance cyclically
	core.Impulse.Loop(AdjustFrequency, when.Always, false)

	// Set the initial resistance high
	core.Impulse.Resistance = 10000000

	core.WhileAlive()
}

func Toggle(ctx core.Context) {
	if stim.Muted {
		fmt.Printf("[%d] Unmuting\n", ctx.Beat)
	} else {
		fmt.Printf("[%d] Muting\n", ctx.Beat)
	}
	stim.Muted = !stim.Muted
}

func Stimulate(ctx core.Context) {
	fmt.Printf("[%d] Stimulated\n", ctx.Beat)
}

func AdjustFrequency(ctx core.Context) {
	time.Sleep(time.Second * 5)
	fmt.Printf("[%d] Adjusting frequency\n", ctx.Beat)
	core.Impulse.MaxFrequency += 1
}
