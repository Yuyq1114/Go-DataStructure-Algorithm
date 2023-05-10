package datastructure

type NodeQ struct {
	value int
	next  *NodeQ
}

type QueueQ struct {
	//头结点
	front *NodeQ
	//尾结点
	rear *NodeQ
}

// 入队
func (q *QueueQ) Enqueue(element int) {
	newNode := &NodeQ{value: element}
	if q.front == nil {
		q.front = newNode
		q.rear = newNode
	} else {
		//新结点连接上队
		q.rear.next = newNode
		//队尾指针后移
		q.rear = newNode
	}
}

// 出队
func (q *QueueQ) Dequeue() int {
	if q.front == nil {
		panic("queue is empty")
	}
	//队头指向的值出队
	element := q.front.value
	//队头指针后移
	q.front = q.front.next
	if q.front == nil {
		q.rear = nil
	}
	return element
}

// 获取队列头部元素
func (q *QueueQ) Front() int {
	if q.front == nil {
		panic("queue is empty")
	}
	return q.front.value
}

// 判断队列是否为空
func (q *QueueQ) IsEmpty() bool {
	//判断队头指针
	return q.front == nil
}

// 获取队列大小
func (q *QueueQ) Size() int {
	size := 0
	curr := q.front
	//遍历
	for curr != nil {
		size++
		curr = curr.next
	}
	return size
}

// func main() {
//     q := QueueQ{}
//     q.Enqueue(1)
//     q.Enqueue(2)
//     q.Enqueue(3)
//     fmt.Println(q.Size())  // 输出：3
//     fmt.Println(q.Front()) // 输出：1
//     fmt.Println(q.Dequeue())  // 输出：1
//     fmt.Println(q.IsEmpty())  // 输出：false
//     q.Dequeue()
//     q.Dequeue()
//     fmt.Println(q.IsEmpty())  // 输出：true
//     fmt.Println(q.Size())  // 输出：0
// }
