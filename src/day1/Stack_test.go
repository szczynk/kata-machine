package main

import (
	"testing"
)

func TestStack(t *testing.T) {
	list := NewStack()

	list.Push(5)
	list.Push(7)
	list.Push(9)

	val := list.Pop()
	if val != 9 {
		t.Errorf("Expected popped value to be 9, but got %v", val)
	}

	if list.length != 2 {
		t.Errorf("Expected queue length to be 2, but got %d", list.length)
	}

	list.Push(11)

	val = list.Pop()
	if val != 11 {
		t.Errorf("Expected popped value to be 11, but got %v", val)
	}

	val = list.Pop()
	if val != 7 {
		t.Errorf("Expected popped value to be 7, but got %v", val)
	}

	val = list.Peek()
	if val != 5 {
		t.Errorf("Expected peek value to be 5, but got %v", list.Peek())
	}

	val = list.Pop()
	if val != 5 {
		t.Errorf("Expected popped value to be 5, but got %v", val)
	}

	val = list.Pop()
	if val != nil {
		t.Errorf("Expected popped value to be nil, but got %v", val)
	}

	list.Push(69)

	val = list.Peek()
	if val != 69 {
		t.Errorf("Expected peek value to be 69, but got %v", list.Peek())
	}

	if list.length != 1 {
		t.Errorf("Expected queue length to be 1, but got %d", list.length)
	}

}
