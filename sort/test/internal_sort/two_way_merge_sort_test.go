package internal_sort_test

import (
	"data-structure-golang-implement/sort/internal_sort"
	"testing"
)

func TestTwoWayMergeSort(t *testing.T) {
	unsortedArray := []int{3, 1, 9, -1, 0}
	sortedArray := internal_sort.MergeSort(unsortedArray)

	resultArray := []int{-1, 0, 1, 3, 9}
	for index, element := range sortedArray {
		if resultArray[index] != element {
			t.Error("internal_sort.MergeSort() error: the sorted array is invalid.")
		}
	}
}
