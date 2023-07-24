package main

// LRU
//
// # Least Recently Used
//
// what that means it's a caching mechanism that says that
// "we will evict / pull out the least recently used item"
//
// how we would do that?
// first thing first, cache really is some sort of container
// We're gonna build some sort of node container based system in which contains some sort of value.
//
// Now we want this value to store on cache.
// But at LRU cache has one more item in which it has a linked list that has other values inside of it
// and we keep track of used values in LRU cache, e.g.
//
// ┌──┐     ┌──┐     ┌──┐     ┌──┐     ┌───┐
// │V0│◄---►│V1│◄---►│V2│◄---►│V3│◄---►│...│
// └──┘     └──┘     └──┘     └──┘     └───┘
//
// when we want to take a value, we pull out that value and move it to the head of the queue
// e.g. we want to take V2, so the LRU become this
//
// ┌──┐     ┌──┐     ┌──┐     ┌──┐     ┌───┐
// │V2│◄---►│V0│◄---►│V1│◄---►│V3│◄---►│...│
// └──┘     └──┘     └──┘     └──┘     └───┘
// most recently used        ...       least recently used
//
// so to implement this, we need a doubly linked list and also a hashmap,
//
//   - The HashMap is used for fast key-based lookups,
//     It stores the keys of the items as keys and
//     the corresponding nodes in the Doubly Linked List as values
//
//   - The Doubly Linked List is used to maintain the order of the items
//     based on their usage. The list keeps track of the most recently used item at the front(head)
//     and the least recently used item at the back(tail).
//
// Here's how the LRU data structure works in general:
//
//  1. When a new item is accessed or added to the cache, it is checked in the HashMap.
//
//     - If the item exists in the HashMap, we moved the item to the front (MRU position) to mark it
//     as the most recently used item on the corresponding nodes in the Doubly Linked List.
//
//     - If the item doesn't exist in the HashMap, we create a new node, we add the item to
//     the HashMap with the key pointing to the new node, and the new node is inserted
//     at the front of the Doubly Linked List.
//
//  2. When the cache reaches its capacity and a new item needs to be added,
//     the least recently used item at the back of the Doubly Linked List(tail) is evicted.
//     This item is also removed from the HashMap to keep the cache in sync.
//
//  3. The process of moving an item to the front (MRU) and evicting
//     the least recently used item(LRU) ensures that the most frequently accessed items
//     stay in the cache, while the least used items are removed when space is needed.
//
// the running time of LRU is 7 O(1), which is 4 O(1) operations for breaking the link and
// 3 O(1) operation for prepend the item
//
// that is an LRU. It's the combination of two data structures. We understood,
// a map + a linked list to be able to create something that can actually cache data
//
// the moment you cross the boundary from just simple UI features into something that the UI uses,
// you will often run into any of these things.
type LRU struct {
	length        int
	head, tail    *Node
	lookup        map[interface{}]*Node
	reverseLookup map[*Node]interface{}
	capacity      int
}

func NewLRU(capacity int) *LRU {
	return &LRU{
		length: 0,
		head:   nil,
		tail:   nil,

		// The lookup table. It maps keys to corresponding nodes in the doubly linked list.
		// this map is to provide efficient access to nodes in the linked list based on their keys
		lookup: make(map[interface{}]*Node),

		// The reverse lookup table. It maps nodes in the doubly linked list back to their corresponding keys.
		// this map is to facilitate efficient removal of nodes from the linked list when the cache
		// needs to be trimmed(i.e., when its capacity is exceeded).
		//
		// when we need to delete the node from the linked list,
		// we can get the key from the node in `reverseLookup` map, then delete that out and
		// using the key we got from `reverseLookup` map, we can delete the node from the `lookup` map
		//
		// This ensures that both maps are kept in sync and prevent memory leaks or incorrect lookup results.
		reverseLookup: make(map[*Node]interface{}),

		capacity: capacity,
	}
}

// why there's no `insert` method? Cuz we don't know if you have the value in the cache
// Now we're just gonna say, hey, update this thing
func (lru *LRU) Update(key, value interface{}) {
	// let get the cache by using get(), does it exist?
	node, exists := lru.lookup[key]

	if !exists {
		// if it doesn't exist, we need to create a new node and insert the node

		node = createNode(value)
		lru.length++ // bookeeping
		lru.prepend(node)

		// It's a function in which will ensure that our cache remains no greater than
		// our capacity by removing the least recently used items from the cache.
		lru.trimCache()

		// update our beautiful `lookup` and `reverseLookup`
		lru.lookup[key] = node
		lru.reverseLookup[node] = key
	} else {
		// if it does exist, we need to update the value we found and
		// move it to the front of the list
		lru.detach(node)
		lru.prepend(node)
		node.value = value
	}
}

func (lru *LRU) Get(key interface{}) interface{} {
	// check the cache for existence
	node, exists := lru.lookup[key]

	if !exists {
		return nil
	}

	// update the value we found and move it to the front
	lru.detach(node)
	lru.prepend(node)

	// return out the value found or undefined if not exist
	return node.value
}

func createNode(value interface{}) *Node {
	return &Node{value: value}
}

func (lru *LRU) detach(node *Node) {
	//   ___________
	//  ↙           \
	// A <!=> B <!=> C
	//  \___________↗
	//
	//
	// if the current node have prev (the current node's prev) exist
	// we need (the current node prev's next) point to (the current node's next)
	if node.prev != nil {
		node.prev.next = node.next
	}
	// the other direction as well
	if node.next != nil {
		node.next.prev = node.prev
	}

	// the tail and the head checking
	//
	// if this.length === 1, it means that the cache contains only one key-value pair,
	// and if that pair is removed, the cache will become empty,
	// which means that our head and our tail should become undefined
	// because there are no nodes left in the linked list.
	if lru.length == 1 {
		lru.tail, lru.head = nil, nil
	}
	// if this.head === node, We need to shift our head over one before we break all the links.
	// Cuz if we break all the links, we just cut off our own head.
	if lru.head == node {
		lru.head = lru.head.next
	}
	// same as the tail as well
	if lru.tail == node {
		lru.tail = lru.tail.prev
	}

	// and then break the chain
	node.prev, node.next = nil, nil
}

func (lru *LRU) prepend(node *Node) {
	// If the linked list is empty, set this node as both the head and the tail.
	if lru.head == nil {
		lru.head, lru.tail = node, node
		return
	}

	// node's next point to the head
	node.next = lru.head
	// the head's prev point to the node
	lru.head.prev = node
	// the head point to the node
	lru.head = node
}

func (lru *LRU) trimCache() {
	if lru.length <= lru.capacity {
		return
	}

	// if (this.length > this.capacity), we need to do some removals.
	// first, Let's remove the tail from the linked list
	//
	// We hold on to a reference to our tail,
	tail := lru.tail
	// and remove our tail from the linked list
	lru.detach(tail)

	// second, Let's remove the tail from the `lookup` and `reverseLookup`
	//
	// we're gonna get the key from `reverseLookup`
	key := lru.reverseLookup[tail]
	// delete the tail from `lookup`
	delete(lru.lookup, key)
	// and delete the key from `reverseLookup`
	delete(lru.reverseLookup, tail)

	lru.length-- //bookeeping
}
