package tmpl3

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	left, right := 0, len(nums)-1
	for left+1 < right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return gotRange(nums, mid)
		} else if nums[mid] < target {
			left = mid
		} else {
			right = mid
		}
	}

	if nums[left] == target {
		return gotRange(nums, left)
	}
	if nums[right] == target {
		return gotRange(nums, right)
	}

	return []int{-1, -1}
}

func gotRange(nums []int, index int) []int {
	left, right := index, index
	for left >= 0 && nums[left] == nums[index] {
		left--
	}
	for right < len(nums) && nums[right] == nums[index] {
		right++
	}

	return []int{left + 1, right - 1}
}
