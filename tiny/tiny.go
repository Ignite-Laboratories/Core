package tiny

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/enum/pad"
	"github.com/ignite-laboratories/core/enum/travel"
	"github.com/ignite-laboratories/core/std"
	"reflect"
	"unsafe"
)

// GetRandomName returns a randomly generated name which conforms to the NameFilter rules.
func GetRandomName() string {
	return core.RandomNameFiltered(NameFilter).Name
}

// GetBitWidth returns the bit width of the provided binary operand.
func GetBitWidth[T std.Operable](operands ...T) uint {
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
			panic("invalid binary type: " + reflect.TypeOf(operand).String())
		}
	}
	return width
}

// BleedEnd returns the ending bits of the operands and the operands missing those bits.
//
// All bleed operations are always returned in their original most→to→least significant order.
func BleedEnd[T std.Operable](width uint, operands ...T) ([][]std.Bit, []T) {
	bits := make([][]std.Bit, 0, len(operands))

	for x := 0; x < int(width); x++ {
		for i, raw := range operands {
			if GetBitWidth(raw) == 0 {
				continue
			}

			switch operand := any(raw).(type) {
			case std.Phrase, std.Complex, std.Index, std.Real, std.Natural:
				// TODO: Implement this
			case std.Measurement:
				var bit std.Bit
				bit, operand = operand.BleedLastBit()
				bits[i] = append([]std.Bit{bit}, bits[i]...)
				operands[i] = any(operand).(T)
			case []byte:
				panic("cannot bleed bits from static width elements")
			case []std.Bit:
				bits[i] = append([]std.Bit{operand[len(operand)-1]}, bits[i]...)
				operands[i] = any(operand[:len(operand)-1]).(T)
			case byte:
				panic("cannot bleed bits from static width elements")
			case std.Bit:
				panic("cannot bleed bits from static width elements")
			default:
				panic("invalid binary type: " + reflect.TypeOf(operand).String())
			}
		}
	}
	return bits, operands
}

// BleedStart returns the first bit of the operands and the operands missing that bit.
//
// All bleed operations are always returned in their original most→to→least significant order.
func BleedStart[T std.Operable](width uint, operands ...T) ([][]std.Bit, []T) {
	bits := make([][]std.Bit, 0, len(operands))

	for x := 0; x < int(width); x++ {
		for i, raw := range operands {
			if GetBitWidth(raw) == 0 {
				continue
			}

			switch operand := any(raw).(type) {
			case std.Phrase, std.Complex, std.Index, std.Real, std.Natural:
				// TODO: Implement this
			case std.Measurement:
				var bit std.Bit
				bit, operand = operand.BleedFirstBit()
				bits[i] = append([]std.Bit{bit}, bits[i]...)
				operands[i] = any(operand).(T)
			case []byte:
				panic("cannot bleed bits from static width elements")
			case []std.Bit:
				bits[i] = append([]std.Bit{operand[0]}, bits[i]...)
				operands[i] = any(operand[1:]).(T)
			case byte:
				panic("cannot bleed bits from static width elements")
			case std.Bit:
				panic("cannot bleed bits from static width elements")
			default:
				panic("invalid binary type: " + reflect.TypeOf(operand).String())
			}
		}
	}
	return bits, operands
}

// GetWidestOperand returns the widest bit width of the provided operands.
func GetWidestOperand[T std.Operable](operands ...T) uint {
	var widest uint
	for _, o := range operands {
		width := GetBitWidth(o)
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
func AlignOperands[T std.Operable](operands []T, width uint, scheme pad.Scheme, traveling travel.Traveling, digits ...std.Bit) []T {
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
				r[limit-ii] = core.ReverseByte(operand[ii])
			}
			reversed[limit-i] = any(r).(T)
		case []std.Bit:
			r := make([]std.Bit, len(operand))
			for ii := len(operand) - 1; ii >= 0; ii-- {
				r[limit-ii] = operand[ii]
			}
			reversed[limit-i] = any(r).(T)
		case byte:
			reversed[limit-i] = any(core.ReverseByte(operand)).(T)
		default:
			reversed[limit-i] = any(operand).(T)
		}
	}

	return reversed
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
	m := std.NewMeasurementOfBytes(core.Measure(value)[0]...)
	m.Endianness = core.GetArchitectureEndianness()
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

// ToType converts a Measurement of binary information into the specified type T.
func ToType[T any](m std.Measurement) T {
	bits := m.GetAllBits()
	var zero T
	typeOf := reflect.TypeOf(zero)

	// Handle slices
	if typeOf.Kind() == reflect.Slice {
		elemType := typeOf.Elem()
		elemSize := elemType.Size()

		numElements := len(bits) / (8 * int(elemSize))
		if numElements == 0 {
			return zero
		}

		sliceVal := reflect.MakeSlice(typeOf, numElements, numElements)
		slicePtr := unsafe.Pointer(sliceVal.UnsafePointer())
		resultBytes := unsafe.Slice((*byte)(slicePtr), numElements*int(elemSize))

		byteI := (len(bits) / 8) - 1
		i := len(bits) - 1
		for i > 0 {
			var currentByte byte
			for ii := 0; ii < 8; ii++ {
				if bits[i] == 1 {
					currentByte |= 1 << ii
				}
				i--
			}

			resultBytes[byteI] = currentByte
			byteI--
		}

		return sliceVal.Interface().(T)
	}

	// Handle non-slices
	size := typeOf.Size()
	if len(bits) > int(size)*8 {
		panic("bit slice too large for target type")
	}

	result := zero
	resultPtr := unsafe.Pointer(&result)
	resultBytes := unsafe.Slice((*byte)(resultPtr), size)

	byteI := (len(bits) / 8) - 1
	i := len(bits) - 1
	for i > 0 {
		var currentByte byte
		for ii := 0; ii < 8; ii++ {
			if bits[i] == 1 {
				currentByte |= 1 << ii
			}
			i--
		}

		resultBytes[byteI] = currentByte
		byteI--
	}

	return result
}

// ShouldReverseLongitudinally indicates if the direction of travel is westerly and the emission should be reversed, otherwise it panics.
//
// NOTE: This is entirely a convenience function for emission passthrough
func ShouldReverseLongitudinally(traveling ...travel.Traveling) bool {
	reverse := false
	if len(traveling) > 0 {
		t := traveling[0]
		switch t {
		case travel.Westbound:
			reverse = true
		case travel.Eastbound:
			reverse = false
		case travel.Inbound, travel.Outbound:
			panic(fmt.Sprintf("cannot emit in multiple directions [%v]", t.StringFull()))
		case travel.Northbound, travel.Southbound:
			panic(fmt.Sprintf("cannot emit latitudinally from a linear binary measurement [%v]", t.StringFull()))
		default:
			panic(fmt.Sprintf("unknown direction of travel [%v]", t))
		}
	}
	return reverse
}
