package core

import (
	"fmt"
	"os"
	"sync/atomic"
	"time"
)

func init() {
	fmt.Println("      JanOS")
	fmt.Println("Ignite Laboratories")
	fmt.Println("-------------------")

	Impulse.Name = "core impulse engine"
	Verbose = true
}

// Alive globally keeps neural activity firing until set to false - it's true by default.
var Alive = true

// Inception provides the moment this operating system was initialized.
var Inception = time.Now()

// Impulse is the host engine.
var Impulse = NewEngine()

// ID is the operating system identifier - it defaults to 1.
var ID uint64 = NextID()

// DefaultWindow is the default dimensional window of observance - 2 seconds.
var DefaultWindow = 2 * time.Second

var TrimFrequency = 1024.0 //hz

var Verbose bool

// currentId holds the last provided identifier.
var currentId uint64

// NextID provides a thread-safe unique identifier to every caller.
func NextID() uint64 {
	return atomic.AddUint64(&currentId, 1)
}

// Shutdown waits a period of time before setting Alive to false.
func Shutdown(period time.Duration, exitCode ...int) {
	Printf(ModuleName, "shutting down in %v\n", period)
	go func() {
		time.Sleep(period)
		ShutdownNow(exitCode...)
	}()
}

// ShutdownNow immediately sets Alive to false.
func ShutdownNow(exitCode ...int) {
	Printf(ModuleName, "shutting down\n")
	Alive = false
	// Give the threads a brief moment to clean themselves up.
	time.Sleep(time.Second)
	if len(exitCode) > 0 {
		os.Exit(exitCode[0])
	} else {
		os.Exit(0)
	}
}

// WhileAlive can be used to hold a main function open.
func WhileAlive() {
	for Alive {
		// Give the host some breathing room.
		time.Sleep(time.Millisecond)
	}
}

// DurationToHertz converts a time.Duration into Hertz.
func DurationToHertz(d time.Duration) float64 {
	if d < 0 {
		d = 0
	}
	s := float64(d) / 1e9
	hz := 1 / s
	return hz
}

// HertzToDuration converts a Hertz value to a time.Duration.
func HertzToDuration(hz float64) time.Duration {
	if hz <= 0 {
		// No division by zero
		hz = 1e-100 //math.SmallestNonzeroFloat64
	}
	s := 1 / hz
	ns := s * 1e9
	return time.Duration(ns)
}

// Verbosef prepends the provided data with a module identifier and then prints it to the console, but only if core.Verbose is true.
func Verbosef(module string, format string, a ...any) {
	if Verbose {
		fmt.Printf("[%v] %v", module, fmt.Sprintf(format, a...))
	}
}

// Printf prepends the provided data with a module identifier and then prints it to the console.
func Printf(module string, format string, a ...any) {
	fmt.Printf("[%v] %v", module, fmt.Sprintf(format, a...))
}

// Fatalf prepends the provided data with a module identifier, prints it to the console, and then calls core.ShutdownNow(1)
func Fatalf(module string, format string, a ...any) {
	fmt.Printf("[%v] %v", module, fmt.Sprintf(format, a...))
	ShutdownNow(1)
}
