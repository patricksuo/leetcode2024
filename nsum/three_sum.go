package nsum

import "sort"

func ThreeSum(nums []int) (result [][]int) {

	tripleSet := make(map[[3]int]struct{})

	elm2idx := make(map[int]map[int]struct{}, len(nums))
	for idx, elm := range nums {
		set, ok := elm2idx[elm]
		if ok {
			set[idx] = struct{}{}
		} else {
			set = make(map[int]struct{})
			set[idx] = struct{}{}
			elm2idx[elm] = set
		}
	}

	for i := 0; i < len(nums); i++ {
		first := nums[i]

		for j := i + 1; j < len(nums); j++ {
			second := nums[j]
			third := 0 - first - second

			idxSet := elm2idx[third]
			for idx := range idxSet {
				if idx <= j {
					continue
				}

				record := []int{first, second, third}
				sort.Ints(record)
				recordArr := [3]int{
					record[0], record[1], record[2],
				}
				if _, ok := tripleSet[recordArr]; !ok {
					tripleSet[recordArr] = struct{}{}
					result = append(result, record)

				}

			}
		}

	}

	return result
}

func ThreeSumBinarySearch(nums []int) (result [][]int) {
	tripleSet := make(map[[3]int]struct{})

	sort.Ints(nums)

	for firstIdx, first := range nums {
		target := -first

		subArray := nums[firstIdx+1:]
		for secondIdx, second := range subArray {
			third := target - second

			_, ok := TwoSumSortedArray_binarySearch(subArray, third, 0, len(subArray)-1, secondIdx)
			if !ok {
				continue
			}

			triple := []int{first, second, third}
			sort.Ints(triple)
			_, dup := tripleSet[[3]int{triple[0], triple[1], triple[2]}]
			if dup {
				continue
			}

			tripleSet[[3]int{triple[0], triple[1], triple[2]}] = struct{}{}
			result = append(result, triple)
		}
	}

	return result
}
