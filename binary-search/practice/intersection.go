package practice

import (
	"sort"
)

// Each element in the result must be unique.

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for _, n := range nums1 {
		m[n] = 1
	}

	out := make([]int, 0)
	for _, n := range nums2 {
		if v, ok := m[n]; ok && v == 1 {
			m[n]++
			out = append(out, n)
		}
	}
	return out
}

func intersection1(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	if len(nums2) < len(nums1) {
		nums1, nums2 = nums2, nums1
	}

	out := make([]int, 0)
	for i := range nums1 {
		if i > 0 {
			if nums1[i] != nums1[i-1] && exist(nums2, nums1[i]) {
				out = append(out, nums1[i])
			}
		} else {
			if exist(nums2, nums1[i]) {
				out = append(out, nums1[i])
			}
		}
	}
	return out
}

func exist(nums []int, tar int) bool {
	l, r := 0, len(nums)-1
	for l < r {
		m := l + (r-l)>>1
		if nums[m] == tar {
			return true
		} else if nums[m] > tar {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return nums[l] == tar
}
