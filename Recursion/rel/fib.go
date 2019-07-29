package rel

// quick with memorization
func fib(N int) int {
	m := make(map[int]int)
	var f func(int) int
	f = func(n int) int {
		if n <= 1 {
			return n
		}
		if v, ok := m[n]; ok {
			return v
		}
		m[n] = f(n-1) + f(n-2)
		return m[n]
	}
	return f(N)
}

// slow
func fib1(N int) int {
	if N <= 1 {
		return N
	}
	return fib1(N-1) + fib1(N-2)
}
