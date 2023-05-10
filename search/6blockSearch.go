package search

// 分块查找算法（Block Search）：
// 将查找区间划分成若干个块，每个块内部元素有序，
// 块与块之间元素无序，通过快速定位目标值所在的块，
// 然后在该块内部使用顺序查找算法进行查找。

import (
	"math"
)

func BlockSearch(arr []int, x int) int {
	n := len(arr)
	bSize := int(math.Sqrt(float64(n)))
	blocks := make([]int, bSize)

	// divide the array into blocks
	for i := 0; i < n; i++ {
		blocks[i/bSize] += arr[i]
	}

	// find the block where x is located
	blockIndex := -1
	for i := 0; i < bSize; i++ {
		if x <= blocks[i] {
			blockIndex = i
			break
		}
	}

	// if x is not found in any block, return -1
	if blockIndex == -1 {
		return -1
	}

	// search for x in the block using linear search
	start := blockIndex * bSize
	end := int(math.Min(float64(start+bSize), float64(n)))
	for i := start; i < end; i++ {
		if arr[i] == x {
			return i
		}
	}

	// x not found in the block, return -1
	return -1
}

// func main() {
//     arr := []int{4, 3, 6, 8, 9, 1, 2, 5, 7}
//     x := 6
//     index := BlockSearch(arr, x)
//     if index != -1 {
//         fmt.Printf("%d is found at index %d\n", x, index)
//     } else {
//         fmt.Printf("%d is not found in the array\n", x)
//     }
// }
