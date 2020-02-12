package optimized_goroutine_k_way_merge_sort

import (
	"errors"
	"fmt"
)

type MinHeap struct {
	elements []int64
	elementsArrayIds []int
	length int
	capacity int
}

func NewMinHeap(capacity int) *MinHeap {
	return &MinHeap{
		elements: make([]int64, capacity, capacity),
		elementsArrayIds: make([]int, capacity, capacity),
		length: 0,
		capacity: capacity,
	}
}

func (minHeap MinHeap) Traverse() {
	for _, element := range minHeap.elements {
		fmt.Printf("%d, ", element)
	}
	println()
}

func (minHeap MinHeap) Len() int {
	return minHeap.length
}

func (minHeap MinHeap) IsEmpty() bool {
	return minHeap.length == 0
}

func (minHeap MinHeap) IsFull() bool {
	return minHeap.length == minHeap.capacity
}

func (minHeap *MinHeap) AppendNode(element int64, elementArrayId int) {
	if minHeap.IsFull() {
		return
	}

	// 在二叉堆末尾添加新节点
	minHeap.elements[minHeap.length] = element
	minHeap.elementsArrayIds[minHeap.length] = elementArrayId
	minHeap.length++

	// 上浮新节点，调整堆
	minHeap.siftUp(minHeap.length - 1)
}

func (minHeap *MinHeap) PopRoot() (int64, bool) {
	if minHeap.IsEmpty() {
		return 0, false
	}
	rootNodeElement := minHeap.elements[0]

	// 删除根节点，并将二叉堆的最后一个叶节点放到根节点上
	minHeap.elements[0] = minHeap.elements[minHeap.length - 1]
	minHeap.elementsArrayIds[0] = minHeap.elementsArrayIds[minHeap.length - 1]
	minHeap.length--

	// 下沉根节点元素，调整堆
	minHeap.siftDown(0)

	return rootNodeElement, true
}

func (minHeap *MinHeap) siftUp(nodeIndex int) {
	if minHeap.length <= 1 {
		return
	}

	for {
		parentIndex := getParentIndex(nodeIndex)
		if less(minHeap.elements[nodeIndex], minHeap.elements[parentIndex]) {
			minHeap.exchangeTwoNodes(nodeIndex, parentIndex)
			nodeIndex = parentIndex
		} else {
			break
		}
	}

}

func (minHeap *MinHeap) siftDown(nodeIndex int) {
	if minHeap.length <= 1 {
		return
	}

	for {
		lessNodeIndex := nodeIndex

		leftNodeIndex := getLeftNodeIndex(nodeIndex)
		if !minHeap.isValidIndex(leftNodeIndex) {
			return
		}

		rightNodeIndex := getRightNodeIndex(nodeIndex)
		if !minHeap.isValidIndex(rightNodeIndex) {
			lessNodeIndex = leftNodeIndex
		} else {
			lessNodeIndex = minHeap.getLessNodeIndex(leftNodeIndex, rightNodeIndex)
		}

		if !less(minHeap.elements[nodeIndex], minHeap.elements[lessNodeIndex]) {
			minHeap.exchangeTwoNodes(nodeIndex, lessNodeIndex)
			nodeIndex = lessNodeIndex
		} else {
			return
		}
	}
}

func getParentIndex(nodeIndex int) int {
	if nodeIndex == 0 {
		return 0
	}

	return (nodeIndex - 1) / 2
}

func getLeftNodeIndex(parentIndex int) int {
	return 2 * parentIndex + 1
}

func getRightNodeIndex(parentIndex int) int {
	return 2 * parentIndex + 2
}

func (minHeap MinHeap) getLessNodeIndex(nodeIndex1, nodeIndex2 int) int {
	if nodeIndex1 >= minHeap.length || nodeIndex2 >= minHeap.length {
		panic(errors.New(fmt.Sprintf("index out of range: nodeIndex1=%d, nodeIndex2=%d\n", nodeIndex1, nodeIndex2)))
	}

	if less(minHeap.elements[nodeIndex1], minHeap.elements[nodeIndex2]) {
		return nodeIndex1
	}

	return nodeIndex2
}

func (minHeap MinHeap) isValidIndex(nodeIndex int) bool {
	return nodeIndex < minHeap.length
}

func less(value1, value2 int64) bool {
	return value1 < value2
}

func (minHeap *MinHeap) exchangeTwoNodes(nodeIndex1, nodeIndex2 int) {
	minHeap.elements[nodeIndex1], minHeap.elements[nodeIndex2] = minHeap.elements[nodeIndex2],
	minHeap.elements[nodeIndex1]
	
	minHeap.elementsArrayIds[nodeIndex1], minHeap.elementsArrayIds[nodeIndex2] = minHeap.elementsArrayIds[nodeIndex2],
	minHeap.elementsArrayIds[nodeIndex1]
}
