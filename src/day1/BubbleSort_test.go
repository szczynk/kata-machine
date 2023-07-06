package main

import (
	"testing"
)

func Test_bubble_sort(t *testing.T) {
	type args struct {
		arr []int
	}
	type wants struct {
		arr []int
	}
	tests := []struct {
		name  string
		args  args
		wants wants
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				arr: []int{9, 3, 7, 4, 69, 420, 42},
			},
			wants: wants{
				arr: []int{3, 4, 7, 9, 42, 69, 420},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bubble_sort(tt.args.arr)
			t.Log("\ntt.args.arr:", tt.args.arr)
			t.Log("\ntt.wants.arr:", tt.wants.arr)

			if !equal(tt.args.arr, tt.wants.arr) {
				t.Errorf("bubble_sort() = %v, want %v", tt.args.arr, tt.wants.arr)
			}
		})
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
