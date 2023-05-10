package datastructure

type Stack []int

// 将元素压入栈中
func (s *Stack) Push(element int) {
	//向s中添加元素
	//*s 取s指向的值
	*s = append(*s, element)
}

// 弹出栈顶元素
func (s *Stack) Pop() int {
	if len(*s) == 0 {
		panic("stack is empty")
	}
	//不为空则
	index := len(*s) - 1
	//取出末尾，即栈顶元素
	element := (*s)[index]
	//栈减一
	*s = (*s)[:index]
	return element
}

// 获取栈顶元素
func (s *Stack) Peek() int {
	if len(*s) == 0 {
		panic("stack is empty")
	}
	index := len(*s) - 1
	//栈不需要减
	return (*s)[index]
}

// 判断栈是否为空
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// 获取栈的大小
func (s *Stack) Size() int {
	return len(*s)
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
