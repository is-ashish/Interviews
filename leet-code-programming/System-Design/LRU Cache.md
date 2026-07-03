1 - Create LRU Cache -: 

Design a data structure that follows the constraints of a Least Recently Used Cache.

Implement the LRUCache class:
(1) LRUCache(int capacity) Initialize the LRU cache with positive size capacity.

(2) int get(int key) Return the value of the key if the key exists, otherwise return -1.

(3) void put(int key, int value) Update the value of the key if the key exists. Otherwise, add the key-value pair to the cache. 

If the number of keys exceeds the capacity from this operation, evict the least recently used key.
The functions get and put must each run in O(1) average time complexity.



// 1 - Create LRU Cache -:

// Design a data structure that follows the constraints of a Least Recently Used Cache.

// Implement the LRUCache class:
// (1) LRUCache(int capacity) Initialize the LRU cache with positive size capacity.

// (2) int get(int key) Return the value of the key if the key exists, otherwise return -1.

// (3) void put(int key, int value) Update the value of the key if the key exists. Otherwise,
// add the key-value pair to the cache.

// If the number of keys exceeds the capacity from this operation, evict the least recently used key.
// The functions get and put must each run in O(1) average time complexity.

package main

import "fmt"

type Node struct {
	key, value int
	prev, next *Node
}

type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node
	tail     *Node
}

func Initialise(capacity int) *LRUCache {
	head := &Node{}
	tail := &Node{}

	head.next = tail
	tail.prev = head

	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     head,
		tail:     tail,
	}
}

func (l *LRUCache) Get(key int) int {
	if node, exists := l.cache[key]; exists {
		l.deleteNode(node)
		l.moveToHead(node)
		return node.value
	}
	return -1
}

func (l *LRUCache) Put(key, value int) {
	if node, exists := l.cache[key]; exists {
		node.value = value
		l.deleteNode(node)
		l.moveToHead(node)
	} else {
		newNode := &Node{key: key, value: value}
		l.moveToHead(newNode)
		l.cache[key] = newNode
		if len(l.cache) > l.capacity {
			node := l.getTail()
			delete(l.cache, node.key)
			l.deleteNode(node)
		}
	}
}

func (l *LRUCache) getTail() *Node {
	node := l.tail.prev
	return node
}

func (l *LRUCache) deleteNode(node *Node) {
	back := node.prev
	front := node.next
	back.next = front
	front.prev = back
}

func (l *LRUCache) moveToHead(node *Node) {
	node.prev = l.head
	node.next = l.head.next
	l.head.next.prev = node
	l.head.next = node
}

func main() {
	lru := Initialise(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	fmt.Println(lru.Get(1)) // 1
	lru.Put(3, 3)           // evicts key 2
	fmt.Println(lru.Get(2)) // -1
	lru.Put(4, 4)           // evicts key 1
	fmt.Println(lru.Get(1)) // -1
	fmt.Println(lru.Get(3)) // 3
	fmt.Println(lru.Get(4)) // 4

}
