package substring

func MaxErasureValue(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	resultForSubsetX := make([]int, len(nums))
	resultIncludeX := make([]int, len(nums))
	seqSizeIncludeX := make([]int, len(nums))

	resultForSubsetX[0] = nums[0]
	resultIncludeX[0] = nums[0]
	seqSizeIncludeX[0] = 1

	for idx, num := range nums[1:] {
		idx += 1

		seqSize := seqSizeIncludeX[idx-1]
		foundDup := false
		sumWithoutDup := 0
		preIdx := idx - 1
		nElm := 0
		for seqSize > 0 {
			if nums[preIdx] == num {
				foundDup = true
				break
			} else {
				sumWithoutDup += nums[preIdx]
				preIdx -= 1
				seqSize -= 1
				nElm += 1
			}
		}
		if foundDup {
			seqSizeIncludeX[idx] = nElm + 1
		} else {
			seqSizeIncludeX[idx] = seqSizeIncludeX[idx-1] + 1
		}

		resultIncludeX[idx] = sumWithoutDup + num
		resultForSubsetX[idx] = max(resultForSubsetX[idx-1], resultIncludeX[idx])
	}

	return resultForSubsetX[len(resultForSubsetX)-1]
}
