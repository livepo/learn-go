package main

import "fmt"


type Node struct { left, right *Node; key int }


type List struct { head, tail *Node }


type Value struct { value int; ptr *Node }


type LRUCache struct {
	cache map[int]Value
	capacity int
	list List
	length int
}


func Constructor(capacity int) LRUCache {
	head := new(Node)
	tail := head
	list := List{head: head, tail: tail}
	cache := make(map[int]Value, 0)
	return LRUCache{cache: cache, capacity: capacity, list: list, length: 0}
}


func (this *LRUCache) Get(key int) int {
	v, ok := this.cache[key]
	if !ok { return -1 }

	if v.ptr == this.list.tail {
		this.list.tail = v.ptr.left
		v.ptr.left.right = v.ptr.right
	} else {
		v.ptr.left.right = v.ptr.right
		v.ptr.right.left = v.ptr.left

	}
	this.list.tail.right = v.ptr
	v.ptr.left = this.list.tail
	v.ptr.right = nil
	this.list.tail = v.ptr
	// fmt.Println("GET", key)
	// fmt.Println("MAP", *this)
	return v.value
}


func (this *LRUCache) Put(key int, value int)  {
	if this.capacity == 0 { return }
	if v, ok := this.cache[key]; ok {
		this.cache[key] = Value{value: value, ptr: v.ptr}
		this.Get(key)

		// fmt.Println("PUT", key, value)
		// fmt.Println("MAP", *this)
		return
	}
	if this.length == this.capacity {
		// rm most left node
		first := this.list.head.right
		this.list.head.right = first.right
		if first.right != nil {
			first.right.left = this.list.head
		} else {
			this.list.tail = first.left
		}
		delete(this.cache, first.key)
		this.length --
	}
	ptr := new(Node)
	ptr.key = key
	ptr.left = this.list.tail
	this.list.tail.right = ptr
	this.list.tail = ptr
	this.cache[key] = Value{value: value, ptr: ptr}
	this.length ++
	// fmt.Println("PUT", key, value)
	// fmt.Println("MAP", *this)
}


func main() {
	cache := Constructor( 2 /* capacity */ )
	cache.Put(2, 6)
	cache.Put(1, 5)
	cache.Put(1, 2)
	cache.Get(1)
	cache.Get(2)
}
