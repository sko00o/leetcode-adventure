package practice_hard

// 切分数组 nums 为 m 个连续非空子数组，
// 算法实现最小化这 m 个子数组的中的最大和。

// 搜索范围时数组最大值到数组元素之和，
// 查找子数组元素和的上界，按给定上界
// 是否可以切分成 m 个子数组，有且分出
// 子数组的数量决定下一搜索区间。
func splitArray(nums []int, m int) int {
	var max, sum int
	for i, num := range nums {
		if i == 0 {
			max = num
		} else if num > max {
			max = num
		}
		sum += num
	}
	l, r := max, sum

	for l <= r {
		mid := l + (r-l)>>1

		if tooSmall(nums, m, mid) { // means mid is too small
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return l
}

func tooSmall(nums []int, m int, average int) bool {
	ct, subSum := 1, 0
	for _, num := range nums {
		subSum += num
		if subSum > average {
			subSum = num
			ct++
			if ct > m {
				return true
			}
		}
	}
	return false
}

// 暴力解法 => Memory Limit Exceeded
func splitArray0(nums []int, m int) int {

	// 一遍扫描拿到前 n 项和 s[n]
	s := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			s[i] = nums[0]
		} else {
			s[i] = s[i-1] + nums[i]
		}
	}

	// 用于计算任意区间的元素和
	// sum [i, j) == nums[i] + ... nums[j-1]
	sum := func(i, j int) int {
		if j < i {
			i, j = j, i
		}

		if i == 0 {
			return s[j-1]
		}
		return s[j-1] - s[i-1]
	}

	// 列出所有切分点
	//  0 1 2 3 ... n-1 n
	// ^ ^ ^ ^ ... ^   ^ ^
	// 0 1 2 3 ... n-1 n n+1
	sIdxs := splitIndexes(len(nums), m)
	result := s[len(nums)-1]
	for _, idxs := range sIdxs {
		var largest int
		for idx, i := range idxs {
			var t int
			if idx == 0 {
				largest = sum(0, i)
			} else {
				if t = sum(idxs[idx-1], i); t > largest {
					largest = t
				}
			}

			if idx == len(idxs)-1 {
				if t = sum(i, len(nums)); t > largest {
					largest = t
				}
			}
		}

		if largest < result {
			result = largest
		}
	}

	return result
}

// 长度为 n 的数组切分成 m 个非空子数组，多少种切法
// 从 n-1 个点中取 m-1 个点
func splitIndexes(n, m int) [][]int {
	if n <= 0 || m <= 1 || n < m {
		return nil
	}

	out := make([][]int, 0)
	cut := make([]int, m-1)
	var f func(n, m int)
	f = func(n, m int) {
		if n <= 0 || m <= 0 || n < m {
			return
		}
		if m == 1 {
			o := make([]int, len(cut))
			copy(o, cut)
			out = append(out, o)
			return
		}

		// 当前能切下的位置
		for i := 0; i < n-(m-1); i++ {
			// 切掉了 i+1 个值，
			// 剩余部分交给下一级递归
			next := n - (i + 1)
			cut[m-2] = next
			f(next, m-1)
		}
	}

	f(n, m)

	return out
}
