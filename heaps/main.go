package main

import "fmt"

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if h.array[l] > h.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}

	}
}

func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	last := len(h.array) - 1
	if last == -1 {
		fmt.Println("Empty Array!")
	}
	h.array[0] = h.array[last]
	h.array = h.array[:last]
	h.maxHeapifyDown(0)
	return extracted
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return i*2 + 1
}

func right(i int) int {
	return i*2 + 2
}

func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func main() {
	myHeap := &MaxHeap{}
	buildArr := []int{10, 20, 30, 45, 23, 13, 56, 78, 33, 21}
	for _, val := range buildArr {
		myHeap.Insert(val)
	}
	fmt.Println(myHeap)
	fmt.Println(myHeap.Extract())
	fmt.Println(myHeap)
}
