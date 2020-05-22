package problems

// Runtime: 8 ms
// Memory Usage: 3.9 MB
// time complexity: O(N)
// space complexity: O(N)
func thirdMax(nums []int) int {
	var top [3]int
	var set = make(map[int]struct{})
	var cnt int
	for _, n := range nums {
		if _, ok := set[n]; ok {
			continue
		}
		set[n] = struct{}{}
		cnt++

		switch {
		case cnt == 1:
			top[0] = n
		case cnt == 2:
			switch {
			case n > top[0]:
				top[1], top[0] = top[0], n
			default:
				top[1] = n
			}
		case cnt == 3:
			switch {
			case n > top[0]:
				top[2], top[1], top[0] = top[1], top[0], n
			case n > top[1]:
				top[2], top[1] = top[1], n
			default:
				top[2] = n
			}
		default:
			switch {
			case n > top[0]:
				top[2], top[1], top[0] = top[1], top[0], n
			case n > top[1]:
				top[2], top[1] = top[1], n
			case n > top[2]:
				top[2] = n
			}
		}
	}

	if cnt < 3 {
		return top[0]
	}
	return top[2]
}

func thirdMax1(nums []int) int {
	var top [3]int
	var got [3]bool
	var set = make(map[int]struct{})
	for _, n := range nums {
		if _, ok := set[n]; ok {
			continue
		}
		set[n] = struct{}{}

		if !got[0] {
			got[0] = true
			top[0] = n
			continue
		}
		if !got[1] {
			got[1] = true
			switch {
			case n > top[0]:
				top[1], top[0] = top[0], n
			default:
				top[1] = n
			}
			continue
		}
		if !got[2] {
			got[2] = true
			switch {
			case n > top[0]:
				top[2], top[1], top[0] = top[1], top[0], n
			case n > top[1]:
				top[2], top[1] = top[1], n
			default:
				top[2] = n
			}
			continue
		}

		switch {
		case n > top[0]:
			top[2], top[1], top[0] = top[1], top[0], n
		case n > top[1]:
			top[2], top[1] = top[1], n
		case n > top[2]:
			top[2] = n
		}
	}

	if got[2] {
		return top[2]
	}
	return top[0]
}

// pass 3 times
// Runtime: 4 ms
// Memory Usage: 3.5 MB
// time complexity: O(N)
// space complexity: O(N)
func thirdMax2(nums []int) int {
	var res = make([]int, 0, 3)
	for i := 0; i < 3; i++ {
		var max int
		if len(nums) == 0 {
			break
		}
		for idx, n := range nums {
			if idx == 0 || n > max {
				max = n
			}
		}
		res = append(res, max)

		var nnums = make([]int, 0, len(nums))
		for _, n := range nums {
			if n != max {
				nnums = append(nnums, n)
			}
		}
		nums = nnums
	}

	if len(res) == 3 {
		return res[2]
	}

	return res[0]
}
