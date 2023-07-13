package main

import (
	"testing"
)

func TestBFS(t *testing.T) {
	if found := BFS(Tree, 45); !found {
		t.Errorf("BFS() = %v, want true", found)
	}

	if found := BFS(Tree, 7); !found {
		t.Errorf("BFS() = %v, want true", found)
	}

	if found := BFS(Tree, 69); found {
		t.Errorf("BFS() = %v, want false", found)
	}
}
