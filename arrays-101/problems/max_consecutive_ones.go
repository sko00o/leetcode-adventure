package problems

func findMaxConsecutiveOnes(nums []int) int {
	var maxLen, currLen int
	for _, n := range append(nums, 0) {
		if n != 1 {
			if maxLen < currLen {
				maxLen = currLen
			}
			currLen = 0
			continue
		}

		currLen++
	}

	return maxLen
}
