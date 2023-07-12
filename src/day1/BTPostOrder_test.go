package main

import (
	"reflect"
	"testing"
)

func TestPostOrderSearch(t *testing.T) {
	want := []int{
		7,
		5,
		15,
		10,
		29,
		45,
		30,
		100,
		50,
		20,
	}

	result := PostOrderSearch(Tree)

	if !reflect.DeepEqual(result, want) {
		t.Errorf("PostOrderSearch() = %v, want %v", result, want)
	}
}
