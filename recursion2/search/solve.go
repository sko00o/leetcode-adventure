package search

/*
This is a fun problem to solve, there exists at least 4 different solutions, each with different time complexity.
As an exercise, you can solve this recursively by dividing it into multiple subproblems, and analyze its time complexity.
*/

// Time: O(n*log(n))
func searchMatrix(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		if binarySearch(matrix[i], target) >= 0 {
			return true
		}
	}
	return false
}
func binarySearch(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1
	for left+1 < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid
		} else {
			right = mid
		}
	}

	if nums[left] == target {
		return left
	}
	if nums[right] == target {
		return right
	}
	return -1
}

func searchMatrix0(matrix [][]int, target int) bool {
	rows := len(matrix)
	if rows == 0 {
		return false
	}
	cols := len(matrix[0]) - 1

	minRow, exist := search(SearchCol{matrix, cols}, target)
	if exist {
		return true
	}
	maxRow, exist := search(SearchCol{matrix, 0}, target)
	if exist {
		return true
	}

	if minRow < rows {
		for r := minRow; r < maxRow; r++ {
			if _, ok := search(SearchRow{matrix, r}, target); ok {
				return true
			}
		}
	}
	return false
}

func searchMatrix1(matrix [][]int, target int) bool {
	rows := len(matrix)
	if rows == 0 {
		return false
	}
	cols := len(matrix[0])

	// begin point first row, last column
	r, c := 0, cols-1
	for r < rows && c >= 0 {
		if target < matrix[r][c] {
			c--
		} else if target > matrix[r][c] {
			r++
		} else {
			return true
		}
	}
	return false
}
