package binary_tree

type BTreeNode struct {
    Data int
    Left, Right *BTreeNode
}

func NewBTreeNode(data int) *BTreeNode {
    return &BTreeNode{
        Data:  data,
        Left:  nil,
        Right: nil,
    }
}

func (node *BTreeNode) IsNilNode() bool {
    if node == nil {
        return true
    }
    return false
}

/*
    二叉树遍历方式：深度优先遍历和广度优先遍历。深度优先遍历又分为前序、中序及后序遍历三种方式，广度优先遍历又称为层序遍历。
    深度优先遍历有递归何非递归实现方式，非递归方式借用栈来缓存未访问节点；广度优先遍历借用队列来缓存未访问节点。
 */

func PreOrder(node *BTreeNode) {
    if node.IsNilNode() {
        return
    }
    print(node.Data, ", ")
    PreOrder(node.Left)
    PreOrder(node.Right)
}

func MiddleOrder(node *BTreeNode) {
    if node.IsNilNode() {
        return
    }
    MiddleOrder(node.Left)
    print(node.Data, ", ")
    MiddleOrder(node.Right)
}

func PostOrder(node *BTreeNode) {
    if node.IsNilNode() {
        return
    }
    PostOrder(node.Left)
    PostOrder(node.Right)
    print(node.Data, ", ")
}

func PreOrderUnrecursive(parentNode *BTreeNode) {
    stack := NewBTreeNodeStack()

    for parentNode != nil || !stack.isEmpty() {
        // 一直纵向延伸到左子树的最左叶子节点
        if parentNode != nil {
            print(parentNode.Data, ", ")
            stack.push(parentNode)
            parentNode = parentNode.Left
        // 到达最左叶子节点后开始访问右子树的节点，以同样的方式访问右子树
        } else {
            parentNode = stack.pop()
            parentNode = parentNode.Right
        }
    }
}

func MiddleOrderUnrecursive(parentNode *BTreeNode) {
    stack := NewBTreeNodeStack()
    for parentNode != nil || !stack.isEmpty() {
        // 一直纵向延伸到左子树的最左叶子节点
        if parentNode != nil {
            stack.push(parentNode)
            parentNode = parentNode.Left
        // 到达最左叶子节点后开始访问节点，并切换到右子树
        } else {
            parentNode = stack.pop()
            print(parentNode.Data, ", ")
            parentNode = parentNode.Right
        }
    }
}

func PostOrderUnrecursive(parentNode *BTreeNode) {
    stack := NewBTreeNodeStack()
    var lastVisitNode *BTreeNode
    for parentNode != nil || !stack.isEmpty() {
        if parentNode != nil {
            stack.push(parentNode)
            parentNode = parentNode.Left
        } else {
            // 如果当前节点的右节点已经被访问过或者没有右节点，则说明它是当前子树的左节点，并访问它
            if parentNode.Right == nil || parentNode.Right == lastVisitNode {
                print(parentNode.Data, ", ")
                lastVisitNode = parentNode
            // 访问右子树
            } else {
                stack.push(parentNode)
                parentNode = parentNode.Right
                // 从右子树的最左孩子节点还是访问
                for parentNode != nil {
                    parentNode = parentNode.Left
                }
            }
        }
    }
}

func BTreeBFS(node *BTreeNode) {
    if node.IsNilNode() {
        return
    }
    queue := NewBTreeNodeQueue()
    queue.enqueue(node)
    var curNode *BTreeNode
    for !queue.isEmpty() {
        curNode = queue.dequeue()
        print(curNode.Data, ", ")
        queue.enqueue(curNode.Left)
        queue.enqueue(curNode.Right)
    }
}
