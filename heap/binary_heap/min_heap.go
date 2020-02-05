package binary_heap

import "errors"

type MinHeap struct {
	BinaryHeap
}

func NewMinHeap() *MinHeap {
	return &MinHeap{}
}

// Insert: 插入新元素
func (minHeap *MinHeap) AppendNodes(values... Element) {
	for _, value := range values {
		minHeap.elements = append(minHeap.elements, value)
		minHeap.length++
		// 插入新节点后为新节点执行上浮操作
		minHeap.siftUp(minHeap.length - 1)
	}
}

// PopRoot: 取出根节点
func (minHeap *MinHeap) PopRoot() Element {
	rootValue := minHeap.elements[0]

	minHeap.elements[0] = minHeap.elements[minHeap.length - 1]
	minHeap.elements = minHeap.elements[:minHeap.length - 1]
	minHeap.length--

	// 取出根节点后，执行下沉操作
	minHeap.siftDown(0)

	return rootValue
}

// siftUp: 节点上浮操作
func (minHeap *MinHeap) siftUp(childIndex int) {
	if childIndex == 0 {
		return
	}
	parentIndex := minHeap.getParentIndex(childIndex)

	for minHeap.elements[childIndex] < minHeap.elements[parentIndex] {
		minHeap.exchangeTwoNodes(childIndex, parentIndex)
		childIndex = parentIndex
		parentIndex = minHeap.getParentIndex(childIndex)
	}
}

// siftDown: 节点下沉操作
func (minHeap *MinHeap) siftDown(targetIndex int) {
	if !minHeap.isValidIndex(targetIndex) {
		panic(errors.New("Parameter 'targetIndex' value is invalid"))
	}
	if minHeap.IsEmpty() || minHeap.length == 1 {
		return
	}

	for {
		// 获取左子节点索引并校验左子节点是否存在
		leftChildIndex := minHeap.getLeftChildIndex(targetIndex)
		if !minHeap.isValidIndex(leftChildIndex) {
			return
		}

		lowerNodeIndex := targetIndex
		// 获取右子节点索引并校验右子节点是否存在
		rightChildIndex := minHeap.getRightChildIndex(targetIndex)
		if !minHeap.isValidIndex(rightChildIndex) {
			lowerNodeIndex = leftChildIndex
		} else {
			lowerNodeIndex = minHeap.getLowerNodeIndex(leftChildIndex, rightChildIndex)
		}

		// 用左右子节点的较小者与父节点比较
		if minHeap.elements[targetIndex] > minHeap.elements[lowerNodeIndex] {
			minHeap.exchangeTwoNodes(targetIndex, lowerNodeIndex)
			targetIndex = lowerNodeIndex
		} else {
			return
		}
	}
}
