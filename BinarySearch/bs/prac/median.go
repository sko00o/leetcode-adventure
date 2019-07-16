package prac

// https://leetcode.com/problems/median-of-two-sorted-arrays/solution/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	m, n := len(nums1), len(nums2)
	iMin, iMax, halfLen := 0, m, (m+n+1)/2

	for iMin <= iMax {
		i := iMin + (iMax-iMin)>>1
		j := halfLen - i

		if i < m && nums2[j-1] > nums1[i] {
			iMin = i + 1
		} else if i > 0 && nums1[i-1] > nums2[j] {
			iMax = i - 1
		} else { // nums2[j-1] <= nums1[i] && nums1[i-1] <= nums2[j]
			var maxLeft float64
			if j == 0 || (i > 0 && nums1[i-1] > nums2[j-1]) {
				maxLeft = float64(nums1[i-1])
			} else {
				maxLeft = float64(nums2[j-1])
			}

			if (m+n)&1 == 1 {
				return maxLeft
			} else {
				var minRight float64
				if j == n || (i != m && nums1[i] < nums2[j]) {
					minRight = float64(nums1[i])
				} else {
					minRight = float64(nums2[j])
				}
				return (maxLeft + minRight) / 2
			}
		}
	}
	return -1
}
