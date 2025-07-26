package tiny

import (
	"github.com/ignite-laboratories/core/internal"
	"github.com/ignite-laboratories/core/std"
	"reflect"
)

// GetOperableBitWidth returns the bit width of the provided binary operand, or panics if not a std.Operable type.
func GetOperableBitWidth[T any](operands ...T) uint {
	width := uint(0)
	for _, raw := range operands {
		switch operand := any(raw).(type) {
		case std.Phrase:
			width += operand.BitWidth()
		case std.Index:
			width += operand.BitWidth()
		case std.Real:
			width += operand.BitWidth()
		case std.Complex:
			width += operand.Real.BitWidth()
			width += operand.Imaginary.BitWidth()
		case std.Natural:
			width += operand.BitWidth()
		case std.Measurement:
			width += operand.BitWidth()
		case []byte:
			width += uint(len(operand) * 8)
		case []std.Bit:
			width += uint(len(operand))
		case byte:
			width += 8
		case std.Bit:
			width += 1
		default:
			panic("invalid operable type: " + reflect.TypeOf(operand).String())
		}
	}
	return width
}

// SanityCheck ensures the provided bits are all either Zero or One - as Bit is a byte underneath.  In the land of
// binary, that can break all logic without you ever knowing - thus, this intentionally hard panics with ErrorNotABit.
//
// NOTE: This does not account for a 'nil' bit - for that, please use SanityCheckWithNil.
func SanityCheck(bits ...std.Bit) {
	for _, b := range bits {
		if b != Zero && b != One {
			panic(ErrorNotABit)
		}
	}
}

// SanityCheckWithNil ensures the provided bits are all either Zero, One, or Nil - as Bit is a byte underneath.  In the land of
// binary, that can break all logic without you ever knowing - thus, this intentionally hard panics with ErrorNotABit.
//
// NOTE: This accounts for a 'nil' bit - if you wish to work with "traditional" bits, please use SanityCheck.
func SanityCheckWithNil(bits ...std.Bit) {
	for _, b := range bits {
		if b != Zero && b != One && b != Nil {
			panic(ErrorNotABitWithNil)
		}
	}
}

// Measure takes a Measurement of any sized object at runtime.
func Measure[T any](value T) std.Measurement {
	m := std.NewMeasurementOfBytes(internal.Measure(value)[0]...)
	m.Endianness = internal.GetArchitectureEndianness()
	return m
}

// MeasureMany takes measurements of many objects at runtime and returns the result as a single Phrase.
func MeasureMany[T any](values ...T) std.Phrase {
	p := std.NewPhrase()
	for _, v := range values {
		p = p.AppendMeasurement(Measure(v))
	}
	return p
}
