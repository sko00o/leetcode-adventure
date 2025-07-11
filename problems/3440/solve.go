package problems

// Still TLE
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
	gapList := make([]gap, 0, len(startTime)+1)
	{
		gapStart := 0
		for i := 0; i < len(startTime); i++ {
			gapVal := startTime[i] - gapStart
			gapList = append(gapList, gap{idx: i, val: gapVal})
			gapStart = endTime[i]
		}
		gapList = append(gapList, gap{idx: len(startTime), val: eventTime - gapStart})
	}
	quickSort(gapList, 0, len(gapList)-1)

	maxFreeTime := 0
	for i := range startTime {
		evLen := endTime[i] - startTime[i]
		freeTime := gapLen(i) + gapLen(i+1)
		if hasBiggerGap(gapList, evLen, i) {
			freeTime += evLen
		}
		if maxFreeTime < freeTime {
			maxFreeTime = freeTime
		}
	}
	return maxFreeTime
}

type gap struct {
	idx int
	val int
}

// first index of arr val that is bigger than val, notI is excluded
func hasBiggerGap(arr []gap, val int, notI int) bool {
	hi := len(arr) - 1
	if arr[hi].val < val {
		return false
	}
	for hi >= 0 && (arr[hi].idx == notI || arr[hi].idx == notI+1) {
		hi--
	}
	if hi < 0 {
		return false
	}
	return arr[hi].val >= val
}

func quickSort(arr []gap, lo, hi int) {
	if lo >= hi || lo < 0 {
		return
	}
	p := partition(arr, lo, hi)
	quickSort(arr, lo, p-1)
	quickSort(arr, p+1, hi)
}
func partition(arr []gap, lo, hi int) int {
	pivot := arr[hi]
	i := lo
	for j := lo; j < hi; j++ {
		if arr[j].val < pivot.val {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[hi] = arr[hi], arr[i]
	return i
}
