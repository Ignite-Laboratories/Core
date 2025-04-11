package main

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/when"
)

func main() {
	core.Impulse.Loop(sample, when.Always, false)
	core.Impulse.Spark()
}

func sample(ctx core.Context) {
	fmt.Println("here")
}
