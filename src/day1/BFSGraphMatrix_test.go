package main

import (
	"reflect"
	"testing"
)

func TestBFSGraphMatrix(t *testing.T) {
	matrix2 := WeightedAdjacencyMatrix{
		{0, 3, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0},
		{0, 0, 7, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 5, 0, 2, 0},
		{0, 0, 18, 0, 0, 0, 1},
		{0, 0, 0, 1, 0, 0, 1},
	}

	expectedPath1 := []int{0, 1, 4, 5, 6}
	path1 := BFSGraphMatrix(matrix2, 0, 6)
	if !reflect.DeepEqual(path1, expectedPath1) {
		t.Errorf("Unexpected path for BFSGraphMatrix(matrix2, 0, 6). Got %v, expected %v", path1, expectedPath1)
	}

	expectedPath2 := []int(nil)
	path2 := BFSGraphMatrix(matrix2, 6, 0)
	if !reflect.DeepEqual(path2, expectedPath2) {
		t.Errorf("Unexpected path for BFSGraphMatrix(matrix2, 6, 0). Got %v, expected %v", path2, expectedPath2)
	}
}
