// heap or “priority queue” is a binary tree where every child and grand-child is smaller (MaxHeap),
// or larger (MinHeap) than the current node
//
// we now implement minHeap
//
// - whenever a node is added or deleted, we must adjust the tree
// - there is no traversing the tree
// - it must also be a full or complete tree (no imbalanced nodes)
// - heaps maintain a weak (imperfect) ordering, in other words,
//   the ordering is only true at each node(MaxHeap / MinHeap), e.g.:
//
//            50
//         /     \
//       71       50
//     /   \     /   \
//   101   80   200   101
//
//   if, for instance, we add 3 , we go to the leaf (101) and then heapify up
//   until the ordering condition (node >= parent), then we can continue adding e.g. 200
//
//             50                                    3 -> ~~50~~
//          /     \                                   ↖   /     \
//        71       50         =>          50 -> ~~3~~  ~~71~~     50
//      /   \     /   \                            ↖   /   \     /   \
//    101   80   200   101            71 -> ~~3~~  ~~101~~  80  200   101
//    /                                         ↖    /
//   3                                   101  ->  ~~3~~
//
//
//             3           then we can             3
//          /     \      adding e.g. 200        /     \
//        50       50         =>              50        50
//      /   \     /   \                     /   \     /   \
//     71   80   200   101                 71   80  200   101
//    /                                   /  \
//   101                                 101  200
//
//
//   if, for instance, we delete 3, we need to take the top of the tree, 3, delete it,
//   and take our final node, 200, and put it at the top, then we can heapify down
//
//             3                          200 -> ~~3~~
//          /     \                             /     \
//        50       50         =>              50        50
//      /   \     /   \                     /   \     /   \
//     71   80   200   101                 71   80  200   101
//    /  \                               /   \
//   101  200                         101  ~~200~~
//
//
//                   50 -> ~~200~~                                 50
//                    ↖   /       \                             /     \
//       71 -> ~~200~~  ~~50~~      50         =>             71        50
//                  ↖   /    \     /   \                     /   \     /   \
//   101 -> ~~200~~  ~~71~~  80   200   101                101   80  200   101
//                 ↖   /                                   /
//          200 -> ~~101~~                               200
//
//   let's add index to every node from left to right
//   then we can store this in an array => [50, 71, 100, 101, 80, 200, 101, 200]
//
// - time-complexity => O(log n)
export default class MinHeap {
    public length: number;
    private data: number[];

    constructor() {
        this.data = [];
        this.length = 0;
    }

    insert(value: number): void {
        // Add the new value at the end of the array
        this.data[this.length] = value;

        // heapifyUp
        this.heapifyUp(this.length);
        this.length++;
    }

    delete(): number {
        // If the heap is empty
        if (this.length === 0) {
            return -1; // as a sentinel value
        }

        const out = this.data[0];
        this.length--;

        // If there are no more elements
        if (this.length === 0) {
            this.data = []; // reset the data array
            return out;
        }

        // Replace the root with the last element
        this.data[0] = this.data[this.length];

        // heapifyDown
        this.heapifyDown(0);
        return out;
    }

    private heapifyDown(idx: number): void {
        const leftIdx = this.leftChild(idx);
        const rightIdx = this.rightChild(idx);

        // check if we have more children, because we always fill in from left to right
        // If the current index or left child index is out of bounds
        if (idx >= this.length || leftIdx >= this.length) {
            return; // stop heapifying
        }

        // find the minimum between two children
        const leftValue = this.data[leftIdx];
        const rightValue = this.data[rightIdx];
        const value = this.data[idx];

        if (leftValue > rightValue && value > rightValue) {
            // swap if rightValue is smaller
            this.data[idx] = rightValue;
            this.data[rightIdx] = value;

            this.heapifyDown(rightIdx);
        } else if (rightValue > leftValue && value > leftValue) {
            // swap if leftValue is smaller
            this.data[idx] = leftValue;
            this.data[leftIdx] = value;

            this.heapifyDown(leftIdx);
        }
    }

    private heapifyUp(idx: number): void {
        // we can not heapify beyond 0th index, we stop there
        // If the current index is the root
        if (idx === 0) {
            return; // stop heapifying
        }

        const parentIdx = this.parent(idx);
        const parentValue = this.data[parentIdx];
        const value = this.data[idx];

        if (parentValue > value) {
            // swap if current value is smaller
            this.data[idx] = parentValue;
            this.data[parentIdx] = value;

            this.heapifyUp(parentIdx);
        }
    }

    // parent-child relationship is then implemented by a simple formula
    //
    //              parent
    //                 |
    //               node
    //             /      \
    //         left        right
    //
    // generic parent  => ⌊ (i - 1) / 2 ⌋
    // left child      => 2 * i + 1
    // right child     => 2 * i + 2
    //
    // so you don't need to implement relationships, it's just accessible by a formula!
    private parent(idx: number): number {
        // because javascript 5 / 2 = 2.5, we need Math.floor()
        return Math.floor((idx - 1) / 2);
    }

    private leftChild(idx: number): number {
        return idx * 2 + 1;
    }

    private rightChild(idx: number): number {
        return idx * 2 + 2;
    }
}
