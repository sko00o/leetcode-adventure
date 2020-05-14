package practice

func dailyTemperatures(T []int) []int {
	out := make([]int, len(T))
	// no newly allocated array in append
	// will get better performance.
	stack := make([]int, 0, len(T))
	for idx, t := range T {
		for len(stack) != 0 {
			ls := len(stack)
			last := stack[ls-1]
			if T[last] >= t {
				break
			}

			stack = stack[:ls-1]
			out[last] = idx - last
		}

		stack = append(stack, idx)
	}
	return out
}
