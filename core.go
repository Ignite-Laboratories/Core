package core

import (
	"sync/atomic"
	"time"
)

func init() {
	// Initialize the core neural impulse engine.
	Impulse.Initialize()
}

// Action functions are provided temporal context when invoked.
type Action func(ctx Context)

// Potential functions are provided temporal context when invoked in order to make decisions.
type Potential func(ctx Context) bool

// Alive globally keeps activations looping until set to false - it's true by default.
var Alive = true

// Inception provides the moment this operating system was initialized.
var Inception = time.Now()

// Impulse is the core neural impulse engine.
var Impulse Engine

// ID is the operating system identifier - it defaults to 1.
var ID uint64 = NextID()

// currentId holds the last provided identifier.
var currentId uint64

// NextID provides a thread-safe unique identifier to every caller.
func NextID() uint64 {
	return atomic.AddUint64(&currentId, 1)
}

// Shutdown waits a period of time before setting Alive to false.
func Shutdown(period time.Duration) {
	go func() {
		time.Sleep(period)
		Alive = false
	}()
}
