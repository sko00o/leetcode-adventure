package conclusion

func isPerfectSquare(num int) bool {
	// positive integer num
	if num <= 1 {
		return true
	}

	low, high := 1, num
	for low <= high {
		mid := low + (high-low)>>1
		mul := mid * mid
		if mul == num {
			return true
		} else if mul > num || mul <= 0 { // mul <= 0 means overflow
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return false
}
