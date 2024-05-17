package number

import "math"

func Atoi(s string) int {
	var (
		result     int
		idx        int
		isNegative bool
	)

	// 1. read leading empty
	for idx < len(s) && isEmpty(s[idx]) {
		idx += 1
	}

	// 2. read sign
	if idx < len(s) && isSign(s[idx]) {
		isNegative = s[idx] == '-'
		idx += 1
	}

	// 3. read one digit and detect overflow
	for idx < len(s) {
		digit, ok := tryReadDigit(s[idx])
		if !ok {
			break
		}

		if isNegative {
			if MulOverflow32(int32(result), 10) {
				return int(math.MinInt32)
			}

			result = result * 10

			if AddOverflow32(int32(result), int32(-1*digit)) {
				return int(math.MinInt32)
			}

			result -= digit
		} else {
			if MulOverflow32(int32(result), 10) {
				return int(math.MaxInt32)
			}

			result = result * 10

			if AddOverflow32(int32(result), int32(digit)) {
				return int(math.MaxInt32)
			}

			result += digit
		}

		idx += 1
	}

	return result
}

func tryReadDigit(char byte) (int, bool) {
	if char >= '0' && char <= '9' {
		return int(char) - int('0'), true
	}

	return -1, false
}

func isEmpty(char byte) bool {
	return char == ' '
}

func isSign(char byte) bool {
	return char == '-' || char == '+'
}
