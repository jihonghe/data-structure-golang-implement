package binary_tree

type btreeNodeQueue struct {
    elems []*BTreeNode
    length int
}

func NewBTreeNodeQueue() *btreeNodeQueue {
    return new(btreeNodeQueue)
}

func (btQueue *btreeNodeQueue) isEmpty() bool {
    if btQueue.length == 0 {
        return true
    }
    return false
}

func (btQueue *btreeNodeQueue) enqueue(node *BTreeNode) {
    if node.IsNilNode() {
        return
    }
    btQueue.elems = append(btQueue.elems, node)
    btQueue.length++
}

func (btQueue *btreeNodeQueue) dequeue() *BTreeNode {
    if btQueue.isEmpty() {
        return nil
    }
    node := btQueue.elems[0]
    btQueue.elems = btQueue.elems[1:]
    btQueue.length--

    return node
}

func (btQueue *btreeNodeQueue) peek() *BTreeNode {
    if btQueue.isEmpty() {
        return nil
    }
    return btQueue.elems[0]
}
