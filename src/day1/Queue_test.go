package main

import (
	"testing"
)

func TestQueue(t *testing.T) {
	list := NewQueue()

	// test Enqueue
	list.Enqueue(5)
	list.Enqueue(7)
	list.Enqueue(9)

	val := list.Dequeue()
	if val != 5 {
		t.Errorf("Expected dequeued value to be 5, but got %v", val)
	}

	if list.length != 2 {
		t.Errorf("Expected queue length to be 2, but got %d", list.length)
	}

	list.Enqueue(11)

	val = list.Dequeue()
	if val != 7 {
		t.Errorf("Expected dequeued value to be 7, but got %v", val)
	}

	val = list.Dequeue()
	if val != 9 {
		t.Errorf("Expected dequeued value to be 9, but got %v", val)
	}

	val = list.Peek()
	if val != 11 {
		t.Errorf("Expected peek value to be 11, but got %v", list.Peek())
	}

	val = list.Dequeue()
	if val != 11 {
		t.Errorf("Expected dequeued value to be 11, but got %v", val)
	}

	val = list.Dequeue()
	if val != nil {
		t.Errorf("Expected dequeued value to be nil, but got %v", val)
	}

	if list.length != 0 {
		t.Errorf("Expected queue length to be 0, but got %d", list.length)
	}

	list.Enqueue(69)

	val = list.Peek()
	if val != 69 {
		t.Errorf("Expected peek value to be 69, but got %v", list.Peek())
	}

	if list.length != 1 {
		t.Errorf("Expected queue length to be 1, but got %d", list.length)
	}

	// test empty queue
	list = NewQueue()

	val = list.Dequeue()

	if val != nil {
		t.Errorf("Expected dequeued value to be nil, but got %v", val)
	}

	if list.Peek() != nil {
		t.Errorf("Expected peek value to be nil, but got %v", list.Peek())
	}
}
