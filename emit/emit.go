// Package emit provides access to bit expression from binary types. This process walks a cursor across the binary information
// and selectively yields bits according to the rules defined by logical expressions. Expressions follow Go slice index accessor
// rules, meaning the low side is inclusive and the high side is exclusive.
package emit

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/enum/travel"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/tiny"
)

type Target[T std.Operable] struct {
	target T
}

func From[T std.Operable](target T) Target[T] {
	return Target[T]{
		target: target,
	}
}

// Until keeps reading the provided bit width until the continue function returns false in most→to→least significant order
func (t Target[T]) Until(continueFn std.ContinueFunc, traveling ...travel.Traveling) ([]std.Bit, error) {
	reverse := tiny.ShouldReverseLongitudinally(traveling...)
	return tiny.Emit(std.Expression{
		Continue: &continueFn,
		Reverse:  &reverse,
	}, t.target)
}

// Positions [𝑛₀,𝑛₁,𝑛₂...] creates a std.Expression which will read the provided index positions of your binary information in most→to→least significant order - regardless of the provided variadic order.
func (t Target[T]) Positions(positions []uint, traveling ...travel.Traveling) ([]std.Bit, error) {
	reverse := tiny.ShouldReverseLongitudinally(traveling...)
	return tiny.Emit(std.Expression{
		Positions: &positions,
		Reverse:   &reverse,
	}, t.target)
}

// Width [𝑛] creates a std.Expression which will read the provided bit width in most→to→least significant order.
//
// Expression operations happen in most→to→least significant order - if you would like least←to←most order, please indicate "reverse".
func (t Target[T]) Width(width uint, traveling ...travel.Traveling) ([]std.Bit, error) {
	reverse := tiny.ShouldReverseLongitudinally(traveling...)
	return tiny.Emit(std.Expression{
		Low:     &tiny.Start,
		High:    &width,
		Reverse: &reverse,
	}, t.target)
}

// First [0] creates a std.Expression which will read the first index position of your binary information.
func (t Target[T]) First() ([]std.Bit, error) {
	return tiny.Emit(std.Expression{
		Positions: &tiny.Initial,
	}, t.target)
}

// Last [𝑛 - 1] creates a std.Expression which will read the last index position of your binary information.
func (t Target[T]) Last() ([]std.Bit, error) {
	return tiny.Emit(std.Expression{
		Last: &core.True,
	}, t.target)
}

// Low [low:] creates a std.Expression which will read from the provided index to the end of your binary information.
func (t Target[T]) Low(low uint, traveling ...travel.Traveling) ([]std.Bit, error) {
	reverse := tiny.ShouldReverseLongitudinally(traveling...)
	return tiny.Emit(std.Expression{
		Low:     &low,
		Reverse: &reverse,
	}, t.target)
}

// High [:high] creates a std.Expression which will read to the provided index from the start of your binary information.
func (t Target[T]) High(high uint, traveling ...travel.Traveling) ([]std.Bit, error) {
	reverse := tiny.ShouldReverseLongitudinally(traveling...)
	return tiny.Emit(std.Expression{
		High:    &high,
		Reverse: &reverse,
	}, t.target)
}

// Between [low:high:*] creates a std.Expression which will read between the provided indexes of your binary information up to the provided maximum in most→to→least significant order.
//
// Expression operations happen in most→to→least significant order - if you would like least←to←most order, please indicate "reverse".
func (t Target[T]) Between(low uint, high uint, traveling ...travel.Traveling) ([]std.Bit, error) {
	reverse := tiny.ShouldReverseLongitudinally(traveling...)
	return tiny.Emit(std.Expression{
		Low:     &low,
		High:    &high,
		Reverse: &reverse,
	}, t.target)
}

// All [:] creates a std.Expression which will read the entirety of your binary information.
//
// Expression operations happen in most→to→least significant order - if you would like least←to←most order, please indicate "reverse".
func (t Target[T]) All(traveling ...travel.Traveling) ([]std.Bit, error) {
	reverse := tiny.ShouldReverseLongitudinally(traveling...)
	return tiny.Emit(std.Expression{
		Reverse: &reverse,
	}, t.target)
}

/**
Logic Gates
*/

// Gate creates a std.Expression which will apply the provided logic gate against every input bit.
//
// Expression operations happen in most→to→least significant order - if you would like least←to←most order, please indicate "reverse".
func (t Target[T]) Gate(logic std.BitLogicFunc, traveling ...travel.Traveling) ([]std.Bit, error) {
	reverse := tiny.ShouldReverseLongitudinally(traveling...)
	return tiny.Emit(std.Expression{
		BitLogic: &logic,
		Reverse:  &reverse,
	}, t.target)
}

// NOT creates a std.Expression which will apply the below truth table against every input bit.
//
// NOTE: If no bits are provided, Zero is returned.
//
//	"The NOT Truth Table"
//
//	        𝑎 | 𝑜𝑢𝑡
//	        0 | 1
//	        1 | 0
func (t Target[T]) NOT(traveling ...travel.Traveling) ([]std.Bit, error) {
	return t.Gate(tiny.Logic.NOT, traveling...)
}

/**
Pattern Emission
*/

// Pattern creates a std.Expression which will XOR the provided pattern against the input bits in most→to→least significant order.
func (t Target[T]) Pattern(pattern []std.Bit, traveling ...travel.Traveling) ([]std.Bit, error) {
	return t.Gate(patternLogic(pattern...), traveling...)
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
