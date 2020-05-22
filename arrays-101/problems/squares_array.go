package problems

import (
	"sort"
)

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

// time complexity: O(N)
// space complexity: O(N)
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

// "This time, you must solve it in O(N) time and O(1) space.
// No using built-in sort, and no creating a new array."
// I tried but didn't make it.

/*
i = 0, j = len(A)-1

-4 -1  0  3 10  A[j]**=2, j--
 ^           j
-4 -1  0  3 100 tmp=A[i]*A[i], copy(A[i,j],A[i+1:j+1]), A[j]=tmp, j--
 ^        j
-1  0  3 16 100 A[j]**=2, j--
 ^     j
-1  0  9 16 100
 ^  j
 0  1  9 16 100 A[0]**=2
*/
// time complexity: O(N)~O(N^2)
// space complexity: O(1)
func sortedSquares3(A []int) []int {
	for j := len(A) - 1; j > 0; j-- {
		ii := A[0] * A[0]
		jj := A[j] * A[j]

		if ii < jj {
			A[j] = jj
			continue
		}

		copy(A[:j], A[1:j+1])
		A[j] = ii
	}
	A[0] = A[0] * A[0]

	return A
}

/*
i=0,j=first positive index
A[j] < A[j-1] ?
-4 -1  0  3 10 tmp=A[j]**2, copy(A[i+1:j+1], A[i:j]), A[i]=tmp, i++, j++
 i	   j
 0 -4 -1  3 10 tmp=A[j-1]**2, copy(A[i+1:j], A[i:j-1]), A[i]=tmp, i++
	i     j
 0  1 -4  3 10
	   i  j
 0  1  9 -4 10
		  i  j
 0  1  9 16 10
			ij
 0  1  9 16 100

 -4, -3, -2, -1 -> O(N)
 1, 2, 3, 4 -> O(N)
 worst case
-2, -2, -2, -2, 3 -> O(N^2)
*/
// time complexity: O(N)~O(N^2)
// space complexity: O(1)
func sortedSquares4(A []int) []int {
	var i, j int
	for j < len(A) && A[j] < 0 {
		j++
	}

	for j < len(A) && i < j {
		l := A[j-1] * A[j-1]
		r := A[j] * A[j]

		if l > r {
			copy(A[i+1:j+1], A[i:j])
			A[i] = r
			i++
			j++
			continue
		}

		if i+1 < j {
			copy(A[i+1:j], A[i:j-1])
		}
		A[i] = l
		i++
	}

	// from i to j , need reverse and square
	for r := i + j; i <= r-i-1; i++ {
		e := r - i - 1
		A[i], A[e] = A[e]*A[e], A[i]*A[i]
	}
	// from j to the end, just square
	for j < len(A) {
		A[j] = A[j] * A[j]
		j++
	}

	return A
}
