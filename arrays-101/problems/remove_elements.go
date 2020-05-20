package problems

func removeElement(nums []int, val int) int {
	var nLen int

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[nLen] = nums[i]
			nLen++
		}
	}

	return nLen
}

// less assignment
func removeElement1(nums []int, val int) int {
	i, nLen := 0, len(nums)
	for i < nLen {
		if nums[i] == val {
			nums[i] = nums[nLen-1]
			nLen--
		} else {
			i++
		}
	}

	return nLen
}
