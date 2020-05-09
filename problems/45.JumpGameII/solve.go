package problems

func jump(nums []int) int {
	lastIdx := len(nums) - 1
	if lastIdx < 1 {
		return 0
	}

	var maxIdx int
	for idx, gap := range nums {
		if d := idx + gap; d > maxIdx {
			maxIdx = d
		}

		if idx >= maxIdx {
			break
		}
		if maxIdx >= lastIdx {
			return jump(nums[:idx+1]) + 1
		}
	}

	return -1
}
