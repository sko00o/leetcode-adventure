package tmpl1

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	if n <= 1 {
		return n
	}

	low, high := 1, n
	for low < high {
		mid := low + ((high - low) >> 1)
		switch guess(mid) {
		case 0:
			return mid
		case 1:
			low = mid + 1
		case -1:
			high = mid - 1
		}
	}

	return low
}

var secret int

func guess(n int) int {
	if n < secret {
		return 1
	}
	if n > secret {
		return -1
	}
	return 0
}
