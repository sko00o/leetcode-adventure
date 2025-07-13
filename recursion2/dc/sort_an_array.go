package dc

// Given an array of integers nums, sort the array in ascending order and return it.
// You must solve the problem without using any built-in functions in O(nlog(n)) time complexity
// and with the smallest space complexity possible.
func sortArray(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	clonedNums := append([]int{}, nums...)
	OptimizedMergeSort(clonedNums)
	return clonedNums
}

// time complexity: O(NlogN)
// space complexity: O(1)
func OptimizedMergeSort(nums []int) {
	// using insert sort in small blocks
	const blockSize = 3
	var i int
	for i+blockSize < len(nums) {
		insertSort(nums[i : i+blockSize])
		i += blockSize
	}
	insertSort(nums[i:])

	// bottom-up merge sort
	for step := blockSize; step < len(nums); step <<= 1 {
		for i := 0; i+step <= len(nums); i += step << 1 {
			end := i + (step << 1)
			if end > len(nums) {
				end = len(nums)
			}
			merge(nums[i:end], step)
		}
	}
}

func insertSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		deal := nums[i]
		var j int
		for j = i; j > 0 && deal < nums[j-1]; j-- {
			nums[j] = nums[j-1]
		}
		if j != i {
			nums[j] = deal
		}
	}
}

// in-place merge
func merge(arr []int, mid int) {
	l, r := 0, mid
	for l < r && r < len(arr) {
		for l < r && arr[l] <= arr[r] {
			l++
		}
		mid := r
		for r < len(arr) && arr[r] <= arr[l] {
			r++
		}
		rotate(arr[l:r], mid-l)
		l += r - mid
	}
}

// rotate
// input:  arr=[a,b,C,D], mid=2
// output: arr=[C,D,a,b]
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
