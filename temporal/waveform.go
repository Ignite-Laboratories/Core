package temporal

import "github.com/ignite-laboratories/core"

// Waveform is a type of Dimension that has been constrained to only integer and floating point types.
type Waveform[TValue core.Numeric, TCache any] *Dimension[TValue, TValue]
