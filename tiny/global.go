package tiny

import (
	"fmt"
	"github.com/ignite-laboratories/core/std"
	"strconv"
)

/**
Global Constants
*/

// EmptyPhrase represents a raw Phrase with no data named "3MP7Y".
var EmptyPhrase = std.NewPhraseNamed("3MP7Y")

// Start is a constantly referencable uint{0}.
//
// For a slice, please use Initial.
//
// See Initial, Zero, One, Nil, ZeroOne, OneZero, SingleZero, DoubleZero, SingleOne, and DoubleOne.
var Start uint = 0

// Initial is a constantly referencable []uint{0}.
//
// For a non-slice, please use Start.
//
// See Start, Zero, One, Nil, ZeroOne, OneZero, SingleZero, DoubleZero, SingleOne, and DoubleOne.
var Initial []uint = []uint{0}

// Zero is an implicit Bit{0}.
//
// See Start, Initial, One, Nil, ZeroOne, OneZero, SingleZero, DoubleZero, SingleOne, and DoubleOne.
const Zero std.Bit = 0

// One is an implicit Bit{1}.
//
// See Start, Initial, Zero, Nil, ZeroOne, OneZero, SingleZero, DoubleZero, SingleOne, and DoubleOne.
const One std.Bit = 1

// OneZero is an implicit Bit{1, 0}.
//
// See Start, Initial, Zero, Nil, ZeroOne, OneZero, SingleZero, DoubleZero, SingleOne, and DoubleOne.
const OneZero std.Bit = 1

// ZeroOne is an implicit Bit{0, 1}.
//
// See Start, Initial, Zero, Nil, ZeroOne, OneZero, SingleZero, DoubleZero, SingleOne, and DoubleOne.
const ZeroOne std.Bit = 1

// SingleZero is an implicit []Bit{0}.
//
// See Start, Initial, Zero, One, Nil, ZeroOne, OneZero, DoubleZero, SingleOne, and DoubleOne.
var SingleZero = []std.Bit{Zero}

// DoubleZero is an implicit []Bit{0, 0}.
//
// See Start, Initial, Zero, One, Nil, ZeroOne, OneZero, SingleZero, SingleOne, and DoubleOne.
var DoubleZero = []std.Bit{Zero, Zero}

// SingleOne is an implicit []Bit{1}.
//
// See Start, Initial, Zero, One, Nil, ZeroOne, OneZero, SingleZero, DoubleZero, and DoubleOne.
var SingleOne = []std.Bit{One}

// DoubleOne is an implicit []Bit{1, 1}.
//
// See Start, Initial, Zero, One, Nil, ZeroOne, OneZero, SingleZero, DoubleZero, and SingleOne.
var DoubleOne = []std.Bit{One, One}

// Nil is an implicit Bit{219} - this allows bits to intentionally be left out of slices and still stand out visibly amongst
// the other bits, as our Bit type is technically a byte in memory.  For example -
//
//	[ 0 0 0 0 0 0 0 0 ]   (0) ← A zero bit
//	[ 0 0 0 0 0 0 0 1 ]   (1) ← A one bit
//	[ 1 1 0 1 1 0 1 1 ] (219) ← A nil bit
//	    ⬑ Darkness is instantly recognizable =)
//
// This also makes logical sense!  If you accidentally overflow or underflow your bit's value by ±1 or ±2, the system won't
// consider it to be in a logically acceptable "nil" state - instead, it -should- panic immediately from a sanity check.
//
// NOTE: Nil is not used in low-level calculations, only in higher level abstractions.
//
// See Start, Initial, Zero, One, ZeroOne, OneZero, SingleZero, DoubleZero, SingleOne, and DoubleOne.
const Nil std.Bit = 219

// WordWidth is the bit width of a standard int, which for all reasonable intents and purposes matches the architecture's word width.
const WordWidth = strconv.IntSize // NOTE: While this could mismatch on exotic hardware, this is just a convenience value.

/**
Errors
*/

// ErrorNotABit is emitted whenever a method expecting a Bit is provided with any other byte value than 1, 0 - as Bit is a byte underneath.
var ErrorNotABit = fmt.Errorf("bits must be 0 or 1 in value")

// ErrorNotABitWithNil is emitted whenever a method expecting a Bit is provided with any other byte value than 1, 0, or 219 (nil) - as Bit is a byte underneath.
var ErrorNotABitWithNil = fmt.Errorf("bits must be 0, 1, or 219 (nil) in value")

// ErrorOutOfData is emitted whenever an emission operation requested more bits than could be emitted.
var ErrorOutOfData = fmt.Errorf("ran out of data")
