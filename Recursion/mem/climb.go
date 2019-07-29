package mem

func climbStairs(n int) int {
	return fib(n + 1)
}

func climbStairs1(n int) int {
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
