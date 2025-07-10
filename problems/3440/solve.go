package problems

// TODO(Time Limit Exceeded)
func maxFreeTime(eventTime int, startTime []int, endTime []int) int {
	maxFreeTime := currMaxFreeTime(eventTime, startTime, endTime)
	for i := range startTime {
		t := tryReschedule(eventTime, i, startTime, endTime)
		if t > maxFreeTime {
			maxFreeTime = t
		}
	}

	return maxFreeTime
}

func remove(i int, arr []int) (res []int) {
	if i > 0 {
		res = append(res, arr[:i]...)
	}
	if i < len(arr) {
		res = append(res, arr[i+1:]...)
	}
	return
}

// insert before i
func insert(i int, arr []int, val int) (res []int) {
	if i < 0 {
		i = 0
	}
	if i >= len(arr) {
		return append(arr, val)
	}

	res = make([]int, len(arr)+1)
	copy(res, arr[:i])
	res[i] = val
	copy(res[i+1:], arr[i:])
	return
}

func tryReschedule(eventTime int, i int, start, end []int) int {
	iLen := end[i] - start[i]
	remStart := remove(i, start)
	remEnd := remove(i, end)

	maxFreeTime := 0
	var newStart, newEnd []int

	extStart := append(append([]int{0}, start...), eventTime)
	extEnd := append(append([]int{0}, end...), eventTime)
	newI := i + 1

	// start[i] move to other endTime[j], j != i
	for j := range extEnd {
		if j == newI {
			continue
		}
		newS := extEnd[j]
		newE := newS + iLen
		if newE > eventTime {
			continue
		}
		if j < len(extEnd)-1 {
			next := j + 1
			if next == newI {
				next++
			}
			if next < len(extStart) && newE > extStart[next] {
				continue
			}
		}

		/*
			    i
			  0 1 2 3
			0 1 2 3 4 5
			    j
		*/

		// insert after, so offset +1
		if j < newI {
			newStart = insert(j, remStart, newS)
			newEnd = insert(j, remEnd, newE)
		} else {
			newStart = insert(j-1, remStart, newS)
			newEnd = insert(j-1, remEnd, newE)
		}
		t := currMaxFreeTime(eventTime, newStart, newEnd)
		if maxFreeTime < t {
			maxFreeTime = t
		}
	}

	// endTime[i] move to other startTime[j], j != i
	for j := range extStart {
		if j == newI {
			continue
		}
		newE := extStart[j]
		newS := newE - iLen
		if newS < 0 {
			continue
		}
		if j > 1 {
			prev := j - 1
			if prev == newI {
				prev--
			}
			if prev >= 0 && newS < extEnd[prev] {
				continue
			}
		}

		/*
			    i
			  0 1 2 3
			0 1 2 3 4 5
			    j
		*/

		// insert before
		if j < newI {
			newStart = insert(j-1, remStart, newS)
			newEnd = insert(j-1, remEnd, newE)
		} else {
			newStart = insert(j-2, remStart, newS)
			newEnd = insert(j-2, remEnd, newE)
		}
		t := currMaxFreeTime(eventTime, newStart, newEnd)
		if maxFreeTime < t {
			maxFreeTime = t
		}
	}

	return maxFreeTime
}

func currMaxFreeTime(eventTime int, startTime []int, endTime []int) int {
	maxGap := 0
	gapStart := 0
	for i := range startTime {
		if gap := startTime[i] - gapStart; gap > 0 && maxGap < gap {
			maxGap = gap
		}
		gapStart = endTime[i]
	}
	if gap := eventTime - gapStart; gap > 0 && maxGap < gap {
		maxGap = gap
	}
	return maxGap
}
