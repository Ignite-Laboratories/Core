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

type ExpressionTarget[T any] struct {
	Expression istd.FluentExpression[T]
	target     []T
}

// Express expresses the expression.
func (t ExpressionTarget[T]) Express(traveling ...traveling.Traveling) istd.Expressed[T] {
	t.Expression.Reverse = itiny.ShouldReverseLinearData(traveling...)
	return itiny.Select(t.Expression, t.target...)
}

// From starts a fluent expression chain against the provided target.
func From[T any](target ...T) ExpressionTarget[T] {
	return ExpressionTarget[T]{
		Expression: istd.FluentExpression[T]{},
		target:     target,
	}
}

// FromWhile calls the provided function until it returns false while collecting the returned values.
func FromWhile[T any](while func(uint) (T, bool)) ExpressionTarget[T] {
	target := make([]T, 0)
	i := uint(0)
	for element, keepGoing := while(i); keepGoing; element, keepGoing = while(i) {
		target = append(target, element)
		i++
	}
	return From(target...)
}

// While keeps reading your data until the continue function returns false.
func (t ExpressionTarget[T]) While(continueFn std.ContinueFunc) ExpressionTarget[T] {
	t.Expression = t.Expression.Add(istd.TypedExpression[T]{
		Expression: istd.Expression{
			Continue: &continueFn,
		},
	})
	return t
}

// Positions [𝑛₀,𝑛₁,𝑛₂...] creates a std.Expression which will read the provided index positions of your data.
func (t ExpressionTarget[T]) Positions(positions []uint) ExpressionTarget[T] {
	t.Expression = t.Expression.Add(istd.TypedExpression[T]{
		Expression: istd.Expression{
			Positions: &positions,
		},
	})
	return t
}

// Take [𝑛] creates a std.Expression which will read the provided number of elements from your data.
func (t ExpressionTarget[T]) Take(count uint) ExpressionTarget[T] {
	t.Expression = t.Expression.Add(istd.TypedExpression[T]{
		Expression: istd.Expression{
			Low:  &tiny.Start,
			High: &count,
		},
	})
	return t
}

// First [0] creates a std.Expression which will read the first index position of your data.
func (t ExpressionTarget[T]) First() ExpressionTarget[T] {
	t.Expression = t.Expression.Add(istd.TypedExpression[T]{
		Expression: istd.Expression{
			Positions: &tiny.Initial,
		},
	})
	return t
}

// Last [𝑛 - 1] creates a std.Expression which will read the last index position of your data.
func (t ExpressionTarget[T]) Last() ExpressionTarget[T] {
	t.Expression = t.Expression.Add(istd.TypedExpression[T]{
		Expression: istd.Expression{
			Last: &core.True,
		},
	})
	return t
}

// Low [low:] creates a std.Expression which will read from the provided index to the end of your data.
func (t ExpressionTarget[T]) Low(low uint) ExpressionTarget[T] {
	t.Expression = t.Expression.Add(istd.TypedExpression[T]{
		Expression: istd.Expression{
			Low: &low,
		},
	})
	return t
}

// High [:high] creates a std.Expression which will read to the provided index from the start of your data.
func (t ExpressionTarget[T]) High(high uint) ExpressionTarget[T] {
	t.Expression = t.Expression.Add(istd.TypedExpression[T]{
		Expression: istd.Expression{
			High: &high,
		},
	})
	return t
}

// Between [low:high:*] creates a std.Expression which will read between the provided indexes of your data.
func (t ExpressionTarget[T]) Between(low uint, high uint) ExpressionTarget[T] {
	t.Expression = t.Expression.Add(istd.TypedExpression[T]{
		Expression: istd.Expression{
			Low:  &low,
			High: &high,
		},
	})
	return t
}

// All [:] creates a std.Expression which will read the entirety of your data.
func (t ExpressionTarget[T]) All() ExpressionTarget[T] {
	t.Expression = t.Expression.Add(istd.TypedExpression[T]{
		Expression: istd.Expression{},
	})
	return t
}

/**
Querying
*/

// Where creates a std.Expression which will call the provided predicate for each element of your data and return only the elements in which the predicate returned true.
func (t ExpressionTarget[T]) Where(predicate std.SelectionFunc[T]) ExpressionTarget[T] {
	t.Expression = t.Expression.Add(istd.TypedExpression[T]{
		Expression: istd.Expression{},
		Where:      predicate,
	})
	return t
}
