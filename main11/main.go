package main

import (
	"github.com/ignite-laboratories/core"
)

func main() {
	core.Impulse.GivenName = *core.LookupName("Eve")
	core.Impulse.Spark()
}
