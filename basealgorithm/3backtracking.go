package basealgorithm

// 回溯算法（Backtracking Algorithm）：
// 在问题的解空间中搜索问题的解，通过回溯的方式寻找最优解的一种算法。

// 1确定问题的解空间（即问题的所有可能解集）。

// 2确定搜索起点，开始从搜索起点进入解空间进行搜索。

// 3确定搜索方向和策略，按照一定的规则对解空间进行搜索，找到符合条件的解。

// 4判断搜索到的解是否符合要求，如果符合要求，则保存解并结束搜索；如果不符合要求，则撤销当前步骤，回溯到上一步，重新选择其他可能的路径，继续搜索。

// 5继续重复步骤3和步骤4，直到找到所有符合要求的解或者搜索完所有可能的路径。
func SolveSudoku(board [][]int) bool {
	// 遍历数独中的每个位置
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			// 如果当前位置是空白位置
			if board[i][j] == empty {
				// 遍历数字1到9
				for num := 1; num <= 9; num++ {
					// 如果数字num可以放在当前位置
					if isValid(board, i, j, num) {
						// 放置数字num在当前位置
						board[i][j] = num

						// 递归尝试放置数字num在下一个空白位置
						if SolveSudoku(board) {
							return true
						}

						// 如果无法解决数独，回溯并尝试下一个数字
						board[i][j] = empty
					}
				}

				// 如果所有数字都尝试过了还无法解决数独，回溯
				return false
			}
		}
	}

	// 如果遍历完整个数独都没有找到空白位置，数独已经解决
	return true
}

func isValid(board [][]int, row int, col int, num int) bool {
	// 检查行是否有重复数字
	for i := 0; i < n; i++ {
		if board[row][i] == num {
			return false
		}
	}

	// 检查列是否有重复数字
	for i := 0; i < 9; i++ {
		if board[i][col] == num {
			return false
		}
	}
	// 判断在该3x3宫格中是否已经出现过num
	r := (row / 3) * 3
	c := (col / 3) * 3
	for i := r; i < r+3; i++ {
		for j := c; j < c+3; j++ {
			if board[i][j] == num {
				return false
			}
		}
	}
	// 如果符合数独规则，返回true
	return true
}

const n = 9     // 数独的大小
const empty = 0 // 空白位置的值

// func main() {
// 	// 数独的初始状态，0代表空白位置
// 	board := [][]int{
// 		{5, 3, 0, 0, 7, 0, 0, 0, 0},
// 		{6, 0, 0, 1, 9, 5, 0, 0, 0},
// 		{0, 9, 8, 0, 0, 0, 0, 6, 0},
// 		{8, 0, 0, 0, 6, 0, 0, 0, 3},
// 		{4, 0, 0, 8, 0, 3, 0, 0, 1},
// 		{7, 0, 0, 0, 2, 0, 0, 0, 6},
// 		{0, 6, 0, 0, 0, 0, 2, 8, 0},
// 		{0, 0, 0, 4, 1, 9, 0, 0, 5},
// 		{0, 0, 0, 0, 8, 0, 0, 7, 9},
// 	}

// 	// 解决数独
// 	SolveSudoku(board)

// 	// 输出解决后的数独
// 	for _, row := range board {
// 		fmt.Println(row)
// 	}
// }
