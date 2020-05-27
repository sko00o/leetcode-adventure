package sort

import (
	"github.com/sko00o/leetcode-adventure/queue-stack/stack"
)

/*
Notes:
选定枢轴，将数组分成大于枢轴和小于枢轴的两个区间，递归完成排序。
优化：
1. 小规模时用直接插入排序
2. 对于大量重复值的优化：三向切分，维护小于/等于/大于枢轴的三个区间。见 quickSort1 。
3. 栈空间优化：递归优化，先排序较小区间。见 quickSort2 。
非递归版本：使用辅助栈。见 quickSort3 。
*/

func quickSort(nums []int) {
	if len(nums) == 0 {
		return
	}

	loc := partition(nums)
	quickSort(nums[:loc])
	quickSort(nums[loc+1:])
}

// partition will return pivot index.
func partition(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	// use first one as pivot
	pivot := arr[0]

	// move it to the end
	arr[0], arr[len(arr)-1] = arr[len(arr)-1], arr[0]

	var sIdx int
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < pivot {
			arr[sIdx], arr[i] = arr[i], arr[sIdx]
			sIdx++
		}
	}

	// move pivot to its final place
	arr[sIdx], arr[len(arr)-1] = arr[len(arr)-1], arr[sIdx]
	return sIdx
}

func quickSort1(nums []int) {
	if len(nums) <= 4 {
		insertSort(nums)
		return
	}

	lIdx, gIdx := partition1(nums)
	quickSort1(nums[:lIdx])
	quickSort1(nums[gIdx:])
}

func partition1(arr []int) (int, int) {
	if len(arr) == 0 {
		return -1, -1
	}

	pivot := arr[0]

	lIdx, gIdx := 0, len(arr)
	for i := 0; i < gIdx; {
		if arr[i] < pivot {
			arr[lIdx], arr[i] = arr[i], arr[lIdx]
			lIdx++
			i++
		} else if arr[i] > pivot {
			arr[gIdx-1], arr[i] = arr[i], arr[gIdx-1]
			gIdx--
		} else {
			i++
		}
	}

	return lIdx, gIdx
}

func quickSort2(nums []int) {
	for begin, end := 0, len(nums); begin < end; {
		offset := partition(nums[begin:end])
		loc := begin + offset
		if offset < end-loc {
			quickSort(nums[begin:loc])
			begin = loc + 1
		} else {
			quickSort(nums[loc+1 : end])
			end = loc
		}
	}
}

func quickSort3(nums []int) {
	stk := new(stack.LinkStack)

	// push `end` and `begin` in stack
	stk.Push(len(nums))
	stk.Push(0)

	for !stk.IsEmpty() {
		begin := stk.Top().(int)
		stk.Pop()
		end := stk.Top().(int)
		stk.Pop()

		loc := begin + partition(nums[begin:end])

		// save left range
		if loc+1 < end {
			stk.Push(end)
			stk.Push(loc + 1)
		}

		// save right range
		if begin < loc-1 {
			stk.Push(loc - 1)
			stk.Push(begin)
		}
	}
}

func quickSort4(nums []int) {
	stk := new(stack.LinkStack)

	// push `end` and `begin` in stack
	stk.Push(len(nums))
	stk.Push(0)

	for !stk.IsEmpty() {
		begin := stk.Top().(int)
		stk.Pop()
		end := stk.Top().(int)
		stk.Pop()

		if begin < end {
			offset := partition(nums[begin:end])
			loc := begin + offset

			// left range is small, push smaller one first
			if offset < end-loc {
				if begin < loc {
					stk.Push(loc)
					stk.Push(begin)
				}
				if loc+1 < end {
					stk.Push(end)
					stk.Push(loc + 1)
				}
			} else {
				if loc+1 < end {
					stk.Push(end)
					stk.Push(loc + 1)
				}
				if begin < loc {
					stk.Push(loc)
					stk.Push(begin)
				}
			}
		}
	}
}
