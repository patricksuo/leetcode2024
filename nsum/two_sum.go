package nsum

// TwoSum .
// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.
func TwoSum(array []int, target int) []int {
	pop := func(set map[int]struct{}, exclude int) (int, bool) {
		for key := range set {
			if key != exclude {
				return key, true
			}
		}
		return 0, false
	}

	elm2idx := make(map[int]map[int]struct{}, len(array))
	for idx, elm := range array {
		set, ok := elm2idx[elm]
		if ok {
			set[idx] = struct{}{}
		} else {
			set = make(map[int]struct{})
			set[idx] = struct{}{}
			elm2idx[elm] = set
		}
	}

	for idx, elm := range array {
		other := target - elm
		set, ok := elm2idx[other]
		if !ok {
			continue
		}

		otherIdx, ok := pop(set, idx)
		if ok {
			return []int{idx, otherIdx}
		}
	}

	panic("no solution")
}

// TwoSumSortedArray .
// Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order, find two numbers such that they add up to a specific target number. Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1 < index2 <= numbers.length.
// Return the indices of the two numbers, index1 and index2, added by one as an integer array [index1, index2] of length 2.
// The tests are generated such that there is exactly one solution. You may not use the same element twice.
// Your solution must use only constant extra space.
// binary search
func TwoSumSortedArray(array []int, target int) []int {
	for idx, elm := range array {
		otherElm := target - elm
		otherIdx, ok := TwoSumSortedArray_binarySearch(array, otherElm, 0, len(array)-1, idx)
		if !ok {
			continue
		}
		return []int{idx, otherIdx}
	}

	panic("no solution")
}

func TwoSumSortedArray_binarySearch(array []int, target, left, right, exclude int) (int, bool) {
	// stop condition
	if left > right {
		return 0, false
	}

	if left == exclude {
		return TwoSumSortedArray_binarySearch(array, target, left+1, right, exclude)
	}

	if right == exclude {
		return TwoSumSortedArray_binarySearch(array, target, left, right-1, exclude)
	}

	mid := (left + right) / 2
	midElm := array[mid]
	if target == midElm {
		// found, but this element is used.
		if mid == exclude {

			// try search left
			if targetIdx, ok := TwoSumSortedArray_binarySearch(array, target, left, mid-1, exclude); ok {
				return targetIdx, ok
			}

			// try search right
			if targetIdx, ok := TwoSumSortedArray_binarySearch(array, target, mid+1, right, exclude); ok {
				return targetIdx, ok
			}

			// no found
			return 0, false
		}

		return mid, true
	}

	if target > midElm {
		// mid + 1 > left?
		// mid >= left ---> mid + 1 > left
		return TwoSumSortedArray_binarySearch(array, target, mid+1, right, exclude)
	} else {
		// mid - 1 < right ?
		// mid <= right ---> mid - 1 < right
		return TwoSumSortedArray_binarySearch(array, target, left, mid-1, exclude)
	}
}
