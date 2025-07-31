package measurement

import (
	"fmt"
	"github.com/ignite-laboratories/core/enum/endian"
	"github.com/ignite-laboratories/core/enum/traveling"
	"github.com/ignite-laboratories/core/internal"
	"github.com/ignite-laboratories/core/std"
	"strings"
)

// Measurement is a variable-width slice of bits and is used to efficiently store them in operating memory.
// As Go inherently requires at least 8 bits to store custom types, storing each bit individually
// would need 8 times the size of every bit - thus, the measurement was born.
//
// NOTE: ALL measurements are processed in standard endian.Big form - however, at the time of measurement we
// ALSO capture the original endianness of the stored value.  It's ENTIRELY informational and can be ignored - but
// it's still quite interesting if you care to investigate =)
type Measurement struct {
	// Endianness indicates the endian.Endianness of the data as it was originally stored before being measured in standard endian.Big form.
	//
	// NOTE: Don't get too caught up here - it's purely informational and has absolutely no bearing on tiny operations.
	endian.Endianness

	// Bytes holds complete byte data.
	Bytes []byte

	// Bits holds any remaining bits.
	Bits []std.Bit
}

/**
From Functions
*/

// From creates a new Measurement of the provided Bit slice.
func From(bits ...std.Bit) Measurement {
	std.BitSanityCheck(bits...)
	return Measurement{
		Bits:       bits,
		Endianness: endian.Big,
	}.RollUp()
}

// FromBytes creates a new Measurement of the provided byte slice.
func FromBytes(bytes ...byte) Measurement {
	return Measurement{
		Bytes:      bytes,
		Endianness: endian.Big,
	}.RollUp()
}

// OfZeros creates a new Measurement of the provided bit-width consisting entirely of 0s.
func OfZeros(width int) Measurement {
	return Measurement{
		Bytes:      make([]byte, width/8),
		Bits:       make([]std.Bit, width%8),
		Endianness: endian.Big,
	}.RollUp()
}

// OfOnes creates a new Measurement of the provided bit-width consisting entirely of 1s.
func OfOnes(width int) Measurement {
	zeros := OfZeros(width)
	for i := range zeros.Bytes {
		zeros.Bytes[i] = 255
	}
	for i := range zeros.Bits {
		zeros.Bits[i] = 1
	}
	return zeros.RollUp()
}

// FromPattern creates a new Measurement of the provided bit-width consisting of the pattern emitted across it in the direction.Direction of travel.Traveling.
//
// Inward and outward travel directions are supported and work from the midpoint of the width, biased towards the west.
func FromPattern(w uint, t traveling.Traveling, pattern ...std.Bit) Measurement {
	if w <= 0 || len(pattern) == 0 {
		return Measurement{
			Endianness: endian.Big,
		}
	}

	if t == traveling.Northbound || t == traveling.Southbound {
		panic(fmt.Sprintf("cannot take a latitudinal binary measurement [%v]", t.StringFull(true)))
	}

	printer := func(width uint, tt traveling.Traveling) []std.Bit {
		bits := make([]std.Bit, width)
		patternI := 0
		for i := 0; i < int(width); i++ {
			ii := i
			if tt == traveling.Westbound {
				ii = int(width) - 1 - i
			}

			bits[ii] = pattern[patternI]
			patternI = (patternI + 1) % len(pattern)
		}
		return bits
	}

	if t == traveling.Inbound || t == traveling.Outbound {
		leftWidth := w / 2
		rightWidth := w - leftWidth

		if t == traveling.Inbound {
			left := From(printer(leftWidth, traveling.Eastbound)...)
			right := From(printer(rightWidth, traveling.Westbound)...)
			return left.AppendMeasurements(right)
		}
		return From(printer(leftWidth, traveling.Westbound)...).Append(printer(rightWidth, traveling.Eastbound)...)
	}
	return From(printer(w, t)...)
}

// FromString creates a new Measurement from the provided binary input string.
//
// NOTE: This will panic if anything but a 1 or 0 is found in the input string.
func FromString(s string) Measurement {
	bits := make([]std.Bit, len(s))
	for i := 0; i < len(s); i++ {
		bits[i] = std.Bit(s[i])
	}
	return From(bits...)
}

/**
Methods
*/

// BitWidth gets the total bit width of this Measurement's recorded data.
func (a Measurement) BitWidth() uint {
	return uint((len(a.Bytes) * 8) + len(a.Bits))
}

// GetAllBits returns a slice of the Measurement's individual bits.
func (a Measurement) GetAllBits() []std.Bit {
	a = a.sanityCheck()
	var byteBits []std.Bit
	for _, b := range a.Bytes {
		bits := make([]std.Bit, 8)
		for i := 7; i >= 0; i-- {
			bits[7-i] = std.Bit((b >> i) & 1)
		}
		byteBits = append(byteBits, bits...)
	}
	return append(byteBits, a.Bits...)
}

// Append places the provided bits at the end of the Measurement.
func (a Measurement) Append(bits ...std.Bit) Measurement {
	a = a.sanityCheck(bits...)

	a.Bits = append(a.Bits, bits...)
	return a.RollUp()
}

// AppendBytes places the provided bits at the end of the Measurement.
func (a Measurement) AppendBytes(bytes ...byte) Measurement {
	a = a.sanityCheck()

	lastBits := a.Bits
	for _, b := range bytes {
		bits := make([]std.Bit, 8)

		ii := 0
		for i := byte(7); i < 8; i-- {
			bits[ii] = std.Bit((b >> i) & 1)
			ii++
		}

		blended := append(lastBits, bits[:8-len(lastBits)]...)
		lastBits = bits[8-len(lastBits):]

		var newByte byte
		ii = 0
		for i := byte(7); i < 8; i-- {
			newByte |= byte(blended[ii]) << i
			ii++
		}

		a.Bytes = append(a.Bytes, newByte)
	}

	a.Bits = lastBits
	return a.RollUp()
}

// AppendMeasurements places the provided measurement at the end of the measurement.
func (a Measurement) AppendMeasurements(m ...Measurement) Measurement {
	for _, mmt := range m {
		a = a.Append(mmt.GetAllBits()...)
	}
	return a.RollUp()
}

// Prepend places the provided bits at the start of the Measurement.
func (a Measurement) Prepend(bits ...std.Bit) Measurement {
	a = a.sanityCheck(bits...)

	oldBits := a.Bits
	oldBytes := a.Bytes
	a.Bytes = []byte{}
	a.Bits = []std.Bit{}
	a = a.Append(bits...)
	a = a.AppendBytes(oldBytes...)
	a = a.Append(oldBits...)
	return a.RollUp()
}

// PrependBytes places the provided bytes at the start of the Measurement.
func (a Measurement) PrependBytes(bytes ...byte) Measurement {
	a = a.sanityCheck()

	oldBits := a.Bits
	oldBytes := a.Bytes
	a.Bytes = bytes
	a.Bits = []std.Bit{}
	a = a.AppendBytes(oldBytes...)
	a = a.Append(oldBits...)
	return a.RollUp()
}

// PrependMeasurements places the provided measurement at the start of the measurement.
func (a Measurement) PrependMeasurements(m ...Measurement) Measurement {
	if len(m) == 0 {
		return a
	}

	result := m[len(m)-1]
	for i := len(m) - 2; i >= 0; i-- {
		result = m[i].AppendBytes(result.Bytes...).Append(result.Bits...)
	}
	result = result.AppendBytes(a.Bytes...).Append(a.Bits...)
	return result.RollUp()
}

// Reverse reverses the order of all bits in the measurement.
func (a Measurement) Reverse() Measurement {
	reversedBytes := make([]byte, len(a.Bytes))
	reversedBits := make([]std.Bit, len(a.Bits))

	ii := 0
	for i := len(a.Bytes) - 1; i >= 0; i-- {
		reversedBytes[ii] = internal.ReverseByte(a.Bytes[i])
		ii++
	}

	ii = 0
	for i := len(a.Bits) - 1; i >= 0; i-- {
		reversedBits[ii] = a.Bits[i]
		ii++
	}

	a.Bytes = reversedBytes
	a.Bits = make([]std.Bit, 0)
	return a.Prepend(reversedBits...)
}

// BleedLastBit returns the last bit of the measurement and a measurement missing that bit.
func (a Measurement) BleedLastBit() (std.Bit, Measurement) {
	if a.BitWidth() == 0 {
		panic("cannot bleed the last bit of an empty measurement")
	}
	// TODO: Implement this
	return 0, a
}

// BleedFirstBit returns the first bit of the measurement and a measurement missing that bit.
func (a Measurement) BleedFirstBit() (std.Bit, Measurement) {
	if a.BitWidth() == 0 {
		panic("cannot bleed the first bit of an empty measurement")
	}

	// TODO: Implement this
	return 0, a
}

// RollUp combines the currently measured bits into the measured bytes if there is enough recorded.
func (a Measurement) RollUp() Measurement {
	for len(a.Bits) >= 8 {
		var b byte
		for i := byte(7); i < 8; i-- {
			if a.Bits[i] == 1 {
				b |= 1 << (7 - i)
			}
		}
		a.Bits = a.Bits[8:]
		a.Bytes = append(a.Bytes, b)
	}
	return a
}

/**
Arithmetic
*/

// NonZero returns true if the underlying measurement holds a non-zero value.
func (a Measurement) NonZero() bool {
	for _, b := range a.Bytes {
		if b > 0 {
			return true
		}
	}
	for _, b := range a.Bits {
		if b > 0 {
			return true
		}
	}
	return false
}

/**
Utilities
*/

// sanityCheck ensures the provided bits are all 1s and 0s and rolls the currently measured bits into bytes, if possible.
func (a Measurement) sanityCheck(bits ...std.Bit) Measurement {
	if a.Bytes == nil {
		a.Bytes = []byte{}
	}
	if a.Bits == nil {
		a.Bits = []std.Bit{}
	}
	std.BitSanityCheck(bits...)
	return a.RollUp()
}

// String converts the measurement to a binary string entirely consisting of 1s and 0s.
func (a Measurement) String() string {
	bits := a.GetAllBits()

	builder := strings.Builder{}
	builder.Grow(len(bits))
	for _, b := range bits {
		builder.WriteString(b.String())
	}
	return builder.String()
}

// StringPretty returns a measurement-formatted string of the current binary information. Measurements
// are simply formatted with a single space between digits.
func (a Measurement) StringPretty() string {
	bits := a.GetAllBits()

	if len(bits) == 0 {
		return ""
	}

	builder := strings.Builder{}
	builder.Grow(len(bits)*2 - 1)

	builder.WriteString(bits[0].String())

	if len(bits) > 1 {
		for _, bit := range bits[1:] {
			builder.WriteString(" ")
			builder.WriteString(bit.String())
		}
	}

	return builder.String()
}
