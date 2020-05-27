package sort

/*
Notes:
简单概括，相邻两元素作比较，遍历区间逐步缩小。
效率低下。
稳定排序。
*/

// time complexity: O(N) ~ O(N^2)
func bubbleSort(nums []int) {
	var swapped bool
	for i := len(nums) - 1; i > 0; i-- {
		swapped = false

		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				swapped = true
			}
		}

		// optimizae
		if !swapped {
			return
		}
	}
}
