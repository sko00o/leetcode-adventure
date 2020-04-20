package prac

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left] < nums[right] {
			return nums[left]
		} else {
			mid := left + (right-left)>>1
			if nums[mid] > nums[right] {
				left = mid + 1
			} else if nums[mid] < nums[left] {
				right = mid
			} else {
				// this condition for
				// `Find Minimum in Rotated Sorted Array II`
				// nums[left] == nums[mid] == nums[right]
				left++
				right--
			}
		}
	}
	return nums[left]
}

func findMin1(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left] < nums[right] {
			return nums[left]
		} else {
			mid := left + (right-left)>>1
			if nums[mid] > nums[right] {
				left = mid + 1
			} else if nums[mid] == nums[right] {
				right--
			} else {
				right = mid
			}
		}
	}
	return nums[left]
}
