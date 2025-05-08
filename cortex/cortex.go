package cortex

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/temporal"
)

type Cortex struct {
	core.Engine
	temporal.Universe
}

func (c Cortex) Spark() {
	c.Engine.Spark()
}
