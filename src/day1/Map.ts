// # What are maps?
//
// its too easy to think of {} or new Map() and just assume we know them.
// The reality is that we know how to use them,
// most of the time we don't actually know they work.
//
// {} this is not a map, this is an object that has different properties
// associated with it
//
// (we wont demistify {}, those are different data structures than a hash map)
//
// # what a map actually is
// ## Terms
// - load factor: The amount of data points vs the amount of storage (data.len / storage.capacity),
//   e.g. 7 items in map, 10 possible slots = 0.7 load factor
// - key: a value that is hashable and We use the key to create something in which we can access
//   our value.The hash has to be consistent.
// - value: a value that is associated with a key
// - collision: when 2 keys map to the same cell.
//
// Actually, joke's on you, it's an ArrayList or LinkedList.
//
// a map is a data structure that maps keys to values.
// And what the key is mapping to a value has to be, is a consistent hash to work correctly.
//
// thus we need a “hashing function” that takes our key as input and produces
// a consistent hash (some sort of unique number) as output.
//
// this hash needs to be a number, so we can take the modulo (%) of it with
// our storage capacity / the length of our data storage, so the number always
// produces an index of one of the available slot inside of our array structure
// and store the key plus the value inside of that structure right there.
// e.g.
// the length of data storage   => 10
// key                          => 123456789
// index                        => 123456789 % 10 = 9
// array structure              => [...,...,{ key: 123456789, value: "test" },...]
//
// however, since we have limited slots, the chance of them colliding is uniformal.
// this immediately shows the danger of collision
// e.g.
// key                          => 123456719
// index                        => 123456719 % 10 = 9
//
// what happens when we collide? Well, we need a way to be able to store both of them.
// What we used to do to handle collision with a “linear” or an “exponential” backoff.
// Meaning we would literally go down to the next slot and put it in there.
// Which means that if you did this, you will fill up more and more slots.
// So as things become less and less efficient, or more and more full,
// you're gonna need a larger amount of storage with a smaller amount of load factor
// to prevent collisions
//
// a more modern approach is that the slot is not a unique key-value pair,
// but you use a list underneath the hood, so upon collision you can just add an item to the list
//    ```
//    const hashtable =
//    [
//        [],
//        [{ k, v }],
//        [{ k, v }, { k, v }],
//        [],
//    ];
//    ```
//
// - So you have an ArrayList, that has an ArrayList, that can add an items and
//   then we just walk the ArrayList.
// - We could technically also use a LinkedList underneath the hood
//   because you do linearly walk it every single time
//
// for retrieval / deletion:
// We pass our key, we go through a hashing function and produces an index,
// we access the storage unit by the index,
// then we can go look for the assosiated value using our key
// and finally get or delete that.
//
// the "ideal" load factor is 0.7 at this point as long as you do this
// in ArrayLists or LinkedList style storing for efficiency.
//
// as maps grow or shrink and exceed the "ideal" load factor.
// To maintain the "ideal" load factor, we can take our current data storage unit,
// iterate through all of the keys that we can find in that,
// then we can rehash them and restore them into a larger storage arena.
//
// Always a trade off between how much do you want to store
// versus how efficient you want to be with memory.
//
// So it's always a game. That's why often you'll see maps come with some sort of capacity.
// Because they want you to give them the hint of how big we size should be so that way
// the storage and retrieval is as efficient as possible.
// Because the more amount of collisions, the less O(1) it is.
//
// time complexity of hashing and retrieval = O(1)
// assuming we have very good hashing algorithm and
// as long as our collisions don't gather up into one single point
//
// implementing a map in JavaScript might be challenging due to limited tools
// and might not be necessary due to the built-in map functionality in the language.
export default class Map<T extends string | number, V> {
    constructor() {}

    get(key: T): V | undefined {}
    set(key: T, value: V): void {}
    delete(key: T): V | undefined {}
    size(): number {}
}
