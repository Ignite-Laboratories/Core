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
	Data *T
}

// SynchroSend sends the provided data over the bridge and waits for a result.
func SynchroSend[T any](bridge chan *Synchro[T], message *T) *T {
	synchro := &Synchro[T]{Data: message}
	synchro.Add(1)
	bridge <- synchro
	synchro.Wait()
	return synchro.Data
}

// SynchroEngage bridges the action against the provided channel and then calls Done() on the Synchro.
func SynchroEngage[T any](bridge chan *Synchro[T], action func(*T)) {
	for {
		select {
		case synchro := <-bridge:
			action(synchro.Data)
			synchro.Done()
		default:
			return
		}
	}
}
