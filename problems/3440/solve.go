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
	gapList := make([]int, 0, len(startTime)+1)
	{
		gapStart := 0
		for i := 0; i < len(startTime); i++ {
			gapVal := startTime[i] - gapStart
			gapList = append(gapList, gapVal)
			gapStart = endTime[i]
		}
		gapList = append(gapList, eventTime-gapStart)
	}

	n := len(gapList)
	prefixMax := make([]int, n)
	suffixMax := make([]int, n)
	prefixMax[0] = gapList[0]
	for i := 1; i < n; i++ {
		if gapList[i] > prefixMax[i-1] {
			prefixMax[i] = gapList[i]
		} else {
			prefixMax[i] = prefixMax[i-1]
		}
	}
	suffixMax[n-1] = gapList[n-1]
	for i := n - 2; i >= 0; i-- {
		if gapList[i] > suffixMax[i+1] {
			suffixMax[i] = gapList[i]
		} else {
			suffixMax[i] = suffixMax[i+1]
		}
	}

	maxFreeTime := 0
	for i := range startTime {
		evLen := endTime[i] - startTime[i]
		freeTime := gapLen(i) + gapLen(i+1)
		maxOtherGap := 0
		if i-1 >= 0 {
			maxOtherGap = prefixMax[i-1]
		}
		if i+2 < n && suffixMax[i+2] > maxOtherGap {
			maxOtherGap = suffixMax[i+2]
		}
		if maxOtherGap >= evLen {
			freeTime += evLen
		}
		if maxFreeTime < freeTime {
			maxFreeTime = freeTime
		}
	}
	return maxFreeTime
}
