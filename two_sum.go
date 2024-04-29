package leetcode2024

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
