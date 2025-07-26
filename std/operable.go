package std

// Operable represents the basic logically operable types.
//
// See Bit, Measurement, Phrase, Natural, Real, Complex, and Index
type Operable interface {
	Bit | byte | Measurement | Phrase | Natural | Real | Complex | Index
}
