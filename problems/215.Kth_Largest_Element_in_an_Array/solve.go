package problems

import (
	"github.com/sko00o/leetcode-adventure/sort/heap"
)

/*
Notes:
○ 找到数组中第 k 大的的数
	§ 排序
	§ 维护一个容量为 k 的小顶堆，最后返回堆顶元素，
	  时间复杂度 O(Nlogk)，空间复杂度 O(k)
	§ 快速选择算法 findKthLargest2 findKthLargest3
*/

// use a min heap with fixed size k
func findKthLargest(nums []int, k int) int {
	h := heap.NewMinHeap(k + 1)
	for _, n := range nums {
		h.Push(n)
		if h.Size() > k {
			h.Pop()
		}
	}
	return h.Top()
}

func findKthLargest1(nums []int, k int) int {
	// build a heap with fixed size k
	for i := k / 2; i >= 0; i-- {
		minHeapify(nums, i, k)
	}

	// put rest in heap
	for i := len(nums) - 1; i >= k; i-- {
		if nums[i] < nums[0] {
			continue
		}

		nums[0] = nums[i]
		minHeapify(nums, 0, k)
	}
	return nums[0]
}

func minHeapify(nums []int, start, end int) {
	root := start
	child := 2*root + 1
	for child < end {
		if child+1 < end && nums[child] > nums[child+1] {
			child++
		}

		if nums[root] <= nums[child] {
			break
		}

		nums[root], nums[child] = nums[child], nums[root]
		root = child
		child = 2*root + 1
	}
}
