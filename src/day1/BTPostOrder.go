package main

func PostOrderSearchWalk(current *BinaryNode, path []int) []int {
	if current == nil {
		return path
	}

	// recurse
	// pre

	// recurse
	path = PostOrderSearchWalk(current.Left, path)

	path = PostOrderSearchWalk(current.Right, path)

	path = append(path, current.Value)

	// post
	return path
}

//           7             <- root / head
//         /   \               |
//       23      3             |
//     /   \   /   \           |
//    5    4   18   21     <- leaf

// post_order_search => [5, 4, 23, 18, 21, 3, 7]
func PostOrderSearch(head *BinaryNode) []int {
	return PostOrderSearchWalk(head, []int{})
}
