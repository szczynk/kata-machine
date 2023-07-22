package main

import (
	"reflect"
	"testing"
)

func TestDijkstraList(t *testing.T) {
	list1 := WeightedAdjacencyList{
		{
			{To: 1, Weight: 3},
			{To: 2, Weight: 1},
		},
		{
			{To: 0, Weight: 3},
			{To: 2, Weight: 4},
			{To: 4, Weight: 1},
		},
		{
			{To: 1, Weight: 4},
			{To: 3, Weight: 7},
			{To: 0, Weight: 1},
		},
		{
			{To: 2, Weight: 7},
			{To: 4, Weight: 5},
			{To: 6, Weight: 1},
		},
		{
			{To: 1, Weight: 1},
			{To: 3, Weight: 5},
			{To: 5, Weight: 2},
		},
		{
			{To: 6, Weight: 1},
			{To: 4, Weight: 2},
			{To: 2, Weight: 18},
		},
		{
			{To: 3, Weight: 1},
			{To: 5, Weight: 1},
		},
	}

	expectedPath := []int{0, 1, 4, 5, 6}
	path := DijkstraList(0, 6, list1)
	if !reflect.DeepEqual(path, expectedPath) {
		t.Errorf("Unexpected path for DijkstraList(0, 6, list1). Got %v, expected %v", path, expectedPath)
	}
}
