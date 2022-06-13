package main

import (
	"container/list"
	"sync"
)

type LRUCache struct {
	capacity int
	lru      *list.List
	cache    map[int]*list.Element
	sync.Mutex
}

type elem struct {
	val int
	key int
}

func Constructor(capacity int) LRUCache {

	return LRUCache{
		capacity: capacity,
		lru:      list.New(),
		cache:    make(map[int]*list.Element),
	}
}

func (this *LRUCache) Get(key int) int {
	this.Lock()
	defer this.Unlock()
	if v, ok := this.cache[key]; ok {
		this.lru.MoveToBack(v)
		return v.Value.(elem).val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	this.Lock()
	defer this.Unlock()
	v, ok := this.cache[key]
	if ok {
		i := v.Value.(elem)
		i.val = value
		v.Value = i
		this.lru.MoveToBack(v)

		return
	} else {
		if len(this.cache) == this.capacity {
			k := this.lru.Remove(this.lru.Front()).(elem).key
			delete(this.cache, k)
			newElem := elem{
				val: value,
				key: key,
			}
			this.cache[key] = this.lru.PushBack(newElem)

		} else {
			newElem := elem{
				val: value,
				key: key,
			}
			this.cache[key] = this.lru.PushBack(newElem)
		}
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
