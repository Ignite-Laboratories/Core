package main

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/condition"
	"time"
)

var stimFreq = 16.0
var muteFreq = 1.0

var stim = core.Impulse.Loop(Stimulate, condition.Frequency(&stimFreq), false)

func main() {
	// Mute/Unmute the stimulation every three seconds
	core.Impulse.Loop(Toggle, condition.Frequency(&muteFreq), false)

	// Trim down the resistance cyclically
	core.Impulse.Loop(TrimResistance, condition.Always, false)

	// Set the initial resistance high
	core.Impulse.Resistance = 10000000

	// Make it so
	core.Impulse.Spark()
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

func TrimResistance(ctx core.Context) {
	time.Sleep(time.Millisecond * 5000)
	fmt.Printf("Trimming resistance\n")
	core.Impulse.Resistance /= 2
}
