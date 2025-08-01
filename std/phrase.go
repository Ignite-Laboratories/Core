package std

import (
	"github.com/ignite-laboratories/core"
	"strings"
)

// Phrase represents a collection of raw binary measurements at the time of recording.
type Phrase struct {
	Name string
	Data []Measurement
}

/**
New Functions
*/

func NewPhrase(m ...Measurement) Phrase {
	return Phrase{
		Name: core.RandomTinyName().Name,
		Data: m,
	}
}

// NewPhraseNamed creates a named Phrase of the provided measurements and name.
func NewPhraseNamed(name string, m ...Measurement) Phrase {
	return Phrase{
		Name: name,
		Data: m,
	}
}

// NewPhraseNamedFromBits creates a named Phrase of the provided bits and name.
func NewPhraseNamedFromBits(name string, bits ...Bit) Phrase {
	return Phrase{
		Name: name,
		Data: []Measurement{NewMeasurement(bits...)},
	}
}

/**
Methods
*/

// GetData returns the phrase's measurement data.  This is exposed as a method to guarantee
// the encoded accessors for any derived types are grouped together in your IDE's type-ahead.
func (a Phrase) GetData() []Measurement {
	return a.Data
}

// GetAllBits returns a slice of the Phrase's individual bits.
func (a Phrase) GetAllBits() []Bit {
	out := make([]Bit, 0, a.BitWidth())
	for _, m := range a.Data {
		out = append(out, m.GetAllBits()...)
	}
	return out
}

// BitWidth gets the total bit width of this Phrase's recorded data.
func (a Phrase) BitWidth() uint {
	total := uint(0)
	for _, m := range a.Data {
		total += m.BitWidth()
	}
	return uint(total)
}

// Append places the provided bits at the end of the Phrase.
func (a Phrase) Append(bits ...Bit) Phrase {
	if len(a.Data) == 0 {
		a.Data = append(a.Data, NewMeasurement(bits...))
		return a
	}

	last := len(a.Data) - 1
	a.Data[last] = a.Data[last].Append(bits...)
	return a.RollUp()
}

// AppendBytes places the provided bits at the end of the Phrase.
func (a Phrase) AppendBytes(bytes ...byte) Phrase {
	if len(a.Data) == 0 {
		a.Data = append(a.Data, NewMeasurementOfBytes(bytes...))
		return a
	}

	last := len(a.Data) - 1
	a.Data[last] = a.Data[last].AppendBytes(bytes...)
	return a.RollUp()
}

// AppendMeasurement places the provided measurement at the end of the Phrase.
func (a Phrase) AppendMeasurement(m ...Measurement) Phrase {
	a.Data = append(a.Data, m...)
	return a
}

// Prepend places the provided bits at the start of the Phrase.
func (a Phrase) Prepend(bits ...Bit) Phrase {
	if len(a.Data) == 0 {
		a.Data = append(a.Data, NewMeasurement(bits...))
		return a
	}

	a.Data[0] = a.Data[0].Prepend(bits...)
	return a.RollUp()
}

// PrependBytes places the provided bytes at the start of the Phrase.
func (a Phrase) PrependBytes(bytes ...byte) Phrase {
	if len(a.Data) == 0 {
		a.Data = append(a.Data, NewMeasurementOfBytes(bytes...))
		return a
	}

	a.Data[0] = a.Data[0].PrependBytes(bytes...)
	return a.RollUp()
}

// PrependMeasurement places the provided measurement at the start of the Phrase.
func (a Phrase) PrependMeasurement(m ...Measurement) Phrase {
	a.Data = append(m, a.Data...)
	return a
}

// Align ensures all Measurements are of the same width, with the last being smaller if measuring an uneven bit-width.
//
// NOTE: If no width is provided, a standard alignment of 8-bits-per-byte will be used.
//
// NOTE: If you provide a negative width, the phrase will be aligned as a single measurement.
//
// For example -
//
//	let a = an un-aligned logical Phrase
//
//	| 0 1 - 0 1 0 - 0 1 1 0 1 0 0 0 - 1 0 1 1 0 - 0 0 1 0 0 0 0 1 |  ← Raw Bits
//	|  M0 -  M1   -  Measurement 2  -     M3    -  Measurement 4  |  ← Un-aligned Measurements
//
//	a.Align()
//
//	| 0 1 0 1 0 0 1 1 - 0 1 0 0 0 1 0 1 - 1 0 0 0 1 0 0 0 - 0 1 |  ← Raw Bits
//	|  Measurement 0  -  Measurement 1  -  Measurement 2  - M3  |  ← Aligned Measurements
//
//	a.Align(16)
//
//	| 0 1 0 1 0 0 1 1 0 1 0 0 0 1 0 1 - 1 0 0 0 1 0 0 0 0 1 |  ← Raw Bits
//	|          Measurement 0          -    Measurement 1    |  ← Aligned Measurements
func (a Phrase) Align(width ...int) Phrase {
	w := 8
	if len(width) > 0 {
		w = int(width[0])
	}
	if w < 0 {
		w = int(a.BitWidth())
	}

	out := make([]Measurement, 0, a.BitWidth())
	current := make([]Bit, 0, w)
	i := 0

	for _, m := range a.Data {
		for _, b := range m.GetAllBits() {
			current = append(current, b)
			i++
			if i == w {
				i = 0
				out = append(out, NewMeasurement(current...))
				current = make([]Bit, 0, w)
			}
		}
	}

	if len(current) > 0 {
		out = append(out, NewMeasurement(current...))
	}

	a.Data = out
	return a
}

// BleedLastBit returns the last bit of the phrase and a phrase missing that bit.
//
// NOTE: This is a destructive operation to the phrase's encoding scheme and returns a Raw Phrase.
func (a Phrase) BleedLastBit() (Bit, Phrase) {
	if a.BitWidth() == 0 {
		panic("cannot bleed the last bit of an empty phrase")
	}

	lastBit, lastMeasurement := a.Data[len(a.Data)-1].BleedLastBit()
	a.Data[len(a.Data)-1] = lastMeasurement
	return lastBit, a
}

// BleedFirstBit returns the first bit of the phrase and a phrase missing that bit.
//
// NOTE: This is a destructive operation to the phrase's encoding scheme and returns a Raw Phrase.
func (a Phrase) BleedFirstBit() (Bit, Phrase) {
	if a.BitWidth() == 0 {
		panic("cannot bleed the first bit of an empty phrase")
	}

	firstBit, firstMeasurement := a.Data[0].BleedFirstBit()
	a.Data[0] = firstMeasurement
	return firstBit, a
}

// RollUp calls Measurement.RollUp for every measurement in the phrase.
func (a Phrase) RollUp() Phrase {
	for i, m := range a.Data {
		a.Data[i] = m.RollUp()
	}
	return a
}

// Reverse reverses the order of all bits in the phrase.
func (a Phrase) Reverse() Phrase {
	reversed := make([]Measurement, len(a.Data))
	ii := 0
	for i := len(a.Data) - 1; i >= 0; i-- {
		reversed[ii] = a.Data[i].Reverse()
		ii++
	}
	a.Data = reversed
	return a
}

// String returns a string consisting entirely of 1s and 0s.
func (a Phrase) String() string {
	builder := strings.Builder{}
	builder.Grow(int(a.BitWidth()))

	for _, m := range a.Data {
		builder.WriteString(m.String())
	}

	return builder.String()
}

// StringPretty returns a phrase-formatted string of the current measurements. Phrases are formatted as:
//
//  0. Pipes wrap the phrase data.
//  1. Dashes break apart measurements.
//  2. A single space between all characters.
//  3. An empty phrase returns two spaces between pipes.
//
// NOTE: The reason for a two-space empty string is to allow the printing of both outward arrows when displaying index sizes.
func (a Phrase) StringPretty() string {
	if len(a.Data) == 0 {
		return "|  |"
	}

	totalSize := 4 + (len(a.Data)-1)*3
	for _, m := range a.Data {
		totalSize += int(m.BitWidth()) * 2 // ← Approximately account for Measurement's StringPretty() spacing
	}

	builder := strings.Builder{}
	builder.Grow(totalSize)

	builder.WriteString("| ")

	builder.WriteString(a.Data[0].StringPretty())

	for _, m := range a.Data[1:] {
		builder.WriteString(" - ")
		builder.WriteString(m.StringPretty())
	}

	builder.WriteString(" | ")

	return builder.String()
}
