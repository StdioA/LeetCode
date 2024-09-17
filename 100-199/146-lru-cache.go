package main

import "fmt"

type LRUNode struct {
	key, value int
	prev, next *LRUNode
}

type LRUCache struct {
	head, tail *LRUNode
	cache      map[int]*LRUNode
	Cap, Len   int
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		head:  new(LRUNode),
		tail:  new(LRUNode),
		cache: make(map[int]*LRUNode),
		Cap:   capacity,
		Len:   0,
	}
	cache.head.next = cache.tail
	cache.tail.prev = cache.head
	return cache
}

func (this *LRUCache) elevate(node *LRUNode) {
	// disconnect
	prev, next := node.prev, node.next
	if prev != nil && next != nil {
		prev.next, next.prev = next, prev
	}

	// Connect with head
	head, next := this.head, this.head.next
	head.next, next.prev = node, node
	node.next, node.prev = next, head
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.elevate(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.cache[key]
	if ok {
		node.value = value
	} else {
		// Put new node to the head
		node = &LRUNode{
			key:   key,
			value: value,
		}
		this.cache[key] = node
		if len(this.cache) > this.Cap {
			// Evict tail node
			tail := this.tail.prev
			delete(this.cache, tail.key)
			prev := tail.prev
			this.tail.prev, prev.next = prev, this.tail
		}
	}
	this.elevate(node)
}

func (cache *LRUCache) print() {
	fmt.Println("-------")
	fmt.Println("head & tail", cache.head, cache.tail)
	for p := cache.head; p != nil; p = p.next {
		fmt.Println(p)
	}
	fmt.Println("-------")
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
func main() {
	cache := Constructor(2)
	cache.Put(2, 1)
	cache.Put(1, 1)
	cache.Put(2, 3)
	cache.Put(4, 1)
	cache.print()
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(2))
}
