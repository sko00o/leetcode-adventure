package practice

// Brute Force: Recursion
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

// Recursion with Memoization
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

// 2D Dynamic Programming
// 二维动态规划
// dp[i][j] 表示从第 i 个索引能到达总和为 j 的方案数量，有如下递推关系：
// dp[i][sum+nums[i]] += dp[i-1][sum]
// dp[i][sum-nums[i]] += dp[i-1][sum]
// 由题目条件可知总和 sum 的范围是 [-1000,1000] ，将 sum + 1000，便于用作数组索引。
// Runtime: 8 ms
// Memory Usage: 6.6 MB
func findTargetSumWays2(nums []int, S int) int {
	dp := make([][2001]int, len(nums))
	for i, num := range nums {
		if i == 0 {
			dp[i][num+1000]++
			dp[i][-num+1000]++
			continue
		}
		for sum := -1000; sum <= 1000; sum++ {
			if dp[i-1][sum+1000] > 0 {
				dp[i][sum+1000+num] += dp[i-1][sum+1000]
				dp[i][sum+1000-num] += dp[i-1][sum+1000]
			}
		}
	}

	if S > 1000 || S < -1000 {
		return 0
	}
	if len(nums) == 0 {
		return 1
	}
	return dp[len(nums)-1][S+1000]
}

// 1D Dynamic Programming
// 观察上一个方案，只有最后一行 dp 是有用的。所以我们可以用一维动态规划数组来节省
// 空间。唯一的区别是每遍历一行就要更新同一个 dp 数组。
// Runtime: 8 ms
// Memory Usage: 2.3 MB
func findTargetSumWays31(nums []int, S int) int {
	var dp [2001]int
	for i, num := range nums {
		if i == 0 {
			dp[num+1000]++
			dp[-num+1000]++
			continue
		}
		var next [2001]int
		for sum := -1000; sum <= 1000; sum++ {
			if dp[sum+1000] > 0 {
				next[sum+num+1000] += dp[sum+1000]
				next[sum-num+1000] += dp[sum+1000]
			}
		}
		dp = next
	}

	if S > 1000 || S < -1000 {
		return 0
	}
	if len(nums) == 0 {
		return 1
	}
	return dp[S+1000]
}

// Runtime: 60 ms
// Memory Usage: 6.5 MB
func findTargetSumWays32(nums []int, S int) int {
	dp := make(map[int]int)
	dp[0] = 1
	for _, num := range nums {
		ndp := make(map[int]int)
		for n, cnt := range dp {
			if cnt > 0 {
				ndp[n+num] += cnt
				ndp[n-num] += cnt
			}
		}
		dp = ndp
	}
	return dp[S]
}

// Runtime: 0 ms
// Memory Usage: 2.4 MB
func findTargetSumWays33(nums []int, S int) int {
	// 全取正值
	var s int
	for _, num := range nums {
		s += num
	}
	if s < S || (s+S)%2 != 0 {
		return 0
	}

	target := (s + S) / 2
	// dp[i] 数组中和为 i 的数量。
	dp := make([]int, target+1)
	dp[0] = 1
	for _, num := range nums {
		for i := target; i >= num; i-- {
			dp[i] += dp[i-num]
		}
	}

	return dp[target]
}

// Runtime: 0 ms
// Memory Usage: 2.4 MB
func findTargetSumWays34(nums []int, S int) int {
	s := S
	for _, num := range nums {
		s += num
	}
	target := s / 2
	if target < S || s%2 != 0 {
		return 0
	}

	// dp[i] 数组中和为 i 的数量。
	dp := make([]int, target+1)
	dp[0] = 1
	for _, num := range nums {
		for i := target; i >= num; i-- {
			dp[i] += dp[i-num]
		}
	}

	return dp[target]
}
