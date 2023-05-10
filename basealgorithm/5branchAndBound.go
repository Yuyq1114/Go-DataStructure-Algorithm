package basealgorithm

// 分支限界算法（Branch and Bound Algorithm）：
// 通过扩展状态空间树的方式，
// 将问题的解空间分为若干个互不重叠的子集，
// 限制每个子集的搜索范围，从而找到最优解的一种算法。

// 1定义问题的搜索状态，并设计合适的状态空间。

// 2设计合适的优先队列，每个队列元素包含一个搜索状态和一个下界（或优先级），用于确定下一步搜索的方向。

// 3初始化优先队列，加入初始状态。

// 4从队列中取出下界最小的状态进行扩展，生成子状态并计算它们的下界或优先级。

// 5将子状态加入优先队列中，根据下界或优先级排序。

// 6如果找到一个解或者优先队列为空，则停止搜索。

// 7回到第4步，继续搜索。

type Node struct {
	level  int
	value  int
	weight int
}

//在上面的代码中，我们使用一个Node结构体表示每个节点的状态，其中包含了当前节点的层数、价值和重量。
//背包问题

func BranchAndBound(capacity int, values []int, weights []int) int {
	n := len(values)
	maxVal := 0
	queue := make([]Node, 0)
	root := Node{0, 0, 0}
	queue = append(queue, root)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.level == n {
			if node.value > maxVal {
				maxVal = node.value
			}
			continue
		}

		// 不选当前物品的节点
		if node.weight+weights[node.level] <= capacity {
			queue = append(queue, Node{node.level + 1, node.value + values[node.level], node.weight + weights[node.level]})
		}

		// 选择当前物品的节点
		queue = append(queue, Node{node.level + 1, node.value, node.weight})
	}

	return maxVal
}

// func main() {
// 	capacity := 10
// 	values := []int{10, 40, 30, 50}
// 	weights := []int{5, 4, 6, 3}
// 	maxVal := BranchAndBound(capacity, values, weights)
// 	fmt.Println(maxVal)
// }
