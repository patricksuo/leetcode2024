package number

import "math"

// ReverseIntegerx .
// Given a signed 32-bit integer x, return x with its digits reversed.
// If reversing x causes the value to go outside the signed 32-bit integer range [-2^31, 2^31 - 1],
// then return 0.
func ReverseIntegerx(xx int) int {
	x := int32(xx)
	var (
		result   int32
		negative = x < 0
	)
	if negative {
		for x < 0 {
			if MulOverflow32(result, 10) {
				return 0
			}
			newResult := result * 10

			dig := x % -10
			if AddOverflow32(newResult, dig) {
				return 0
			}

			newResult += dig

			x = x / 10

			result = newResult
		}
	} else {
		for x > 0 {
			if MulOverflow32(result, 10) {
				return 0
			}
			newResult := result * 10

			dig := x % 10
			if AddOverflow32(newResult, dig) {
				return 0
			}

			newResult += dig
			x = x / 10

			result = newResult
		}
	}

	return int(result)
}

func AddOverflow32(x, y int32) bool {
	// minint32<= x + y <= maxint32

	if x > 0 && y > 0 {
		return math.MaxInt32-x < y
	}

	if x < 0 && y < 0 {
		return math.MinInt32-x > y
	}

	return false
}

// MulOverflow32 判断是否溢出
// 有更多边界需要处理
func MulOverflow32(x, y int32) bool {
	if x == 0 || y == 0 {
		return false
	}

	return x*y/y != x
}
