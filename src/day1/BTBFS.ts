//    Breadth First Search / Traversal
//    using implicit Data Structure -> Queue
//           7             <- root / head
//         /   \               |
//       23      3             |
//     /   \   /   \           |
//    5    4   18   21     <- leaf
// 
// in_order_search => [7, 23, 3, 5, 4, 18, 21]
export default function bfs(head: BinaryNode<number>, needle: number): boolean {

    const queue: (BinaryNode<number> | null)[] = [head]
    
    while (queue.length) {

        const current = queue.shift() as BinaryNode<number> | undefined | null;
        if (!current) {
            continue
        }

        // search
        if (current.value === needle) {
            return true
        }

        queue.push(current.left)
        queue.push(current.right)
    }

    return false;
}