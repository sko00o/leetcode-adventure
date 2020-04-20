package ana

// non-tail-recursion
func sum(ls []int) int {
	if len(ls) == 0 {
		return 0
	}
	// not a tail recursion because it does some computation after the recursive
	// call returned.
	return ls[0] + sum(ls[1:])
}

// tail-recursion
func sum1(ls []int) int {
	var helper func([]int, int) int
	helper = func(ls []int, acc int) int {
		if len(ls) == 0 {
			return acc
		}

		// this is a tail recursion because the final instruction is a recursive
		// call.
		return helper(ls[1:], ls[0]+acc)
	}
	return helper(ls, 0)
}
