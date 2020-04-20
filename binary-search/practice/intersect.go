package practice

import (
	"sort"
)

// Each element in the result should appear as many times as it shows in both arrays.

func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for _, n := range nums1 {
		m[n]++
	}

	out := make([]int, 0)
	for _, n := range nums2 {
		if m[n] > 0 {
			m[n]--
			out = append(out, n)
		}
	}
	return out
}

/*
Think about 3rd follow up question:
What if elements of nums2 are stored on disk, and the memory is
limited such that you cannot load all elements into the memory
at once?

External sort (https://en.wikipedia.org/wiki/External_sorting)
MapReduce (https://en.wikipedia.org/wiki/MapReduce)
*/

func intersect1(nums1 []int, nums2 []int) []int {
	// longer slice will be nums2.
	if len(nums2) < len(nums1) {
		nums1, nums2 = nums2, nums1
	}

	sort.Ints(nums1)
	sort.Ints(nums2)

	out := make([]int, 0)
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] == nums2[j] {
			out = append(out, nums1[i])
			i++
			j++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			i++
		}
	}
	return out
}
