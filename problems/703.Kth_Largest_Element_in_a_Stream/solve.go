package problems

// KthLargest is a min heap.
type KthLargest struct {
	nums []int
	cap  int
	size int
}

// Constructor return a fixed size min heap.
func Constructor(k int, nums []int) KthLargest {
	var size = len(nums)
	if k <= size {
		for i := k / 2; i >= 0; i-- {
			minHeapify(nums, i, k)
		}
		for i := len(nums) - 1; i >= k; i-- {
			if nums[0] > nums[i] {
				continue
			}

			nums[0] = nums[i]
			minHeapify(nums, 0, k)
		}
		size = k
	} else {
		nnums := make([]int, k)
		copy(nnums, nums)
		nums = nnums
	}

	return KthLargest{
		nums: nums,
		cap:  k,
		size: size,
	}
}

// Add value in heap.
func (kl *KthLargest) Add(val int) int {
	if kl.size < kl.cap {
		kl.nums[kl.size] = val
		kl.size++
	} else if val > kl.nums[0] {
		kl.nums[0] = val
	}

	minHeapify(kl.nums, 0, kl.size)
	return kl.nums[0]
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
