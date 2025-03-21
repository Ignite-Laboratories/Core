package core

import (
	"time"
)

type System struct {
	Entity
	active     bool
	activation *Activation
	duration   time.Duration
	LoopFunc   func(ctx Context)
	PaceFunc   func(ctx Context) bool
}

func (s *System) GetActive() bool {
	return s.active
}

// Activate adds the system to the impulse engine in an asynchronous fashion, if it isn't already active.
func (s *System) Activate() {
	if s.active {
		return
	}
	s.activation = Impulse.Loop(s.LoopFunc, s.PaceFunc)
}

// ActivateSynchronously adds the system to the impulse engine in a blocking fashion, if it isn't already active.
func (s *System) ActivateSynchronously() {
	if s.active {
		return
	}
	s.activation = Impulse.Block(s.LoopFunc, s.PaceFunc)
}

func (s *System) GetActivation() *Activation {
	return s.activation
}
