package temporal

import "github.com/ignite-laboratories/core/std"

func init() {
	u := make(Universe)
	Nexus = &u
}

// Universe is a collection of named worlds.
type Universe std.FuzzyMap[string, World]

// World is a collection of named dimensional slices.
type World std.FuzzyMap[string, []*Dimension[any, any]]

// Nexus is the global universe of all dimensions.
var Nexus *Universe
