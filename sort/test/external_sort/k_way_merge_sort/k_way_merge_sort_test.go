package k_way_merge_sort_test

import (
	"data-structure-golang-implement/sort/external_sort/k_way_merge_sort"
	"data-structure-golang-implement/utils"
	"fmt"
	"testing"
	"time"
)

func TestKWayMerge(t *testing.T) {
	arrays := []([]int) {
		[]int{-1, 6, 9, 15, 20, 30, 40, 50, 60, 70, 80, 90},
		[]int{3, 8, 9, 77, 80},
		[]int{99, 121, 200, },
		[]int{-3, 80, 95, 101, 345, 666},
	}

	start := time.Now()
	sortedArray := k_way_merge_sort.KWayMerge(4, arrays)
	end := time.Now()
	println("Time costs:", end.Nanosecond() - start.Nanosecond())
	for _, value := range sortedArray {
		fmt.Printf("%d, ", value)
	}
	println()
}

func TestGetExponentialOfNum(t *testing.T) {
	groupCount := k_way_merge_sort.GetGroupCount(9, 8)
	fmt.Printf("Group Count: %d\n", groupCount)
}

func TestGetRandomSlice(t *testing.T) {
	nums := k_way_merge_sort.GetRandomSlice(10, 10)
	k_way_merge_sort.TraverseIntArray(nums)
}

func TestKWayMergeSort(t *testing.T) {
	nums := utils.ReadIntsFromFile("nums")

	//k_way_merge_sort.TraverseIntArray(nums)
	start := time.Now().UnixNano()
	nums = k_way_merge_sort.KWayMergeSort(nums)
	end := time.Now().UnixNano()
	println("Merge sort time costs:", int(float64(end - start) / 1e9))
}
