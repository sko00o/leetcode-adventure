package con

// recursive
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		return 1 / myPow(x, -n)
	}

	r := myPow(x, n>>1)
	return r * r * myPow(x, n-n>>1<<1)
}

// non-recursive
func myPow1(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	neg, res := false, float64(1)
	if n < 0 {
		neg = true
		n = -n
	}

	// this `i` can be reused
	// so declare outside of for-loop,
	// maybe save 4kb here.
	var i int
	for n > 0 {
		p := x
		for i = 1; i<<1 < n; i <<= 1 {
			p *= p
		}
		n -= i
		res *= p
	}

	if neg {
		return 1 / res
	}
	return res
}
