package problems

import (
	"math/rand"
	"time"

	"github.com/sko00o/leetcode-adventure/sort/heap"
)

/*
Notes:
○ 找到数组中第 k 大的的数
	§ 排序
	§ 维护一个容量为 k 的小顶堆，最后返回堆顶元素，
	  时间复杂度 O(Nlogk)，空间复杂度 O(k)
	§ 快速选择算法

Ref:
https://leetcode-cn.com/problems/kth-largest-element-in-an-array/solution/shu-zu-zhong-de-di-kge-zui-da-yuan-su-by-leetcode/

快速选择算法
1. 随机选择一个枢轴 pivot 。
2. pivot 放到数组中合适的位置 pos ，将小于 pivot 的元素移到左边，大于 pivot 的
   元素移到右边。
3. 比较 pos 和 n-k ，决定在哪边递归。
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

func init() {
	rand.Seed(time.Now().Unix())
}

func findKthLargest1(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, left, right int, kSmall int) int {
	if left == right {
		return nums[left]
	}

	pIdx := left + rand.Intn(right-left)
	pIdx = partition(nums, left, right, pIdx)

	if kSmall == pIdx {
		return nums[kSmall]
	} else if kSmall < pIdx {
		return quickSelect(nums, left, pIdx-1, kSmall)
	} else {
		return quickSelect(nums, pIdx+1, right, kSmall)
	}
}

func partition(nums []int, left, right int, pIdx int) int {
	pivot := nums[pIdx]
	// move pivot to the end
	nums[pIdx], nums[right] = nums[right], nums[pIdx]

	// move smaller elements to the left
	var sIdx = left
	for i := left; i < right; i++ {
		if nums[i] < pivot {
			nums[sIdx], nums[i] = nums[i], nums[sIdx]
			sIdx++
		}
	}

	// move pivot to its final place
	nums[right], nums[sIdx] = nums[sIdx], nums[right]

	return sIdx
}

func findKthLargest2(nums []int, k int) int {
	return quickSelect1(nums, 0, len(nums)-1, k-1)
}

func quickSelect1(nums []int, left, right int, k int) int {
	if left == right {
		return nums[left]
	}

	pIdx := left + rand.Intn(right-left)
	pIdx = partition1(nums, left, right, pIdx)

	if k == pIdx {
		return nums[k]
	} else if k < pIdx {
		return quickSelect1(nums, left, pIdx-1, k)
	} else {
		return quickSelect1(nums, pIdx+1, right, k)
	}
}

func partition1(nums []int, left, right int, pIdx int) int {
	pivot := nums[pIdx]
	// move pivot to the end
	nums[pIdx], nums[right] = nums[right], nums[pIdx]

	// move larger elements to the left
	var sIdx = left
	for i := left; i < right; i++ {
		if nums[i] > pivot {
			nums[sIdx], nums[i] = nums[i], nums[sIdx]
			sIdx++
		}
	}

	// move pivot to its final place
	nums[right], nums[sIdx] = nums[sIdx], nums[right]

	return sIdx
}
