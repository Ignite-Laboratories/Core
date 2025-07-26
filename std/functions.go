package std

/**
Tiny functions
*/

// BitLogicFunc takes in many bits and their collectively shared index and returns an output bit plus a nilable artifact.
type BitLogicFunc func(uint, ...Bit) ([]Bit, *Phrase)

// ArtifactFunc applies the artifact from a single round of calculation against the provided operand bits.
type ArtifactFunc func(i uint, artifact Phrase, operands ...Phrase) []Phrase

// ContinueFunc is called after every Bit is read with the currently read bits - if it returns false, the emission terminates traversal.
type ContinueFunc func(i uint, data []Bit) bool

// SelectionFunc acts as a selection predicate for selection queries.
type SelectionFunc[T any] func(T) bool
