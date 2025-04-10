package std

import "sync"

// Synchro represents a way to synchronize data between two threads.
//
// The mechanic is simple - one thread creates a contextual "packet" of data to send to another thread
// over a channel.  The other thread should handle messages on that channel and manipulate the packet
// as desired.  The channel acts as a "bridge" between two spinning threads that need periodic
// synchronization.
//
// To send data using a synchro, use:
//
//	msg = std.SynchroSend(bridge, message)
//
// To process data using a synchro, use:
//
//	std.SynchroEngage(bridge, func(data) { ... })
type Synchro[T any] struct {
	sync.WaitGroup
	Action func()
}

// SynchroSend sends the provided data over the bridge and waits for a result.
func SynchroSend[T any](bridge chan *Synchro[T], action func()) {
	synchro := &Synchro[T]{Action: action}
	synchro.Add(1)
	bridge <- synchro
	synchro.Wait()
}

// SynchroEngage handles incoming messages on the provided channel and then calls Done() on the Synchro.
func SynchroEngage[T any](bridge chan *Synchro[T]) {
	for {
		select {
		case synchro := <-bridge:
			synchro.Action()
			synchro.Done()
		default:
			return
		}
	}
}
