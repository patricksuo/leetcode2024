package sortalg

// FindKthNum .
// https://leetcode.com/problems/kth-largest-element-in-an-array/description/
func FindKthNum(nums []int, k int) int {
	/*
		思路1 排序
		space O(1), time O(n * log(n))

		思路2 维护一个大顶堆
		space O(k), O(n * log(k))
	*/
	heap := MakeMiniHeap(k)
	for _, num := range nums {
		heap.Add(num)
	}

	miniElm, ok := heap.TryGetTop()
	if !ok {
		panic("not enough nums")
	}
	return miniElm

}

/*
					parent array[n]
	left child array[2n+1]   right child array[2n+2]
	parent <= left child && parent <= right child

	1. 如果 heap.size() < k 那么 找一个位置放进去
		把 新元素放到堆末尾，和 其parent交换，以维护堆的性质
	2. 如果 heap.size() == k
		2.1 如果新元素比 array[0] 目前堆中最小的元素小，那么替换，并维护 heap 的性质
			和其 child 交换，以维护堆的性质
		2.2 否则丢弃这个新元素
*/

type MiniHeap struct {
	cap   int
	size  int
	array []int
}

func MakeMiniHeap(cap int) *MiniHeap {
	if cap <= 0 {
		panic("cap must be postive")
	}

	h := &MiniHeap{
		cap:   cap,
		size:  0,
		array: make([]int, cap),
	}

	return h
}

func (h *MiniHeap) TryGetTop() (int, bool) {
	if h.size > 0 {
		return h.array[0], true
	}

	return 0, false
}

func (h *MiniHeap) Add(elm int) {
	if h.size < h.cap {
		h.array[h.size] = elm
		h.size += 1

		h.heapnifyChild2Parent(h.size - 1)
	} else {
		if elm <= h.array[0] {
			return
		}

		h.array[0] = elm
		h.heapnifyParent2Child(0)
	}
}

func (h *MiniHeap) heapnifyChild2Parent(idx int) {
	/*
		leftChildIdx := parentIdx * 2 + 1
		rightChildIdx := parentIdx * 2 + 2

		0 => 1 2
		1 -> 3 ,4
		2 -> 5, 6
	*/

	var (
		parentIdx int
	)
	if idx&1 == 1 {
		parentIdx = idx / 2
	} else {
		parentIdx = idx/2 - 1
	}

	if parentIdx < 0 {
		return
	}

	if h.array[idx] < h.array[parentIdx] {
		h.array[idx], h.array[parentIdx] = h.array[parentIdx], h.array[idx]
	}

	h.heapnifyChild2Parent(parentIdx)
}

func (h *MiniHeap) heapnifyParent2Child(idx int) {
	leftChild := 2*idx + 1
	rightChild := 2*idx + 2

	var (
		miniIdx = idx
	)

	if leftChild < h.cap && h.array[leftChild] < h.array[miniIdx] {
		miniIdx = leftChild
	}

	if rightChild < h.cap && h.array[rightChild] < h.array[miniIdx] {
		miniIdx = rightChild
	}

	// 无需交换
	if miniIdx == idx {
		return
	}

	h.array[idx], h.array[miniIdx] = h.array[miniIdx], h.array[idx]

	h.heapnifyParent2Child(miniIdx)
}
