package datastructure

import "fmt"

// 定义链表节点结构体
type Node struct {
	Data int
	next *Node
}

// 定义链表结构体
type LinkedList struct {
	head *Node
}

// 在链表尾部插入新节点
func (l *LinkedList) InsertAtEnd(data int) {
	//定义一个新结点
	newNode := &Node{data, nil}
	//如果传入结点空，则插入结点为新结点
	if l.head == nil {
		l.head = newNode
		return
	}
	// 不为空的话，定义一个当前结点=传入结点
	curr := l.head
	//循环到结点尾部
	for curr.next != nil {
		curr = curr.next
	}
	//插入结点
	curr.next = newNode
}

// 在链表头部插入新节点
func (l *LinkedList) InsertAtBeginning(data int) {
	newNode := &Node{data, nil}
	//如果传入结点空，则插入结点为新结点
	if l.head == nil {
		l.head = newNode
		return
	}
	//新建结点的next指向传入结点head
	newNode.next = l.head
	//head指针移动到新建结点头
	l.head = newNode

}

// 删除指定值的节点
func (l *LinkedList) DeleteNode(data int) {
	if l.head == nil {
		return
	}
	//如果头结点值等于要删除的值，头指针后移动
	if l.head.Data == data {
		l.head = l.head.next
		return
	}
	//新建prev=传入链表头，当前curr=传入链表下一个结点
	prev := l.head
	curr := l.head.next
	//循环直到链表结束
	for curr != nil {
		//找到删除的结点
		if curr.Data == data {
			//删除找到的结点
			prev.next = curr.next
			return
		}
		//如果不是则向后移动
		prev = curr
		curr = curr.next
	}
}

// 查找指定值的节点
func (l *LinkedList) SearchNode(data int) *Node {
	curr := l.head
	for curr != nil {
		if curr.Data == data {
			//返回当前结点地址
			return curr
		}
		curr = curr.next
	}
	return nil
}

// 打印链表中的所有节点值
func (l *LinkedList) PrintList() {
	curr := l.head
	for curr != nil {
		fmt.Printf("%d ", curr.Data)
		curr = curr.next
	}
	fmt.Println()
}

// func main() {
// 	list := LinkedList{}
// 	list.InsertAtEnd(1)
// 	list.InsertAtEnd(2)
// 	list.InsertAtEnd(3)
// 	list.InsertAtBeginning(0)
// 	list.PrintList() // 输出：0 1 2 3
// 	node := list.SearchNode(2)
// 	if node != nil {
// 		node.Data = 4
// 	}
// 	list.PrintList() // 输出：0 1 4 3
// 	list.DeleteNode(1)
// 	list.PrintList() // 输出：0 4 3
// }
