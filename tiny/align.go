package tiny

import (
	"fmt"
	"github.com/ignite-laboratories/core/enum/direction"
	"github.com/ignite-laboratories/core/enum/traveling"
	"github.com/ignite-laboratories/core/std"
)

func middlePadOperands[T std.Operable](width uint, d direction.Direction, t traveling.Traveling, digits []std.Bit, operands ...T) []T {
	// TODO: Implement north/south padding
	out := make([]T, len(operands))

	for i, o := range operands {
		l := GetOperableBitWidth(o)
		toPad := width - l
		left := toPad / 2
		right := toPad - left

		if t == traveling.Outbound {
			out[i] = padOperands(left, d, traveling.Westbound, digits, o)[0]
			out[i] = padOperands(right, d, traveling.Eastbound, digits, out[i])[0]
		} else if t == traveling.Inbound {
			out[i] = padOperands(left, d, traveling.Eastbound, digits, o)[0]
			out[i] = padOperands(right, d, traveling.Westbound, digits, out[i])[0]
		} else {
			out[i] = padOperands(left, d, t, digits, o)[0]
			out[i] = padOperands(right, d, t, digits, out[i])[0]
		}
	}

	return out
}

func padOperands[T std.Operable](width uint, d direction.Direction, t traveling.Traveling, digits []std.Bit, operands ...T) []T {
	out := make([]T, len(operands))

	if d == direction.North || d == direction.South {
		// TODO: Implement north/south padding
		return operands
	}

	for i, raw := range operands {
		paddingWidth := width - GetOperableBitWidth(raw)
		if paddingWidth == 0 {
			out[i] = raw
			continue
		}

		padding := std.NewMeasurementOfPattern(paddingWidth, t, digits...).GetAllBits()

		switch operand := any(raw).(type) {
		case std.Phrase, std.Complex, std.Index, std.Real, std.Natural:
		case std.Measurement:
			switch d {
			case direction.West:
				out[i] = any(operand.Prepend(padding...)).(T)
			case direction.East:
				out[i] = any(operand.Append(padding...)).(T)
			default:
				panic(fmt.Sprintf("cannot pad to '%v' direction - please use the cardinal directions North, West, South, and East", d))
			}
		case []std.Bit:
			switch d {
			case direction.West:
				out[i] = any(append(padding, operand...)).(T)
			case direction.East:
				out[i] = any(append(operand, padding...)).(T)
			default:
				panic(fmt.Sprintf("cannot pad to '%v' direction - please use the cardinal directions North, West, South, and East", d))
			}
		case []byte:
		case byte:
		case std.Bit:
			panic("cannot pad static width elements")
		default:
			panic("unknown operand type")
		}
	}
	return out
}
