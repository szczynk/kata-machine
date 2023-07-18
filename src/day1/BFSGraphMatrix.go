package main

// * graph => series of nodes with some amount of connections around it
//
// hardest & largest part of all algo
// bestest data structure
// longest problem to how we draw our graph to this day
//
// e.g. Euler draw a graph to king tour but the city does not supported
//
// so all trees are graphs, but not all graphs are trees
//
//   - Terminology of Graphs
//
//   - Graph Terms
//     cycle      : if you can visit x amount of nodes and find a return path to the starting node
//     acyclic    : graph that does not contain any cycles
//     connected  : every node can reach every other node
//     directed   : When there is a direction to the connections. Think Twitter
//     undirected : not directed. Think Facebook (i haven't been on in 10 years, it may have changed)
//     weighted   : the edges have a “weight” or “value” associated with them
//     (e.g.a traffic map where the same road can have more traffic going
//     North than South). Think Maps
//     dag        : Directed, acyclic graph.
//
//   - Implementation Terms
//     node or vertex   : a point or vertex on the graph
//     edge or link     : the connection betxit two nodes
//
//   - Big O
//     BigO is commonly stated in terms of V and E where V stands for vertices and E stands for edges
//     So O(V * E) means that we will check every vertex, and on every vertex we check every edge
//
// for instance, we have a graph like this
//
//	┌─┐ 10  ┌─┐
//	│0├────►│1│
//	└─┘\    └─┘
//	 ▲  \ 5  ▲
//	 │   \   │
//	┌┴┐   ◄ ┌┴┐
//	│2│◄────┤3│
//	└─┘     └─┘
//
// implementing graphs can be done in two ways:
// 1) adjacency list    (popular)
//
//   - the index of the list maps to the nodes
//
//   - these are specified by a list of edges (where you could also specify a weight)
//
//     become like this,
//     ```
//     const AdjacencyList = [
//     [
//     { to: 1, weight: 10 },
//     { to: 3, weight: 5 },
//     ],                          // index 0 = node `0`
//     [],                         // index 1 = node `1` (has no edges)
//     [{ to: 0 }],
//     [{ to: 1 }, { to: 2 }],
//     ];
//     ```
//
// 2) adjacency matrix  (memory-intensive, O(V²))
//
//   - the rows represent nodes and
//
//   - contain a list of indices that represent the nodes
//
//     become like this,
//     ```
//     const AdjacencyMatrix =
//     [
//     // node 0, 1, 2, 3
//     [0, 10, 0, 5],  // node 0
//     [0,  0, 0, 0],  // node 1
//     [...],  // node 2
//     [...],  // node 3
//     ]
//     ```
//     you can already see (with node 1 ), if you have a sparsely connected map,
//     than you are going to lose a lot of memory, because each row needs to be
//     filled regardless of the edges
//
// basic search like DFS and BFS still on a graph,
// and they are virtually no different than on a tree
//
// in a tree, there no path, because it always top to bottom
// but in a graph, it need the path.
//
// to represent path in a graph
// -> in DFS, we just build our path as we recurse
//
//	(pre-recurse -> push, post-recurse -> pop)
//
// -> in BFS, we don't have somethiong to maintain the shape / structure
//
//	of our search, so we need to maintain ourself. The best way to do it,
//	create previous array = [-1,...] and seen array = [true, false, ...]
//
// now we are implement BFS on Adjacency Matrix based graph.
func BFSGraphMatrix(graph [][]int, source, needle int) []int {
	seen := make([]bool, len(graph))
	prev := make([]int, len(graph))
	for i := range prev {
		prev[i] = -1
	}

	seen[source] = true
	queue := []int{source}

	for len(queue) > 0 {
		// grab 1st thing out of our queue
		current := queue[0]
		queue = queue[1:]

		if current == needle {
			break
		}

		// anything we select, we grab out the row (the connections to other nodes from the current node.)
		// by grabbing out the row, that represent our connection to other node
		adjs := graph[current]
		// for every adjs in current node
		for i := 0; i < len(graph); i++ {
			// if there is no edge between the current node and the adjacent node
			if adjs[i] == 0 {
				continue
			}

			if seen[i] {
				continue
			}

			// since we have seen this child node, we mark it as visited
			seen[i] = true
			// set this child as the parent
			prev[i] = current

			// push the node to queue
			queue = append(queue, i)
		}
	}

	// if we have not found a path to the needle.
	if prev[needle] == -1 {
		return nil
	}

	// let's build it backward,
	// now we have to walk our previouses until we get to a -1
	// we need to start at the point we need to search from
	// or where we hoping to find
	curr := needle

	// our out represents the path from the needle back to the source
	// start at the needle back to the source
	out := []int{}

	// keep searching until we found a point that has no more parent
	// if we never found needle (no path exists), the loop will terminate when prev[curr] becomes -1.
	for prev[curr] != -1 {
		// push current to the out
		out = append(out, curr)
		// set this current to the parent
		curr = prev[curr]
	}

	// return our out and reverse it
	// finally add our source
	out = append(out, source)
	reverse(out)

	return out
}

// reverse reverses the elements in the given slice.
func reverse(arr []int) {
	i := 0
	j := len(arr) - 1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}
