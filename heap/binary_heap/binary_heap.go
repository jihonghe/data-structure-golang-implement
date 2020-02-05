package binary_heap

import "fmt"

type Element int

type BinaryHeap struct {
	length int
	elements []Element
}

func (binaryHeap *BinaryHeap) IsEmpty() bool {
	return binaryHeap.length == 0
}

func (binaryHeap *BinaryHeap) GetSize() int {
	return binaryHeap.length
}

func (binaryHeap *BinaryHeap) getParentIndex(childIndex int) int {
	if childIndex == 0 {
		return 0
	}

	return (childIndex - 1) / 2
}

func (binaryHeap *BinaryHeap) getLeftChildIndex(parentIndex int) int {
	return 2 * parentIndex + 1
}

func (binaryHeap *BinaryHeap) getRightChildIndex(parentIndex int) int {
	return 2 * parentIndex + 2
}

func (binaryHeap *BinaryHeap) isValidIndex(index int) bool {
	return index < binaryHeap.length
}

func (binaryHeap *BinaryHeap) getBiggerNodeIndex(leftChildIndex, rightChildIndex int) int {
	if binaryHeap.elements[leftChildIndex] > binaryHeap.elements[rightChildIndex] {
		return leftChildIndex
	}

	return rightChildIndex
}

func (binaryHeap *BinaryHeap) getLowerNodeIndex(leftChildIndex, rightChildIndex int) int {
	if binaryHeap.elements[leftChildIndex] < binaryHeap.elements[rightChildIndex] {
		return leftChildIndex
	}

	return rightChildIndex
}

func (binaryHeap *BinaryHeap) exchangeTwoNodes(nodeIndex1, nodeIndex2 int) {
	binaryHeap.elements[nodeIndex1], binaryHeap.elements[nodeIndex2] = binaryHeap.elements[nodeIndex2], binaryHeap.elements[nodeIndex1]
}

func TraverseHeap(binaryHeap BinaryHeap) {
	for _, element := range binaryHeap.elements {
		fmt.Printf("%d, ", element)
	}
	println()
}
