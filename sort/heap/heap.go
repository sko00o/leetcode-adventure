package heap

// Heap is a specialized tree-based data structure which is essentially an
// almost complete. We use a slice to store elements.
type Heap struct {
	data []int
	size int
}

// Top get top element in heap.
// Heap can not be empty.
func (h *Heap) Top() int {
	if h.size == 0 {
		return -1
	}
	return h.data[0]
}

// IsEmpty will return true if the heap is empty.
func (h *Heap) IsEmpty() bool {
	return h.size == 0
}

// MaxHeap defines a max heap.
type MaxHeap struct {
	Heap
}

// NewMaxHeap will make a max heap with a fixed size.
func NewMaxHeap(n int) *MaxHeap {
	return &MaxHeap{
		Heap: Heap{
			data: make([]int, n),
		},
	}
}

// Push one element into the heap.
// time complexity: O(1) ~ O(nlogn)
func (h *MaxHeap) Push(x int) {
	if h.size == 0 {
		h.data[0] = x
		h.size++
		return
	}

	curr := h.size
	for curr > 0 {
		parent := (curr - 1) / 2
		// 当前值比上一级小不需要上浮
		if x <= h.data[parent] {
			break
		}
		h.data[curr] = h.data[parent]
		curr = parent
	}
	h.data[curr] = x
	h.size++
}

// Pop top element out of the heap.
// Return true if the operation is successful.
// time complexity: O(1) ~ O(logn)
func (h *MaxHeap) Pop() bool {
	if h.size == 0 {
		return false
	}

	h.size--
	x := h.data[h.size]
	parent := 0
	for 2*parent+1 < h.size {
		child := 2*parent + 1
		rchild := 2*parent + 2
		// 选一个较大的孩子
		if rchild < h.size && h.data[rchild] > h.data[child] {
			child = rchild
		}

		// 已经比其孩子大，不需要下沉
		if x >= h.data[child] {
			break
		}

		h.data[parent] = h.data[child]
		parent = child
	}

	h.data[parent] = x
	return true
}

// MinHeap defines a min heap
type MinHeap struct {
	Heap
}

// NewMinHeap will make a min heap with a fixed size.
func NewMinHeap(n int) *MinHeap {
	return &MinHeap{
		Heap: Heap{
			data: make([]int, n),
		},
	}
}

// Push one element into the heap.
func (h *MinHeap) Push(x int) {
	if h.size == 0 {
		h.data[0] = x
		h.size++
		return
	}

	child := h.size
	for child > 0 {
		parent := (child - 1) / 2
		// 上一级已经是最小，不需要上浮
		if x >= h.data[parent] {
			break
		}
		h.data[child] = h.data[parent]
		child = parent
	}
	h.data[child] = x
	h.size++
}

// Pop top element out of the heap.
// Return true if the operation is successful.
func (h *MinHeap) Pop() bool {
	if h.size == 0 {
		return false
	}

	h.size--
	x := h.data[h.size]
	parent := 0
	for 2*parent+1 < h.size {
		child := 2*parent + 1
		rchild := 2*parent + 2
		// 选一个较小的孩子
		if rchild < h.size && h.data[rchild] < h.data[child] {
			child = rchild
		}

		// 已经比其孩子小，不需要下沉
		if x <= h.data[child] {
			break
		}

		h.data[parent] = h.data[child]
		parent = child
	}

	h.data[parent] = x
	return true
}
