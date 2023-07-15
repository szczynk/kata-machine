package main

// * DFSOnBST
// a depth-first search (DFS) implementation for searching
// a binary search tree(BST) for a specific value
//
// * Depth-first Find
// Binary Search Tree (BST): binary tree, but at each node,
// everything to the left is ≤ than the node,
// while everything to the right is > than the node
//
//	<--<=--- 17 -- < -->
//	       /    \
//	     15      50
//	   /   \    /
//	  5    16   25
//	           /  \
//	range  18-25  26-50
//
// effectively looks a lot like quicksort
// find on BST is effectively a simplified version of binary search on an array
// time-complexity: range of O(height) to O(log n)
//
// * Depth-first Insert
// insert is a lot like find; we need to look for the appropriate place to insert the value
// e.g if you insert 17 and 18 in this example,
//
//	        17
//	     /      \
//	   15        50
//	 /   \       /
//	5    16     25
//	       \    /
//	       17  18
//
// * Depth-first Delete
// removing a node with no children does not affect the validity of the tree structure
// case 1)  no child, delete
// case 2)  if it has one child, you set the node parent reference directly to the child
// case 3)  if it has more than one child, you need to consider the smallest child of the node
//
//	on the right side or the largest child of the node on the left side, and move those
//	up the tree
//
// ideally in the case 3, you choose the option that has the least height of the tree,
// so that the tree is kept as clean as possible
// e.g if you delete 50 in this example,
//
//	        17
//	     /     \
//	   15       50
//	 /   \     /   \
//	5    16   25   100
//	         /  \
//	        23  27
//
// you can do it like this (largest element on the smaller side is 27 )
//
//	        17
//	     /     \
//	   15       27
//	 /   \     /   \
//	5    16   25   100
//	         /
//	        23
//
// or like this (smallest element on the larger side is 100 )
//
//	        17
//	     /     \
//	   15       100
//	 /   \     /
//	5    16   25
//	         /  \
//	        23  27
//
// (notice how the shape of the tree changes!)
//
// we just implement Depth-first Find
func DFSOnBSTSearch(current *BinaryNode, needle int) bool {
	if current == nil {
		return false
	}

	if current.Value == needle {
		return true
	}

	// traversal
	if current.Value < needle {
		return DFSOnBSTSearch(current.Right, needle)
	}

	return DFSOnBSTSearch(current.Left, needle)
}

func DFSOnBST(head *BinaryNode, needle int) bool {
	return DFSOnBSTSearch(head, needle)
}
