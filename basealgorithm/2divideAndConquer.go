package basealgorithm

// 分治算法（Divide and Conquer Algorithm）：
// 将问题分解为若干个规模较小但类似于原问题的子问题，
// 递归求解，然后将子问题的解合并起来，得到原问题的解。

// 1分解问题：将原问题分解成若干个规模较小的子问题，并解决这些子问题。这一步通常是通过递归来实现的。
// 2解决问题：对于子问题，若其规模小到足以直接解决，则直接解决，否则递归处理子问题。
// 3合并问题：将子问题的解合并成原问题的解。
// 返回 arr 中的最大值和最小值
func FindMaxAndMin(arr []int) (int, int) {
	n := len(arr)
	if n == 0 {
		return 0, 0
	} else if n == 1 {
		return arr[0], arr[0]
	} else if n == 2 {
		if arr[0] < arr[1] {
			return arr[1], arr[0]
		} else {
			return arr[0], arr[1]
		}
	} else {
		mid := n / 2
		leftMax, leftMin := FindMaxAndMin(arr[:mid])
		rightMax, rightMin := FindMaxAndMin(arr[mid:])
		if leftMax > rightMax {
			return leftMax, min(leftMin, rightMax)
		} else {
			return rightMax, min(leftMax, rightMin)
		}
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

// func main() {
// 	arr := []int{3, 2, 1, 5, 6, 4, 8, 9, 7}
// 	max, min := findMaxAndMin(arr)
// 	fmt.Println("Max:", max, "Min:", min)
// }
