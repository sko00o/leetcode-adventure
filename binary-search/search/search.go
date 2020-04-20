package search

func search(nums []int, target int) int {
	maxIdx := len(nums) - 1
	for l, h := 0, maxIdx; l <= h; {
		if l >= 0 && l <= maxIdx && nums[l] == target {
			return l
		}
		if h <= maxIdx && h >= 0 && nums[h] == target {
			return h
		}

		m := l + ((h - l) >> 1)
		mid := nums[m]
		if mid == target {
			return m
		} else if mid > target {
			l = l + 1
			h = m - 1
		} else {
			l = m + 1
			h = h - 1
		}
	}
	return -1
}
