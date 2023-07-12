package main

import (
	"reflect"
	"testing"
)

func TestInOrderSearch(t *testing.T) {
	want := []int{
		5,
		7,
		10,
		15,
		20,
		29,
		30,
		45,
		50,
		100,
	}

	result := InOrderSearch(Tree)

	if !reflect.DeepEqual(result, want) {
		t.Errorf("InOrderSearch() = %v, want %v", result, want)
	}
}
