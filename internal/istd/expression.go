package istd

import (
	"github.com/ignite-laboratories/core/std"
)

// Expression represents the standard slice index accessor pattern and handles bit-level expressions.
type Expression struct {
	Positions *[]uint
	Low       *uint
	High      *uint
	Last      *bool
	Reverse   *bool
	BitLogic  *std.BitLogicFunc
	Artifact  *std.ArtifactFunc
	Continue  *std.ContinueFunc
	WhereBit  *std.SelectionFunc[std.Bit]
	Limit     uint
}

// TypedExpression adds typing to an Expression and handles slice-level expressions.
type TypedExpression[T any] struct {
	Expression
	Where *std.SelectionFunc[T]
}

// FluentExpression provides a way to chain expressions sequentially.
type FluentExpression[T any] struct {
	Expressions []TypedExpression[T]
	Reverse     bool
}

func (e FluentExpression[T]) Add(expression TypedExpression[T]) FluentExpression[T] {
	e.Expressions = append(e.Expressions, expression)
	return e
}

// Expressed represents the results of an Expression operation.
type Expressed[T any] struct {
	// Yield represents the collected expression yield.
	Yield []T

	// Error indicates if there was any issue in the expression, but can be generally ignored.
	Error error
}

func NewExpressed[T any](err error, yield ...T) Expressed[T] {
	return Expressed[T]{
		Yield: yield,
		Error: err,
	}
}
