package sort

/*
Notes:
分组进行的插入排序。
分组增量序列的选取影响算法效率。
Hibbard 增量序列：1，3，7，...，2n-1 被证明可广泛应用，时间复杂度 O(N^1.5)
希尔排序复杂度范围大约 O(N^1.3) ~ O(N^2)。
不稳定排序。
*/

func shellSort(nums []int) {
	for step := len(nums) / 2; step >= 1; step /= 2 {
		// insert sort
		for i := step; i < len(nums); i += step {
			for j := i - step; j >= 0; j -= step {
				if nums[j+step] >= nums[j] {
					break
				}
				nums[j+step], nums[j] = nums[j], nums[j+step]
			}
		}
	}
}
