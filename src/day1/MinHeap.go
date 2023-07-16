package main

// heap or “priority queue” is a binary tree where every child and grand-child is smaller (MaxHeap),
// or larger (MinHeap) than the current node
//
// we now implement minHeap
//
// whenever a node is added or deleted, we must adjust the tree
// there is no traversing the tree
// it must also be a full or complete tree (no imbalanced nodes)
// heaps maintain a weak (imperfect) ordering, in other words,
// the ordering is only true at each node(MaxHeap / MinHeap), e.g.:
//
//	         50
//	      /     \
//	    71       50
//	  /   \     /   \
//	101   80   200   101
//
// if, for instance, we add 3 , we go to the leaf (101) and then heapify up
// until the ordering condition (node >= parent), then we can continue adding e.g. 200
//
//	          50                                    3 -> ~~50~~
//	       /     \                                   ↖   /     \
//	     71       50         =>          50 -> ~~3~~  ~~71~~     50
//	   /   \     /   \                            ↖   /   \     /   \
//	 101   80   200   101            71 -> ~~3~~  ~~101~~  80  200   101
//	 /                                         ↖    /
//	3                                   101  ->  ~~3~~
//
//
//	          3           then we can             3
//	       /     \      adding e.g. 200        /     \
//	     50       50         =>              50        50
//	   /   \     /   \                     /   \     /   \
//	  71   80   200   101                 71   80  200   101
//	 /                                   /  \
//	101                                 101  200
//
// if, for instance, we delete 3, we need to take the top of the tree, 3, delete it,
// and take our final node, 200, and put it at the top, then we can heapify down
//
//	          3                          200 -> ~~3~~
//	       /     \                             /     \
//	     50       50         =>              50        50
//	   /   \     /   \                     /   \     /   \
//	  71   80   200   101                 71   80  200   101
//	 /  \                               /   \
//	101  200                         101  ~~200~~
//
//
//	                50 -> ~~200~~                                 50
//	                 ↖   /       \                             /     \
//	    71 -> ~~200~~  ~~50~~      50         =>             71        50
//	               ↖   /    \     /   \                     /   \     /   \
//	101 -> ~~200~~  ~~71~~  80   200   101                101   80  200   101
//	              ↖   /                                   /
//	       200 -> ~~101~~                               200
//
// let's add index to every node from left to right
// then we can store this in an array => [50, 71, 100, 101, 80, 200, 101, 200]
//
// time-complexity => O(log n)
type MinHeap struct {
	length int
	data   []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		data:   make([]int, 0),
		length: 0,
	}
}

func (h *MinHeap) Insert(value int) {
	// Add the new value at the end of the array
	h.data = append(h.data, value)

	// heapifyUp
	h.heapifyUp(h.length)
	h.length++
}

func (h *MinHeap) Delete() int {
	// If the heap is empty
	if h.length == 0 {
		return -1 // as a sentinel value
	}

	out := h.data[0]
	h.length--

	// If there are no more elements
	if h.length == 0 {
		h.data = []int{} // reset the data array
		return out
	}

	// Replace the root with the last element
	h.data[0] = h.data[h.length]

	// heapifyDown
	h.heapifyDown(0)
	return out
}

func (h *MinHeap) heapifyDown(idx int) {
	leftIdx := h.leftChild(idx)
	rightIdx := h.rightChild(idx)

	// check if we have more children, because we always fill in from left to right
	// If the current index or left child index is out of bounds
	if idx >= h.length || leftIdx >= h.length {
		return
	}

	// find the minimum between two children
	leftValue := h.data[leftIdx]
	rightValue := h.data[rightIdx]
	value := h.data[idx]

	if leftValue < rightValue && value > leftValue {
		// swap if rightValue is smaller
		h.data[idx] = leftValue
		h.data[leftIdx] = value

		h.heapifyDown(leftIdx)
	} else if rightValue < leftValue && value > rightValue {
		// swap if leftValue is smaller
		h.data[idx] = rightValue
		h.data[rightIdx] = value

		h.heapifyDown(rightIdx)
	}
}

func (h *MinHeap) heapifyUp(idx int) {
	// we can not heapify beyond 0th index, we stop there
	// If the current index is the root
	if idx == 0 {
		return
	}

	parentIdx := h.parent(idx)
	parentValue := h.data[parentIdx]
	value := h.data[idx]

	if parentValue > value {
		// swap if current value is smaller
		h.data[idx] = parentValue
		h.data[parentIdx] = value

		h.heapifyUp(parentIdx)
	}
}

// parent-child relationship is then implemented by a simple formula
//
//	     parent
//	        |
//	      node
//	    /      \
//	left        right
//
// generic parent  => ⌊ (i - 1) / 2 ⌋
// left child      => 2 * i + 1
// right child     => 2 * i + 2
//
// so you don't need to implement relationships, it's just accessible by a formula!
func (h *MinHeap) parent(idx int) int {
	return (idx - 1) / 2
}

func (h *MinHeap) leftChild(idx int) int {
	return idx*2 + 1
}

func (h *MinHeap) rightChild(idx int) int {
	return idx*2 + 2
}
