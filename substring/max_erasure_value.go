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

func MaxErasureValueTwoPointer(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 0 {
		return nums[0]
	}

	left, right := 0, 1
	resultForSubsetX := make([]int, len(nums))
	resultIncludeX := make([]int, len(nums))
	set := make(map[int]struct{})

	resultForSubsetX[0] = nums[0]
	resultIncludeX[0] = nums[0]
	set[nums[0]] = struct{}{}

	for right < len(nums) {
		rnum := nums[right]
		_, dup := set[rnum]
		if !dup {
			set[rnum] = struct{}{}
			resultIncludeX[right] = resultIncludeX[right-1] + rnum
		} else {
			var popSum int
			for left < right {
				lnum := nums[left]
				popSum += lnum
				delete(set, lnum)
				left += 1
				if lnum == rnum {
					break
				}
			}
			set[rnum] = struct{}{}
			resultIncludeX[right] = resultIncludeX[right-1] + rnum - popSum
		}

		resultForSubsetX[right] = max(resultForSubsetX[right-1], resultIncludeX[right])

		right += 1
	}

	return resultForSubsetX[len(nums)-1]
}
