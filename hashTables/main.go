package main

import (
	"fmt"
)

const ArraySize = 7

type HashTable struct {
	array [ArraySize]*Bucket
}

type Bucket struct {
	headNode *BucketNode
}

type BucketNode struct {
	key  string
	next *BucketNode
}

func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].Insert(key)
}

func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].Search(key)
}

func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].Delete(key)
}

func (b *Bucket) Insert(key string) {
	if !b.Search(key) {
		newNode := &BucketNode{key: key}
		newNode.next = b.headNode
		b.headNode = newNode
	} else {
		fmt.Println("Key alreadey exists!")
	}
}

func (b *Bucket) Search(key string) bool {
	currentNode := b.headNode
	for currentNode != nil {
		if currentNode.key == key {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

func (b *Bucket) Delete(key string) {
	if b.headNode.key == key {
		b.headNode = b.headNode.next
		return
	}
	prevNode := b.headNode
	for prevNode.next != nil {
		if prevNode.next.key == key {
			prevNode.next = prevNode.next.next
		}
	}

}

func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

func Init() *HashTable {
	hashTable := &HashTable{}
	for i := range hashTable.array {
		hashTable.array[i] = &Bucket{}
	}
	return hashTable
}

func main() {
	myHashTable := Init()
	arr := []string{
		"olma",
		"anor",
		"olcha",
		"gilos", "behi", "o'rik", "shaftoli",
	}
	for _, val := range arr {
		myHashTable.Insert(val)
	}
	myHashTable.Delete("shaftoli")
	fmt.Println(myHashTable.Search("shaftoli"))
}
