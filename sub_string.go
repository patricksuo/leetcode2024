package leetcode2024

// LengthOfLongestSubstring .
// Given a string s, find the length of the longest substring without repeating characters.
func LengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	rune2idx := make(map[rune]int)
	subStringLenFromX := make([]int, len(s))
	subStringLenIncludeX := make([]int, len(s))

	subStringLenFromX[0] = 1
	subStringLenIncludeX[0] = 1
	rune2idx[rune(s[0])] = 0

	for idx, char := range s[1:] {
		idx += 1

		if subStringLenFromX[idx-1] == subStringLenIncludeX[idx-1] {
			dupIdx, dup := rune2idx[char]
			if !dup {
				subStringLenFromX[idx] = subStringLenFromX[idx-1] + 1
				subStringLenIncludeX[idx] = subStringLenIncludeX[idx-1] + 1
			} else {
				distance := idx - dupIdx
				subStringLenIncludeX[idx] = min(distance, subStringLenIncludeX[idx-1]+1)
				subStringLenFromX[idx] = max(subStringLenFromX[idx-1], subStringLenIncludeX[idx])
			}
		} else {
			dupIdx, dup := rune2idx[char]
			if !dup {
				subStringLenIncludeX[idx] = subStringLenIncludeX[idx-1] + 1
			} else {
				distance := idx - dupIdx
				subStringLenIncludeX[idx] = min(distance, subStringLenIncludeX[idx-1]+1)
			}
			subStringLenFromX[idx] = subStringLenFromX[idx-1]
		}

		rune2idx[char] = idx
	}

	return subStringLenFromX[len(subStringLenFromX)-1]
}
