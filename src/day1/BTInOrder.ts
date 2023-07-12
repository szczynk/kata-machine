function walk(current: BinaryNode<number> | null, path: number[]): number[]{
    if (!current) {
        return path
    }

    // recurse
    // pre

    // recurse
    walk(current.left, path)

    path.push(current.value)

    walk(current.right,path)

    // post
    return path
}

//           7             <- root / head
//         /   \               |
//       23      3             |
//     /   \   /   \           |
//    5    4   18   21     <- leaf
    
//    post_order_search => [5, 23, 4, 7, 18, 3, 21]
export default function in_order_search(head: BinaryNode<number>): number[] {
    return walk(head, [])
}