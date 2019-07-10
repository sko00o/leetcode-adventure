package problems

// Brute Force
// time: O(n^2)
// space: O(1)
func twoSum(nums []int, target int) []int {
	for idx, num := range nums {
		expect := target - num
		for idx1 := idx + 1; idx1 < len(nums); idx1++ {
			if nums[idx1] == expect {
				return []int{idx, idx1}
			}
		}
	}
	return nil
}

// Two-pass Hash Table
// time: O(n)
// space: O(n)
func twoSum1(nums []int, target int) []int {
	m := make(map[int]int)
	for idx, num := range nums {
		m[num] = idx
	}
	for idx, num := range nums {
		if idx1, ok := m[target-num]; ok && idx != idx1 {
			return []int{idx, idx1}
		}
	}
	return nil
}

// One-pass Hash Table
// time: O(n)
// space: O(n)
func twoSum2(nums []int, target int) []int {
	m := make(map[int]int)
	for idx, num := range nums {
		if idx1, ok := m[target-num]; ok && idx != idx1 {
			return []int{idx1, idx}
		}
		m[num] = idx
	}
	return nil
}
