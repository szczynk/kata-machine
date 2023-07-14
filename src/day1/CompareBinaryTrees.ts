// compare two binary trees to see if they are equal in value and in structure, e.g.
// 
//           5            5
//         /   \        /   \
//        3     45     3     45
// 
// using breadth-first: compare 5 to 5, 3 to 3, 45 to 45 -> both trees are equal,
// 
// BUT consider the following tree,
// 
//           5
//         /
//       3
//     /
//    45 
// 
// when you visit this breadth-first, it also yields 5, 3, 45! -> both trees are not equal,
// 
// using depth-first, on the other hand, preserves the shape of the tree!
export default function compare(a: BinaryNode<number> | null, b: BinaryNode<number> | null): boolean {
    // * base case
    // Check if both trees have reached the end (leaf nodes) 
    // (structural check)
    if (a === null && b === null) {
        return true;
    }

    // Check if one tree has reached the end while the other has remaining nodes 
    // (structural check)
    if (a === null || b === null) {
        return false;
    }

    // Check if the values of the current nodes are equal 
    // (value check)
    if (a.value !== b.value) {
        return false;
    }

    // * recurse
    // Recursively compare the left subtrees and the right subtrees
    return compare(a.left, b.left) && compare(a.right, b.right);
}