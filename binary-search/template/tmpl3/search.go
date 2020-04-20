package tmpl3

func binarySearch(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1
	for left+1 < right {
		// Prevent (left + right) overflow
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid
		} else {
			right = mid
		}
	}

	// Post-processing:
	// End Condition: left + 1 == right
	if nums[left] == target {
		return left
	}
	if nums[right] == target {
		return right
	}
	return -1
}
