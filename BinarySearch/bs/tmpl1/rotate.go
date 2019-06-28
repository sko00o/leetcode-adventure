package tmpl1

// upgrade from `Binary Search Template I`
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1
	for left <= right {
		// Prevent (left + right) overflow
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			if target > nums[right] {
				// search left
				if t := search(nums[left:mid], target); t != -1 {
					return t + left
				}
			}
			left = mid + 1
		} else { // nums[mid] > target
			if target < nums[left] {
				// search right
				if t := search(nums[mid+1:right+1], target); t != -1 {
					return t + mid + 1
				}
			}
			right = mid - 1
		}
	}

	// End Condition: left > right
	return -1
}
