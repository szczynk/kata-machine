function walk(current: BinaryNode<number> | null, path: number[]): number[]{
    if (!current) {
        return path
    }

    // recurse
    // pre

    // recurse
    walk(current.left, path)

    walk(current.right, path)

    path.push(current.value)

    // post
    
    return path
}

//           7             <- root / head
//         /   \               |
//       23      3             |
//     /   \   /   \           |
//    5    4   18   21     <- leaf
    
//    post_order_search => [5, 4, 23, 18, 21, 3, 7]
export default function post_order_search(head: BinaryNode<number>): number[] {
    return walk(head, [])
}