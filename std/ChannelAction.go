package std

import (
	"github.com/ignite-laboratories/core"
	"sync"
)

// ChannelAction provides either the impulse context or an action to perform.
//
// It provides some common mechanisms for such a design.
type ChannelAction struct {
	Context   core.Context
	Action    func()
	WaitGroup *sync.WaitGroup
}
