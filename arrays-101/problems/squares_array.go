package problems

import "sort"

// Runtime: 36 ms
// Memory Usage: 6.5 MB
// time complexity: O(N*log(N))
// space complexity: O(N)
func sortedSquares(A []int) []int {
	for i := range A {
		A[i] = A[i] * A[i]
	}
	sort.Sort(sort.IntSlice(A))
	return A
}

// Two Pointer
// time complexity: O(N)
// space complexity: O(N)
func sortedSquares1(A []int) []int {
	var res = make([]int, len(A))

	abs := func(n int) int {
		if n < 0 {
			n = -n
		}
		return n
	}

	for i, j, k := 0, len(A)-1, len(A)-1; i <= j; k-- {
		if abs(A[i]) < abs(A[j]) {
			res[k] = A[j] * A[j]
			j--
		} else {
			res[k] = A[i] * A[i]
			i++
		}
	}
	return res
}

func sortedSquares2(A []int) []int {
	var res = make([]int, len(A))

	for i, j, k := 0, len(A)-1, len(A)-1; i <= j; k-- {
		if A[i]*A[i] < A[j]*A[j] {
			res[k] = A[j] * A[j]
			j--
		} else {
			res[k] = A[i] * A[i]
			i++
		}
	}
	return res
}
