package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_quickSort(t *testing.T) {
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
			quickSort(tt.args.arr)
			fmt.Println("tt.args.arr:", tt.args.arr)
			fmt.Println("tt.wants.arr:", tt.wants.arr)

			if !reflect.DeepEqual(tt.args.arr, tt.wants.arr) {
				t.Errorf("quick_sort() = %v, want %v", tt.args.arr, tt.wants.arr)
			}
		})
	}
}
