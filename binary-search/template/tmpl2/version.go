package tmpl2

func firstBadVersion(n int) int {
	if n < 1 {
		return -1
	}

	low, high := 1, n
	for low < high {
		mid := low + (high-low)>>1
		if isBadVersion(mid) {
			high = mid
		} else {
			low = mid + 1
		}
	}

	for low <= n && isBadVersion(low) {
		return low
	}
	return -1
}

var badVersion int

func isBadVersion(ver int) bool {
	return ver >= badVersion
}
