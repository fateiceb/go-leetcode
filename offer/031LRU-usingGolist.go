package main

import (
	"container/list"
)

type entry struct {
	key, val int
}
type LRUCache struct {
	Map map[int]*list.Element
	L   *list.List
	cap int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Map: make(map[int]*list.Element),
		L:   list.New(),
		cap: capacity,
	}
}
func (this *LRUCache) Get(key int) int {
	//更新,移动至最后
	if e := this.Map[key]; e != nil {
		this.L.MoveToBack(e)
		return e.Value.(entry).val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if e := this.Map[key]; e != nil {
		e.Value = entry{key, value}
		this.L.MoveToBack(e)
		return
	}
	//添加新元素至链表最后
	e := entry{key, value}
	this.L.PushBack(e)
	this.Map[key] = this.L.Back()
	//达到容量限制
	if len(this.Map) > this.cap {
		e := this.L.Front()
		this.L.Remove(e)
		delete(this.Map, e.Value.(entry).key)
	}
}
func main() {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	lru.Get(1)
	lru.Put(3, 3)
	lru.Get(2)
}
