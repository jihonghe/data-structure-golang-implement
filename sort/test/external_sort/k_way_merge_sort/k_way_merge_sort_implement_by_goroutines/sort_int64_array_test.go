package k_way_merge_sort_implement_by_goroutines

import (
	"data-structure-golang-implement/sort/external_sort/k_way_merge_sort/k_way_merge_sort_implement_by_goroutines"
	"sort"
	"testing"
)

func TestInt64ArraySort(t *testing.T) {
	//int64Array := k_way_merge_sort_implement_by_goroutines.NewInt64Array(3, 1, 5, 18, -1, 0, 33, 12, 8, 0)
	ints64 := []int64{2, 1, 3, 8, -1, 8, 10, 0, 5}
	ints642 := k_way_merge_sort_implement_by_goroutines.Int64Array(ints64)
	sort.Sort(ints642)
	for _, value := range ints64 {
		print(value, ", ")
	}
}
