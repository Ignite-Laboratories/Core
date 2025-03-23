package main

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/when"
	"time"
)

var stimFreq = time.Millisecond * 500
var muteFreq = time.Second * 3

func main() {
	// Stimulate every half second
	stim := core.Impulse.Loop(Stimulate, when.After.Period(&stimFreq))
	stim.Muted = true

	// Mute/Unmute the stimulation every three seconds
	core.Impulse.Loop(func(ctx core.Context) {
		stim.Muted = !stim.Muted
	}, when.After.Period(&muteFreq))

	// Trim down the resistance cyclically
	core.Impulse.Loop(TrimResistance, when.Always)

	// Print out the current beat on every impulse
	core.Impulse.Stimulate(PrintBeat, when.Always)

	// Set the initial resistance high
	core.Impulse.Resistance = 1000000000

	// Make it so
	core.Impulse.Spark()
}

func Stimulate(ctx core.Context) {
	fmt.Printf("Stimulated on %d\n", ctx.Beat)
}

func TrimResistance(ctx core.Context) {
	for core.Impulse.Beat < 22 {
		// Hold the impulse window open
	}
	core.Impulse.Resistance /= 2
}

func PrintBeat(ctx core.Context) {
	fmt.Printf("Beat %d\n", ctx.Beat)
}
