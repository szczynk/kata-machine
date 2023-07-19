// now we are implement DFS on Adjacency List-based graph.
//
// recursion function
// we did not consider our source
// same as our MazeSolver walk function
function walk(
    graph: WeightedAdjacencyList,
    current: number,
    needle: number,
    seen: boolean[],
    path: number[],
): boolean {
    // base case
    if (seen[current]) {
        return false;
    }

    // since we have seen this node, we mark it as visited
    seen[current] = true;

    // recurse
    // pre
    path.push(current);
    // if the needle is found in any of the paths.
    if (current === needle) {
        return true;
    }

    // recurse
    // 1st thing we need is the list of edges in our node
    const list = graph[current];
    // for every edges in the list of edges that we grab
    for (let i = 0; i < list.length; ++i) {
        const edge = list[i];

        // if we successfully walk this path and find the needle
        if (walk(graph, edge.to, needle, seen, path)) {
            // we need to break out of this. we have found it. we are good
            return true;
        }
    }

    // post
    path.pop();

    return false;
}
export default function dfs(
    graph: WeightedAdjacencyList,
    source: number,
    needle: number,
): number[] | null {
    const seen: boolean[] = new Array(graph.length).fill(false);
    const path: number[] = [];

    walk(graph, source, needle, seen, path);

    if (path.length === 0) {
        return null;
    }

    return path;
}
