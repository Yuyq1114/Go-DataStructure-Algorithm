package datastructure

type Queue struct {
	elements []int
}

// 入队
func (q *Queue) Enqueue(element int) {
	//切片添加元素
	q.elements = append(q.elements, element)
}

// 出队
func (q *Queue) Dequeue() int {

	if len(q.elements) == 0 {
		panic("queue is empty")
	}
	//从切片首元素出队
	element := q.elements[0]
	//减一[1:]
	q.elements = q.elements[1:]
	return element
}

// 获取队列头部元素
func (q *Queue) Front() int {
	if len(q.elements) == 0 {
		panic("queue is empty")
	}
	return q.elements[0]
}

// 判断队列是否为空
func (q *Queue) IsEmpty() bool {
	return len(q.elements) == 0
}

// 获取队列大小
func (q *Queue) Size() int {
	return len(q.elements)
}

// func main() {
// 	q := Queue{}
// 	q.Enqueue(1)
// 	q.Enqueue(2)
// 	q.Enqueue(3)
// 	fmt.Println(q.Size())    // 输出：3
// 	fmt.Println(q.Front())   // 输出：1
// 	fmt.Println(q.Dequeue()) // 输出：1
// 	fmt.Println(q.IsEmpty()) // 输出：false
// 	q.Dequeue()
// 	q.Dequeue()
// 	fmt.Println(q.IsEmpty()) // 输出：true
// 	fmt.Println(q.Size())    // 输出：0
// }
