package misc

// TopK is a min heap
type TopK struct {
	val  []int
	size int
}

func New(k int) *TopK {
	return &TopK{
		val: make([]int, k),
	}
}
func (h *TopK) heapifyUp() {
	child := h.size - 1
	for child > 0 {
		par := (child - 1) >> 1
		if h.val[child] >= h.val[par] {
			break
		}
		h.val[child], h.val[par] = h.val[par], h.val[child]
		child = par
	}
}
func (h *TopK) heapifyDown() {
	par := 0
	childs := func() (int, int) {
		return 2*par + 1, 2*par + 2
	}
	for child, bro := childs(); bro < h.size; child, bro = childs() {
		if h.val[child] > h.val[bro] {
			child = bro
		}
		if h.val[child] >= h.val[par] {
			break
		}
		h.val[par], h.val[child] = h.val[child], h.val[par]
		par = child
	}
}
func (h *TopK) Insert(x int) {
	if h.size < len(h.val) {
		h.val[h.size] = x
		h.size++
		h.heapifyUp()
		return
	}
	if h.val[0] < x {
		h.val[0] = x
		h.heapifyDown()
	}
}
func (h *TopK) Sum() int {
	sum := 0
	for _, v := range h.val {
		sum += v
	}
	return sum
}
