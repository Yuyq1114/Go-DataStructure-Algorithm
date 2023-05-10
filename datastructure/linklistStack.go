package datastructure

type NodeS struct {
	Data int
	next *NodeS
}

type StackL struct {
	top *NodeS
}

// 将元素压入栈中
func (s *StackL) Push(element int) {
	node := &NodeS{Data: element, next: s.top}
	// 栈顶指针移到插入元素
	s.top = node
}

// 弹出栈顶元素
func (s *StackL) Pop() int {
	if s.top == nil {
		panic("stack is empty")
	}
	//获取元素
	element := s.top.Data
	//指针下移
	s.top = s.top.next
	return element
}

// 获取栈顶元素
func (s *StackL) Peek() int {
	if s.top == nil {
		panic("stack is empty")
	}
	//不需要移动指针
	return s.top.Data
}

// 判断栈是否为空
func (s *StackL) IsEmpty() bool {
	//判断栈顶元素
	return s.top == nil
}

// 获取栈的大小
func (s *StackL) Size() int {
	size := 0
	node := s.top
	//遍历栈
	for node != nil {
		size++
		node = node.next
	}
	return size
}

// func main() {
//     s := Stack{}
//     s.Push(1)
//     s.Push(2)
//     s.Push(3)
//     fmt.Println(s.Size())  // 输出：3
//     fmt.Println(s.Pop())   // 输出：3
//     fmt.Println(s.Peek())  // 输出：2
//     fmt.Println(s.IsEmpty())  // 输出：false
//     s.Pop()
//     s.Pop()
//     fmt.Println(s.IsEmpty())  // 输出：true
//     fmt.Println(s.Size())  // 输出：0
// }
