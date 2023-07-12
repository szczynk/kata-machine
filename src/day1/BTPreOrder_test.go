package main

import (
	"reflect"
	"testing"
)

func TestPreOrderSearch(t *testing.T) {
	want := []int{
		20,
		10,
		5,
		7,
		15,
		50,
		30,
		29,
		45,
		100,
	}

	result := PreOrderSearch(Tree)

	if !reflect.DeepEqual(result, want) {
		t.Errorf("PreOrderSearch() = %v, want %v", result, want)
	}
}
