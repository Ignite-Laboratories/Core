package main

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/after"
	"github.com/ignite-laboratories/core/when"
	"time"
)

var stimFreq = time.Millisecond * 500
var muteFreq = time.Second * 3

// Stimulate every half second
var stim = core.Impulse.Loop(Stimulate, after.Period(&stimFreq), false)

func main() {
	// Mute/Unmute the stimulation every three seconds
	core.Impulse.Loop(Toggle, after.Period(&muteFreq), false)

	// Trim down the resistance cyclically
	core.Impulse.Loop(TrimResistance, when.Always, false)

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
