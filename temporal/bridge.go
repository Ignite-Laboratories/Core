package temporal

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/when"
	"time"
)

// Bridge creates a dimension that is driven by a channel.  Because bridges are not driven by
// an impulse engine, they don't have a potential mechanic.  However, their observational trimmers
// still require an engine to function.
func Bridge[TValue any](engine *core.Engine, channel chan TValue) *Dimension[TValue, any] {
	d := Dimension[TValue, any]{}
	d.ID = core.NextID()
	d.Window = core.DefaultWindow
	d.Trimmer = engine.Loop(d.Trim, when.Frequency(&core.TrimFrequency), false)

	var beat int
	var lastMoment time.Time

	go func() {
		for core.Alive && !d.Destroyed {
			value := <-channel
			now := time.Now()

			ctx := core.Context{
				Beat:   beat,
				Moment: now,
				Period: now.Sub(lastMoment),
			}

			data := std.Data[TValue]{
				Context: ctx,
				Point:   value,
			}
			d.Mutex.Lock()
			d.Timeline = append(d.Timeline, data)
			d.Current = &data
			d.Mutex.Unlock()

			beat++
			lastMoment = now
		}
	}()
	return &d
}
