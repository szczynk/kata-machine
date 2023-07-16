package main

import "testing"

func TestMinHeap(t *testing.T) {
	heap := NewMinHeap()

	if heap.length != 0 {
		t.Errorf("Expected heap length to be 0, got %d", heap.length)
	}

	heap.Insert(5)
	heap.Insert(3)
	heap.Insert(69)
	heap.Insert(420)
	heap.Insert(4)
	heap.Insert(1)
	heap.Insert(8)
	heap.Insert(7)

	if heap.length != 8 {
		t.Errorf("Expected heap length to be 8, got %d", heap.length)
	}

	if deleted := heap.Delete(); deleted != 1 {
		t.Errorf("Expected deleted element to be 1, got %d", deleted)
	}

	if deleted := heap.Delete(); deleted != 3 {
		t.Errorf("Expected deleted element to be 3, got %d", deleted)
	}

	if deleted := heap.Delete(); deleted != 4 {
		t.Errorf("Expected deleted element to be 4, got %d", deleted)
	}

	if deleted := heap.Delete(); deleted != 5 {
		t.Errorf("Expected deleted element to be 5, got %d", deleted)
	}

	if heap.length != 4 {
		t.Errorf("Expected heap length to be 4, got %d", heap.length)
	}

	if deleted := heap.Delete(); deleted != 7 {
		t.Errorf("Expected deleted element to be 7, got %d", deleted)
	}

	if deleted := heap.Delete(); deleted != 8 {
		t.Errorf("Expected deleted element to be 8, got %d", deleted)
	}

	if deleted := heap.Delete(); deleted != 69 {
		t.Errorf("Expected deleted element to be 69, got %d", deleted)
	}

	if deleted := heap.Delete(); deleted != 420 {
		t.Errorf("Expected deleted element to be 420, got %d", deleted)
	}

	if heap.length != 0 {
		t.Errorf("Expected heap length to be 0, got %d", heap.length)
	}
}
