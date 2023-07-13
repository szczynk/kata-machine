function walk(current: BinaryNode<number> | null, path: number[]): number[]{
    if (!current) {
        return path
    }

    // recurse
    // pre
    
    // recurse
    path.push(current.value)

    walk(current.left, path)

    walk(current.right, path)
    
    // post
    return path
}

//    Depth First Search / Traversal
//    using implicit Data Structure -> Stack
// 
//           7             <- root / head
//         /   \               |
//       23      3             |
//     /   \   /   \           |
//    5    4   18   21     <- leaf
// 
//    pre_order_search => [7, 23, 5, 4, 3, 18, 21]
export default function pre_order_search(head: BinaryNode<number>): number[] {
    return walk(head, [])
}