package k_way_merge_sort_test

import (
	"data-structure-golang-implement/sort/external_sort/k_way_merge_sort"
	"fmt"
	"testing"
)

func TestNewTraceableElementMinHeap(t *testing.T) {
	minHeap := k_way_merge_sort.NewMinHeap(4)

	if !minHeap.IsEmpty() {
		t.Error("k_way_merge_sort.NewMinHeap() error: initialize min heap failed, the length is invalid.")
	}
}

func TestAppendNode(t *testing.T) {
	minHeap := k_way_merge_sort.NewMinHeap(4)
	element := k_way_merge_sort.NewTraceableElement(9, 1)
	element1 := k_way_merge_sort.NewTraceableElement(0, 1)
	element2 := k_way_merge_sort.NewTraceableElement(-1, 1)
	element3 := k_way_merge_sort.NewTraceableElement(8, 1)
	element4 := k_way_merge_sort.NewTraceableElement(-3, 1)
	minHeap.AppendNode(*element)
	minHeap.AppendNode(*element1)
	minHeap.AppendNode(*element2)
	minHeap.AppendNode(*element3)
	minHeap.AppendNode(*element4)

	if minHeap.IsEmpty() || minHeap.GetLength() != 4 {
		t.Error("AppendNode() error: failed.")
	}
	minHeap.Traverse()
}

func TestPopRoot(t *testing.T) {
	minHeap := k_way_merge_sort.NewMinHeap(4)
	element := k_way_merge_sort.NewTraceableElement(9, 1)
	element1 := k_way_merge_sort.NewTraceableElement(0, 1)
	element2 := k_way_merge_sort.NewTraceableElement(-1, 1)
	element3 := k_way_merge_sort.NewTraceableElement(8, 1)
	minHeap.AppendNode(*element)
	minHeap.AppendNode(*element1)
	minHeap.AppendNode(*element2)
	minHeap.AppendNode(*element3)

	rootNodeValue := minHeap.PopRoot().GetElementValue()
	if minHeap.GetLength() != 3 {
		t.Error(fmt.Sprintf("PopRoot() error: failed."))
	}
	fmt.Printf("rootNodeValue=%d\n", rootNodeValue)
	minHeap.Traverse()

}
