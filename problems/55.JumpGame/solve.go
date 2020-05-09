package problems

func canJump(nums []int) bool {
	var maxIdx int
	lastIdx := len(nums) - 1
	for idx, gap := range nums {
		if d := idx + gap; d > maxIdx {
			maxIdx = d
		}

		if idx >= maxIdx {
			break
		}
		if maxIdx >= lastIdx {
			return true
		}
	}

	return maxIdx >= lastIdx
}
