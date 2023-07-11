package main

import "fmt"

type Node struct {
	value interface{}
	prev  *Node
	next  *Node
}

type DoublyLinkedList struct {
	length int
	head   *Node
	tail   *Node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		length: 0,
		head:   nil,
		tail:   nil,
	}
}

func (dll *DoublyLinkedList) Prepend(item interface{}) {
	dll.length++

	node := &Node{value: item}

	if dll.head == nil {
		dll.head = node
		dll.tail = node
		return
	}

	node.next = dll.head
	dll.head.prev = node
	dll.head = node
}

func (dll *DoublyLinkedList) InsertAt(item interface{}, idx int) {
	if idx > dll.length {
		panic("Index out of bounds")
	}

	if idx == dll.length {
		dll.Append(item)
		return
	} else if idx == 0 {
		dll.Prepend(item)
		return
	}

	dll.length++

	current := dll.getAt(idx)

	node := &Node{value: item}

	node.next = current
	node.prev = current.prev
	current.prev = node

	if node.prev != nil {
		node.prev.next = node
	}
}

func (dll *DoublyLinkedList) Append(item interface{}) {
	dll.length++

	node := &Node{value: item}

	if dll.tail == nil {
		dll.head = node
		dll.tail = node
		return
	}

	node.prev = dll.tail
	dll.tail.next = node
	dll.tail = node
}

func (dll *DoublyLinkedList) Remove(item interface{}) interface{} {
	current := dll.head

	for i := 0; current != nil && i < dll.length; i++ {
		if current.value == item {
			break
		}
		current = current.next
	}

	if current == nil {
		return nil
	}

	return dll.removeNode(current)
}

func (dll *DoublyLinkedList) RemoveAt(idx int) interface{} {
	node := dll.getAt(idx)

	if node == nil {
		return nil
	}

	return dll.removeNode(node)
}

func (dll *DoublyLinkedList) removeNode(node *Node) interface{} {
	dll.length--

	if dll.length == 0 {
		value := dll.head.value
		dll.head = nil
		dll.tail = nil
		return value
	}

	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}

	if node == dll.head {
		dll.head = node.next
	}
	if node == dll.tail {
		dll.tail = node.prev
	}

	node.prev = nil
	node.next = nil

	return node.value
}

func (dll *DoublyLinkedList) Get(idx int) interface{} {
	node := dll.getAt(idx)
	if node != nil {
		return node.value
	}
	return nil
}

func (dll *DoublyLinkedList) getAt(idx int) *Node {
	current := dll.head

	for i := 0; current != nil && i < idx; i++ {
		current = current.next
	}

	return current
}

func (dll *DoublyLinkedList) Debug() {
	current := dll.head
	for i := 0; current != nil; i++ {
		fmt.Printf("%d => %v, ", i, current.value)
		current = current.next
	}
	fmt.Println()
}
