package problems

import (
	"sort"
)

// Brute Force, compare after sort.
func heightChecker(heights []int) int {
	var sorted = make([]int, len(heights))
	copy(sorted, heights)
	sort.Ints(sorted)

	var cnt int
	for i := range heights {
		if heights[i] != sorted[i] {
			cnt++
		}
	}
	return cnt
}

// bucket sort
func heightChecker1(heights []int) int {
	// 1 <= heights[i] <= 100
	hCnt := make([]int, 101)

	for _, n := range heights {
		if n >= 1 && n <= 100 {
			hCnt[n]++
		}
	}

	var cnt, curr int
	for _, n := range heights {
		for hCnt[curr] == 0 {
			curr++
		}

		if curr != n {
			cnt++
		}
		hCnt[curr]--
	}
	return cnt
}
