package basealgorithm

// 动态规划算法（Dynamic Programming Algorithm）：
// 用于处理具有重叠子问题和最优子结构性质的问题，
// 如背包问题、最长公共子序列等。

// 1定义问题的状态：确定要优化什么指标，以及这个指标与哪些变量相关，将相关的变量定义为状态。

// 2定义状态转移方程：确定各个状态之间的转移关系，也就是某个状态下的最优解与其他状态的最优解的关系。这个关系通常是一个递推式。

// 3确定边界状态：边界状态是问题中最简单的状态，通常是问题中的某个限制导致的。

// 4确定问题的最优解：通过递推计算出所有可能的状态的最优解，找到问题的最优解。

// 5可选：回溯计算最优解的路径。

func KnapsackD(weights []int, values []int, capacity int) int {
	n := len(weights)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, capacity+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= capacity; j++ {
			if weights[i-1] <= j {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weights[i-1]]+values[i-1])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[n][capacity]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func main() {
// 	weights := []int{2, 3, 4, 5}
// 	values := []int{3, 4, 5, 6}
// 	capacity := 8

// 	fmt.Println(Knapsack(weights, values, capacity))
// }
