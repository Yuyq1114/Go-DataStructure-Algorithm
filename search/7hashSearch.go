package search

// 哈希查找算法（Hash Search）：
// 将每个元素的关键字作为索引存储在哈希表中，通过计算目标值的哈希值，
// 快速定位目标值所在的位置，
// 如果该位置有多个元素，则使用其他算法解决冲突问题，如拉链法或开放定址法。

type Node struct {
	Key   int
	Value interface{}
	Next  *Node
}

type HashTable struct {
	Table []*Node
	Size  int
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		Table: make([]*Node, size),
		Size:  size,
	}
}

func (ht *HashTable) hashFunction(key int) int {
	return key % ht.Size
}

func (ht *HashTable) Insert(key int, value interface{}) {
	index := ht.hashFunction(key)
	node := &Node{Key: key, Value: value}

	if ht.Table[index] == nil {
		ht.Table[index] = node
	} else {
		// Insert the new node at the beginning of the linked list
		node.Next = ht.Table[index]
		ht.Table[index] = node
	}
}

func (ht *HashTable) Search(key int) interface{} {
	index := ht.hashFunction(key)
	node := ht.Table[index]

	// Traverse the linked list at this index until we find the node with the matching key
	for node != nil {
		if node.Key == key {
			return node.Value
		}
		node = node.Next
	}

	// The key was not found in the hash table
	return nil
}

// func main() {
// 	ht := NewHashTable(10)

// 	ht.Insert(1, "One")
// 	ht.Insert(2, "Two")
// 	ht.Insert(3, "Three")
// 	ht.Insert(4, "Four")
// 	ht.Insert(5, "Five")

// 	fmt.Println(ht.Search(3)) // Output: Three
// 	fmt.Println(ht.Search(6)) // Output: nil
// }
