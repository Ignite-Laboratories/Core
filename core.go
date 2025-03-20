package core

import (
	"sync/atomic"
	"time"
)

func init() {
	Alive = true
	Inception = time.Now()
	ID = NextID() // Default ID 1
	Impulse.activations = make(map[uint64]*Activation)
}

// Alive globally keeps activations looping until it is set to false.
var Alive bool

// Inception provides the moment this operating system was initialized.
var Inception time.Time

// ID is the operating system identifier - it defaults to 1.
var ID uint64

// Impulse is the core neural impulse engine.
var Impulse engine

// currentId holds the last provided identifier.
var currentId uint64

// NextID provides a thread-safe unique identifier to every caller.
func NextID() uint64 {
	return atomic.AddUint64(&currentId, 1)
}

// Shutdown waits a period of time before setting Alive to false.
func Shutdown(period time.Duration) {
	time.Sleep(period)
	Alive = false
}
