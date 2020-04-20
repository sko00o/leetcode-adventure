package mem

import "math"

// Fibonacci Number
// time O(n)
// space O(1)
func climbStairs(n int) int {
	return fibHelper(n)
}

func fibHelper(n int) int { // <=> fib(n+1)
	// n is positive integer
	if n <= 2 {
		return n
	}

	c1, c2 := 1, 2
	for i := 3; i <= n; i++ {
		c1, c2 = c2, c1+c2
	}

	return c2
}

// Dynamic Programming
// time O(n)
// space O(n)
func climbStairs1(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// Brute Force
// time O(n^2)
// space O(n)
func climbStairs2(n int) int {
	return helper(0, n)
}

func helper(i int, n int) int {
	// i defines the current step,
	// n defines the destination step.
	if i > n {
		return 0
	}
	if i == n {
		return 1
	}
	return helper(i+1, n) + helper(i+2, n)
}

// Recursion with Memoization
// time O(n)
// space O(n)
func climbStairs3(n int) int {
	m := make(map[int]int)
	return helper1(0, n, m)
}

func helper1(i int, n int, m map[int]int) int {
	if i > n {
		return 0
	}
	if i == n {
		return 1
	}
	if m[i] > 0 {
		return m[i]
	}
	m[i] = helper1(i+1, n, m) + helper1(i+2, n, m)
	return m[i]
}

// Fibonacci Formula
// time O(logn) pow method takes log n time
// space O(1)
func climbStairs4(n int) int {
	sqrt5 := math.Sqrt(5)
	fibn := math.Pow((1+sqrt5)/2, float64(n+1)) - math.Pow((1-sqrt5)/2, float64(n+1))
	return int(fibn / sqrt5)
}

// Binets Method (*)
// ref: https://leetcode.com/explore/learn/card/recursion-i/255/recursion-memoization/2377/
// This is an interesting solution which uses matrix multiplication to obtain
// the nth Fibonacci Number (F(n)).
// Q(n) = [[F(n+1),F(n)],[F(n),F(n-1)]]
// Q(1) = [[1,1],[1,0]]
// F(n) = Q(1)^(n-1)[0,0]
// time O(logn)
// space O(1)
func climbStairs5(n int) int {
	q := [][]int{{1, 1}, {1, 0}}
	res := pow(q, n)
	return res[0][0]
}

func pow(a [][]int, n int) [][]int {
	ret := [][]int{{1, 0}, {0, 1}} // this is E (identity matrix)
	for n > 0 {
		if n&1 == 1 {
			ret = multiply(ret, a) // we have E * M = M
		}
		n >>= 1
		a = multiply(a, a)
	}
	return ret
}

func multiply(a, b [][]int) [][]int {
	c := make([][]int, 2)
	for i := range c {
		c[i] = make([]int, 2)
		for j := range c[i] {
			c[i][j] = a[i][0]*b[0][j] + a[i][1]*b[1][j]
		}
	}
	return c
}
