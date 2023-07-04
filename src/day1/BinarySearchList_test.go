package main

import "testing"

func Test_bs_list(t *testing.T) {
	type args struct {
		haystack []int
		needle   int
	}
	foo := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "find 69",
			args: args{haystack: foo, needle: 69},
			want: true,
		},
		{
			name: "find 1336, not found",
			args: args{haystack: foo, needle: 1336},
			want: false,
		},
		{
			name: "find 69420",
			args: args{haystack: foo, needle: 69420},
			want: true,
		},
		{
			name: "find 69421, not found",
			args: args{haystack: foo, needle: 69421},
			want: false,
		},
		{
			name: "find 1",
			args: args{haystack: foo, needle: 1},
			want: true,
		},
		{
			name: "find 0, not found",
			args: args{haystack: foo, needle: 0},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bs_list(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("bs_list() = %v, want %v", got, tt.want)
			}
		})
	}
}
