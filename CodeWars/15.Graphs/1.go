package main

import "fmt"

type Graph struct {
	vertices []*Vertex
	mapper   map[int]*Vertex
}

type Vertex struct {
	key      int
	adjacent []*Vertex
	mapper   map[int]struct{}
}

func (g *Graph) AddVertex(k int) {
	//if k already in graph return
	if _, ok := g.mapper[k]; ok {
		return
	}
	m := make(map[int]struct{})
	v := &Vertex{key: k, mapper: m}
	g.vertices = append(g.vertices, v)
	g.mapper[k] = v
}

func (g *Graph) AddEdge(from, to int) {
	if _, ok := g.mapper[from]; !ok {
		return
	}
	vF := g.mapper[from]
	if _, ok := g.mapper[to]; !ok {
		return
	}
	vT := g.mapper[to]

	//check if vT already in adjacent list of vF
	if _, ok := vF.mapper[to]; ok {
		return
	}
	//adding edge to vF
	vF.adjacent = append(vF.adjacent, vT)
	vF.mapper[to] = struct{}{}
}

func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("Vertex %d:", v.key)
		for _, v := range v.adjacent {
			fmt.Printf(" %d ", v.key)
		}
		fmt.Println("")
	}
}

func NewGraph() *Graph {
	return &Graph{
		mapper: make(map[int]*Vertex),
	}
}

func main() {
	fmt.Println("hello")
	g := NewGraph()
	for i := 0; i < 5; i++ {
		g.AddVertex(i)
	}
	g.AddEdge(0, 1)
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(1, 7)
	g.Print()
}
