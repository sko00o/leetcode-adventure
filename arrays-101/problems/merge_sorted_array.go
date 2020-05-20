package problems

func merge(nums1 []int, m int, nums2 []int, n int) {
	if len(nums1)-m < n {
		return
	}

	for i, j := 0, 0; j < n; {
		if i < m && nums1[i] > nums2[j] {
			copy(nums1[i+1:], nums1[i:m])
			nums1[i] = nums2[j]
			j++
			m++
		} else if i < m {
			i++
		} else {
			nums1[i] = nums2[j]
			j++
			i++
		}
	}
}

func merge1(nums1 []int, m int, nums2 []int, n int) {
	if len(nums1)-m < n {
		return
	}

	last := m + n - 1
	for i, j := m-1, n-1; j >= 0; {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[last] = nums1[i]
			last--
			i--
		} else {
			nums1[last] = nums2[j]
			last--
			j--
		}
	}
}
