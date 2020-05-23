package problems

// time complexity: O(N)
// space complexity: O(N)
func findDisappearedNumbers(nums []int) []int {
	nLen := len(nums)
	var chk = make([]bool, nLen)
	for _, n := range nums {
		if n >= 1 && n <= nLen {
			chk[n-1] = true
		}
	}

	var out = make([]int, 0, nLen)
	for n, c := range chk {
		if c == false {
			out = append(out, n+1)
		}
	}

	return out
}

// in-place operation
// time complexity: O(N)
// space complexity: O(1)
// returned list does not count as extra space
/*
|
4 3 2 7 8 2 3 1
      |
|
7 3 2 4 8 2 3 1
	  ^     |
|
3 3 2 4 8 2 7 1
	| ^     ^
|
2 3 3 4 8 2 7 1
  | ^ ^     ^
|
3 2 3 4 8 2 7 1
  ^ ^ ^     ^
        |
3 2 3 4 8 2 7 1
  ^ ^ ^     ^ |
        |
3 2 3 4 1 2 7 8
| ^ ^ ^     ^ ^
        |
1 2 3 4 3 2 7 8
^ ^ ^ ^     ^ ^
*/
func findDisappearedNumbers1(nums []int) []int {
	for i := 0; i < len(nums); {
		// we expected nums[i] == i+1
		if nums[i] != i+1 {
			n := nums[i] - 1
			if nums[n] != n+1 {
				// move nums[i] to the right index=nums[i]-1
				nums[n], nums[i] = nums[i], nums[n]
				continue
			}
		}

		i++
	}

	var out = make([]int, 0, len(nums))
	for i, n := range nums {
		if n != i+1 {
			out = append(out, i+1)
		}
	}

	return out
}

// in-place operation
// time complexity: O(N)
// space complexity: O(1)
// returned list does not count as extra space
/*
 |
 4  3  2  7  8  2  3  1
 	      ^
    |
 4  3  2 -7  8  2  3  1
 	   ^
       |
 4  3 -2 -7  8  2  3  1
    ^
          |
 4 -3 -2 -7  8  2  3  1
 				   ^
             |
 4 -3 -2 -7  8  2 -3  1
 					  ^
                |
 4 -3 -2 -7  8  2 -3 -1
    '
                   |
 4 -3 -2 -7  8  2 -3 -1
 	   '
 	                  |
 4 -3 -2 -7  8  2 -3 -1
 ^

-4 -3 -2 -7  8  2 -3 -1
*/
func findDisappearedNumbers2(nums []int) []int {
	var cnt int
	for _, n := range nums {
		if n > 0 && nums[n-1] > 0 {
			nums[n-1] = -nums[n-1]
			cnt++
		} else if n < 0 && nums[-n-1] > 0 {
			nums[-n-1] = -nums[-n-1]
			cnt++
		}
	}

	var out = make([]int, 0, len(nums)-cnt)
	for i, n := range nums {
		if n > 0 {
			out = append(out, i+1)
		}
	}

	return out
}
