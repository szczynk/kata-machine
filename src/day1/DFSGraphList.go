package main

func DFSGraphListWalk(graph WeightedAdjacencyList, current, needle int, seen []bool, path *[]int) bool {
	if seen[current] {
		return false
	}

	seen[current] = true
	*path = append(*path, current)

	if current == needle {
		return true
	}

	for _, edge := range graph[current] {
		if DFSGraphListWalk(graph, edge.To, needle, seen, path) {
			return true
		}
	}

	*path = (*path)[:len(*path)-1]

	return false
}

func DFSGraphList(graph WeightedAdjacencyList, source, needle int) []int {
	seen := make([]bool, len(graph))
	path := []int{}

	DFSGraphListWalk(graph, source, needle, seen, &path)

	if len(path) == 0 {
		return nil
	}

	return path
}
