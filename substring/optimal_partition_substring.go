package substring

// OptimalPartitionString .
// https://leetcode.com/problems/optimal-partition-of-string/
// Given a string s, partition the string into one or more substrings such that the characters in each substring are unique. That is, no letter appears in a single substring more than once.
// Return the minimum number of substrings in such a partition.
// Note that each character should belong to exactly one substring in a partition.
func OptimalPartitionString(s string) int {
	// 初始状态：取第一个char作为第一个substring
	// 每次迭代往右取一个新char，如果与当前substring的成员重复，那么独立作为一个新substring
	// 为什么这样是最优解？这样保证第i个substring长度最大，并且第i+1个substring包含的元素越少那么长度越大

	if len(s) <= 1 {
		return len(s)
	}

	set := make(map[byte]struct{})
	set[s[0]] = struct{}{}
	result := 1

	for cursor := 1; cursor < len(s); cursor++ {
		char := s[cursor]
		_, dup := set[char]
		if dup {
			clear(set)
			result += 1
		}

		set[char] = struct{}{}
	}

	return result
}
