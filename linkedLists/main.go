package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (l *LinkedList) prepend(n *Node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l LinkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
}

func main() {
	myList := LinkedList{}
	node1 := &Node{data: 12}
	node2 := &Node{data: 99}
	node3 := &Node{data: 45}
	node4 := &Node{data: 39}
	node5 := &Node{data: 22}

	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	myList.prepend(node5)

	myList.printListData()

}
