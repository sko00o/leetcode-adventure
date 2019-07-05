package tmpl3

// time O(nlog(n)) slow
func findClosestElements(arr []int, k int, x int) []int {
	if len(arr) == 0 {
		return nil
	}

	low, high := 0, len(arr)-1
	for low+1 < high {
		mid := low + (high-low)>>1
		if arr[mid] == x {
			return fillClosest(arr, k, x, mid)
		} else if arr[mid] < x {
			low = mid
		} else {
			high = mid
		}
	}

	if arr[high] <= x {
		return fillClosest(arr, k, x, high)
	}
	return fillClosest(arr, k, x, low)
}

// note: k must smaller than len(arr)
func fillClosest(arr []int, k int, x int, idx int) []int {
	if idx == 0 {
		return arr[:k]
	}
	if idx == len(arr)-1 {
		return arr[len(arr)-k:]
	}

	left, right := idx, idx+1
	for right-left < k {
		if right == len(arr) ||
			(left != 0 && x-arr[left-1] <= arr[right]-x) {
			left--
		} else {
			right++
		}
	}
	return arr[left:right]
}

func findClosestElements1(arr []int, k int, x int) []int {
	// out array head will in [left,right) range
	left, right := 0, len(arr)-k
	for left < right {
		mid := left + (right-left)>>1
		if x-arr[mid] > arr[mid+k]-x {
			left = mid + 1 // IMPORTANT
		} else {
			right = mid
		}
	}
	// end condition is left == right
	return arr[left : left+k]
}
