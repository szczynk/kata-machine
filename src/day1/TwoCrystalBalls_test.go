package main

import (
	"math/rand"
	"testing"
)

func Test_two_crystal_balls(t *testing.T) {
	type args struct {
		breaks []bool
	}

	idx := rand.Intn(10000)
	data := make([]bool, 10000)
	for i := idx; i < 10000; i++ {
		data[i] = true
	}

	data2 := make([]bool, 821)

	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "two crystal balls found",
			args: args{breaks: data},
			want: idx,
		},
		{
			name: "two crystal balls not found",
			args: args{breaks: data2},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := two_crystal_balls(tt.args.breaks); got != tt.want {
				t.Errorf("two_crystal_balls() = %v, want %v", got, tt.want)
			}
		})
	}
}
