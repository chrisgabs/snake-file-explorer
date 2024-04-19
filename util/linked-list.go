package util

import "fmt"

// Node

type Node[T any] struct {
	Data T
	Next *Node[T]
	Prev *Node[T]
}

func (node *Node[T]) String() string {
	prev := "nil"
	next := "nil"
	if node.Prev != nil {
		prev = fmt.Sprintf("%v", node.Prev.Data)
	}
	if node.Next != nil {
		next = fmt.Sprintf("%v", node.Next.Data)
	}
	return fmt.Sprintf("(%v, %v, %v)", prev, node.Data, next)
}

// LinkedList

type LinkedList[T any] struct {
	Head *Node[T]
	Tail *Node[T]
	Size int
}

func (linkedList *LinkedList[T]) InsertFront(data T) {
	prevHead := linkedList.Head
	newNode := &Node[T]{
		Data: data,
		Next: nil,
		Prev: prevHead,
	}
	if linkedList.isEmpty() {
		linkedList.Head = newNode
		linkedList.Tail = newNode
	}else {
		prevHead.Next = newNode
		linkedList.Head = newNode
	}
	linkedList.Size++
}

func (linkedList *LinkedList[T]) InsertBack(data T) {
	prevTail := linkedList.Tail
	newNode := &Node[T]{
		Data: data,
		Next: prevTail,
		Prev: nil,
	}
	if linkedList.isEmpty() {
		linkedList.Head = newNode
		linkedList.Tail = newNode
	}else {
		prevTail.Prev = newNode
		linkedList.Tail = newNode
	}
	linkedList.Size++
}

func (linkedList *LinkedList[T]) isEmpty() bool {
	return linkedList.Size == 0
}

func (linkedList *LinkedList[T]) String() string {
	if linkedList.isEmpty() {
		return "[empty]"
	}

	curr := linkedList.Head
	s := "[\n"
	for curr != nil {

		s += fmt.Sprintf("%v\n", curr)
		curr = curr.Prev
	}
	s += "]"

	return s
}