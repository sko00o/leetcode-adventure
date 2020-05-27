package sort

/*
Notes:
切分再合并两个有序数组。
自顶向下合并，递归的合并左右两个有序数组。
自底向上合并，相邻两个有序数组进行合并。
稳定排序。

优化：
1. 小规模数据用直接插入排序。
2. 原地排序，省去辅助数组。使用一种反转算法 —— 手摇算法。
见 mergeSort2
*/

// top to bottom
// time complexity: O(NlogN)
// space complexity: O(N)
func mergeSort(nums []int) {
	if len(nums) > 1 {
		mid := len(nums) / 2
		mergeSort(nums[:mid])
		mergeSort(nums[mid:])
		merge(nums, mid)
	}
}

// time complexity: O(N)
// space complexity: O(N)
func merge(arr []int, mid int) {
	result := make([]int, 0, len(arr))
	l, r := 0, mid
	for l < mid && r < len(arr) {
		if arr[l] < arr[r] {
			result = append(result, arr[l])
			l++
			continue
		}

		result = append(result, arr[r])
		r++
	}
	result = append(result, arr[l:mid]...)
	result = append(result, arr[r:]...)

	copy(arr, result)
}

// bottom to top
// time complexity: O(NlogN)
// space complexity: O(N)
func mergeSort1(nums []int) {
	for step := 1; step < len(nums); step *= 2 {
		for i := 0; i+step <= len(nums); i += 2 * step {
			var end = i + step*2
			if end > len(nums) {
				end = len(nums)
			}
			merge(nums[i:end], step)
		}
	}
}

// time complexity: O(NlogN)
// space complexity: O(1)
func mergeSort2(nums []int) {
	// insert sort for small blocks
	var blockSize = 3
	var i int
	for i+blockSize < len(nums) {
		insertSort(nums[i : i+blockSize])
		i += blockSize
	}
	insertSort(nums[i:])

	for step := blockSize; step < len(nums); step *= 2 {
		for i := 0; i+step <= len(nums); i += 2 * step {
			var end = i + step*2
			if end > len(nums) {
				end = len(nums)
			}
			merge1(nums[i:end], step)
		}
	}
}

// in-place merge
func merge1(arr []int, mid int) {
	var i, j = 0, mid
	for i < j && j < len(arr) {
		for i < j && arr[i] <= arr[j] {
			i++
		}

		mid := j
		for j < len(arr) && arr[j] <= arr[i] {
			j++
		}

		rotate(arr[i:j], mid-i)
		i += j - mid
	}
}

// rotate: arr[a,b,C,D],mid=2 - > arr[C,D,a,b]
func rotate(arr []int, mid int) {
	if mid < 0 || mid >= len(arr) {
		return
	}

	reverse(arr[:mid])
	reverse(arr[mid:])
	reverse(arr)
}

func reverse(arr []int) {
	for l, i := len(arr), 0; i < l-i-1; i++ {
		j := l - i - 1
		arr[i], arr[j] = arr[j], arr[i]
	}
}
