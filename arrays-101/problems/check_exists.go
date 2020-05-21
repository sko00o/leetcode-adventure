package problems

func checkIfExist(arr []int) bool {
	var set = make(map[int]bool)

	for _, n := range arr {
		if set[n*2] {
			return true
		}
		if n%2 == 0 && set[n/2] {
			return true
		}
		set[n] = true
	}

	return false
}
