// trie => re"trie"val tree or prefix tree or digital tree
// 
// think like an autocomplete
// 
// implementation of a Trie
// - autocomplete
// - cache mechanism (like Folker)
// 
// for instance, English language spellchecker/autocomplete as a trie
// with a root that have empty value and
// 26 possible children that have 26 characters in the alphabet
// if we want to add word "cat", 
// we need to insert every characters for every node, e.g.
// 
//                root            root
//                 /               /
//                C               C
//               /               /
//               A               A 
//                \               \
//                 T               T[isWord: true]
//                 |
//                 *
// 
// there are two method to determine is a word
// 1. add asterisk to denote this path is a word
// 2. add isWord (boolean) to the node itself (recommended)
// 
// then we can add word "cats", "cattle", "card", and "mark"
// 
//                root
//                 / \
//                C   M
//               /     \
//              A       A
//             / \       \
//            R   T       R
//           /   / \       \
//          D    S  T       K
//                  /
//                 L
//                /
//               E
// 
// how we do find?
// if someone type "c", do we have "c" branch?
// if yes -> traverse the trie, recursively check is a word (isWord: true)
//    if yes -> out that out, 
//    if no  -> iterate over all children and recursively check is a word to next character
// 
// how we can optimize to find the right word in the middle of huge words?
// we can add a score / frequency to the node itself to get most selected
// words that always on the top
// 
// how we do insertion?
// easy, 
// for every characters in the string
//   if that character exist, traverse that node
//   else, create a new node, fill that node with new value
// mark the last character is a word (isWord: true)
// 
// how we do deletion?
// deletion need recursion
// traverse trie to end of the node and recursively check is a word (isWord: true)
//   if yes -> mark it so it's not a word anymore, check if it has children
//     if yes -> we can safely delete
//   if no  -> recursively check is a word to next character
// 
// running time => O(height)
export default class Trie {
    

    

    constructor() {
    }

    insert(item: string): void {

}
    delete(item: string): void {

}
    find(partial: string): string[] {

}
}