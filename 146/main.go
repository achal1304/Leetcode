package main

import "fmt"

// REVISIT = used doubly linked list and dict for O(1) complexity for get and put
type LRUNode struct {
	Key  int
	Val  int
	Prev *LRUNode
	Next *LRUNode
}

type LRUCache struct {
	head     *LRUNode
	tail     *LRUNode
	cache    map[int]*LRUNode
	capacity int
}

func Constructor(capacity int) LRUCache {
	lruCache := LRUCache{
		head:     &LRUNode{},
		tail:     &LRUNode{},
		cache:    make(map[int]*LRUNode),
		capacity: capacity,
	}
	lruCache.head.Next = lruCache.tail
	lruCache.tail.Prev = lruCache.head
	return lruCache
}

func (this *LRUCache) MoveToHead(key int) {
	node := this.cache[key]
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev

	nextHead := this.head.Next
	this.head.Next = node
	nextHead.Prev = node
	node.Next = nextHead
	node.Prev = this.head
}

func (this *LRUCache) RemoveTail() int {
	node := this.tail.Prev

	node.Prev.Next = this.tail
	this.tail.Prev = node.Prev
	node.Next = nil
	node.Prev = nil
	return node.Key
}

func (this *LRUCache) Insert(node *LRUNode) {
	tmp := this.head.Next
	this.head.Next = node
	tmp.Prev = node
	node.Next = tmp
	node.Prev = this.head
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.MoveToHead(key)
		return node.Val
	}
	return -1
}

func (this *LRUCache) PrintList() {
	curr := this.head.Next
	fmt.Println("new print")
	for curr != nil {
		fmt.Println(curr.Prev.Val, curr.Val)
		curr = curr.Next
	}
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.Val = value
		this.MoveToHead(key)
	} else {
		if this.capacity <= len(this.cache) {
			keyToDel := this.RemoveTail()
			delete(this.cache, keyToDel)
		}
		this.cache[key] = &LRUNode{Val: value, Key: key}
		this.Insert(this.cache[key])

	}
}

func main() {
	lruCache := Constructor(2)
	lruCache.PrintList()

	lruCache.Put(2, 1)
	lruCache.PrintList()

	lruCache.Put(1, 1)
	lruCache.PrintList()

	// fmt.Println("GetCall ", lruCache.Get(1))
	// lruCache.PrintList()

	lruCache.Put(2, 3)
	lruCache.PrintList()

	// fmt.Println("GetCall ", lruCache.Get(2))
	// lruCache.PrintList()

	lruCache.Put(4, 1)
	lruCache.PrintList()

	fmt.Println("GetCall ", lruCache.Get(1))
	lruCache.PrintList()
	fmt.Println("GetCall ", lruCache.Get(2))
	lruCache.PrintList()
	// fmt.Println("GetCall ", lruCache.Get(3))
	// lruCache.PrintList()
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
