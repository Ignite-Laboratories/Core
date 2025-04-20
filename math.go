package core

import (
	"math"
	"math/rand"
)

// MaxIntegerValue returns the maximum value of the provided type.
func MaxIntegerValue[T Integer]() uint64 {
	switch any(T(0)).(type) {
	case int8:
		return math.MaxInt8
	case uint8:
		return uint64(math.MaxUint8)
	case int16:
		return uint64(math.MaxInt16)
	case uint16:
		return uint64(math.MaxUint16)
	case int32:
		return uint64(math.MaxInt32)
	case uint32:
		return uint64(math.MaxUint32)
	case int64:
		return uint64(math.MaxInt64)
	case uint64:
		return math.MaxUint64
	case int:
		return math.MaxInt
	case uint:
		return math.MaxUint
	default:
		panic("unsupported numeric type")
	}
}

// NormalizeToFloat64 returns a normalized value of the provided type in the range [0.0, 1.0].
func NormalizeToFloat64[T Integer](value T) float64 {
	return float64(value) / float64(MaxIntegerValue[T]())
}

// NormalizeToFloat32 returns a normalized value of the provided type in the range [0.0, 1.0].
func NormalizeToFloat32[T Integer](value T) float32 {
	return float32(value) / float32(MaxIntegerValue[T]())
}

// ScaleFloat64ToType returns a scaled value of the provided type in the range [0, T.MaxValue].
//
// NOTE: This will panic if the provided value is greater than the maximum value of the provided type.
func ScaleFloat64ToType[T Integer](value float64) T {
	if value < 0.0 || value > 1.0 {
		panic("value must be in range [0.0, 1.0]")
	}
	return T(value * float64(MaxIntegerValue[T]()))
}

// ScaleFloat32ToType returns a scaled value of the provided type in the range [0, T.MaxValue].
//
// NOTE: This will panic if the provided value is greater than the maximum value of the provided type.
func ScaleFloat32ToType[T Integer](value float32) T {
	if value < 0.0 || value > 1.0 {
		panic("value must be in range [0.0, 1.0]")
	}
	return T(value * float32(MaxIntegerValue[T]()))
}

// RandomNumber returns a non-negative pseudo-random number of the provided type.
//
// If requesting a floating point type, the resulting number will be bounded
// in the fully closed interval [0.0, 1.0]
//
// If requesting an integer type, the resulting number will be bounded
// in the fully closed interval [0, n] - where n is the maximum value of
// the provided type.
func RandomNumber[T Numeric]() T {
	switch any(T(0)).(type) {
	case float32:
		return T(RandomNumberRange[float32](NumericRange[float32]{0.0, 1.0}))
	case float64:
		return T(RandomNumberRange[float64](NumericRange[float64]{0.0, 1.0}))
	case int8:
		return T(RandomNumberRange[int8](NumericRange[int8]{math.MinInt8, math.MaxInt8}))
	case uint8:
		return T(RandomNumberRange[uint8](NumericRange[uint8]{0, math.MaxUint8}))
	case int16:
		return T(RandomNumberRange[int16](NumericRange[int16]{math.MinInt16, math.MaxInt16}))
	case uint16:
		return T(RandomNumberRange[uint16](NumericRange[uint16]{0, math.MaxUint16}))
	case int32:
		return T(RandomNumberRange[int32](NumericRange[int32]{math.MinInt32, math.MaxInt32}))
	case uint32:
		return T(RandomNumberRange[uint32](NumericRange[uint32]{0, math.MaxUint32}))
	case int64:
		return T(RandomNumberRange[int64](NumericRange[int64]{math.MinInt64, math.MaxInt64}))
	case int:
		return T(RandomNumberRange[int](NumericRange[int]{math.MinInt, math.MaxInt}))
	case uint64:
		return T(RandomNumberRange[uint64](NumericRange[uint64]{0, math.MaxUint64}))
	case uint:
		return T(RandomNumberRange[uint](NumericRange[uint]{0, math.MaxUint}))
	default:
		panic("unsupported numeric type")
	}
}

// RandomNumberRange returns a pseudo-random number of the provided type bounded in the closed interval [min, max].
//
// NOTE: This uses a 0.01% chance to return exactly max.
func RandomNumberRange[T Numeric](numericRange NumericRange[T]) T {
	if numericRange.Start >= numericRange.Stop {
		return numericRange.Start
	}
	switch any(T(0)).(type) {
	case float32, float64:
		// 0.1% chance to return exactly max
		if rand.Float64() < 0.001 {
			return numericRange.Stop
		}
		return T(float64(numericRange.Start) + (float64(numericRange.Stop)-float64(numericRange.Start))*rand.Float64())
	case int8, int16, int32, int64, int, uint8, uint16, uint32, uint64, uint:
		range64 := uint64(numericRange.Stop) - uint64(numericRange.Start)
		return T(uint64(numericRange.Start) + uint64(rand.Int63n(int64(range64+1))))
	default:
		panic("unsupported numeric type")
	}
}

type NumericRange[T Numeric] struct {
	Start T
	Stop  T
}

func Range[T Numeric](start T, stop T) NumericRange[T] {
	return NumericRange[T]{
		Start: start,
		Stop:  stop,
	}
}
