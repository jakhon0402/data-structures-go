package main

import "fmt"

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(key int) {
	if n.Key < key {
		if n.Right == nil {
			n.Right = &Node{Key: key}
		} else {
			n.Right.Insert(key)
		}
	} else {
		if n.Left == nil {
			n.Left = &Node{Key: key}
		} else {
			n.Left.Insert(key)
		}
	}
}

func (n *Node) Search(key int) bool {
	if n == nil {
		return false
	}
	if n.Key < key {
		return n.Right.Search(key)
	} else if n.Key > key {
		return n.Left.Search(key)
	}
	return true
}

func main() {
	tree := &Node{Key: 100}
	tree.Insert(123)
	tree.Insert(99)
	tree.Insert(13)
	tree.Insert(1)
	tree.Insert(255)

	fmt.Println(tree)

	fmt.Println(tree.Search(255))

}
