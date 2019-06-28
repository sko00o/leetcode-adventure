package tmpl1

func mySqrt(x int) int {
	// x should be non-negative integer
	if x <= 1 {
		return x
	}

	l, h := 1, x
	for l < h {
		m := l + ((h - l) >> 1)
		mt := m * m
		if mt == x {
			return m
		} else if mt > x || mt <= 0 { // mt <= 0 means overflow
			h = m - 1
		} else {
			l = m + 1
		}
	}

	if l*l > x {
		return l - 1
	}
	return l
}

// Newton's Method
// f(x) = x^m - a
// x(n+1) = x(n) - (x(n) - a/(x(x)^(m-1)))/m
// m = 2
// x(n+1) = (x(n) + a/x(n))/2
func mySqrt1(x int) int {
	if x < 2 {
		return x
	}
	r := x >> 1
	// prevent overflow
	for r > x/r {
		r = (r + x/r) >> 1
	}
	return r
}
