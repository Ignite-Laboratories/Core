package temporal

import "github.com/ignite-laboratories/core"

// Nexus is the global universe of all worldly dimensions.
var Nexus = make(Universe)

// Universe is a collection of named worlds.
type Universe core.FilterableMap[string, World]

// World is a collection of named dimensional slices and their associated entities.
type World struct {
	Dimensions core.FilterableMap[string, core.FilterableSlice[*Dimension[any, any]]]
	Entities   core.FilterableMap[uint64, *core.Entity]
}
