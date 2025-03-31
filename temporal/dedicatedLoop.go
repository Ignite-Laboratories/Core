package temporal

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/when"
	"runtime"
)

// DedicatedLoop is a special kind of Loop - it guarantees the target action is always
// called from the same host thread using runtime.LockOSThread()
func DedicatedLoop(engine *core.Engine, potential core.Potential, muted bool, target core.Action) *Dimension[core.Runtime, chan core.Context] {
	d := Dimension[core.Runtime, chan core.Context]{}
	d.ID = core.NextID()
	d.Window = core.DefaultWindow
	d.Trimmer = engine.Loop(d.Trim, when.Always, false)
	c := make(chan core.Context)
	d.Cache = &c
	f := func(ctx core.Context) {
		data := std.Data[core.Runtime]{
			Context: ctx,
			Point:   d.Stimulator.Last,
		}
		*d.Cache <- ctx
		d.Mutex.Lock()
		d.Timeline = append(d.Timeline, data)
		d.Current = &data
		d.Mutex.Unlock()
	}

	go func() {
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()
		for ctx := range *d.Cache {
			target(ctx)
		}
	}()

	d.Stimulator = engine.Stimulate(f, potential, muted)
	return &d
}
