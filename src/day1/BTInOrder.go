package main

func InOrderSearchWalk(current *BinaryNode, path []int) []int {
	if current == nil {
		return path
	}

	// recurse
	// pre

	// recurse
	path = InOrderSearchWalk(current.Left, path)

	path = append(path, current.Value)

	path = InOrderSearchWalk(current.Right, path)

	// post
	return path
}

//	Depth First Search / Traversal
//	using implicit Data Structure -> Stack
//	       7             <- root / head
//	     /   \               |
//	   23      3             |
//	 /   \   /   \           |
//	5    4   18   21     <- leaf
//
// in_order_search => [5, 23, 4, 7, 18, 3, 21]
func InOrderSearch(head *BinaryNode) []int {
	return InOrderSearchWalk(head, []int{})
}
