package binarysearch

// FindInRotatedSortedArray .
// 数组 nums 由一个递增数组经过旋转变形（也可能不旋转）
// [1, 2, 3, 4, 5] -> [4, 5, 1, 2, 3]
// 用 O(log(n)) 的事件复杂度找到目标元素的index，没有返回 -1
func FindInRotatedSortedArray(nums []int, target int) int {
	/*
		正常二分查找从 left, middle, right 开始找

		如果 left < right 那么意味着 num[left:right] 没有旋转，
			如果目标落在此区间，采用正常二分查找

		如果 left > right 那么 left right 之间存在 旋转点
			那么 num[left:middle] num[middle:right] 之间有一个是没有旋转的子数组，一个是由旋转子数组

		找到递归
	*/

	return binarySearchRotatedSortedArray(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, left, right, target int) int {
	if left < 0 || right < 0 || left >= len(nums) || right >= len(nums) {
		return -1
	}

	if left > right {
		return -1
	}

	if left == right {
		if nums[left] == target {
			return left
		}

		return -1
	}

	middle := (left + right) / 2
	middleVal := nums[middle]

	if target < middleVal {
		return binarySearch(nums, left, middle, target)
	} else if target > middleVal {
		return binarySearch(nums, middle+1, right, target)
	} else {
		return middle
	}
}

func binarySearchRotatedSortedArray(nums []int, left, right, target int) int {
	if left < 0 || right < 0 || left >= len(nums) || right >= len(nums) {
		return -1
	}

	if left == right {
		if nums[left] == target {
			return left
		}

		return -1
	}

	leftVal := nums[left]
	rightVal := nums[right]

	if leftVal <= rightVal {
		return binarySearch(nums, left, right, target)
	}

	// middle 0 1 -> 0
	// middle 0 1 2 -> 1

	middle := (left + right) / 2
	middleVal := nums[middle]

	if left < middle {
		if leftVal < middleVal {
			if leftVal <= target && target <= middleVal {
				return binarySearch(nums, left, middle, target)
			} else {
				return binarySearchRotatedSortedArray(nums, middle+1, right, target)
			}
		} else {
			// leftVal > middleVal
			if middleVal < target && target <= rightVal {
				return binarySearch(nums, middle+1, right, target)
			} else {
				return binarySearchRotatedSortedArray(nums, left, middle, target)
			}
		}

	} else {
		// left == middle
		if target == leftVal {
			return left
		}
		if target == rightVal {
			return right
		}

		return -1
	}

}
