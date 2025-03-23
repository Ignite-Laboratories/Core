package main

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/when"
	"time"
)

type BasicSystem struct {
	core.System
}

func NewBasicSystem(periodicity time.Duration) BasicSystem {
	var sys BasicSystem
	sys.System = core.Impulse.CreateSystem(true, sys.Loop, when.After.Period(&periodicity))
	return sys
}

func (s *BasicSystem) Loop(ctx core.Context) {
	fmt.Printf("Stimulated on %d\n", ctx.Beat)
}
