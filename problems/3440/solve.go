package problems

func maxFreeTime(eventTime int, startTime []int, endTime []int) int {
	gapLen := func(n int) int {
		if n == 0 {
			return startTime[0]
		}
		if n == len(startTime) {
			return eventTime - endTime[n-1]
		}
		return startTime[n] - endTime[n-1]
	}
	n := len(startTime) + 1
	gapList := make([]int, 0, n)
	{
		gapStart := 0
		for i := 0; i < len(startTime); i++ {
			gapVal := startTime[i] - gapStart
			gapList = append(gapList, gapVal)
			gapStart = endTime[i]
		}
		gapList = append(gapList, eventTime-gapStart)
	}

	prefixMax := make([]int, n)
	suffixMax := make([]int, n)
	prefixMax[0] = gapList[0]
	for i := 1; i < n; i++ {
		prefixMax[i] = max(gapList[i], prefixMax[i-1])
	}
	suffixMax[n-1] = gapList[n-1]
	for i := n - 2; i >= 0; i-- {
		suffixMax[i] = max(gapList[i], suffixMax[i+1])
	}
	hasOtherGap := func(i int, evLen int) bool {
		maxOtherGap := 0
		if i-1 >= 0 {
			maxOtherGap = prefixMax[i-1]
		}
		if i+2 < n && suffixMax[i+2] > maxOtherGap {
			maxOtherGap = suffixMax[i+2]
		}
		return maxOtherGap >= evLen
	}

	maxFreeTime := 0
	for i := range startTime {
		evLen := endTime[i] - startTime[i]
		freeTime := gapLen(i) + gapLen(i+1)
		if hasOtherGap(i, evLen) {
			freeTime += evLen
		}
		if maxFreeTime < freeTime {
			maxFreeTime = freeTime
		}
	}
	return maxFreeTime
}

// Greedy
func maxFreeTime1(eventTime int, startTime []int, endTime []int) int {
	n := len(startTime)
	// q[i] means if meeting i has non-adjacent gap to move
	q := make([]bool, n)
	// left to right, max duration t1 of any non-adjacent gap
	// right to left, max duration t2 of any non-adjacent gap
	t1, t2 := 0, 0
	for i := range n {
		if endTime[i]-startTime[i] <= t1 {
			q[i] = true
		}
		if i == 0 {
			t1 = max(t1, startTime[i])
		} else {
			t1 = max(t1, startTime[i]-endTime[i-1])
		}

		if endTime[n-1-i]-startTime[n-1-i] <= t2 {
			q[n-1-i] = true
		}
		if i == 0 {
			t2 = max(t2, eventTime-endTime[n-1])
		} else {
			t2 = max(t2, startTime[n-i]-endTime[n-1-i])
		}
	}

	maxFreeTime := 0
	for i := range n {
		left := 0
		if i != 0 {
			left = endTime[i-1]
		}
		right := eventTime
		if i != n-1 {
			right = startTime[i+1]
		}
		freeTime := right - left
		if !q[i] {
			freeTime -= endTime[i] - startTime[i]
		}
		maxFreeTime = max(maxFreeTime, freeTime)
	}
	return maxFreeTime
}
