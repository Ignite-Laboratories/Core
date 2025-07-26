package tiny

import (
	"github.com/ignite-laboratories/core/internal"
	"github.com/ignite-laboratories/core/std"
	"reflect"
	"unsafe"
)

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
