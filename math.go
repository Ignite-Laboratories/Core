package core

import (
	"math"
	"math/rand"
)

// MaxValue returns the maximum value of the provided type.
func MaxValue[T Numeric]() T {
	switch any(T(0)).(type) {
	case float32:
		return T(float32(math.MaxFloat32))
	case float64:
		return T(float64(math.MaxFloat64))
	case int8:
		return T(int8(math.MaxInt8))
	case uint8:
		return T(uint8(math.MaxUint8))
	case int16:
		return T(int16(math.MaxInt16))
	case uint16:
		return T(uint16(math.MaxUint16))
	case int32:
		return T(int32(math.MaxInt32))
	case uint32:
		return T(uint32(math.MaxUint32))
	case int64:
		return T(int64(math.MaxInt64))
	case int:
		return T(int(math.MaxInt))
	case uint64:
		return T(uint64(math.MaxUint64))
	case uint:
		return T(uint(math.MaxUint))
	default:
		panic("unsupported numeric type")
	}
}

// NormalizeToFloat64 returns a normalized value of the provided type in the range [0.0, 1.0].
func NormalizeToFloat64[T Numeric](value T) float64 {
	return float64(value) / float64(MaxValue[T]())
}

// NormalizeToFloat32 returns a normalized value of the provided type in the range [0.0, 1.0].
func NormalizeToFloat32[T Numeric](value T) float32 {
	return float32(value) / float32(MaxValue[T]())
}

// ScaleFloat64ToType returns a scaled value of the provided type in the range [0, T.MaxValue].
//
// NOTE: This will panic if the provided value is greater than the maximum value of the provided type.
func ScaleFloat64ToType[T Numeric](value float64) T {
	if value < 0.0 || value > 1.0 {
		panic("value must be in range [0.0, 1.0]")
	}
	return T(value * float64(MaxValue[T]()))
}

// ScaleFloat32ToType returns a scaled value of the provided type in the range [0, T.MaxValue].
//
// NOTE: This will panic if the provided value is greater than the maximum value of the provided type.
func ScaleFloat32ToType[T Numeric](value float32) T {
	if value < 0.0 || value > 1.0 {
		panic("value must be in range [0.0, 1.0]")
	}
	return T(value * float32(MaxValue[T]()))
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
		return RandomNumberRange[T](0.0, 1.0)
	case float64:
		return RandomNumberRange[T](0.0, 1.0)
	case int8:
		return RandomNumberRange[T](math.MinInt8, math.MaxInt8)
	case uint8:
		return RandomNumberRange[T](0, math.MaxUint8)
	case int16:
		return RandomNumberRange[T](math.MinInt16, math.MaxInt16)
	case uint16:
		return RandomNumberRange[T](0, math.MaxUint16)
	case int32:
		return RandomNumberRange[T](math.MinInt32, math.MaxInt32)
	case uint32:
		return RandomNumberRange[T](0, math.MaxUint32)
	case int64:
		return RandomNumberRange[T](math.MinInt64, math.MaxInt64)
	case int:
		return RandomNumberRange[T](math.MinInt, math.MaxInt)
	case uint64:
		return RandomNumberRange[T](0, math.MaxUint64)
	case uint:
		return RandomNumberRange[T](0, math.MaxUint)
	default:
		panic("unsupported numeric type")
	}
}

// RandomNumberRange returns a pseudo-random number of the provided type bounded in the closed interval [min, max].
//
// NOTE: This uses a 0.01% chance to return exactly max.
func RandomNumberRange[T Numeric](min, max T) T {
	if min >= max {
		return min
	}
	switch any(T(0)).(type) {
	case float32, float64:
		// 1% chance to return exactly max
		if rand.Float64() < 0.01 {
			return max
		}
		return T(float64(min) + (float64(max)-float64(min))*rand.Float64())
	case int8, int16, int32, int64, int, uint8, uint16, uint32, uint64, uint:
		range64 := uint64(max) - uint64(min)
		return T(uint64(min) + uint64(rand.Int63n(int64(range64+1))))
	default:
		panic("unsupported numeric type")
	}
}
