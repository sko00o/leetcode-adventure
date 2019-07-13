package prac

func twoSum(numbers []int, target int) []int {
	for idx1, n := range numbers {
		if idx2 := index(numbers, target-n, idx1); idx2 != -1 {
			return []int{idx1 + 1, idx2 + 1}
		}
	}
	return []int{-1, -1}
}

func index(nums []int, target int, skipIdx int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)>>1
		if nums[m] == target {
			if m == skipIdx {
				if m-1 >= 0 && nums[m-1] == target {
					return m - 1
				} else if m+1 < len(nums) && nums[m+1] == target {
					return m + 1
				}
				return -1
			}
			return m
		} else if nums[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return -1
}

func twoSum1(nums []int, target int) []int {
	m := make(map[int]int)
	for idx, num := range nums {
		if idx1, ok := m[target-num]; ok && idx != idx1 {
			return []int{idx1 + 1, idx + 1}
		}
		m[num] = idx
	}
	return nil
}
