package substring

// LongestPalindrome Given a string s, return the longest
// https://leetcode.com/problems/longest-palindromic-substring/description/
// 暴力接发，找出所有子串的排列组合，💡判断是否满足会文
// 时间复杂度 n^3
func LongestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	var result string

	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			subString := s[i : j+1]
			if len(subString) > len(result) && isPalidrome(subString) {
				result = subString
			}
		}
	}

	return result
}

func isPalidrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// LongestPalindromeFast .
func LongestPalindromeFast(s string) string {
	var result string

	for i := 0; i < len(s); i++ {
		maxExpand := 0
		for expand := 0; i-expand >= 0 && i+expand < len(s); expand++ {
			if s[i-expand] == s[i+expand] {
				maxExpand = expand
			} else {
				break
			}
		}
		subString := s[i-maxExpand : i+maxExpand+1]
		if len(subString) > len(result) {
			result = subString
		}
	}

	for i := 1; i < len(s); i++ {
		maxExpand := 0
		for expand := 1; i-expand >= 0 && i+expand-1 < len(s); expand++ {
			if s[i-expand] == s[i+expand-1] {
				maxExpand = expand
			} else {
				break
			}
		}

		subString := s[i-maxExpand : i+maxExpand]
		if len(subString) > len(result) {
			result = subString
		}
	}

	return result
}
