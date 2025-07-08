package problems

func strStr(S string, W string) int {
	j, k := 0, 0
	T := kmpTable(W)
	for j < len(S) {
		if S[j] == W[k] {
			j++
			k++
			// occurrence found one
			if k == len(W) {
				return j - k
			}
		} else {
			if T[k] < 0 {
				j++
				k = 0
			} else {
				k = T[k]
			}
		}
	}
	return -1
}

func kmpTable(W string) []int {
	T := make([]int, len(W))
	T[0] = -1
	for cnd, pos := 0, 1; pos < len(W); pos++ {
		if W[pos] == W[cnd] {
			T[pos] = T[cnd]
		} else {
			T[pos] = cnd
			for cnd >= 0 && W[pos] != W[cnd] {
				cnd = T[cnd]
			}
		}
		cnd++
	}
	return T
}

// -----------------------------------------------------------------------------

func strStr1(haystack string, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}
