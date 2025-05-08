package temporal

import "github.com/ignite-laboratories/core"

// Nexus is the global universe of all worldly dimensions.
//
// Abstractly, this represents: *Universe[string, World[string, []*Dimension[any, any]]]
var Nexus = &Universe{
	Worlds:   make(core.FilterableMap[string, World]),
	Entities: make(core.FilterableMap[uint64, *core.Entity]),
}

// Universe is a collection of named worlds and their associated entities.
type Universe struct {
	Worlds   core.FilterableMap[string, World]
	Entities core.FilterableMap[uint64, *core.Entity]
}

// World is a collection of named dimensional slices.
type World core.FilterableMap[string, core.FilterableSlice[*Dimension[any, any]]]
