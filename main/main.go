package main

import "github.com/ignite-laboratories/core"

func main() {
	core.Self.Loop(printBeat)
	_ = core.Self.Ignite()
}

func printBeat(ctx core.Context) {
	println(ctx.Beat)
}
