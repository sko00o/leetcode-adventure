package heap

// complexity: O(logN)
func maxHeapify(arr []int, curr, end int) {
	// left and right childs
	lchild, rchild := 2*curr+1, 2*curr+2

	// find max node
	var maxIdx = curr
	if lchild < end && arr[lchild] > arr[curr] {
		maxIdx = lchild
	}
	if rchild < end && arr[rchild] > arr[maxIdx] {
		maxIdx = rchild
	}

	if maxIdx != curr {
		arr[curr], arr[maxIdx] = arr[maxIdx], arr[curr]
		maxHeapify(arr, maxIdx, end)
	}
}

// iteratively
func maxHeapify1(arr []int, start, end int) {
	root := start
	child := 2*root + 1
	for child < end {
		if child+1 < end && arr[child] < arr[child+1] {
			child++
		}

		if arr[root] >= arr[child] {
			return
		}

		arr[root], arr[child] = arr[child], arr[root]
		root = child
		child = 2*root + 1
	}
}

// complexity: O(N)
func buildMaxHeap(arr []int) {
	n := len(arr)
	for i := n / 2; i >= 0; i-- {
		maxHeapify(arr, i, n)
	}
}

func minHeapify(arr []int, curr, end int) {
	// left and right childs
	lchild, rchild := 2*curr+1, 2*curr+2

	// find min node
	var minIdx = curr
	if lchild < end && arr[lchild] < arr[curr] {
		minIdx = lchild
	}
	if rchild < end && arr[rchild] < arr[minIdx] {
		minIdx = rchild
	}

	if minIdx != curr {
		arr[curr], arr[minIdx] = arr[minIdx], arr[curr]
		minHeapify(arr, minIdx, end)
	}
}

// iteratively
func minHeapify1(arr []int, start, end int) {
	root := start
	child := 2*root + 1
	for child < end {
		if child+1 < end && arr[child] > arr[child+1] {
			child++
		}

		if arr[root] <= arr[child] {
			return
		}

		arr[root], arr[child] = arr[child], arr[root]
		root = child
		child = 2*root + 1
	}
}

func buildMinHeap(arr []int) {
	n := len(arr)
	for i := n / 2; i >= 0; i-- {
		minHeapify(arr, i, n)
	}
}

func heapSort(arr []int) {
	buildMaxHeap(arr)
	for i := len(arr) - 1; i > 0; i-- {
		// move max to the end
		arr[0], arr[i] = arr[i], arr[0]
		maxHeapify(arr, 0, i)
	}
}

func heapSort1(arr []int) {
	n := len(arr)
	for i := n / 2; i >= 0; i-- {
		maxHeapify1(arr, i, n)
	}
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		maxHeapify1(arr, 0, i)
	}
}

func heapSortRevert(arr []int) {
	buildMinHeap(arr)
	for i := len(arr) - 1; i >= 1; i-- {
		// move min to the end
		arr[0], arr[i] = arr[i], arr[0]
		minHeapify(arr, 0, i)
	}
}

func heapSortRevert1(arr []int) {
	n := len(arr)
	for i := n / 2; i >= 0; i-- {
		minHeapify1(arr, i, n)
	}
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		minHeapify1(arr, 0, i)
	}
}
