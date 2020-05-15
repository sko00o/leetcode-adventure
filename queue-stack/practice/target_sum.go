package practice

// Runtime: 628 ms
// Memory Usage: 2.1 MB
func findTargetSumWays(nums []int, S int) int {
	if len(nums) == 0 {
		if S == 0 {
			return 1
		}
		return 0
	}

	f := findTargetSumWays
	return f(nums[1:], S+nums[0]) + f(nums[1:], S-nums[0])
}

// Runtime: 96 ms
// Memory Usage: 7 MB
func findTargetSumWays1(nums []int, S int) int {
	return dfs(nums, S, make(map[[2]int]int))
}

// DFS with memorize
func dfs(nums []int, S int, mem map[[2]int]int) int {
	if len(nums) == 0 {
		if S == 0 {
			return 1
		}
		return 0
	}
	curr := [2]int{len(nums), S}
	if n, ok := mem[curr]; ok {
		return n
	}
	mem[curr] = dfs(nums[1:], S-nums[0], mem) + dfs(nums[1:], S+nums[0], mem)
	return mem[curr]
}
