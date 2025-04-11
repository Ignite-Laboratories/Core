package main

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/temporal"
	"github.com/ignite-laboratories/core/when"
)

func main() {
	temporal.Loop(core.Impulse, when.Frequency(std.HardRef(4.0).Ref), false, sample)
}

func sample(ctx core.Context) {
	fmt.Println("here")
}
