package main

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/temporal"
	"github.com/ignite-laboratories/core/when"
)

func main() {
	temporal.Loop(core.Impulse, when.Always, false, sample)
	core.Impulse.Spark()
}

func sample(ctx core.Context) {
	fmt.Println("here")
}
