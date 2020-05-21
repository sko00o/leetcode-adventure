package problems

func sortArrayByParity(A []int) []int {
	for i, j := 0, len(A)-1; i < j; {
		if A[i]%2 == 0 {
			i++
			continue
		}
		if A[j]%2 != 0 {
			j--
			continue
		}

		A[i], A[j] = A[j], A[i]
		i++
		j--
	}
	return A
}

func sortArrayByParity1(A []int) []int {
	for i, j := 0, len(A)-1; i < j; {
		if A[i]%2 > A[j]%2 {
			A[i], A[j] = A[j], A[i]
		}
		if A[i]%2 == 0 {
			i++
		}
		if A[j]%2 == 1 {
			j--
		}
	}
	return A
}
