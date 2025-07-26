// Package emit provides access to bit expression from binary types. This process walks a cursor across the binary information
// and selectively yields bits according to the rules defined by logical expressions. Expressions follow Go slice index accessor
// rules, meaning the low side is inclusive and the high side is exclusive.
package emit

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/enum/traveling"
	"github.com/ignite-laboratories/core/internal/istd"
	"github.com/ignite-laboratories/core/internal/itiny"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/tiny"
)

type Target[T std.Operable] struct {
	target []T
}

// From starts a fluent expression chain against the provided target.
func From[T std.Operable](target ...T) Target[T] {
	return Target[T]{
		target: target,
	}
}

// FromAny takes a measurement of the provided target, then starts a fluent expression chain against the provided target.
func FromAny[T any](target T) Target[std.Measurement] {
	return From(tiny.Measure[T](target))
}

// Until keeps reading your binary information until the continue function returns false while traveling.Eastbound, unless otherwise specified.
func (t Target[T]) Until(continueFn std.ContinueFunc, travel ...traveling.Traveling) ([]std.Bit, error) {
	reverse := tiny.ShouldReverseLongitudinally(travel...)
	return itiny.Emit(istd.Expression{
		Continue: &continueFn,
		Reverse:  &reverse,
	}, t.target...)
}

// Positions [𝑛₀,𝑛₁,𝑛₂...] creates a std.Expression which will read the provided index positions of your binary information while traveling.Eastbound, unless otherwise specified.
func (t Target[T]) Positions(positions []uint, travel ...traveling.Traveling) ([]std.Bit, error) {
	reverse := tiny.ShouldReverseLongitudinally(travel...)
	return itiny.Emit(istd.Expression{
		Positions: &positions,
		Reverse:   &reverse,
	}, t.target...)
}

// Width [𝑛] creates a std.Expression which will read the provided bit width while traveling.Eastbound, unless otherwise specified.
func (t Target[T]) Width(width uint, travel ...traveling.Traveling) ([]std.Bit, error) {
	reverse := tiny.ShouldReverseLongitudinally(travel...)
	return itiny.Emit(istd.Expression{
		Low:     &tiny.Start,
		High:    &width,
		Reverse: &reverse,
	}, t.target...)
}

// First [0] creates a std.Expression which will read the first index position of your binary information.
func (t Target[T]) First() ([]std.Bit, error) {
	return itiny.Emit(istd.Expression{
		Positions: &tiny.Initial,
	}, t.target...)
}

// Last [𝑛 - 1] creates a std.Expression which will read the last index position of your binary information.
func (t Target[T]) Last() ([]std.Bit, error) {
	return itiny.Emit(istd.Expression{
		Last: &core.True,
	}, t.target...)
}

// Low [low:] creates a std.Expression which will read from the provided index to the end of your binary information while traveling.Eastbound, unless otherwise specified.
func (t Target[T]) Low(low uint, travel ...traveling.Traveling) ([]std.Bit, error) {
	reverse := tiny.ShouldReverseLongitudinally(travel...)
	return itiny.Emit(istd.Expression{
		Low:     &low,
		Reverse: &reverse,
	}, t.target...)
}

// High [:high] creates a std.Expression which will read to the provided index from the start of your binary information while traveling.Eastbound, unless otherwise specified.
func (t Target[T]) High(high uint, travel ...traveling.Traveling) ([]std.Bit, error) {
	reverse := tiny.ShouldReverseLongitudinally(travel...)
	return itiny.Emit(istd.Expression{
		High:    &high,
		Reverse: &reverse,
	}, t.target...)
}

// Between [low:high:*] creates a std.Expression which will read between the provided indexes of your binary information up to the provided maximum while traveling.Eastbound, unless otherwise specified.
func (t Target[T]) Between(low uint, high uint, travel ...traveling.Traveling) ([]std.Bit, error) {
	reverse := tiny.ShouldReverseLongitudinally(travel...)
	return itiny.Emit(istd.Expression{
		Low:     &low,
		High:    &high,
		Reverse: &reverse,
	}, t.target...)
}

// All [:] creates a std.Expression which will read the entirety of your binary information while traveling.Eastbound, unless otherwise specified.
func (t Target[T]) All(travel ...traveling.Traveling) ([]std.Bit, error) {
	reverse := tiny.ShouldReverseLongitudinally(travel...)
	return itiny.Emit(istd.Expression{
		Reverse: &reverse,
	}, t.target...)
}

/**
Logic Gates
*/

// Gate creates a std.Expression which will apply the provided logic gate against every input bit while traveling.Eastbound, unless otherwise specified.
func (t Target[T]) Gate(logic std.BitLogicFunc, travel ...traveling.Traveling) ([]std.Bit, error) {
	reverse := tiny.ShouldReverseLongitudinally(travel...)
	return itiny.Emit(istd.Expression{
		BitLogic: &logic,
		Reverse:  &reverse,
	}, t.target...)
}

// NOT creates a std.Expression which will apply the below truth table against every input bit while traveling.Eastbound, unless otherwise specified.
//
// NOTE: If no bits are provided, Zero is returned.
//
//	"The NOT Truth Table"
//
//	        𝑎 | 𝑜𝑢𝑡
//	        0 | 1
//	        1 | 0
func (t Target[T]) NOT(travel ...traveling.Traveling) ([]std.Bit, error) {
	return t.Gate(tiny.Logic.NOT, travel...)
}

/**
Pattern Emission
*/

// Pattern creates a std.Expression which will XOR the provided pattern against the input bits while traveling.Eastbound, unless otherwise specified.
func (t Target[T]) Pattern(pattern []std.Bit, travel ...traveling.Traveling) ([]std.Bit, error) {
	return t.Gate(patternLogic(pattern...), travel...)
}

func patternLogic(pattern ...std.Bit) std.BitLogicFunc {
	limit := len(pattern)
	step := 0
	return func(i uint, operands ...std.Bit) ([]std.Bit, *std.Phrase) {
		for _, b := range pattern {
			operands[i] = b ^ pattern[i]
		}
		step++
		if step >= limit {
			step = 0
		}
		return operands, nil
	}
}
