package conclusion

func nextGreatestLetter(letters []byte, target byte) byte {
	low, high := 0, len(letters)-1
	for low < high {
		mid := low + (high-low)>>1
		if letters[mid] == target {
			low = mid + 1
		} else if letters[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	if letters[low] > target {
		return letters[low]
	}
	return letters[(low+1)%len(letters)]
}
