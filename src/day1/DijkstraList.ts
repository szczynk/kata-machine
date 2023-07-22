// Dijkstra’s Shortest Path => greedy algorithm
// 
// Dijkstra’s Shortest Path is algorithm that calculate
// the shortest path from one node to all other nodes in the graph
// (we can find an individual node as well)
// 
// if every node has been visited, it is very easy to find shortest path
// from A to B as long as you have previous list 
// which mean we need to implement *BFS*, coz it required previous list 
// 
// 
// if we implement Dijkstra’s Shortest Path using MinHeap which
// have better running time O(log V(V+E))
//
// this is an example of how using a different data-structure alone
// 
// 
// but now we implement Dijkstra’s Shortest Path using BFS which runtime is O(V²+E)
function hasUnvisited(seen: boolean[], distances: number[]): boolean {
    // if it's not seen and it's distance less than infinity
    return seen.some((s, i) => !s && distances[i] < Infinity);
}

function getLowestUnvisited(seen: boolean[], distances: number[]): number {
    let idx = -1;
    let lowestDistance = Infinity;

    // walk through all of our distances and
    // get the lowest distance and unseen node
    for (let i = 0; i < seen.length; ++i) {
        if (seen[i]) {
            continue;
        }

        if (distances[i] < lowestDistance) {
            // update it
            lowestDistance = distances[i];
            idx = i;
        }
    }

    return idx;
}

export default function dijkstra_list(
    source: number,
    sink: number,
    arr: WeightedAdjacencyList,
): number[] {
    const seen = new Array(arr.length).fill(false);
    const prev = new Array(arr.length).fill(-1);

    const distances = new Array(arr.length).fill(Infinity);
    // smallest distances possible at the source
    distances[source] = 0;

    // get the unvisited node with the lowest distances
    while (hasUnvisited(seen, distances)) {
        console.log(distances)
        const current = getLowestUnvisited(seen, distances);

        seen[current] = true;

        // anything we select, we grab out the row 
        // (the connections to other nodes from the current node.)
        // by grabbing out the row, that represent our connection to other node
        const adjs = arr[current];
        // for every adjs in current node
        for (let i = 0; i < adjs.length; ++i) {
            const edge = adjs[i];

            if (seen[edge.to]) {
                continue;
            }

            // if we have not seen this edge,
            // we will calculate its distance,
            // which is based on my distance plus the distance to that edge
            const distance = distances[current] + edge.weight;
            // if less than its lowest known distance,
            if (distance < distances[edge.to]) {
                // we are gonna push/update
                // both the previous and distances array with new value
                prev[edge.to] = current;
                distances[edge.to] = distance;
            }
        }
    }

    // our out represents the path from the needle back to the source
    // start at the needle back to the source
    const out: number[] = [];

    // let's build it backward,
    // now we have to walk our previouses until we get to a -1
    // we need to start at the point we need to search from
    // or where we hoping to find
    let curr = sink;

    // keep searching until we found a point that has no more parent
    // if we never found needle (no path exists), the loop will terminate when prev[curr] becomes -1.
    while (prev[curr] !== -1) {
        // push current to the out
        out.push(curr);
        // set this current to the parent
        curr = prev[curr];
    }

    // add our source
    // finally return our out and reverse it
    out.push(source);
    return out.reverse();
}
