package main

import "testing"

func TestCompareBinaryTrees(t *testing.T) {
	if result := CompareBinaryTrees(Tree, Tree); !result {
		t.Errorf("InOrderSearch() = %v, want true", result)
	}

	if result := CompareBinaryTrees(Tree, Tree2); result {
		t.Errorf("InOrderSearch() = %v, want false", result)
	}
}
