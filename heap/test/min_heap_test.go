package test

import (
	"data-structure-golang-implement/heap/binary_heap"
	"testing"
)

func TestNewMaxHeap(t *testing.T) {
	maxHeap := binary_heap.NewMaxHeap()

	if !maxHeap.IsEmpty() || maxHeap.GetSize() != 0 {
		t.Error("binary_heap.NewMaxHeap() error: got non-empty initialized min heap.")
	}
}

func TestMaxHeapAppendNodes(t *testing.T) {
	maxHeap := binary_heap.NewMaxHeap()

	maxHeap.AppendNodes(1, 9, -1, 33, 0)

	if maxHeap.IsEmpty() || maxHeap.GetSize() != 5 {
		t.Error("maxHeap.AppendNodes() error: nodes amount is invalid.")
	}
}

func TestMaxHeapPopRoot(t *testing.T) {
	maxHeap := binary_heap.NewMaxHeap()

	maxHeap.AppendNodes(1, 9, -1, 33, 0)

	if maxHeap.PopRoot() != 33 {
		t.Error("maxHeap.PopRoot() error: the root value is not the minimum in heap nodes.")
	}
}
