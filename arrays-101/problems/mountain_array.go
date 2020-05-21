package problems

func validMountainArray(A []int) bool {
	if len(A) < 3 {
		return false
	}

	var hasPeak bool

	for i := 1; i < len(A); i++ {
		if A[i] > A[i-1] {
			if hasPeak {
				return false
			}
		} else if A[i] < A[i-1] {
			if i-1 == 0 {
				return false
			}
			if !hasPeak {
				hasPeak = true
			}
		} else {
			return false
		}
	}

	return hasPeak
}

// one pass
func validMountainArray1(A []int) bool {
	var aLen = len(A)
	var i int

	// walk up
	for i+1 < aLen && A[i] < A[i+1] {
		i++
	}

	if i == 0 || i == aLen-1 {
		return false
	}

	// walk down
	for i+1 < aLen && A[i] > A[i+1] {
		i++
	}

	return i == aLen-1
}
