package main

type node struct {
	value interface{} // Value of the node
	prev  *node       // Pointer to the previous node
	next  *node       // Pointer to the next node
}

type Queue struct {
	length int   // Length of the queue
	head   *node // Pointer to the head (front) of the queue
	tail   *node // Pointer to the tail (rear) of the queue
}

func NewQueue() *Queue {
	return &Queue{
		length: 0,
		head:   nil,
		tail:   nil,
	}
}

func (q *Queue) Enqueue(item interface{}) {
	node := &node{value: item} // Create a new node with the given value

	q.length++ // Increment the length of the queue

	if q.tail == nil {
		q.tail = node // If queue is empty, set the new node as both head and tail
		q.head = node
		return
	}

	q.tail.next = node // Set the next node of the tail as the new node
	q.tail = node      // Set the new node as the new tail of the queue
}

func (q *Queue) Dequeue() interface{} {
	if q.head == nil {
		return nil // If queue is empty, return nil
	}

	q.length-- // Decrement the length of the queue

	head := q.head       // Store the current head of the queue
	q.head = q.head.next // Set the next node of the current head as the new head

	head.next = nil // Set the next node of the current head to nil to disconnect it from the queue

	if q.length == 0 {
		q.tail = nil // If queue becomes empty after dequeue, set tail to nil
	}

	return head.value // Return the value of the dequeued node
}

func (q *Queue) Peek() interface{} {
	if q.head == nil {
		return nil // If queue is empty, return nil
	}

	return q.head.value // If queue is not empty, return the value of the head
}
