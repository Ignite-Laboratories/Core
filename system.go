package core

import (
	"time"
)

// System is a neural structure that hosts activation.
type System struct {
	Entity
	active     bool
	activation *Activation
	duration   time.Duration
	LoopFunc   func(ctx Context)
	WhenFunc   func(ctx Context) bool
}

// GetActive returns whether the system is currently active or not.
func (s *System) GetActive() bool {
	return s.active
}

// Activate adds the system to the impulse engine in an asynchronous fashion, if it isn't already active.
func (s *System) Activate(async bool) {
	if s.active {
		return
	}
	s.active = true

	whenActive := func(ctx Context) bool {
		return s.active && s.WhenFunc(ctx)
	}

	if async {
		s.activation = Impulse.Loop(s.LoopFunc, whenActive)
	} else {
		s.activation = Impulse.Block(s.LoopFunc, whenActive)
	}
}

// Mute suppresses system activation.
func (s *System) Mute() {
	s.activation.Muted = true
}

// Unmute un-suppresses system activation.
func (s *System) Unmute() {
	s.activation.Muted = false
}

// GetActivation returns a pointer to the system's activation.
func (s *System) GetActivation() *Activation {
	return s.activation
}
