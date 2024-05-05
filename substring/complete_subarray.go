package substring

// CountCompleteSubarrays .
// https://leetcode.com/problems/count-complete-subarrays-in-an-array/
// You are given an array nums consisting of positive integers.
// We call a subarray of an array complete if the following condition is satisfied:
// The number of distinct elements in the subarray is equal to the number of distinct elements in the whole array.
// Return the number of complete subarrays.
// A subarray is a contiguous non-empty part of an array.
func CountCompleteSubarrays(nums []int) int {
	// 一开始陷入误区，想着用双指针从两边收缩查找 subarray 范围，但这样没有办法穷举所有可能性
	// 最后是暴力穷举所有组合
	if len(nums) <= 1 {
		return len(nums)
	}

	elm2count := make(map[int]struct{})
	for _, elm := range nums {
		elm2count[elm] = struct{}{}
	}
	elm2count2 := make(map[int]struct{}, len(elm2count))

	var result int
	for i := 0; i < len(nums); i++ {
		clear(elm2count2)
		for j := i; j < len(nums); j++ {
			elm2count2[nums[j]] = struct{}{}
			if len(elm2count2) == len(elm2count) {
				result += 1
			}
		}

		if len(elm2count2) < len(elm2count) {
			break
		}
	}

	return result
}
