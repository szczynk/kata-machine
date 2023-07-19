package main

import (
	"reflect"
	"testing"
)

func TestDFSGraphList(t *testing.T) {
	list2 := WeightedAdjacencyList{
		{
			{To: 1, Weight: 3},
			{To: 2, Weight: 1},
		},
		{
			{To: 4, Weight: 1},
		},
		{
			{To: 3, Weight: 7},
		},
		{},
		{
			{To: 1, Weight: 1},
			{To: 3, Weight: 5},
			{To: 5, Weight: 2},
		},
		{
			{To: 2, Weight: 18},
			{To: 6, Weight: 1},
		},
		{
			{To: 3, Weight: 1},
		},
	}

	expectedPath1 := []int{0, 1, 4, 5, 6}
	path1 := DFSGraphList(list2, 0, 6)
	if !reflect.DeepEqual(path1, expectedPath1) {
		t.Errorf("Unexpected path for DFSGraphList(list2, 0, 6). Got %v, expected %v", path1, expectedPath1)
	}

	expectedPath2 := []int(nil)
	path2 := DFSGraphList(list2, 6, 0)
	if !reflect.DeepEqual(path2, expectedPath2) {
		t.Errorf("Unexpected path for DFSGraphList(list2, 6, 0). Got %v, expected %v", path2, expectedPath2)
	}
}
