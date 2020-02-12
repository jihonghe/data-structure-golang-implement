package optimized_goroutine_k_way_merge_sort

import (
    "data-structure-golang-implement/sort/external_sort/k_way_merge_sort/k_way_merge_sort_implement_by_goroutines"
    "data-structure-golang-implement/sort/external_sort/k_way_merge_sort/optimized_goroutine_k_way_merge_sort"
    "data-structure-golang-implement/utils"
    "fmt"
    "sort"
    "testing"
    "time"
)

func TestMergeSort(t *testing.T) {
    numsPath := "/home/jihonghe/gitRepos/data-structure-golang-implement/sort/test/external_sort/k_way_merge_sort/nums"
    nums := utils.IntsToints64(utils.ReadIntsFromFile(numsPath))
    start := time.Now().UnixNano()
    optimized_goroutine_k_way_merge_sort.MergeSort(nums)
    end := time.Now().UnixNano()
    fmt.Printf("MergeSort time cost: %5f\n", float64(end - start) / 1e9)
    
    qkSortedNums  := utils.IntsToints64(utils.ReadIntsFromFile(numsPath))
    start = time.Now().UnixNano()
    sort.Sort(k_way_merge_sort_implement_by_goroutines.Int64Array(qkSortedNums))
    end = time.Now().UnixNano()
    println()
    fmt.Printf("QuickSort time cost: %5f\n", float64(end - start) / 1e9)
    println()
    
    // 校验
    subNums := nums[19999700:]
    for i, value := range qkSortedNums[19999700:] {
        if value != subNums[i] {
            t.Error("failed")
        }
    }
}
