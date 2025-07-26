package itiny

import (
	"fmt"
	"github.com/ignite-laboratories/core/internal"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/tiny"
	"github.com/ignite-laboratories/core/tiny/enum/pad"
	"github.com/ignite-laboratories/core/tiny/enum/traveling"
)

// ShouldReverseLongitudinally indicates if the direction of travel is westerly and the emission should be reversed,
// easterly and remain as is, or panic.
func ShouldReverseLongitudinally(travel ...traveling.Traveling) bool {
	reverse := false
	if len(travel) > 0 {
		t := travel[0]
		switch t {
		case traveling.Westbound:
			reverse = true
		case traveling.Eastbound:
			reverse = false
		case traveling.Inbound, traveling.Outbound:
			panic(fmt.Sprintf("cannot emit in multiple directions [%v]", t.StringFull()))
		case traveling.Northbound, traveling.Southbound:
			panic(fmt.Sprintf("cannot emit latitudinally from a linear binary measurement [%v]", t.StringFull()))
		default:
			panic(fmt.Sprintf("unknown direction of travel [%v]", t))
		}
	}
	return reverse
}

// GetWidestOperand returns the widest bit width of the provided operands.
func GetWidestOperand[T std.Operable](operands ...T) uint {
	var widest uint
	for _, o := range operands {
		width := tiny.GetOperableBitWidth(o)
		if width > widest {
			widest = width
		}
	}
	return widest
}

// AlignOperands applies the provided padding scheme against the operands to align the measured binary information relative to the provided bit width.
//
// You must provide at least one digit to pad the data with, but you may provide a pattern of digits.  The pattern is emitted across the operand starting
// from the West side and working towards the East.  If working latitudinally, the pattern bits are applied longitudinally across each operand in the same way.
//
// NOTE: If you wish for
func AlignOperands[T std.Operable](operands []T, width uint, scheme pad.Scheme, travel traveling.Traveling, digits ...std.Bit) []T {
	// TODO: alignment
	return operands
}

// ReverseOperands reverses the provided input operands.  If they are an Operable type, the internal bits
// are reversed - otherwise, the operands themselves are returned in reverse order.
func ReverseOperands[T any](operands ...T) []T {
	// Put your thing down, flip it, and reverse it
	reversed := make([]T, len(operands))
	limit := len(operands) - 1

	for i, raw := range operands {
		switch operand := any(raw).(type) {
		case std.Real, std.Complex:
			panic(fmt.Errorf("cannot reverse real or complex numbers - please first convert to a phrase"))
		case std.Phrase:
			reversed[limit-i] = any(operand.Reverse()).(T)
		case std.Index:
			reversed[limit-i] = any(operand.Reverse()).(T)
		case std.Natural:
			reversed[limit-i] = any(operand.Reverse()).(T)
		case std.Measurement:
			reversed[limit-i] = any(operand.Reverse()).(T)
		case []byte:
			r := make([]byte, len(operand))
			for ii := len(operand) - 1; ii >= 0; ii-- {
				r[limit-ii] = internal.ReverseByte(operand[ii])
			}
			reversed[limit-i] = any(r).(T)
		case []std.Bit:
			r := make([]std.Bit, len(operand))
			for ii := len(operand) - 1; ii >= 0; ii-- {
				r[limit-ii] = operand[ii]
			}
			reversed[limit-i] = any(r).(T)
		case byte:
			reversed[limit-i] = any(internal.ReverseByte(operand)).(T)
		default:
			reversed[limit-i] = any(operand).(T)
		}
	}

	return reversed
}
