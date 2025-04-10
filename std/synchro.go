package std

import "sync"

// Synchro represents a way to synchronize execution across threads.
//
// To send execution using a synchro, first create one using make.  Then Engage the synchro
// in "main loop" of the thread you wish to execute on.  The calling thread can Send the
// desired action to the other loop for execution.
//
//	 global -
//		 var synchro = make(std.Synchro)
//
//	 main loop -
//	  for {
//		  synchro.Engage()
//	   ...
//	  }
//
//	 sender -
//		 synchro.Send(func() { ... })
type Synchro chan *syncAction

type syncAction struct {
	sync.WaitGroup
	action func()
}

// Send sends the provided action over the synchro channel and waits for it to be executed.
func (s Synchro) Send(action func()) {
	syn := &syncAction{action: action}
	syn.Add(1)
	s <- syn
	syn.Wait()
}

// Engage handles incoming immediate messages on the Synchro channel before returning control.
func (s Synchro) Engage() {
	for {
		select {
		case syn := <-s:
			syn.action()
			syn.Done()
		default:
			return
		}
	}
}
