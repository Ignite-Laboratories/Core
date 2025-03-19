package core

import (
	"fmt"
	"github.com/ignite-laboratories/support/threadSafe"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

// Engine is a source of stimulating Go routines using a synchronized temporal context.
type Engine struct {
	Entity

	// Running indicates if the engine is currently executing kernels or not.
	Running bool

	// beat increments with every step until all kernels have finished execution before looping over.
	beat int

	// impulseStats provides the latest runtimeStats of the engine's main loop.
	impulseStats runtimeStats

	// serialNumber is the current unique identifier to give to the next invocation.
	serialNumber uint64

	// invocations provides a thread-safe slice of activation pointer references.
	invocations *threadSafe.Slice[*activation]

	// invokeStats provides a map of the latest runtimeStats of every activation.
	invokeStats sync.Map
}

// NewEngine creates a new instance of Engine.
func NewEngine() Engine {
	var e Engine
	e.ID = NextID()
	e.invocations = threadSafe.NewSlice[*activation]()
	return e
}

// Block activates the provided Action Potential on every impulse in a blocking fashion.
func (e *Engine) Block(action Action, potential Potential) {
	a := newBlockingActivation(NewActionPotential(action, potential))
	a.serialNumber = e.nextSerialNumber()
	e.invocations.Add(&a)
}

// Stimulate activates the provided Action Potential on every impulse in an asynchronous fashion.
func (e *Engine) Stimulate(action Action, potential Potential) {
	a := newStimulation(NewActionPotential(action, potential))
	a.serialNumber = e.nextSerialNumber()
	e.invocations.Add(&a)
}

// Loop cyclically activates the provided Action Potential on the next possible impulse.
func (e *Engine) Loop(action Action, potential Potential) {
	a := newLoopingActivation(NewActionPotential(action, potential))
	a.serialNumber = e.nextSerialNumber()
	e.invocations.Add(&a)
}

func (e *Engine) Cluster(actionPotentials []Action, wg *sync.WaitGroup) {
}

// nextSerialNumber returns an engine-specific unique identifier to seed invocations with.
//
// This is to keep the global unique identifiers from growing rapidly fast.
func (e *Engine) nextSerialNumber() uint64 {
	return atomic.AddUint64(&c.serialNumber, 1)
}

// Start begins the engine's main loop.
func (e *Engine) Start() error {
	if e.Running {
		return fmt.Errorf("this engine is already running")
	}
	e.Running = true

	// Set up a wait group for blocking operations
	var wg sync.WaitGroup

	// On the first impulse time is oriented to the system's inception moment
	lastFinishTime := Inception
	lastNow := Inception

	// Loop =)
	for Alive && e.Running {
		// Grab 'now' ASAP!
		now := time.Now()

		// Check if any of the kernels are still running...
		hasExecution := e.colonel.kernels.IfAny(func(inv *activation) bool { return inv.executing })
		if !hasExecution {
			// ...if not, loop the beat back to 0
			e.beat = 0
		}

		// Calculate the impulse stats
		e.impulseStats.Inception = lastNow
		e.impulseStats.Start = lastNow
		e.impulseStats.End = lastFinishTime
		e.impulseStats.RefractoryPeriod = now.Sub(lastFinishTime)

		// Build a temporal context
		var ctx Context
		ctx.ID = NextID()
		ctx.Moment = now
		ctx.Delta = now.Sub(lastNow)
		ctx.Beat = e.beat
		ctx.ActivationStats = e.impulseStats

		// Launch all kernels...
		for _, k := range e.colonel.kernels.All() {
			// Don't re-activation a kernel that is still executing
			if k.executing {
				continue
			}

			wg.Add(1)
			go func() {
				defer wg.Done()
				defer func() {
					if r := recover(); r != nil {
						log.Println("Kernel panic - Serial Number: ", k.SerialNumber)
					}
				}()

				// Grab the last stats for this kernel
				lastStats, exists := e.colonel.stats.Load(k.SerialNumber)
				if exists {
					ctx.ActivationStats = lastStats.(runtimeStats)
					// Calculate the remainder value until 'now'
					ctx.ActivationStats.RefractoryPeriod = now.Sub(ctx.ActivationStats.End)
				}

				// Execute the kernel
				kernelStart := time.Now()
				k.function(ctx)
				kernelEnd := time.Now()

				// Save off the kernel statistics
				e.colonel.stats.Store(k.SerialNumber, runtimeStats{
					Inception: now,
					Start:     kernelStart,
					End:       kernelEnd,
				})
			}()
		}
		// ...and wait for them to finish launching
		wg.Wait()
		finishTime := time.Now()

		// Save off 'now'
		lastNow = now
		lastFinishTime = finishTime

		// Increment the beat
		e.beat++
	}
	return nil
}

// Stop terminates the engine's main loop.
func (e *Engine) Stop() {
	e.Running = false
}
