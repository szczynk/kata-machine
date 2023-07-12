package main

func PreOrderSearchWalk(current *BinaryNode, path []int) []int {
	if current == nil {
		return path
	}

	// recurse
	// pre

	// recurse
	path = append(path, current.Value)

	path = PreOrderSearchWalk(current.Left, path)

	path = PreOrderSearchWalk(current.Right, path)

	// post
	return path
}

//           7             <- root / head
//         /   \               |
//       23      3             |
//     /   \   /   \           |
//    5    4   18   21     <- leaf

// pre_order_search => [7, 23, 5, 4, 3, 18, 21]
func PreOrderSearch(head *BinaryNode) []int {
	return PreOrderSearchWalk(head, []int{})
}
