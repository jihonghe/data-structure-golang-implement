package k_way_merge_sort

import (
	"data-structure-golang-implement/sort/external_sort/k_way_merge_sort"
	"data-structure-golang-implement/utils"
	"testing"
	"time"
)

func BenchmarkOptimizedKWayMergeSort(b *testing.B) {
	nums := utils.ReadIntsFromFile("nums")

	//k_way_merge_sort.TraverseIntArray(nums)
	for i := 0; i < b.N; i++ {
		resultNums := make([]int, len(nums), len(nums) + 1)
		start := time.Now().UnixNano()
		k_way_merge_sort.OptimizedKWayMergeSort(nums, resultNums, 0, len(nums) - 1)
		end := time.Now().UnixNano()
		println("Merge sort time costs:", int(float64(end - start) / 1e9))
	}
}
