package main

import "testing"

func TestLRU(t *testing.T) {
	lru := NewLRU(3)

	if lru.Get("foo") != nil {
		t.Error("Expected nil for non-existent key")
	}

	lru.Update("foo", 69)
	if val := lru.Get("foo"); val != 69 {
		t.Errorf("Expected 69, got %v", val)
	}

	lru.Update("bar", 420)
	if val := lru.Get("bar"); val != 420 {
		t.Errorf("Expected 420, got %v", val)
	}

	lru.Update("baz", 1337)
	if val := lru.Get("baz"); val != 1337 {
		t.Errorf("Expected 1337, got %v", val)
	}

	lru.Update("ball", 69420)
	if val := lru.Get("ball"); val != 69420 {
		t.Errorf("Expected 69420, got %v", val)
	}

	if val := lru.Get("foo"); val != nil {
		t.Error("Expected nil for non-existent key")
	}

	if val := lru.Get("bar"); val != 420 {
		t.Errorf("Expected 420, got %v", val)
	}

	lru.Update("foo", 69)
	if val := lru.Get("bar"); val != 420 {
		t.Errorf("Expected 420, got %v", val)
	}

	if val := lru.Get("foo"); val != 69 {
		t.Errorf("Expected 69, got %v", val)
	}

	// Test that baz should have been evicted since "bar" was accessed recently
	if val := lru.Get("baz"); val != nil {
		t.Error("Expected nil for non-existent key")
	}
}
