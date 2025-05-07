package temporal

import "github.com/ignite-laboratories/core"

func init() {
	u := make(Universe)
	Nexus = &u
}

// Nexus is the global universe of all worldly dimensions.
//
// Abstractly, this represents: *Universe[string, World[string, []*Dimension[any, any]]]
var Nexus *Universe

// Universe is a collection of named worlds.
type Universe core.FilterableMap[string, World]

// World is a collection of named dimensional slices.
type World core.FilterableMap[string, core.FilterableSlice[*Dimension[any, any]]]
