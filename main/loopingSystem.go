package main

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"time"
)

type LoopingSystem struct {
	core.System
	duration time.Duration
}

func NewLoopingSystem(duration time.Duration) *LoopingSystem {
	sys := LoopingSystem{
		duration: duration,
	}
	sys.ID = core.NextID()
	sys.LoopFunc = sys.Loop
	sys.PaceFunc = sys.Pace

	return &sys
}

func (s *LoopingSystem) Pace(ctx core.Context) bool {
	return time.Now().Sub(ctx.LastActivation.Inception) > s.duration
}

func (s *LoopingSystem) Loop(ctx core.Context) {
	fmt.Printf("Stimulated on %d\n", ctx.Beat)
}
