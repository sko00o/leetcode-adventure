package practice_hard

import (
	"sort"
)

// https://leetcode.com/explore/learn/card/binary-search/146/more-practices-ii/1041/discuss/109082/Approach-the-problem-using-the-%22trial-and-error%22-algorithm

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	low, high, ans := 0, nums[len(nums)-1]-nums[0], -1
	for low <= high {
		ct := 0
		mid := low + (high-low)>>1

		for i, j := 0, 0; i < len(nums); i++ {
			for j < len(nums) && nums[j]-nums[i] <= mid {
				j++
			}
			ct += j - i - 1
		}

		if ct >= k {
			ans = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return ans
}
