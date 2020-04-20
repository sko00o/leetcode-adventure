package tmpl3

// try to use different method to solve this
func findPeakElement(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1

	for left+1 < right {
		mid := left + (right-left)>>1
		switch getType(nums, mid) {
		case Peak:
			return mid
		case Up:
			left = mid
		case Down:
			right = mid
		}
	}

	if getType(nums, left) == Peak {
		return left
	}
	if getType(nums, right) == Peak {
		return right
	}

	return -1 // will not arrive
}

func greater(nums []int, p1, p2 int) bool {
	if p1 <= -1 || p1 >= len(nums) {
		return false
	}
	if p2 <= -1 || p2 >= len(nums) {
		return true
	}

	return nums[p1] > nums[p2]
}

func getType(nums []int, mid int) Type {
	b1 := greater(nums, mid, mid-1)
	b2 := greater(nums, mid, mid+1)
	if b1 {
		if b2 {
			return Peak
		} else {
			return Up
		}
	} else {
		return Down
	}
}

type Type int

const (
	Peak Type = iota
	Up
	Down
)

// time O(n) is not good
func findPeakElement1(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	for i := 0; i < len(nums); {
		if greater(nums, i, i-1) && greater(nums, i, i+1) {
			return i
		} else {
			if greater(nums, i, i+1) {
				i = i + 2
			} else {
				i++
			}
		}
	}

	return -1 // will not arrive
}
