package std

import "fmt"

// Bit represents one binary place. [0 - 1]
//
// NOTE: This has a memory footprint of 8 bits.
type Bit byte

// String converts the provided Bit to a string "1", "0", or "-" for Nil - or panics if the found value is anything else.
func (b Bit) String() string {
	switch b {
	case 0:
		return "0"
	case 1:
		return "1"
	case 219:
		return "-"
	default:
		panic(fmt.Errorf("not a bit value: %d", b))
	}
}
