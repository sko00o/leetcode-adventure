package problems

func findNumbers(nums []int) int {
	var cnt int
	for _, n := range nums {
		var dCnt int
		for n > 0 {
			n /= 10
			dCnt++
		}
		cnt += (dCnt + 1) % 2
	}
	return cnt
}
