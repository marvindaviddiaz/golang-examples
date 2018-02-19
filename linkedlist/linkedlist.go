package linkedlist

import (
	"fmt"
)

type SSLNode struct {
	value int
	next *SSLNode
}

type SingleLinkedList struct {
	head *SSLNode
	tail *SSLNode
}

func newSingleLinkedList() *SingleLinkedList {
	return new(SingleLinkedList)
}

func (list *SingleLinkedList) Add(v int) {
	newNode := &SSLNode{value: v}
	if list.head == nil {
		list.head = newNode
	} else if list.tail == list.head {
		list.head.next = newNode
	} else if list.tail != nil {
		list.tail.next = newNode
	}

	list.tail = newNode
}

func (list *SingleLinkedList) String() string {
	s := ""
	for n := list.head; n != nil; n = n.next {
		s += fmt.Sprintf(" {%d} ", n.value)
	}
	return s
}

func Run() {
	list := newSingleLinkedList()
	list.Add(3)
	list.Add(4)
	list.Add(5)
	list.Add(6)
	fmt.Println("Linked List:", list)
}