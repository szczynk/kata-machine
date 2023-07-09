package main

import "fmt"

type Point struct {
	x int
	y int
}

var dirs = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func walk(maze []string, wall string, curr Point, end Point, seen [][]bool, path *[]Point) bool {
	// Base case
	// 1. off the map
	if curr.x < 0 || curr.x >= len(maze[0]) || curr.y < 0 || curr.y >= len(maze) {
		return false
	}

	// 2. on the wall
	if string(maze[curr.y][curr.x]) == wall {
		return false
	}

	// 3. it's the end
	if curr.x == end.x && curr.y == end.y {
		// path = append(path, end)
		*path = append(*path, end)
		return true
	}

	// 4. if we have seen it
	if seen[curr.y][curr.x] {
		return false
	}

	// Recurse case
	// pre
	seen[curr.y][curr.x] = true
	*path = append(*path, curr)

	// recurse
	for _, dir := range dirs {
		x, y := dir[0], dir[1]
		if walk(maze, wall, Point{curr.x + x, curr.y + y}, end, seen, path) {
			return true
		}
	}

	// post
	*path = (*path)[:len(*path)-1]

	return false
}

func Solve(maze []string, wall string, start Point, end Point) []Point {
	seen := make([][]bool, len(maze))
	path := []Point{}

	for i := 0; i < len(maze); i++ {
		seen[i] = make([]bool, len(maze[0]))
		// fmt.Println(seen[i])
	}

	walk(maze, wall, start, end, seen, &path)

	for i := 0; i < len(maze); i++ {
		fmt.Printf("seen[%v]: %v\n", i, seen[i])
	}

	return path
}
