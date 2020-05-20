package problems

// Brute Force
func duplicateZeros(arr []int) {
	aLen := len(arr)
	// ignore the last one
	for i := aLen - 2; i >= 0; i-- {
		if arr[i] == 0 {
			copy(arr[i+1:], arr[i:aLen-1])
			arr[i] = 0
		}
	}
}

// 2 pass
func duplicateZeros1(arr []int) {
	var dups int
	var aLen = len(arr)

	for i := 0; i < aLen-dups; i++ {
		if arr[i] == 0 {
			if i == aLen-dups-1 {
				arr[aLen-1] = 0
				aLen--
			} else {
				dups++
			}
		}
	}

	for i := aLen - dups - 1; i >= 0; i-- {
		if arr[i] == 0 {
			arr[i+dups] = 0
			dups--
			arr[i+dups] = 0
		} else {
			arr[i+dups] = arr[i]
		}
	}
}
