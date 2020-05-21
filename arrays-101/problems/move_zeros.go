package problems

func moveZeroes(nums []int) {
	if len(nums) < 2 {
		return
	}

	p1, p2 := 0, 1
	for p2 < len(nums) {
		if nums[p1] != 0 {
			p1++
			p2++
			continue
		}

		// nums[p1] == 0
		if nums[p2] != 0 {
			nums[p1], nums[p2] = nums[p2], nums[p1]
			continue
		}

		p2++
	}
}

func moveZeroes1(nums []int) {
	var zIdx = -1
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 && zIdx != -1 {
			nums[i], nums[zIdx] = nums[zIdx], nums[i]
			zIdx++
		} else if nums[i] == 0 && zIdx == -1 {
			zIdx = i
		}
	}
}

// Optimal
func moveZeroes2(nums []int) {
	for i, j := 0, 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}
