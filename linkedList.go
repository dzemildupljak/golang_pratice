package main

import "fmt"

type node struct {
	next *node
	data int
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(value int) {
	newNode := node{data: value}
	if l.head != nil {
		newNode.next = l.head
		l.head = &newNode
	} else {
		l.head = &newNode
	}
	l.length++
}
func (l *linkedList) printList() {
	if l.head == nil {
		return
	}
	currNode := l.head
	for currNode != nil {
		fmt.Printf("%d - ", currNode.data)
		currNode = currNode.next
	}
	fmt.Println()

}

func mainLL() {

	ls := linkedList{
		head:   nil,
		length: 0,
	}

	ls.prepend(14)
	ls.prepend(23)
	ls.prepend(87)
	ls.prepend(56)

	ls.printList()

	fmt.Println(ls.head.data)
}
