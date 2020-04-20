package prac_hard

/*
Note:

You must not modify the array (assume the array is read only).
-- no sort
You must use only constant, O(1) extra space.
-- no hash map
Your runtime complexity should be less than O(n2).
-- use binary search
There is only one duplicate number in the array, but it could be repeated more than once.

*/

// https://en.wikipedia.org/wiki/Pigeonhole_principle
func findDuplicate(nums []int) int {
	l, r := 1, len(nums)-1
	for l < r {
		m := l + (r-l)>>1

		// count how many element
		// smaller than m
		ct := 0
		for _, n := range nums {
			if n <= m {
				ct++
			}
		}

		// count > m means
		// duplicate number
		// in [l,m] range
		// otherwise in
		// [m+1,r] range
		if ct > m {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

/*
very cool solution at:
https://leetcode.com/explore/learn/card/binary-search/146/more-practices-ii/1039/discuss/72846/My-easy-understood-solution-with-O(n)-time-and-O(1)-space-without-modifying-the-array.-With-clear-explanation.

convert to linked list problem
nums[i]=j means node i next is node j
duplicate element in slice means
different node have same next node
it is cycle, duplicate number must be
entry point of cycle.
we can use two point tech to find it.

*/

func findDuplicate1(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	// It nums have cycle
	sp, fp := nums[0], nums[nums[0]]
	for sp != fp {
		sp, fp = nums[sp], nums[nums[fp]]
	}

	sp = 0
	for fp != sp {
		sp, fp = nums[sp], nums[fp]
	}
	return fp
}
