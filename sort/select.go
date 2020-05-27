package sort

/*
Notes:
简述，每轮遍历未排序区间，取最小值，交换到未排区间头部，缩小未排序区间。
优化做法：每次遍历，找出最大值和最小值，并移到正确的位置。
不稳定排序。
*/

func selectSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		minIdx := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIdx] {
				minIdx = j
			}
		}
		if minIdx != i {
			nums[i], nums[minIdx] = nums[minIdx], nums[i]
		}
	}
}

func selectSort1(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		minIdx := i
		maxIdx := i // optimizae
		for j := i + 1; j < len(nums)-i; j++ {
			if nums[j] > nums[maxIdx] {
				maxIdx = j
				continue
			}
			if nums[j] < nums[minIdx] {
				minIdx = j
			}
		}

		last := len(nums) - i - 1
		if maxIdx != i {
			nums[i], nums[minIdx] = nums[minIdx], nums[i]
			nums[last], nums[maxIdx] = nums[maxIdx], nums[last]
		} else {
			if minIdx != last {
				nums[last], nums[maxIdx] = nums[maxIdx], nums[last]
				nums[i], nums[minIdx] = nums[minIdx], nums[i]
			} else {
				nums[minIdx], nums[maxIdx] = nums[maxIdx], nums[minIdx]
			}
		}
	}
}
