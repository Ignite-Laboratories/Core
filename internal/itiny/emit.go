package itiny

import (
	"fmt"
	"github.com/ignite-laboratories/core/internal/istd"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/tiny"
)

// Select expresses the input operands according to the provided logical expression.
//
// NOTE: This will return a tiny.ErrorOutOfData if the expression either couldn't be satisfied or no input data was provided.
func Select[T any](fluent istd.FluentExpression[T], operands ...T) istd.Expressed[T] {
	totalWidth := uint(len(operands))
	if totalWidth == 0 {
		return istd.NewExpressed[T](tiny.ErrorOutOfData)
	}

	for _, expr := range fluent.Expressions {
		// Evaluate and set the appropriate expression boundary values
		expr.Expression, operands = evaluateBaseExpression(expr.Expression, totalWidth, operands...)

		var yield []T

		if expr.Positions != nil {
			yield = make([]T, expr.Limit)

			for i, pos := range *expr.Positions {
				yield[i] = operands[pos]
			}
			return istd.NewExpressed[T](nil, yield...)
		}
		if expr.Where != nil {
			yield = make([]T, 0, expr.Limit)
			for _, operand := range operands {
				if expr.Where(operand) {
					yield = append(yield, operand)
				}
			}
			return istd.NewExpressed[T](nil, yield...)
		}
		yield = operands[*expr.Low:*expr.High]

		operands = yield
	}
	return istd.NewExpressed[T](nil, operands...)
}

// Emit expresses the underlying bits of the Operable operands according to the provided logical expression.
//
// NOTE: This will return a tiny.ErrorOutOfData if the expression either couldn't be satisfied or no input data was provided.
func Emit[T any](expr istd.Expression, operands ...T) istd.Expressed[std.Measurement] {
	var totalWidth uint
	var test T
	switch any(test).(type) {
	case std.Bit, byte, std.Measurement, std.Phrase, std.Natural, std.Real, std.Complex, std.Index:
		for _, op := range operands {
			totalWidth += tiny.GetOperableBitWidth(op)
		}
	default:
		p := std.NewPhrase()
		for _, op := range operands {
			p = p.AppendMeasurement(tiny.Measure(op))
		}

		return Emit(expr, p)
	}

	// Do nothing if there is no binary information to emit
	if totalWidth == 0 {
		return istd.NewExpressed(tiny.ErrorOutOfData, std.NewMeasurement())
	}

	// Evaluate and set the appropriate expression boundary values
	expr, operands = evaluateBaseExpression(expr, totalWidth, operands...)

	// Switch between matrix or linear logic based upon the presence of an artifact function

	// Matrix logic - Performs logic at the phrase level while emitting out the underlying bits
	if expr.Artifact != nil {
		yield, _ := matrixLogic(0, expr, operands...)
		return istd.NewExpressed(nil, std.NewMeasurement(yield...))
	}

	// Linear logic - Recurses to the bit level before performing logic
	yield, _ := linearLogic(0, expr, operands...)
	if len(yield) < int(expr.Limit) {
		return istd.NewExpressed(tiny.ErrorOutOfData, std.NewMeasurement(yield...))
	}

	return istd.NewExpressed(nil, std.NewMeasurement(yield...))
}

func linearLogic[T any](cursor uint, expr istd.Expression, operands ...T) ([]std.Bit, uint) {
	yield := make([]std.Bit, 0, 1<<10) // Pre-allocate a reasonable chunk of memory

	// Walk through the current operands one at a time
	for _, raw := range operands {
		if tiny.GetOperableBitWidth(raw) == 0 {
			continue
		}

		cycleBits := make([]std.Bit, 0, 1<<10) // Pre-allocate a reasonable chunk of memory

		// Decompose them through recursion
		switch operand := any(raw).(type) {
		case std.Complex:
			panic(fmt.Errorf("cannot perform linear logic on complex numbers as they cannot be implicitly aligned"))
		case std.Phrase:
			// Phrases recurse into their respective measurements
			var bits []std.Bit
			bits, cursor = linearLogic(cursor, expr, operand.GetData()...)
			cycleBits = append(cycleBits, bits...)
		case std.Index:
			// Indexes recurse into their respective measurements
			var bits []std.Bit
			bits, cursor = linearLogic(cursor, expr, operand.GetData()...)
			cycleBits = append(cycleBits, bits...)
		case std.Real:
			// Reals recurse into their phrase form
			var bits []std.Bit
			bits, cursor = linearLogic(cursor, expr, operand.AsPhrase())
			cycleBits = append(cycleBits, bits...)
		case std.Natural:
			// Naturals recurse into their composed measurement
			var bits []std.Bit
			bits, cursor = linearLogic(cursor, expr, operand.Measurement)
			cycleBits = append(cycleBits, bits...)
		case std.Measurement:
			// Measurements recurse into their individual bits
			var bits []std.Bit

			bits, cursor = linearLogic(cursor, expr, operand.GetAllBits()...)
			cycleBits = append(cycleBits, bits...)
		case []byte:
			// Byte slices recurse into their individual bytes
			var bits []std.Bit
			bits, cursor = linearLogic(cursor, expr, operand...)
			cycleBits = append(cycleBits, bits...)
		case byte:
			// Bytes recurse into their individual bits
			bits := make([]std.Bit, 8)
			ii := 0
			for i := 7; i >= 0; i-- {
				bits[ii] = std.Bit((operand >> i) & 1)
				ii++
			}
			bits, cursor = linearLogic(cursor, expr, bits...)
			cycleBits = append(cycleBits, bits...)
		case std.Bit:
			// Bits step the cursor across the bits and select out data
			if expr.BitLogic != nil {
				bits, _ := (*expr.BitLogic)(cursor, operand)
				operand = bits[0]
			}

			if expr.Positions != nil && len(*expr.Positions) > 0 {
				// We are performing explicit position selection
				for _, pos := range *expr.Positions {
					if pos == cursor {

						cycleBits = append(cycleBits, operand)
					}
				}
			} else {
				// We are performing ranged selection
				if cursor >= *expr.Low && cursor < *expr.High {
					cycleBits = append(cycleBits, operand)
				}
			}

			// Increment the cursor's current bit position in the source information
			cursor++
		default:
			panic(fmt.Errorf("invalid binary type: %T", operand))
		}

		// Yield the found bits
		yield = append(yield, cycleBits...)

		// Check if there is a continuation function and whether it has returned false
		if expr.Continue != nil && !(*expr.Continue)(cursor, yield) {
			return yield, cursor
		}

		// Bailout when the pre-calculated limit has been met
		overage := len(yield) - int(expr.Limit)
		if overage > 0 {
			cursor -= uint(overage)
			return yield[:int(expr.Limit)], cursor
		}
	}
	return yield, cursor
}

func matrixLogic[T any](cursor uint, expr istd.Expression, operands ...T) ([]std.Bit, uint) {
	// TODO: start sub-expressions to grab bits and build a matrix for computation

	//if expr._matrix != nil && *expr._matrix {
	//	/**
	//	Matrix Logic
	//	*/
	//
	//	if expr._matrixLogic == nil {
	//		panic("matrix expressions require a logic function")
	//	}
	//
	//	calculate := *expr._matrixLogic
	//
	//	if expr._alignment == nil {
	//		align := PadLeftSideWithZeros
	//		expr._alignment = &align
	//	}
	//
	//	longest := GetWidestOperand(data...)
	//
	//	if longest <= 0 {
	//		return yield, 0
	//	}
	//
	//	subExpr := expr
	//	subExpr._matrix = &False
	//
	//	// The underlying table is ordered [Col][Row]Bit
	//	table := make([][]std.Bit, longest)
	//	for i, raw := range data {
	//		data[i] = AlignOperands(raw, longest, *expr._alignment)
	//		table[i], _ = Emit[T](subExpr, raw)
	//	}
	//
	//	// TODO: We can't walk using longest because longest will grow as we carry - instead we need to just walk until we are out of bits to walk and pass the walk count to the matrix func
	//
	//	for i := 0; i < longest; i++ {
	//		colId := i
	//		if reverse {
	//			colId = longest - i - 1
	//		}
	//
	//		column := make([]std.Bit, len(table))
	//		for rowId, row := range table {
	//			column[rowId] = row[colId]
	//		}
	//		calculated, overflow := calculate(colId, column...)
	//
	//		// TODO: Insert the overflow binary value BELOW the upcoming columns in the direction of calculation
	//
	//		if reverse {
	//			yield = append(yield, calculated)
	//		} else {
	//			yield = append([]std.Bit{b}, yield...)
	//		}
	//	}
	//
	//	linear := make([]std.Bit, 0, len(matrix)*longest)
	//	for _, element := range matrix {
	//		linear = append(linear, element...)
	//	}
	//
	//	yield = linear
	//	count = uint(longest) // TODO: Alignment all the operands and set this to the number of returned operands
	//} else {
	return nil, 0
}

// evaluateBaseExpression performs sanity checks and sets the boundary values for the expression.
func evaluateBaseExpression[T any](expr istd.Expression, totalWidth uint, operands ...T) (istd.Expression, []T) {
	if (expr.Positions != nil && len(*expr.Positions) > 0) && (expr.Low != nil || expr.High != nil) {
		panic("cannot search for an explicit position inside of a range - you can perform that operation with compound emit operations")
	}

	// Set our cursor limit
	if expr.Positions != nil {
		// We are performing point selection
		expr.Limit += uint(len(*expr.Positions))
	} else {
		// We are performing range selection
		expr.Limit = totalWidth

		// Set the boundaries and limit value according to the expression
		if expr.Low != nil {
			expr.Limit -= uint(*expr.Low)

			if expr.High != nil {
				expr.Limit = *expr.High - *expr.Low
			} else {
				// If no high boundary, set it to the last index of the operands
				last := uint(int(totalWidth))
				expr.High = &last
			}
		} else {
			first := uint(0)
			last := uint(int(totalWidth))

			if expr.High != nil {
				last = *expr.High
			}

			expr.Low = &first
			expr.High = &last
		}
	}

	// Calculate the last bit position of the operands, if requested
	if expr.Last != nil {
		added := append(*expr.Positions, totalWidth-1)
		expr.Positions = &added

		expr.Last = nil
	}

	// Check if the data should be reversed at this point
	if expr.Reverse != nil && *expr.Reverse {
		// If so...put your thing down, flip it, and reverse it
		operands = ReverseOperands(operands...)
		expr.Reverse = nil
	}
	return expr, operands
}
