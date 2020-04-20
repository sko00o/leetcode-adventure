package tmpl2

func findMin(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		return nums[0]
	}

	l, r := 0, len(nums)-1
	for l < r {
		m := l + (r-l)>>1
		if m == 0 {
			if nums[m+1] > nums[m] {
				return nums[m]
			}
		}
		if m == len(nums)-1 {
			if nums[m] < nums[m-1] {
				return nums[m]
			}
		}

		if nums[m] < nums[l] { // min in left of m
			r = m
		} else if nums[m] < nums[r] {
			r = m
		} else if nums[m] > nums[r] {
			l = m + 1
		}

	}

	if l > 0 && l < len(nums) {
		return nums[l]
	}
	return -1
}
