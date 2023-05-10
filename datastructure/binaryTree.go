package datastructure

import "fmt"

type NodeT struct {
	value int
	left  *NodeT
	right *NodeT
}

type Tree struct {
	root *NodeT
}

// 插入节点
func (t *Tree) Insert(value int) {
	newNode := &NodeT{value: value}
	if t.root == nil {
		t.root = newNode
	} else {
		t.root.insertNode(newNode)
	}
}

// 递归插入节点
func (n *NodeT) insertNode(newNode *NodeT) {
	//插入的值小于当前结点值，则左子树
	if newNode.value < n.value {
		if n.left == nil {
			//结点的左结点为空则插入
			n.left = newNode
		} else {
			//不为空则递归继续左
			n.left.insertNode(newNode)
		}
	} else {
		//插入的值大于等于当前结点值，则右子树
		if n.right == nil {
			//结点的右结点为空则插入
			n.right = newNode
		} else {
			//不为空则递归继续右
			n.right.insertNode(newNode)
		}
	}
}

// 查找节点
func (t *Tree) Search(value int) bool {
	return t.root.searchNode(value)
}

// 递归查找节点
func (n *NodeT) searchNode(value int) bool {
	if n == nil {
		return false
	}
	if n.value == value {
		return true
	} else if value < n.value {
		return n.left.searchNode(value)
	} else {
		return n.right.searchNode(value)
	}
}

// 遍历
func (t *Tree) InorderTraversal() {
	t.root.inorderTraversal()
}
func (t *Tree) PreorderTraversal() {
	t.root.preorderTraversal()
}
func (t *Tree) PostorderTraversal() {
	t.root.postorderTraversal()
}

// 递归中序遍历
func (n *NodeT) inorderTraversal() {
	if n != nil {
		n.left.inorderTraversal()
		fmt.Printf("%d ", n.value)
		n.right.inorderTraversal()
	}
}

// 先序遍历
func (n *NodeT) preorderTraversal() {
	if n != nil {
		fmt.Printf("%d ", n.value)
		n.left.preorderTraversal()
		n.right.preorderTraversal()
	}
}

// 后序遍历
func (n *NodeT) postorderTraversal() {
	if n != nil {
		n.left.postorderTraversal()
		n.right.postorderTraversal()
		fmt.Printf("%d ", n.value)
	}
}

//求树高度
func (t *Tree) GetHeight() int {
	return t.root.getHeight()
}

//求树高度
func (n *NodeT) getHeight() int {
	if n == nil {
		return 0
	}
	//左子树高度递归
	leftHeight := n.left.getHeight()
	//右子树高度递归
	rightHeight := n.right.getHeight()

	//左大于右 左加1
	if leftHeight > rightHeight {
		return leftHeight + 1
	} else {
		return rightHeight + 1
	}
}

// func main() {
// 	t := Tree{}
// 	t.Insert(5)
// 	t.Insert(3)
// 	t.Insert(7)
// 	t.Insert(1)
// 	t.Insert(9)
// 	fmt.Println(t.Search(7)) // 输出：true
// 	fmt.Println(t.Search(4)) // 输出：false
// 	t.InorderTraversal()     // 输出：1 3 5 7 9
//  t.PostorderTraversal()
//  t.PreorderTraversal()
//  fmt.Printf(t.GetHeight())
// }
