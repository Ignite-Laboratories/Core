package main

import (
	"github.com/ignite-laboratories/core"
)

func main() {
	core.Verbose = true
	//core.Impulse.GivenName = core.GivenName{
	//	Name:        "Arwen",
	//	Description: "Aetheric beauty, elegance, and guidance",
	//}
	core.Impulse.Spark()
}
