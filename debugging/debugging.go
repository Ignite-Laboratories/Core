package debugging

import (
	"bytes"
	"runtime"
)

// GetGoroutineID gets the ID of the currently executing
// goroutine by parsing it from the stack trace buffer.
//
// NOTE: There are no guarantees behind the stack trace's
// formatting!  This may or may not function going forward,
// but is pivotal in understanding -how- goroutines are
// interoperating.
//
//	Works consistently on go 1.24.1
func GetGoroutineID() uint64 {
	// Get the stack trace buffer
	buf := make([]byte, 64)
	buf = buf[:runtime.Stack(buf, false)]

	var parser = func(b []byte) uint64 {
		var n uint64
		for _, c := range b {
			n = n*10 + uint64(c-'0')
		}
		return n
	}

	// Parse the goroutine ID from the stack trace
	// Sample stack trace: "goroutine 10 [running]: ..."
	// Look for the goroutine number after "goroutine "
	fields := bytes.Fields(buf)
	if len(fields) >= 2 {
		return parser(fields[1])
	}
	return 0
}
