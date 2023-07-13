package main

//	Breadth First Search / Traversal
//	using implicit Data Structure -> Queue
//	       7             <- root / head
//	     /   \               |
//	   23      3             |
//	 /   \   /   \           |
//	5    4   18   21     <- leaf
//
// in_order_search => [7, 23, 3, 5, 4, 18, 21]
func BFS(head *BinaryNode, needle int) bool {
	queue := []*BinaryNode{head}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == nil {
			continue
		}

		// search
		if current.Value == needle {
			return true
		}

		queue = append(queue, current.Left)
		queue = append(queue, current.Right)
	}

	return false
}
