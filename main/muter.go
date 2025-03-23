package main

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"time"
)

type Muter struct {
	Pacer

	toMute *core.System
}

func NewMuter(systemToMute *core.System, duration time.Duration) *core.System {
	sys := Muter{
		toMute: systemToMute,
	}
	sys.ID = core.NextID()
	sys.LoopFunc = sys.Loop
	sys.WhenFunc = sys.Pace
	sys.Duration = duration

	return &sys.System
}

func (s *Muter) Loop(ctx core.Context) {
	activation := s.toMute.GetActivation()
	activation.Muted = !activation.Muted
	if activation.Muted {
		fmt.Printf("Muted on %d\n", ctx.Beat)
	} else {
		fmt.Printf("Un-muted on %d\n", ctx.Beat)
	}
}
