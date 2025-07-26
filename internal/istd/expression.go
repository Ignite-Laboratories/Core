package istd

import "github.com/ignite-laboratories/core/std"

// Expression represents the standard slice index accessor pattern, and expressions can be generated from the global Read variable.
type Expression struct {
	Positions *[]uint
	Low       *uint
	High      *uint
	Last      *bool
	Reverse   *bool
	BitLogic  *std.BitLogicFunc
	Artifact  *std.ArtifactFunc
	Continue  *std.ContinueFunc
	Where     *std.SelectionFunc[any]
	Limit     uint
}
