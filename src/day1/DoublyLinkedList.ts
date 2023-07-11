type Node<T> = {
    value: T;
    prev?: Node<T>;
    next?: Node<T>;
};

export default class DoublyLinkedList<T> {
    public length: number;
    private head?: Node<T>;
    private tail?: Node<T>;

    constructor() {
        this.length = 0;
        this.head = undefined;
        this.tail = undefined;
    }

    // The prepend method adds a new node with the specified value at the beginning of the list. 
    // It updates the head property accordingly, as well as the prev and next pointers of the 
    // affected nodes.
    prepend(item: T): void {

        // G point to A, A point to G, head point to G
        // G <-> A
        //  ↖
        //   head

        // bookkeeping
        this.length++;

        // creating the node
        const node = { value: item } as Node<T>;

        // if no head, set this node as the head
        if (!this.head) {
            this.head = this.tail = node;
            return;
        }

        // next point to the head
        node.next = this.head;
        // the head's prev point to the node
        this.head.prev = node;
        // the head point to the node
        this.head = node;
    }

    // The insertAt method inserts a new node with the specified value at the given index in 
    // the list. It handles three cases: inserting at the end of the list(by calling append), 
    // inserting at the beginning of the list(by calling prepend), and inserting at a specific 
    // index.It traverses the list to find the current node at the given index, creates a new node, 
    // updates the prev and next pointers of the affected nodes, and breaks the old links.
    insertAt(item: T, idx: number): void {

        // F need to point to C as the next,
        // F need to point to B as the prev,
        // C's prev point to F
        // B's next point to F
        // break the old link
        // B <!=> C <-> D
        //  ⤡   ⤢
        //    F
        //
        // we need to traverse the list

        // we can really insert at that point so
        if (idx > this.length) {
            throw new Error("oh no, Index out of bounds");
        }

        if (idx === this.length) {
            // we really just appending the list
            this.append(item);
            return;
        } else if (idx === 0) {
            // we really just prepending the list
            this.prepend(item);
            return;
        }

        // bookkeeping
        this.length++;

        // we have current node position, call getAt, and type casting it
        const current = this.getAt(idx) as Node<T>;

        // creating the node
        const node = { value: item } as Node<T>;

        // the node's next point to the current
        node.next = current;
        // the node's prev point to the current's prev
        node.prev = current.prev;
        // now break the old link
        // the current's prev point to the node
        current.prev = node;

        // if the node have prev (the node's prev) exist
        // we need (the node prev's next) point to (the node)
        if (node.prev) {
            node.prev.next = node;
        }
    }

    // The append method adds a new node with the specified value at the end of the list. 
    // It updates the tail property accordingly, as well as the prev and next pointers of the 
    // affected nodes.
    append(item: T): void {

        // A point to G, G point to A, tail point to A
        // G <-> A
        //      ↗
        //   tail

        // bookkeeping
        this.length++;

        // creating the node
        const node = { value: item } as Node<T>;

        // if no tail, set this node as the tail
        if (!this.tail) {
            this.head = this.tail = node;
            return;
        }

        // prev point to the tail
        node.prev = this.tail;
        // the tail's next point to the node
        this.tail.next = node;
        // the tail point to the node
        this.tail = node;
    }

    // The remove method removes the first node with the specified value from the list. 
    // It traverses the list to find the node to be removed, updates the prev and next pointers of 
    // the affected nodes, and breaks the old links.It returns the removed value if found, or 
    // undefined otherwise.
    remove(item: T): T | undefined {

        // find first and then removed it

        // we have current position
        let current = this.head;

        // we need to simply just walk until we get to the value we need
        for (let i = 0; current && i < this.length; ++i) {
            if (current.value === item) {
                break;
            }
            current = current.next;
        }
        
        if (!current) {
            // no item need to be removed
            return undefined;
        }

        return this.removeNode(current);
    }

    // The get method retrieves the value at the given index in the list. It calls the getAt method 
    // to find the node at the specified index and returns its value, or undefined if the index is 
    // out of bounds.
    get(idx: number): T | undefined {
        const node = this.getAt(idx)
        return node?.value
    }
    
    // The removeAt method removes the node at the given index in the list. It calls the getAt 
    // method to find the node at the specified index, updates the prev and next pointers of the 
    // affected nodes, breaks the old links, and returns the removed value if successful, or 
    // undefined otherwise.
    removeAt(idx: number): T | undefined {
        const node = this.getAt(idx)
        
        if (!node) {
            return undefined
        }

        return this.removeNode(node);
    }

    // The private method removeNode is used internally to remove a specific node from the list. 
    // It handles various cases based on the node's position in the list, updates the prev and next 
    // pointers of the affected nodes, breaks the old links, and returns the removed value.
    private removeNode(node: Node<T>): T | undefined {

        // bookkeeping
        this.length--;

        if (this.length === 0) {
            const out = this.head?.value
            this.head = this.tail = undefined;
            return out;
        }

        //   ___________
        //  ↙           \
        // A <!=> B <!=> C
        //  \___________↗
        //   
        // 
        // if the current node have prev (the current node's prev) exist
        // we need (the current node prev's next) point to (the current node's next)
        if (node.prev) {
            node.prev.next = node.next
        }
        // the other direction as well
        if (node.next) {
            node.next.prev = node.prev
        }

        if (node === this.head) {
            this.head = node.next
        }
        if (node === this.tail) {
            this.tail = node.prev
        }

        // and then break the chain
        node.prev = node.next = undefined;

        return node.value
    }
    
    // The private method getAt is used internally to retrieve the node at the given index in 
    // the list. It traverses the list to find the node and returns it, or undefined if the index 
    // is out of bounds.
    private getAt(idx: number): Node<T> | undefined {
        // we have current position
        let current = this.head;

        // we need to simply just walk until we get to the index
        // at that point we will have the node
        // which we need to be replace
        for (let i = 0; current && i < idx; ++i) {
            current = current.next;
        }

        return current
    }

    // The private method debug is not used in the functionality of the doubly linked list. 
    // It appears to be a helper method for debugging purposes, as it prints the values of 
    // the nodes in the list.
    private debug() {
        let current = this.head;
        let out =""
        for (let i = 0; current && i < this.length; ++i) {
            out+=`${i} => ${current.value}, `
            current = current.next;
        }
        console.log(out)
    }
}
