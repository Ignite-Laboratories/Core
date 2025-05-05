package core

import (
	"container/list"
	"sync"
)

// Carousel represetns a FIFO collection of Go routines which can be used to execute code.
//
// The number of Go routines is based upon demand - if none are available when a request is made,
// a new Go routine is spawned.  The TTL value indicates the number of activations a single routine
// should cycle for before destruction - allowing high-stress moments to build more routines, then
// dissipate as the load diminishes.
type Carousel struct {
	list.List

	// TTL represents the number of cycles the Go routines should activate for before destruction.
	TTL int

	mutex sync.Mutex
}

type routine struct {
	TTL     int
	Input   chan func()
	Running bool
	Alive   bool
}

func NewCarousel(ttl int) *Carousel {
	return &Carousel{
		TTL: ttl,
	}
}

func (c *Carousel) getNextAvailable() *routine {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// Look for an available routine...
	if front := c.Front(); front != nil {
		if value, ok := front.Value.(*routine); ok {
			// Move the found routine to the back of the carousel
			c.Remove(front)
			if value.Alive {
				c.PushBack(value)
				if !value.Running {
					// If it's not running, return it
					return value
				}
			}
		}
		// NOTE: The logic is that if the routine at the front of the queue is busy, we need
		// more routines anyway.  Thus, there is no reason to do a more exhaustive search.
	}
	// ...if one wasn't found, spawn a new one
	r := routine{
		TTL:   c.TTL,
		Input: make(chan func()),
		Alive: true,
	}
	go r.Run()
	c.PushBack(r)
	return &r
}

// Step takes the provided action and calls it on an existing Go routine.
//
// If none is currently available, a new Go routine will be spawned.
func (c *Carousel) Step(action func()) {
	r := c.getNextAvailable()
	r.Input <- action
}

// Run starts listening for actions, decrementing the spawned TTL value with each message until 0 before "dying."
func (r *routine) Run() {
	for Alive && r.TTL > 0 {
		action := <-r.Input
		r.Running = true
		action()
		r.Running = false
		r.TTL--
	}
	r.Alive = false
}
