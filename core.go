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

	Impulse.Name = "Eve"
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

// DefaultObservanceWindow is the default dimensional window of observance - 2 seconds.
var DefaultObservanceWindow = 2 * time.Second

// TrimFrequency sets the global frequency for dimensional trimmers.
var TrimFrequency = 1024.0 //hz

// Verbose sets whether the system should emit more verbose logs or not.
var Verbose bool

// currentId holds the last provided identifier.
var currentId uint64

// NextID provides a thread-safe unique identifier to every caller.
func NextID() uint64 {
	return atomic.AddUint64(&currentId, 1)
}

// Shutdown waits a period of time before setting Alive to false.
//
// You may optionally provide an OS exit code, otherwise '0' is implied.
func Shutdown(period time.Duration, exitCode ...int) {
	Printf(ModuleName, "shutting down in %v\n", period)
	go func() {
		time.Sleep(period)
		ShutdownNow(exitCode...)
	}()
}

// ShutdownNow immediately sets Alive to false.
//
// You may optionally provide an OS exit code, otherwise '0' is implied.
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

// WhileAlive can be used to efficiently hold a main function open.
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
		hz = 1e-100 // math.SmallestNonzeroFloat64 <- NOTE: Raspberry Pi doesn't handle this constant well
	}
	s := 1 / hz
	ns := s * 1e9
	return time.Duration(ns)
}

// Verbosef prepends the provided string format with a module identifier and then prints it to the console, but only if core.Verbose is true.
func Verbosef(module string, format string, a ...any) {
	if Verbose {
		fmt.Printf("[%v] %v", module, fmt.Sprintf(format, a...))
	}
}

// Printf prepends the provided string format with a module identifier and then prints it to the console.
func Printf(module string, format string, a ...any) {
	fmt.Printf("[%v] %v", module, fmt.Sprintf(format, a...))
}

// Fatalf prepends the provided string format with a module identifier, prints it to the console, and then calls core.ShutdownNow(1).
func Fatalf(module string, format string, a ...any) {
	fmt.Printf("[%v] %v", module, fmt.Sprintf(format, a...))
	ShutdownNow(1)
}

// FatalfCode prepends the provided string format with a module identifier, prints it to the console, and then calls core.ShutdownNow(exitCode).
func FatalfCode(exitCode int, module string, format string, a ...any) {
	fmt.Printf("[%v] %v", module, fmt.Sprintf(format, a...))
	ShutdownNow(exitCode)
}
