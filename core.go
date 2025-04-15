package core

import (
	"debug/buildinfo"
	"fmt"
	"os"
	"strings"
	"sync/atomic"
	"time"
)

func init() {
	exe, _ := os.Executable()
	exeInfo, _ = buildinfo.ReadFile(exe)

	fmt.Printf("JanOS %v\n", getModuleVersion(ModuleName))
	fmt.Println("© 2025, Ignite Laboratories")
	fmt.Println("---------------------------")

	initializeNameDB()
	Impulse = NewEngine()
}

var exeInfo *buildinfo.BuildInfo

func getModuleVersion(module string) string {
	for _, dep := range exeInfo.Deps {
		if strings.Contains(dep.Path, "github.com/ignite-laboratories/"+ModuleName) {
			return dep.Version
		}
	}
	return "unknown"
}

// ModuleReport reports the version information of a module to the console.
func ModuleReport(module string) {
	fmt.Printf(" - [%v] %v\n", module, getModuleVersion(module))
}

// SubmoduleReport reports the version information of a submodule to the console.
func SubmoduleReport(module string, submodule string) {
	fmt.Printf(" - [%v].[%v]\n", module, submodule)
}

// Alive globally keeps neural activity firing until set to false - it's true by default.
var Alive = true

// Inception provides the moment this operating system was initialized.
var Inception = time.Now()

// Impulse is the host engine.
var Impulse *Engine

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

// Shutdown waits a period of time before calling ShutdownNow
func Shutdown(period time.Duration) {
	Printf(ModuleName, "shutting down in %v\n", period)
	time.Sleep(period)
	ShutdownNow()
}

// ShutdownNow immediately sets Alive to false.
func ShutdownNow() {
	Printf(ModuleName, "shutting down\n")
	Alive = false
}

// Exit briefly pauses to let other threads clean up before calling os.Exit
//
// You may optionally provide an OS exit code, otherwise '0' is implied.
func Exit(exitCode ...int) {
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

// AbsDuration returns the absolute value of the provided duration.
func AbsDuration(d time.Duration) time.Duration {
	if d < 0 {
		d = -d
	}
	return d
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
	ShutdownNow()
	Exit(1)
}

// FatalfCode prepends the provided string format with a module identifier, prints it to the console, and then calls core.ShutdownNow(exitCode).
func FatalfCode(exitCode int, module string, format string, a ...any) {
	fmt.Printf("[%v] %v", module, fmt.Sprintf(format, a...))
	ShutdownNow()
	Exit(exitCode)
}
