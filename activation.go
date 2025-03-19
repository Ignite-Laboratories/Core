package core

import "sync"

// activation is a way of invoking functions on neural impulses.
type activation struct {
	serialNumber uint64
	function     Action
	executing    bool
}

// newBlockingActivation creates a new activation that blocks the impulse.
func newBlockingActivation(function Action) activation {
	var a activation
	a.function = func(ctx Context) {
		a.executing = true
		function(ctx)
		a.executing = false
	}
	return a
}

// newStimulation creates a new activation that fires asynchronously on every impulse.
func newStimulation(function Action) activation {
	var a activation
	a.function = func(ctx Context) {
		go a.function(ctx)
	}
	return a
}

// newLoopingActivation creates a new activation that fires in an asynchronous loop.
func newLoopingActivation(function Action) activation {
	var a activation
	a.function = func(ctx Context) {
		a.executing = true
		go func() {
			function(ctx)
			a.executing = false
		}()
	}
	return a
}

// newClusteredActivation creates a new looping activation that calls Done() on the provided wait group upon completion.
func newClusteredActivation(function Action, wg *sync.WaitGroup) activation {
	var a activation
	a.function = func(ctx Context) {
		a.executing = true
		go func() {
			function(ctx)
			a.executing = false
			wg.Done()
		}()
	}
	return a
}
