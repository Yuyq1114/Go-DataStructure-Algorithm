package datastructure

import (
	"fmt"
)

type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	key      int
	adjacent []*Vertex
}

// 创建图
func NewGraph() *Graph {
	return &Graph{}
}

// 创建点
func NewVertex(key int) *Vertex {
	return &Vertex{key: key}
}

// 加点
func (g *Graph) AddVertex(v *Vertex) {
	g.vertices = append(g.vertices, v)
}

// 加边
func (g *Graph) AddEdge(v1, v2 *Vertex) {
	v1.adjacent = append(v1.adjacent, v2)
	v2.adjacent = append(v2.adjacent, v1)
}

// 深度优先搜索
func (g *Graph) DFS(start *Vertex) {
	visited := make(map[*Vertex]bool)
	g.dfs(start, visited)
}

func (g *Graph) dfs(vertex *Vertex, visited map[*Vertex]bool) {
	visited[vertex] = true
	fmt.Print(vertex.key, " ")

	for _, v := range vertex.adjacent {
		if !visited[v] {
			g.dfs(v, visited)
		}
	}
}

// 广度优先搜索
func (g *Graph) BFS(start *Vertex) {
	queue := make([]*Vertex, 0)
	visited := make(map[*Vertex]bool)

	queue = append(queue, start)
	visited[start] = true

	for len(queue) > 0 {
		v := queue[0]
		fmt.Print(v.key, " ")
		queue = queue[1:]

		for _, adj := range v.adjacent {
			if !visited[adj] {
				queue = append(queue, adj)
				visited[adj] = true
			}
		}
	}
}

// func main() {
// 	v1 := NewVertex(1)
// 	v2 := NewVertex(2)
// 	v3 := NewVertex(3)
// 	v4 := NewVertex(4)

// 	g := NewGraph()

// 	g.AddVertex(v1)
// 	g.AddVertex(v2)
// 	g.AddVertex(v3)
// 	g.AddVertex(v4)

// 	g.AddEdge(v1, v2)
// 	g.AddEdge(v2, v3)
// 	g.AddEdge(v3, v4)

// 	fmt.Println("DFS:")
// 	g.DFS(v1)

// 	fmt.Println("\nBFS:")
// 	g.BFS(v1)
// }
