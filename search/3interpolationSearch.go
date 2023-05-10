package search

// 插值查找算法（Interpolation Search）：
// 要求查找的数组或列表必须是有序的，并且元素的分布应该是均匀的，
// 根据目标值在有序数组或列表中的大致位置，计算出一个插值作为下标进行查找。

func InterpolationSearch(arr []int, target int) int {
	n := len(arr)
	low := 0
	high := n - 1
	for low <= high && target >= arr[low] && target <= arr[high] {
		//相对于二分查找的1/2改为自适应的
		pos := low + int(float64(target-arr[low])/(float64(arr[high]-arr[low]))*float64(high-low))
		if arr[pos] == target {
			return pos
		} else if arr[pos] < target {
			low = pos + 1
		} else {
			high = pos - 1
		}
	}
	return -1
}

// func main() {
// 	arr := []int{1, 2, 3, 4, 5, 7, 8, 10}
// 	target := 7
// 	index := InterpolationSearch(arr, target)
// 	if index == -1 {
// 		fmt.Println("Not found")
// 	} else {
// 		fmt.Printf("Found at index %d\n", index)
// 	}
// }
