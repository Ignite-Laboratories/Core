package std

import (
	"github.com/ignite-laboratories/core"
)

// ChannelAction provides either the impulse context or an action to perform.
type ChannelAction struct {
	Context core.Context
	Action  func()
}
