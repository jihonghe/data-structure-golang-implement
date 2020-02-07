package k_way_merge_sort

import (
	"data-structure-golang-implement/sort/external_sort/k_way_merge_sort"
	"data-structure-golang-implement/utils"
	"fmt"
	"testing"
	"time"
)

func TestOptimizedKWayMergeSort(t *testing.T) {
	nums := utils.ReadIntsFromFile("nums")
	//nums := k_way_merge_sort.GetRandomSlice(10000000, 100)
	//k_way_merge_sort.TraverseIntArray(nums)
	resultNums := make([]int, len(nums), len(nums) + 1)

	start := time.Now().UnixNano()
	k_way_merge_sort.OptimizedKWayMergeSort(nums, resultNums, 0, len(nums) - 1)
	end := time.Now().UnixNano()
	println()
	fmt.Printf("arrayLength: %d, time: %d\n", len(nums), int(float64(end - start) / 1e9))
	println()
	//k_way_merge_sort.TraverseIntArray(resultNums)
}
