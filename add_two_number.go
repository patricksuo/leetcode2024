package leetcode2024

type ListNode struct {
	Val  int
	Next *ListNode
}

// AddTwoNumbers .
// https://leetcode.com/problems/add-two-numbers/
// 需要注意点地方
// 1）循环的终止条件和两个指针的维护
// 2）两个链表各位相加之后可能存在进位
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int
	var result *ListNode
	var current *ListNode

	for l1 != nil || l2 != nil {
		val1 := nodeValue(l1)
		val2 := nodeValue(l2)
		val := carry + val1 + val2

		if val >= 10 {
			carry = 1
			val -= 10
		} else {
			carry = 0
		}

		node := &ListNode{
			Val:  val,
			Next: nil,
		}

		if result == nil {
			result = node
			current = node
		} else {
			current.Next = node
			current = node
		}
		l1 = nextNode(l1)
		l2 = nextNode(l2)
	}

	if carry != 0 {
		node := &ListNode{
			Val:  1,
			Next: nil,
		}

		if result == nil {
			result = node
		} else {
			current.Next = node
		}
	}

	return result
}

func nodeValue(node *ListNode) int {
	if node != nil {
		return node.Val
	}

	return 0
}

func nextNode(node *ListNode) *ListNode {
	if node != nil {
		return node.Next
	}

	return nil
}

// bitSumMatric .
// bitSumMatric[0][1] -> [2]int{result, carry}
var bitSumMatric = [2][2][2][2]int{
	{
		[2][2]int{
			{0, 0},
			{1, 0},
		},
		[2][2]int{
			{1, 0},
			{0, 1},
		},
	},
	{
		[2][2]int{
			{1, 0},
			{0, 1},
		},
		[2][2]int{
			{0, 1},
			{1, 1},
		},
	},
}

func AddTowNumBitOp(aa, bb int) int {
	var result int
	var shift int
	var carry int
	var a, b uint64
	a = uint64(aa)
	b = uint64(bb)

	for a > 0 || b > 0 {
		aBit := a & 1
		bBit := b & 1

		bitSumCache := bitSumMatric[aBit][bBit][carry]
		bit := bitSumCache[0]
		carry = bitSumCache[1]
		result = result | (bit << shift)

		shift += 1
		a = a >> 1
		b = b >> 1
	}

	if carry > 0 {
		result = result | (carry << shift)
	}

	return result
}
