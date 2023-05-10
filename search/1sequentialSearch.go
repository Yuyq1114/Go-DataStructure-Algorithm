package search

// 顺序查找算法（Sequential Search）：
// 遍历数组或列表中的每个元素，依次与目标值比较，
// 直到找到目标值或遍历完整个数组或列表。
func SequentialSearch(arr []int, target int) int {
	n := len(arr)
	for i := 0; i < n; i++ {
		if arr[i] == target {
			return i
		}
	}
	return -1
}

// func main() {
//     arr := []int{5, 3, 8, 4, 2, 1, 10, 7}
//     target := 8
//     index := SequentialSearch(arr, target)
//     if index == -1 {
//         fmt.Println("Not found")
//     } else {
//         fmt.Printf("Found at index %d\n", index)
//     }
// }
