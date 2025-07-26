// Package selection provides a way to create logical expressions against slices of objects.
package selection

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/internal/istd"
	"github.com/ignite-laboratories/core/internal/itiny"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/tiny"
	"github.com/ignite-laboratories/core/tiny/enum/traveling"
)

type Target[T any] struct {
	target []T
}

// From starts a fluent expression chain against the provided target.
func From[T any](target ...T) Target[T] {
	return Target[T]{
		target: target,
	}
}

// Until keeps reading your data until the continue function returns false while traveling.Eastbound, unless otherwise specified.
func (t Target[T]) Until(continueFn std.ContinueFunc, traveling ...traveling.Traveling) ([]T, error) {
	reverse := itiny.ShouldReverseLongitudinally(traveling...)
	return itiny.Select(istd.Expression{
		Continue: &continueFn,
		Reverse:  &reverse,
	}, t.target...)
}

// Positions [𝑛₀,𝑛₁,𝑛₂...] creates a std.Expression which will read the provided index positions of your data while traveling.Eastbound, unless otherwise specified.
func (t Target[T]) Positions(positions []uint, traveling ...traveling.Traveling) ([]T, error) {
	reverse := itiny.ShouldReverseLongitudinally(traveling...)
	return itiny.Select(istd.Expression{
		Positions: &positions,
		Reverse:   &reverse,
	}, t.target...)
}

// Width [𝑛] creates a std.Expression which will read the provided bit width of your data while traveling.Eastbound, unless otherwise specified.
func (t Target[T]) Width(width uint, traveling ...traveling.Traveling) ([]T, error) {
	reverse := itiny.ShouldReverseLongitudinally(traveling...)
	return itiny.Select(istd.Expression{
		Low:     &tiny.Start,
		High:    &width,
		Reverse: &reverse,
	}, t.target...)
}

// First [0] creates a std.Expression which will read the first index position of your data.
func (t Target[T]) First() ([]T, error) {
	return itiny.Select(istd.Expression{
		Positions: &tiny.Initial,
	}, t.target...)
}

// Last [𝑛 - 1] creates a std.Expression which will read the last index position of your data.
func (t Target[T]) Last() ([]T, error) {
	return itiny.Select(istd.Expression{
		Last: &core.True,
	}, t.target...)
}

// Low [low:] creates a std.Expression which will read from the provided index to the end of your data while traveling.Eastbound, unless otherwise specified.
func (t Target[T]) Low(low uint, traveling ...traveling.Traveling) ([]T, error) {
	reverse := itiny.ShouldReverseLongitudinally(traveling...)
	return itiny.Select(istd.Expression{
		Low:     &low,
		Reverse: &reverse,
	}, t.target...)
}

// High [:high] creates a std.Expression which will read to the provided index from the start of your data while traveling.Eastbound, unless otherwise specified.
func (t Target[T]) High(high uint, traveling ...traveling.Traveling) ([]T, error) {
	reverse := itiny.ShouldReverseLongitudinally(traveling...)
	return itiny.Select(istd.Expression{
		High:    &high,
		Reverse: &reverse,
	}, t.target...)
}

// Between [low:high:*] creates a std.Expression which will read between the provided indexes of your data up to the provided maximum while traveling.Eastbound, unless otherwise specified.
func (t Target[T]) Between(low uint, high uint, traveling ...traveling.Traveling) ([]T, error) {
	reverse := itiny.ShouldReverseLongitudinally(traveling...)
	return itiny.Select(istd.Expression{
		Low:     &low,
		High:    &high,
		Reverse: &reverse,
	}, t.target...)
}

// All [:] creates a std.Expression which will read the entirety of your data while traveling.Eastbound, unless otherwise specified.
func (t Target[T]) All(traveling ...traveling.Traveling) ([]T, error) {
	reverse := itiny.ShouldReverseLongitudinally(traveling...)
	return itiny.Select(istd.Expression{
		Reverse: &reverse,
	}, t.target...)
}

/**
Querying
*/

// Where creates a std.Expression which will call the provided predicate for each entry of your data while traveling.Eastbound, unless otherwise specified.
func (t Target[T]) Where(predicate std.SelectionFunc[T], traveling ...traveling.Traveling) ([]T, error) {
	reverse := itiny.ShouldReverseLongitudinally(traveling...)
	p := any(predicate).(std.SelectionFunc[any])
	return itiny.Select(istd.Expression{
		Where:   &p,
		Reverse: &reverse,
	}, t.target...)
}
