package sort 

/*
Notes:
将待排元素，插入到左侧已排序列的合适位置，主要开销是数组元素移动。
待排元素数量较小时，性能很好。
待排序列基本有序时，比较次数少。
稳定排序。
*/

func insertSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		deal := nums[i]

		var j int
		for j = i; j > 0 && deal < nums[j-1]; j-- {
				nums[j] = nums[j-1]
		}
		if j != i {
			nums[j] = deal
		}
	}
}