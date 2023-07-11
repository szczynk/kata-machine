package main

import (
	"testing"
)

func TestDoublyLinkedList(t *testing.T) {
	list := NewDoublyLinkedList()

	list.Append(5)
	list.Append(7)
	list.Append(9)

	if value := list.Get(2); value != 9 {
		t.Errorf("Expected value at index 2 to be 9, but got %v", value)
	}

	if value := list.RemoveAt(1); value != 7 {
		t.Errorf("Expected removed value at index 1 to be 7, but got %v", value)
	}

	if length := list.length; length != 2 {
		t.Errorf("Expected length to be 2, but got %v", length)
	}

	list.Append(11)

	if value := list.RemoveAt(1); value != 9 {
		t.Errorf("Expected removed value at index 1 to be 9, but got %v", value)
	}

	if value := list.Remove(9); value != nil {
		t.Errorf("Expected value 9 to be removed, but got %v", value)
	}

	if value := list.RemoveAt(0); value != 5 {
		t.Errorf("Expected removed value at index 0 to be 5, but got %v", value)
	}

	if value := list.RemoveAt(0); value != 11 {
		t.Errorf("Expected removed value at index 0 to be 11, but got %v", value)
	}

	if length := list.length; length != 0 {
		t.Errorf("Expected length to be 0, but got %v", length)
	}

	list.Prepend(5)
	list.Prepend(7)
	list.Prepend(9)

	if value := list.Get(2); value != 5 {
		t.Errorf("Expected value at index 2 to be 5, but got %v", value)
	}

	if value := list.Get(0); value != 9 {
		t.Errorf("Expected value at index 0 to be 9, but got %v", value)
	}

	if value := list.Remove(9); value != 9 {
		t.Errorf("Expected value 9 to be removed, but got %v", value)
	}

	if length := list.length; length != 2 {
		t.Errorf("Expected length to be 2, but got %v", length)
	}

	if value := list.Get(0); value != 7 {
		t.Errorf("Expected value at index 0 to be 7, but got %v", value)
	}
}
