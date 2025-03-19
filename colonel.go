package core

import (
	"github.com/ignite-laboratories/support/threadSafe"
	"sync"
	"sync/atomic"
)

// colonel is the engine's director of kernels, bluntly enough!
//
// This is where any and all kernels are given to and taken from the engine.
type colonel struct {
	// serialNumber is the current unique identifier to give to the next enlisted kernel.
	serialNumber uint64

	// kernels provides a thread-safe slice of kernel pointer references.
	kernels *threadSafe.Slice[*activation]

	// stats provides a map of the latest runtimeStats of every kernel.
	stats sync.Map
}

// newColonel creates a new instance of colonel.
func newColonel() *colonel {
	return &colonel{
		kernels: threadSafe.NewSlice[*activation](),
	}
}

// nextSerialNumber returns a colonel-specific unique identifier to seed kernels with.
//
// This is to keep the global unique identifiers from growing rapidly fast.
func (c *colonel) nextSerialNumber() uint64 {
	return atomic.AddUint64(&c.serialNumber, 1)
}

// AddBlockingAction adds a blocking Action to every impulse.
//
//	NOTE: This will slow all operation!
func (c *colonel) AddBlockingAction(action Action) activation {
	inv := newActivation(action)
	inv.SerialNumber = c.nextSerialNumber()
	c.kernels.Add(&inv)
	return inv
}

// addAction adds an asynchronous looping Action to every impulse.
func (c *colonel) addAction(action Action) activation {
	inv := newAsyncInvoke(action)
	inv.SerialNumber = c.nextSerialNumber()
	c.kernels.Add(&inv)
	return inv
}

// addWaitGroupAction adds an asynchronous Action to every impulse that calls wg.Done() whenever it finishes.
func (c *colonel) addWaitGroupAction(action Action, wg *sync.WaitGroup) activation {
	inv := newWaitGroupInvoke(action, wg)
	inv.SerialNumber = c.nextSerialNumber()
	c.kernels.Add(&inv)
	return inv
}

// addStimulatedAction adds an asynchronous Action that activates on every impulse.
func (c *colonel) addStimulatedAction(action Action) activation {
	inv := newStimulatedInvoke(action)
	inv.SerialNumber = c.nextSerialNumber()
	c.kernels.Add(&inv)
	return inv
}

// RemoveKernel removes the provided kernel by it's serialNumber from the colonel's kernel pool and stats map.
func (c *colonel) RemoveKernel(id uint64) {
	c.kernels.RemoveIf(func(inv *activation) bool {
		return inv.serialNumber == id
	})
	c.stats.Delete(id)
}
