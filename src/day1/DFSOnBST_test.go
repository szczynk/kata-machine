package main

import (
	"testing"
)

func TestDFSOnBST(t *testing.T) {
	if found := DFSOnBST(Tree, 45); !found {
		t.Errorf("BFS() = %v, want true", found)
	}

	if found := DFSOnBST(Tree, 7); !found {
		t.Errorf("BFS() = %v, want true", found)
	}

	if found := DFSOnBST(Tree, 69); found {
		t.Errorf("BFS() = %v, want false", found)
	}
}
