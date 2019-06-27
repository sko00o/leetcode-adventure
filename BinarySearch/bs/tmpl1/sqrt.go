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

/*func mySqrt1(x int) int {
	r := x
	for r*r > x {
		r = (r + x/r) / 2
	}
	return r
}*/

/*func mySqrt2(x int) int {
	low := 1
	heigh := x/2+1
	mid := 1
	for{
		if low > heigh{
			break
		}
		mid = (low + heigh)/2
		temp := mid*mid
		if temp == x{
			return mid
		}else if temp > x{
			heigh = mid -1
		}else{
			low = mid + 1
		}
	}
	return heigh
}*/

// TODO:
func NewtonMethod(int) int {
	return 0
}
