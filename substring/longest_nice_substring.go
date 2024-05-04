package substring

// LongestNiceSubarray .
// You are given an array nums consisting of positive integers.
// We call a subarray of nums nice if the bitwise AND of every pair of elements that are in different positions in the subarray is equal to 0.
// Return the length of the longest nice subarray.
// A subarray is a contiguous part of an array.
// Note that subarrays of length 1 are always considered nice.
// https://leetcode.com/problems/longest-nice-subarray/
func LongestNiceSubarray(nums []int) int {
	// 原始状态，只有1个数字时，此时 substring is nice
	// 尝试往右扩张 substring
	// 判断是否组成 nice substring
	// 如果不是从左边开始剔除一些元素，直到构成新的 nice substring

	if len(nums) == 0 {
		return 0
	}

	left, right := 0, 1
	result := 1
	for _, value := range nums[1:] {

		isNice := true
		newLeftIdx := -1
		for i := left; i < right; i++ {
			if nums[i]&value != 0 {
				isNice = false
				newLeftIdx = i
			}
		}

		if !isNice {
			left = newLeftIdx + 1
		}

		if isNice && right-left+1 > result {
			result = right - left + 1
		}

		right += 1
	}

	return result
}
