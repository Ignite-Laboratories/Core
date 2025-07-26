// Package emit provides access to bit expression at runtime. This process walks a cursor across the binary information
// and selectively yields bits according to the rules defined by logical expressions. Expressions follow Go slice index
// accessor rules, meaning the low side is inclusive and the high side is exclusive.
package emit

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/enum/traveling"
	"github.com/ignite-laboratories/core/internal/istd"
	"github.com/ignite-laboratories/core/internal/itiny"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/tiny"
)

type Targets[T any] struct {
	targets []T
}

// Pattern creates a std.Expression which will XOR the provided pattern against the input bits while traveling.Eastbound, unless otherwise specified.
func Pattern(width uint, travel traveling.Traveling, pattern ...std.Bit) std.Measurement {
	return std.NewMeasurementOfPattern(width, travel, pattern...)
}

// From starts a fluent expression chain against the provided targets.
//
// NOTE: If providing non-std.Operable types, this will first take individual measurements of each operand.
func From[T any](targets ...T) Targets[T] {
	return Targets[T]{
		targets: targets,
	}
}

// Until keeps reading your binary information until the continue function returns false while traveling.Eastbound, unless otherwise specified.
func (t Targets[T]) Until(continueFn std.ContinueFunc, travel ...traveling.Traveling) (std.Measurement, error) {
	reverse := itiny.ShouldReverseLongitudinally(travel...)
	return itiny.Emit(istd.Expression{
		Continue: &continueFn,
		Reverse:  &reverse,
	}, t.targets...)
}

// Positions [𝑛₀,𝑛₁,𝑛₂...] creates a std.Expression which will read the provided index positions of your binary information while traveling.Eastbound, unless otherwise specified.
func (t Targets[T]) Positions(positions []uint, travel ...traveling.Traveling) (std.Measurement, error) {
	reverse := itiny.ShouldReverseLongitudinally(travel...)
	return itiny.Emit(istd.Expression{
		Positions: &positions,
		Reverse:   &reverse,
	}, t.targets...)
}

// Width [𝑛] creates a std.Expression which will read the provided bit width while traveling.Eastbound, unless otherwise specified.
func (t Targets[T]) Width(width uint, travel ...traveling.Traveling) (std.Measurement, error) {
	reverse := itiny.ShouldReverseLongitudinally(travel...)
	return itiny.Emit(istd.Expression{
		Low:     &tiny.Start,
		High:    &width,
		Reverse: &reverse,
	}, t.targets...)
}

// First [0] creates a std.Expression which will read the first index position of your binary information.
func (t Targets[T]) First() (std.Measurement, error) {
	return itiny.Emit(istd.Expression{
		Positions: &tiny.Initial,
	}, t.targets...)
}

// Last [𝑛 - 1] creates a std.Expression which will read the last index position of your binary information.
func (t Targets[T]) Last() (std.Measurement, error) {
	return itiny.Emit(istd.Expression{
		Last: &core.True,
	}, t.targets...)
}

// Low [low:] creates a std.Expression which will read from the provided index to the end of your binary information while traveling.Eastbound, unless otherwise specified.
func (t Targets[T]) Low(low uint, travel ...traveling.Traveling) (std.Measurement, error) {
	reverse := itiny.ShouldReverseLongitudinally(travel...)
	return itiny.Emit(istd.Expression{
		Low:     &low,
		Reverse: &reverse,
	}, t.targets...)
}

// High [:high] creates a std.Expression which will read to the provided index from the start of your binary information while traveling.Eastbound, unless otherwise specified.
func (t Targets[T]) High(high uint, travel ...traveling.Traveling) (std.Measurement, error) {
	reverse := itiny.ShouldReverseLongitudinally(travel...)
	return itiny.Emit(istd.Expression{
		High:    &high,
		Reverse: &reverse,
	}, t.targets...)
}

// Between [low:high:*] creates a std.Expression which will read between the provided indexes of your binary information up to the provided maximum while traveling.Eastbound, unless otherwise specified.
func (t Targets[T]) Between(low uint, high uint, travel ...traveling.Traveling) (std.Measurement, error) {
	reverse := itiny.ShouldReverseLongitudinally(travel...)
	return itiny.Emit(istd.Expression{
		Low:     &low,
		High:    &high,
		Reverse: &reverse,
	}, t.targets...)
}

// All [:] creates a std.Expression which will read the entirety of your binary information while traveling.Eastbound, unless otherwise specified.
func (t Targets[T]) All(travel ...traveling.Traveling) (std.Measurement, error) {
	reverse := itiny.ShouldReverseLongitudinally(travel...)
	return itiny.Emit(istd.Expression{
		Reverse: &reverse,
	}, t.targets...)
}

/**
Logic Gates
*/

// Gate creates a std.Expression which will apply the provided logic gate against every input bit while traveling.Eastbound, unless otherwise specified.
func (t Targets[T]) Gate(logic std.BitLogicFunc, travel ...traveling.Traveling) (std.Measurement, error) {
	reverse := itiny.ShouldReverseLongitudinally(travel...)
	return itiny.Emit(istd.Expression{
		BitLogic: &logic,
		Reverse:  &reverse,
	}, t.targets...)
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
func (t Targets[T]) NOT(travel ...traveling.Traveling) (std.Measurement, error) {
	return t.Gate(tiny.Logic.NOT, travel...)
}
