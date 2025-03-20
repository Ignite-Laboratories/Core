package core

import (
	"sync/atomic"
	"time"
)

func init() {
	Alive = true
	Inception = time.Now()
	Self.ID = NextID() // Default ID 1
}

// Alive globally keeps activations looping until it is set to false.
var Alive bool

// Inception provides the moment this operating system was initialized.
var Inception time.Time

// Self provides the core engine host structure.
var Self self

// masterId holds the last provided identifier.
var masterId uint64

// NextID provides a thread-safe unique identifier to every caller.
func NextID() uint64 {
	return atomic.AddUint64(&masterId, 1)
}

// Shutdown waits a period of time before setting Alive to false.
func Shutdown(period time.Duration) {
	time.Sleep(period)
	Alive = false
}
