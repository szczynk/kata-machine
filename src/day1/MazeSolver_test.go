package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		maze  []string
		wall  string
		start Point
		end   Point
	}

	maze := []string{
		"xxxxxxxxxx x",
		"x        x x",
		"x        x x",
		"x xxxxxxxx x",
		"x          x",
		"x xxxxxxxxxx",
	}

	mazeResult := []Point{
		{10, 0},
		{10, 1},
		{10, 2},
		{10, 3},
		{10, 4},
		{9, 4},
		{8, 4},
		{7, 4},
		{6, 4},
		{5, 4},
		{4, 4},
		{3, 4},
		{2, 4},
		{1, 4},
		{1, 5},
	}

	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "maze solver",
			args: args{
				maze:  maze,
				wall:  "x",
				start: Point{x: 10, y: 0},
				end:   Point{x: 1, y: 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Solve(tt.args.maze, tt.args.wall, tt.args.start, tt.args.end)
			fmt.Println("result:", result)

			fmt.Println("drawPath(maze, result)")
			for i := 0; i < len(maze); i++ {
				fmt.Println(drawPath(maze, result)[i])
			}

			fmt.Println("drawPath(maze, mazeResult)")
			for i := 0; i < len(maze); i++ {
				fmt.Println(drawPath(maze, mazeResult)[i])
			}

			if !reflect.DeepEqual(drawPath(maze, result), drawPath(maze, mazeResult)) {
				t.Errorf("Solve() failed:\n got: %v\nwant: %v", drawPath(maze, result), drawPath(maze, mazeResult))
			}
		})
	}
}

func drawPath(data []string, path []Point) []string {
	data2 := make([][]rune, len(data))
	for i, row := range data {
		data2[i] = []rune(row)
	}

	for _, p := range path {
		if p.y < len(data2) && p.x < len(data2[p.y]) {
			data2[p.y][p.x] = '*'
		}
	}

	result := make([]string, len(data2))
	for i, row := range data2 {
		result[i] = string(row)
	}
	return result
}
