package core

import (
	"sync/atomic"
	"time"
)

func init() {
	Impulse.activations = make(map[uint64]*activation)
}

// Alive globally keeps activations looping until it is set to false.
var Alive = true

// Inception provides the moment this operating system was initialized.
var Inception = time.Now()

// ID is the operating system identifier - it defaults to 1.
var ID uint64 = NextID()

// Impulse is the neural impulse engine.
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
