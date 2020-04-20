package tmpl2

// time O(log(n))
func findPeakElement(nums []int) int {
	if nums == nil {
		return -1
	}
	if len(nums) == 1 {
		return 0
	}

	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)>>1
		switch whichType(nums, mid) {
		case peak:
			return mid
		case up:
			left = mid + 1
		case down:
			right = mid
		}
	}

	if left < len(nums) && whichType(nums, left) == peak {
		return left
	}
	return -1
}

type T int

const (
	peak T = iota
	up
	down
)

func whichType(nums []int, idx int) T {
	if idx == 0 && nums[idx] > nums[idx+1] {
		return peak
	}
	if idx == len(nums)-1 && nums[idx] > nums[idx-1] {
		return peak
	}

	if len(nums) >= 3 {
		if nums[idx] > nums[idx-1] {
			if nums[idx] > nums[idx+1] {
				return peak
			}
			return up
		}
		return down
	}

	return down
}
