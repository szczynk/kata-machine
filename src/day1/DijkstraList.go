package main

import (
	"fmt"
	"math"
)

// Dijkstra’s Shortest Path => greedy algorithm
//
// Dijkstra’s Shortest Path is algorithm that calculate
// the shortest path from one node to all other nodes in the graph
// (we can find an individual node as well)
//
// if every node has been visited, it is very easy to find shortest path
// from A to B as long as you have previous list
// which mean we need to implement *BFS*, coz it required previous list
//
// now we implement Dijkstra’s Shortest Path using BFS which runtime is O(V²+E)
func DijkstraList(source int, sink int, arr WeightedAdjacencyList) []int {
	seen := make([]bool, len(arr))
	prev := make([]int, len(arr))
	for i := range prev {
		prev[i] = -1
	}

	distances := make([]int, len(arr))
	for i := range distances {
		distances[i] = math.MaxInt32
	}
	// smallest distances possible at the source
	distances[source] = 0

	// get the unvisited node with the lowest distances
	for hasUnvisited(seen, distances) {
		fmt.Printf("%+v\n", distances)
		current := getLowestUnvisited(seen, distances)

		seen[current] = true

		// anything we select, we grab out the row
		// (the connections to other nodes from the current node.)
		// by grabbing out the row, that represent our connection to other node
		adjs := arr[current]
		// for every adjs in current node
		for _, edge := range adjs {
			if seen[edge.To] {
				continue
			}

			// if we have not seen this edge,
			// we will calculate its distance,
			// which is based on my distance plus the distance to that edge
			distance := distances[current] + edge.Weight
			// if less than its lowest known distance,
			if distance < distances[edge.To] {
				// we are gonna push/update
				// both the previous and distances array with new value
				prev[edge.To] = current
				distances[edge.To] = distance
			}
		}
	}

	// our out represents the path from the needle back to the source
	// start at the needle back to the source
	out := make([]int, 0)

	// let's build it backward,
	// now we have to walk our previouses until we get to a -1
	// we need to start at the point we need to search from
	// or where we hoping to find
	curr := sink

	// keep searching until we found a point that has no more parent
	// if we never found needle (no path exists), the loop will terminate when prev[curr] becomes -1.
	for prev[curr] != -1 {
		// push current to the out
		out = append(out, curr)
		// set this current to the parent
		curr = prev[curr]
	}

	// add our source
	out = append(out, source)

	// Reverse the out array
	for i := 0; i < len(out)/2; i++ {
		j := len(out) - i - 1
		out[i], out[j] = out[j], out[i]
	}

	// finally return our out
	return out
}

func hasUnvisited(seen []bool, distances []int) bool {
	for i, s := range seen {
		// if it's not seen and it's distance less than infinity
		if !s && distances[i] < math.MaxInt32 {
			return true
		}
	}
	return false
}

func getLowestUnvisited(seen []bool, distances []int) int {
	idx := -1
	lowestDistance := math.MaxInt32

	// walk through all of our distances and
	// get the lowest distance and unseen node
	for i := 0; i < len(seen); i++ {
		if seen[i] {
			continue
		}

		if distances[i] < lowestDistance {
			idx = i
			lowestDistance = distances[i]
		}
	}

	return idx
}
