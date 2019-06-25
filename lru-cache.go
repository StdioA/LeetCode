package main

import "fmt"

type LRUNode struct {
	key        int
	value      int
	prev, next *LRUNode
}

type LRUCache struct {
	keyMap     map[int]*LRUNode
	head, tail *LRUNode
	Cap, Len   int
}


func Constructor(capacity int) LRUCache {
    vnode := new(LRUNode)
	cache := LRUCache{
		make(map[int]*LRUNode),
		vnode, vnode,
		capacity, 1,
	}
	return cache
}


func (this *LRUCache) elevate(node *LRUNode) {
	prev, next := node.prev, node.next
	if node != this.head {
		if prev != nil {
			prev.next = next
		}
		if next != nil {
			next.prev = prev
		}
		prevHead := this.head
		prevHead.prev, node.next = node, prevHead
		node.prev = nil
		this.head = node
	}
	if node == this.tail && prev != nil {
		this.tail = prev
	}
}


func (this *LRUCache) Get(key int) int {
    var value int
    node, hit := this.keyMap[key]
    if hit {
		value = node.value
		this.elevate(node)
    } else {
        value = -1
    }
    return value
}

func (this *LRUCache) Put(key int, value int)  {
	node, ext := this.keyMap[key]
	if ext {
		node.value = value
		this.elevate(node)
		return
	}

	node = &LRUNode{
		key:   key,
		value: value,
	}
    this.keyMap[key] = node
    prevHead := this.head
    prevHead.prev, node.next = node, prevHead
	this.head = node
	
    if this.Len+1 > this.Cap {
		// eliminate tail node
		tail := this.tail
        delete(this.keyMap, tail.key)
        this.tail = tail.prev
        this.tail.next = nil
	} else {
		this.Len++
	}
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
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(2))
}
