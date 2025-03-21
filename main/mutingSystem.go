package main

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"time"
)

type MutingSystem struct {
	core.System

	toMute   core.Loopable
	duration time.Duration
}

func NewMutingSystem(systemToMute core.Loopable, duration time.Duration) *MutingSystem {
	sys := MutingSystem{
		toMute:   systemToMute,
		duration: duration,
	}
	sys.ID = core.NextID()
	sys.LoopFunc = sys.Loop
	sys.PaceFunc = sys.Pace

	return &sys
}

func (s *MutingSystem) Pace(ctx core.Context) bool {
	return time.Now().Sub(ctx.LastActivation.Inception) > s.duration
}

func (s *MutingSystem) Loop(ctx core.Context) {
	activation := s.toMute.GetActivation()
	activation.Muted = !activation.Muted
	if activation.Muted {
		fmt.Printf("Muted on %d\n", ctx.Beat)
	} else {
		fmt.Printf("Un-muted on %d\n", ctx.Beat)
	}
}
