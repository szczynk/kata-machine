// LRU
//
// Least Recently Used
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
// - The HashMap is used for fast key-based lookups,
//   It stores the keys of the items as keys and
//   the corresponding nodes in the Doubly Linked List as values
// - The Doubly Linked List is used to maintain the order of the items
//   based on their usage. The list keeps track of the most recently used item at the front(head)
//   and the least recently used item at the back(tail).
//
// Here's how the LRU data structure works in general:
// 1. When a new item is accessed or added to the cache, it is checked in the HashMap:
//    - If the item exists in the HashMap, we moved the item to the front (MRU position) to mark it
//      as the most recently used item on the corresponding nodes in the Doubly Linked List
//    - If the item doesn't exist in the HashMap, we create a new node, we add the item to
//      the HashMap with the key pointing to the new node, and the new node is inserted
//      at the front of the Doubly Linked List
// 2. When the cache reaches its capacity and a new item needs to be added,
//    the least recently used item at the back of the Doubly Linked List(tail) is evicted.
//    This item is also removed from the HashMap to keep the cache in sync.
// 3. The process of moving an item to the front (MRU) and evicting
//    the least recently used item(LRU) ensures that the most frequently accessed items
//    stay in the cache, while the least used items are removed when space is needed.
//
// the running time of LRU is 7 O(1), which is 4 O(1) operations for breaking the link and
// 3 O(1) operation for prepend the item
//
// that is an LRU. It's the combination of two data structures. We understood,
// a map + a linked list to be able to create something that can actually cache data
//
// the moment you cross the boundary from just simple UI features into something that the UI uses,
// you will often run into any of these things.
type Node<T> = {
    value: T;
    next?: Node<T>;
    prev?: Node<T>;
};

function createNode<V>(value: V): Node<V> {
    return { value };
}
export default class LRU<K, V> {
    private length: number;
    private head?: Node<V>;
    private tail?: Node<V>;

    // The lookup table. It maps keys to corresponding nodes in the doubly linked list.
    // this map is to provide efficient access to nodes in the linked list based on their keys
    private lookup: Map<K, Node<V>>;

    // The reverse lookup table. It maps nodes in the doubly linked list back to their corresponding keys.
    // this map is to facilitate efficient removal of nodes from the linked list when the cache
    // needs to be trimmed(i.e., when its capacity is exceeded).
    //
    // when we need to delete the node from the linked list,
    // we can get the key from the node in `reverseLookup` map, then delete that out and
    // using the key we got from `reverseLookup` map, we can delete the node from the `lookup` map
    //
    // This ensures that both maps are kept in sync and prevent memory leaks or incorrect lookup results.
    private reverseLookup: Map<Node<V>, K>;

    // For fun, I'll put a little = 10 here just in case my unit test just assumed it was 10
    constructor(private capacity: number = 10) {
        this.length = 0;
        this.head = this.tail = undefined;
        this.lookup = new Map<K, Node<V>>();
        this.reverseLookup = new Map<Node<V>, K>();
    }

    // why there's no `insert` method? Cuz we don't know if you have the value in the cache
    // Now we're just gonna say, hey, update this thing
    update(key: K, value: V): void {
        // let get the cache by using get(), does it exist?
        let node = this.lookup.get(key);

        if (!node) {
            // if it doesn't exist, we need to create a new node and insert the node
            node = createNode(value);
            this.length++; // bookeeping
            this.prepend(node);

            // It's a function in which will ensure that our cache remains no greater than
            // our capacity by removing the least recently used items from the cache.
            this.trimCache();

            // update our beautiful `lookup` and `reverseLookup`
            this.lookup.set(key, node);
            this.reverseLookup.set(node, key);
        } else {
            // if it does exist, we need to update the value we found and
            // move it to the front of the list
            this.detach(node);
            this.prepend(node);
            node.value = value;
        }

        // if it doesn't exist, we need to insert the item
        //      - check the capacity and if it exceed, evict the least recently used item
        //
    }

    get(key: K): V | undefined {
        // check the cache for existence
        const node = this.lookup.get(key);
        if (!node) {
            return undefined;
        }

        // update the value we found and move it to the front
        this.detach(node);
        this.prepend(node);

        // return out the value found or undefined if not exist
        return node.value;
    }

    private detach(node: Node<V>) {
        //   ___________
        //  ↙           \
        // A <!=> B <!=> C
        //  \___________↗
        //
        //
        // if the current node have prev (the current node's prev) exist
        // we need (the current node prev's next) point to (the current node's next)
        if (node.prev) {
            node.prev.next = node.next;
        }
        // the other direction as well
        if (node.next) {
            node.next.prev = node.prev;
        }

        // the tail and the head checking
        //
        // if this.length === 1, it means that the cache contains only one key-value pair,
        // and if that pair is removed, the cache will become empty,
        // which means that our head and our tail should become undefined
        // because there are no nodes left in the linked list.
        if (this.length === 1) {
            this.tail = this.head = undefined;
        }

        // if this.head === node, We need to shift our head over one before we break all the links.
        // Cuz if we break all the links, we just cut off our own head.
        if (this.head === node) {
            this.head = this.head.next;
        }
        // same as the tail as well
        if (this.tail === node) {
            this.tail = this.tail.prev;
        }

        // and then break the chain
        node.prev = undefined;
        node.next = undefined;
    }

    private prepend(node: Node<V>) {
        // If the linked list is empty, set this node as both the head and the tail.
        if (!this.head) {
            this.head = this.tail = node;
            return;
        }

        // node's next point to the head
        node.next = this.head;
        // the head's prev point to the node
        this.head.prev = node;
        // the head point to the node
        this.head = node;
    }

    private trimCache(): void {
        if (this.length <= this.capacity) {
            return;
        }

        // if (this.length > this.capacity), we need to do some removals.
        // first, Let's remove the tail from the linked list
        //
        // We hold on to a reference to our tail,
        const tail = this.tail as Node<V>;
        // and remove our tail from the linked list
        this.detach(this.tail as Node<V>);

        // second, Let's remove the tail from the `lookup` and `reverseLookup`
        //
        // we're gonna get the key from `reverseLookup`
        const key = this.reverseLookup.get(tail) as K;
        // delete the tail from `lookup`
        this.lookup.delete(key);
        // and delete the key from `reverseLookup`
        this.reverseLookup.delete(tail);

        this.length--; //bookeeping
    }
}
