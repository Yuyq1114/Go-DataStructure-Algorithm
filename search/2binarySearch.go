package search

// 二分查找算法（Binary Search）：
// 要求查找的数组或列表必须是有序的，每次将查找区间缩小一半，
// 直到找到目标值或查找区间为空。

func BinarySearch(arr []int, target int) int {
	n := len(arr)
	low := 0
	high := n - 1
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// func main() {
// 	arr := []int{1, 2, 3, 4, 5, 7, 8, 10}
// 	target := 7
// 	index := BinarySearch(arr, target)
// 	if index == -1 {
// 		fmt.Println("Not found")
// 	} else {
// 		fmt.Printf("Found at index %d\n", index)
// 	}
// }
