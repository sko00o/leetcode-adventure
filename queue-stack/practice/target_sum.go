package practice

func findTargetSumWays(nums []int, S int) int {
	return dfs(nums, S)
}

func dfs(nums []int, S int) int {
	if len(nums) == 0 {
		if S == 0 {
			return 1
		} else {
			return 0
		}
	}

	return dfs(nums[1:], S-nums[0]) + dfs(nums[1:], S+nums[0])
}
