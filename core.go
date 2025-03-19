package core

import (
	"sync/atomic"
	"time"
)

func init() {
	Inception = time.Now() // Set the start time of the operating system
}

// Alive globally keeps any persistent routines alive until it is set to false.
var Alive = true

// Inception provides the moment this operating system was initialized.
var Inception time.Time

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
