package substring

// LongestPalindromeSubseq .
// A subsequence is a sequence that can be derived from another sequence by deleting some or no elements without changing the order of the remaining elements.
// https://leetcode.com/problems/longest-palindromic-subsequence/description/
// abc -> 1
// aba -> 3
// aaba -> 3
func LongestPalindromeSubsequence(s string) int {
	/*
		思路一 暴力破解
		排列组合有 2^N 种，每种排列判断是不是回文以及回文长度

		思路二
		取前n位字符构成子串，记其解为 L(n)
		包含第n位字符构成的最长记为 S(n)
		L(n + 1) = max(L(n), S(n))

		S(n) 可以由第 n 位字符单独构成回文，此时 S(n) 等于 1
		S(n) 可以是 aba 形式， 变成2个相同字符间的回文子序列长度
		找到左边的a时间复杂度是n，找n次 n^2

		求解 left a 到 right a 之间的回文子串是一个递归的问题，但看起来复杂度还是很高

		把 left a 与 right a 之间的解存起来？
	*/

	matrix := make([][]int, len(s))
	for i := range matrix {
		matrix[i] = make([]int, len(s))
		for j := range matrix[i] {
			matrix[i][j] = -1
			matrix[i][i] = 1
		}
	}

	return longestPalindromeSubsequence(s, 0, len(s)-1, matrix)
}

func longestPalindromeSubsequence(s string, left, right int, matrix [][]int) int {
	if left > right {
		return 0
	}

	if left == right {
		return 1
	}

	if matrix[left][right] != -1 {
		return matrix[left][right]
	}

	// 不含 第 right 个字符的解
	lpsWihoutRight := longestPalindromeSubsequence(s, left, right-1, matrix)

	// 含 第 right 个字符的解
	lpsWithRight := 1
	for i := left; i < right; i++ {
		if s[i] == s[right] {
			// 求解 aba 模式 b 的长度
			val := longestPalindromeSubsequence(s, i+1, right-1, matrix)
			if val+2 > lpsWithRight {
				lpsWithRight = val + 2
			}
			// 这里只需要找到最左边的a即可
			break
		}
	}

	matrix[left][right] = max(lpsWihoutRight, lpsWithRight)
	return matrix[left][right]
}
