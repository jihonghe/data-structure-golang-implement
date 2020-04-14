package binary_tree

type btreeNodeStack struct {
    elems []*BTreeNode
    length int
    topNodeIndex int
}

func NewBTreeNodeStack() *btreeNodeStack {
    return &btreeNodeStack{
        elems:       nil,
        length:      0,
        topNodeIndex: -1,
    }
}

func (btStack *btreeNodeStack) isEmpty() bool {
    if btStack.length == 0 {
        return true
    }
    return false
}

func (btStack *btreeNodeStack) pop() *BTreeNode {
    if btStack.isEmpty() {
        return nil
    }
    popNode := btStack.elems[btStack.topNodeIndex]
    btStack.elems = btStack.elems[:btStack.topNodeIndex]
    btStack.length--

    return popNode
}

func (btStack *btreeNodeStack) push(node *BTreeNode) {
    if node.IsNilNode() {
        return
    }
    btStack.elems = append(btStack.elems, node)
    btStack.length++
    btStack.topNodeIndex++
}

func (btStack *btreeNodeStack) peek() *BTreeNode {
    if btStack.isEmpty() {
        return nil
    }
    return btStack.elems[btStack.topNodeIndex]
}
