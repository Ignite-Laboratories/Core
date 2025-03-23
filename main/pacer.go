package main

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"time"
)

type Pacer struct {
	core.System
	Duration time.Duration
}

func NewPacer(duration time.Duration) *core.System {
	sys := Pacer{
		Duration: duration,
	}
	sys.ID = core.NextID()
	sys.LoopFunc = sys.Loop
	sys.WhenFunc = sys.Pace

	return &sys.System
}

func (s *Pacer) Pace(ctx core.Context) bool {
	return time.Now().Sub(ctx.LastActivation.Inception) > s.Duration
}

func (s *Pacer) Loop(ctx core.Context) {
	fmt.Printf("Stimulated on %d\n", ctx.Beat)
}
