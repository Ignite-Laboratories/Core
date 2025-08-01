package phrase

import (
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/measurement"
	"github.com/ignite-laboratories/core/std/name"
)

func newPhrase[T any]() std.Phrase[T] {
	return std.Phrase[T]{
		Name: name.Tiny().Name,
		Data: make([]std.Measurement[T], 0),
	}
}

/**
Creation
*/

func From[T any](data ...T) std.Phrase[T] {
	p := newPhrase[T]()
	for _, d := range data {
		p = p.AppendMeasurement(measurement.From(d))
	}
	return p
}

func FromMeasurements[T any](m ...std.Measurement[T]) std.Phrase[T] {
	p := newPhrase[T]()
	p.Data = m
	return p
}

// FromBits creates a named Phrase of the provided bits and name.
func FromBits(name string, bits ...std.Bit) std.Phrase[any] {
	p := newPhrase[any]()
	p.Name = name
	p.Data = []std.Measurement[any]{measurement.FromBits(bits...)}
	return p
}
