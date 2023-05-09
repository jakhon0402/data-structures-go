package main

import "fmt"

type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	key      int
	adjacent []*Vertex
}

func (g *Graph) AddVertex(k int) {
	if contains(g.vertices, k) {
		err := fmt.Errorf("Vertex %v not added because it is already exists!", k)
		fmt.Println(err.Error())
	} else {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}
}

func (g *Graph) AddEdge(fromKey, toKey int) {
	fromVertex, toVertex := g.getVertex(fromKey), g.getVertex(toKey)
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("Invalid edge (%v -> %v)", fromKey, toKey)
		fmt.Println(err.Error())
	} else if contains(fromVertex.adjacent, toKey) {
		err := fmt.Errorf("Existing edge (%v -> %v)", fromKey, toKey)
		fmt.Println(err.Error())
	} else {
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)

	}

}

func (g *Graph) Print() {
	for _, val := range g.vertices {
		fmt.Printf("\nVertex %v: ", val.key)
		for _, v := range val.adjacent {
			fmt.Printf(" %v ", v.key)
		}
	}
}

func (g *Graph) getVertex(k int) *Vertex {
	for i, val := range g.vertices {
		if val.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

func contains(s []*Vertex, key int) bool {
	for _, val := range s {
		if val.key == key {
			return true
		}
	}
	return false
}

func main() {
	myGraph := &Graph{}
	for i := 0; i < 5; i++ {
		myGraph.AddVertex(i)
	}
	myGraph.AddEdge(1, 2)
	myGraph.AddEdge(1, 2)

	myGraph.Print()

}
