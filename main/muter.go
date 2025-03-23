package main

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/when"
	"time"
)

type MutingSystem struct {
	core.System
	toMute *core.System
}

func NewMutingSystem(systemToMute *core.System, periodicity time.Duration) MutingSystem {
	sys := MutingSystem{
		toMute: systemToMute,
	}
	sys.System = core.Impulse.CreateSystem(true, sys.Loop, when.After.Period(&periodicity))
	return sys
}

func (s *MutingSystem) Loop(ctx core.Context) {
	activation := s.toMute
	activation.Muted = !activation.Muted
	if activation.Muted {
		fmt.Printf("Muted on %d\n", ctx.Beat)
	} else {
		fmt.Printf("Un-muted on %d\n", ctx.Beat)
	}
}
