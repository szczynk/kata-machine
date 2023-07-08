package main

type node2 struct {
	value interface{} // Value of the node
	prev  *node2      // Pointer to the previous node
	next  *node2      // Pointer to the next node
}

type Stack struct {
	length int    // Length of the stack
	head   *node2 // Pointer to the head (top) of the stack
}

func NewStack() *Stack {
	return &Stack{
		length: 0,
		head:   nil,
	}
}

func (s *Stack) Push(item interface{}) {
	node := &node2{value: item} // Create a new node with the given value

	s.length++ // Increment the length of the stack

	if s.head == nil {
		s.head = node // If stack is empty, set the new node as the head
	}

	node.prev = s.head // Set the previous node of the new node to the current head
	s.head = node      // Set the new node as the new head of the stack
}

func (s *Stack) Pop() interface{} {
	if s.length == 0 {
		return nil // If stack is empty, return nil
	}

	s.length-- // Decrement the length of the stack

	head := s.head     // Store the current head of the stack
	s.head = head.prev // Set the previous node of the current head as the new head

	head.prev = nil // Set the previous node of the current head to nil to disconnect it from the stack

	return head.value // Return the value of the popped node
}

func (s *Stack) Peek() interface{} {
	if s.head != nil {
		return s.head.value // If stack is not empty, return the value of the head
	}
	return nil // If stack is empty, return nil
}
