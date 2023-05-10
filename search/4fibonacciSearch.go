package search

// 斐波那契查找算法（Fibonacci Search）：
// 要求查找的数组或列表必须是有序的，
// 将查找区间分成两个斐波那契数列的长度，
// 然后在两个区间中选择一个继续查找，直到找到目标值或查找区间为空。

// 生成斐波那契数列
func generateFibonacci(n int) []int {
	fib := []int{1, 1}
	for i := 2; i < n; i++ {
		fib = append(fib, fib[i-1]+fib[i-2])
	}
	return fib
}

func FibonacciSearch(arr []int, target int) int {
	n := len(arr)
	low := 0
	high := n - 1
	k := 0
	fib := generateFibonacci(n)

	// 找到第一个大于等于n的斐波那契数
	for fib[k]-1 < n {
		k++
	}

	// 将数组长度扩展到fib[k]-1的长度
	for i := n; i < fib[k]-1; i++ {
		arr = append(arr, arr[n-1])
	}

	// 使用斐波那契分割数mid将区间划分为F(k-1)-1和F(k-2)-1两部分
	for low <= high {

		//mid被划分为斐波那契的部分
		mid := low + fib[k-1] - 1
		if target < arr[mid] {
			high = mid - 1
			k -= 1
		} else if target > arr[mid] {
			low = mid + 1
			k -= 2
		} else {
			if mid < n {
				return mid
			}
			return n - 1
		}
	}
	return -1
}

// func main() {
//     arr := []int{1, 3, 5, 7, 9, 11, 13}
//     target := 7
//     index := FibonacciSearch(arr, target)
//     if index == -1 {
//         fmt.Println("Not found")
//     } else {
//         fmt.Printf("Found at index %d\n", index)
//     }
// }
