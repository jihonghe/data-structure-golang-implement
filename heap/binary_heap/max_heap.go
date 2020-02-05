package binary_heap

import "errors"

type MaxHeap struct {
	BinaryHeap
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{}
}

// Insert: 插入新元素
func (maxHeap *MaxHeap) AppendNodes(values... Element) {
	for _, value := range values {
		maxHeap.elements = append(maxHeap.elements, value)
		maxHeap.length++
		// 插入新节点后为新节点执行上浮操作
		maxHeap.siftUp(maxHeap.length - 1)
	}
}

// PopRoot: 取出根节点
func (maxHeap *MaxHeap) PopRoot() Element {
	rootValue := maxHeap.elements[0]

	maxHeap.elements[0] = maxHeap.elements[maxHeap.length - 1]
	maxHeap.elements = maxHeap.elements[:maxHeap.length - 1]
	maxHeap.length--

	// 取出根节点后，执行下沉操作
	maxHeap.siftDown(0)

	return rootValue
}

// siftUp: 节点上浮操作
func (maxHeap *MaxHeap) siftUp(childIndex int) {
	if childIndex == 0 {
		return
	}
	parentIndex := maxHeap.getParentIndex(childIndex)

	for maxHeap.elements[childIndex] > maxHeap.elements[parentIndex] {
		maxHeap.exchangeTwoNodes(childIndex, parentIndex)
		childIndex = parentIndex
		parentIndex = maxHeap.getParentIndex(childIndex)
	}
}

// siftDown: 节点下沉操作
func (maxHeap *MaxHeap) siftDown(targetIndex int) {
	if !maxHeap.isValidIndex(targetIndex) {
		panic(errors.New("Parameter 'targetIndex' value is invalid"))
	}
	if maxHeap.IsEmpty() || maxHeap.length == 1 {
		return
	}

	for {
		// 获取左子节点索引并校验左子节点是否存在
		leftChildIndex := maxHeap.getLeftChildIndex(targetIndex)
		if !maxHeap.isValidIndex(leftChildIndex) {
			return
		}

		biggerNodeIndex := targetIndex
		// 获取右子节点索引并校验右子节点是否存在
		rightChildIndex := maxHeap.getRightChildIndex(targetIndex)
		if !maxHeap.isValidIndex(rightChildIndex) {
			biggerNodeIndex = leftChildIndex
		} else {
			biggerNodeIndex = maxHeap.getBiggerNodeIndex(leftChildIndex, rightChildIndex)
		}

		// 用左右子节点的较大者与父节点比较
		if maxHeap.elements[targetIndex] < maxHeap.elements[biggerNodeIndex] {
			maxHeap.exchangeTwoNodes(targetIndex, biggerNodeIndex)
			targetIndex = biggerNodeIndex
		} else {
			return
		}
	}
}
