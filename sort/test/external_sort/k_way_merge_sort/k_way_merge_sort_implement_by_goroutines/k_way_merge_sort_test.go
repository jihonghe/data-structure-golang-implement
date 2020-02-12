package k_way_merge_sort_implement_by_goroutines

import (
	"data-structure-golang-implement/sort/external_sort/k_way_merge_sort/k_way_merge_sort_implement_by_goroutines"
	"data-structure-golang-implement/utils"
	"fmt"
	"sort"
	"testing"
	"time"
)

func TestKWayMergeSort(t *testing.T) {
	numsPath := "/home/jihonghe/gitRepos/data-structure-golang-implement/sort/test/external_sort/k_way_merge_sort/nums"
	nums := utils.IntsToints64(utils.ReadIntsFromFile(numsPath))
	arrayLength := len(nums)

	start := time.Now().UnixNano()
	k_way_merge_sort_implement_by_goroutines.MergeSort(nums)
	end := time.Now().UnixNano()
	println()
	fmt.Printf("arrayLength: %d, time: %5f\n", arrayLength, float64(end - start) / 1e9)
	println()

	qkSortedNums  := utils.IntsToints64(utils.ReadIntsFromFile(numsPath))
	start = time.Now().UnixNano()
	sort.Sort(k_way_merge_sort_implement_by_goroutines.Int64Array(qkSortedNums))
	end = time.Now().UnixNano()
	println()
	fmt.Printf("arrayLength: %d, time: %5f\n", arrayLength, float64(end - start) / 1e9)
	println()

	// 校验
	subNums := nums[19999700:]
	for i, value := range qkSortedNums[19999700:] {
		if value != subNums[i] {
			t.Error("failed")
		}
	}
}

