package problems

// O(n^2) solution
func lengthOfLIS(nums []int) int {
	// dp[i] means the length of the longest increasing subsequence that ends with nums[i]
	dp := make([]int, len(nums))

	// init
	for i := range nums {
		dp[i] = 1
	}

	// dp[i] = max(dp[j] + 1, dp[i]), 0 <= j < i, nums[j] < nums[i]
	for i := range nums {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	// find the max value in dp
	maxVal := 0
	for i := range dp {
		if dp[i] > maxVal {
			maxVal = dp[i]
		}
	}
	return maxVal
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ------------------------------------------------------------

// O(nlogn) solution
func lengthOfLIS1(nums []int) int {
	// tails[i] means the smallest tail of all increasing subsequences with length i+1
	tails := make([]int, 0, len(nums))
	for _, num := range nums {
		if len(tails) == 0 || num > tails[len(tails)-1] {
			tails = append(tails, num)
		} else {
			searchAndUpdate(tails, num)
		}
	}
	return len(tails)
}

// search the first element in tails that is greater than or equal to target, and update it to target
func searchAndUpdate(tails []int, target int) {
	l, r := 0, len(tails)-1
	for l < r {
		mid := (l + r) >> 1
		if tails[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	tails[l] = target
}
