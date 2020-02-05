package test

import (
	"data-structure-golang-implement/heap/binary_heap"
	"testing"
)

func TestNewMinHeap(t *testing.T) {
	minHeap := binary_heap.NewMinHeap()

	if !minHeap.IsEmpty() || minHeap.GetSize() != 0 {
		t.Error("binary_heap.NewMinHeap() error: got non-empty initialized min heap.")
	}
}

func TestMinHeapAppendNodes(t *testing.T) {
	minHeap := binary_heap.NewMinHeap()

	minHeap.AppendNodes(1, 9, -1, 33, 0)

	if minHeap.IsEmpty() || minHeap.GetSize() != 5 {
		t.Error("minHeap.AppendNodes() error: nodes amount is invalid.")
	}
}

func TestMinHeapPopRoot(t *testing.T) {
	minHeap := binary_heap.NewMinHeap()

	minHeap.AppendNodes(1, 9, -1, 33, 0)

	if minHeap.PopRoot() != -1 {
		t.Error("minHeap.PopRoot() error: the root value is not the minimum in heap nodes.")
	}
}
