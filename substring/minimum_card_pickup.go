package substring

import "sort"

func MinimumCardPickup(cards []int) int {
	if len(cards) <= 1 {
		return -1
	}

	for i, val := range cards {
		cards[i] = val << 32
		cards[i] |= i
	}

	sort.Ints(cards)

	preVal := cards[0] >> 32
	preIdx := cards[0] & 0xffffffff
	result := -1

	for _, masked := range cards[1:] {
		val := masked >> 32
		idx := masked & 0xffffffff
		if val == preVal {
			if result == -1 {
				result = idx - preIdx + 1
			} else {
				result = min(idx-preIdx+1, result)
			}
		}

		preVal = val
		preIdx = idx
	}

	return result
}
