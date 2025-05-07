package temporal

func init() {
	u := make(Universe)
	Nexus = &u
}

// Universe is a collection of named worlds.
type Universe map[string]World

// World is a collection of named dimensions.
type World map[string]*Dimension[any, any]

// Nexus is the global universe of all dimensions.
var Nexus *Universe
