package main

import (
	"fmt"
	"github.com/ignite-laboratories/core/enum/greek"
	"github.com/ignite-laboratories/core/std/name"
	"github.com/ignite-laboratories/core/sys/tiny"
)

func main() {
	fmt.Println(greek.Lower.SigmaFinal)
	name.New("bob")

	name.Filtered(tiny.NameFilter)
	name.Tiny()
}
