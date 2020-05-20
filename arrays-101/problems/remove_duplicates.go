package problems

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var nLen = 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[nLen] = nums[i]
			nLen++
		}
	}
	return nLen
}
