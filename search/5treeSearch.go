package search

// 树表查找算法（Tree Search）：
// 利用二叉搜索树或平衡二叉树等数据结构，
// 通过比较目标值和根节点的大小关系来选择左子树或右子树继续查找，
// 直到找到目标值或树为空。

type TreeNode struct {
	Val   int
	left  *TreeNode
	right *TreeNode
}

func (t *TreeNode) Insert(val int) {
	if t == nil {
		return
	}

	if val < t.Val {
		if t.left == nil {
			t.left = &TreeNode{Val: val}
		} else {
			t.left.Insert(val)
		}
	} else if val > t.Val {
		if t.right == nil {
			t.right = &TreeNode{Val: val}
		} else {
			t.right.Insert(val)
		}
	}
}

func (t *TreeNode) Search(val int) *TreeNode {
	if t == nil {
		return nil
	}
	//在左子树右子树查找
	if val == t.Val {
		return t
	} else if val < t.Val {
		return t.left.Search(val)
	} else {
		return t.right.Search(val)
	}
}

// func main() {
// 	root := &TreeNode{val: 5}
// 	root.Insert(3)
// 	root.Insert(8)
// 	root.Insert(1)
// 	root.Insert(4)
// 	root.Insert(7)
// 	root.Insert(9)

// 	result := root.Search(7)
// 	if result != nil {
// 		fmt.Println("Found value:", result.val)
// 	} else {
// 		fmt.Println("Value not found")
// 	}
// }
