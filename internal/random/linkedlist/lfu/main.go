package main

import "container/list"

// 460. LFU 缓存
// https://leetcode.cn/problems/lfu-cache/

// LFUCache 直接使用 Map + LinkedList 来解决会超时
// 在整个 LinkedList 链条上操作非常消费性能，因此需要根据频次，来对 list 进行分组
type LFUCache struct {
	KeyToNodeMap map[int]*list.Element
	LinkedList   map[int]*list.List
	Size         int
	Capacity     int
	MinFreq      int
}

type Node struct {
	Key   int
	Value int
	Freq  int
}

const (
	minFreq         = 1
	defaultCapacity = 10
)

func Constructor(capacity int) LFUCache {
	thisCap := capacity
	if capacity <= 0 {
		thisCap = defaultCapacity
	}

	cache := new(LFUCache)
	cache.Capacity = thisCap
	cache.Size = 0
	cache.KeyToNodeMap = make(map[int]*list.Element, thisCap)
	cache.LinkedList = make(map[int]*list.List, thisCap)
	return *cache
}

func (c *LFUCache) Get(key int) int {
	element, ok := c.KeyToNodeMap[key]
	if !ok {
		// 没有对应的值，返回 -1
		return -1
	}

	node := element.Value.(Node)
	oldFreq := node.Freq
	oldList := c.LinkedList[oldFreq]
	oldList.Remove(element)
	node.Freq++

	// 维护最小 Freq
	if oldFreq == c.MinFreq && oldList.Len() == 0 {
		delete(c.LinkedList, oldFreq)
		c.MinFreq++
	}

	newList, ok := c.LinkedList[node.Freq]
	if !ok {
		c.LinkedList[node.Freq] = list.New()
		newList = c.LinkedList[node.Freq]
	}

	newList.PushFront(node)
	c.KeyToNodeMap[key] = newList.Front()
	return node.Value
}

func (c *LFUCache) Put(key int, value int) {
	// 如果存在
	element, ok := c.KeyToNodeMap[key]
	if ok {
		n := element.Value.(Node)
		n.Value = value
		n.Freq++

		oldFreq := n.Freq - 1
		oldList := c.LinkedList[oldFreq]
		oldList.Remove(element)
		if oldList.Len() == 0 {
			delete(c.LinkedList, oldFreq)
			if c.MinFreq == oldFreq {
				c.MinFreq++
			}
		}

		newFreq := n.Freq
		newList, ok := c.LinkedList[newFreq]
		if !ok {
			c.LinkedList[newFreq] = list.New()
			newList = c.LinkedList[newFreq]
		}

		newList.PushFront(n)
		c.KeyToNodeMap[key] = newList.Front()
		return
	}

	// 如果不存在
	if c.Size == c.Capacity {
		// 超出时，需要移除最不经常使用的项
		minList := c.LinkedList[c.MinFreq]
		backNode := minList.Back()
		minList.Remove(backNode)

		delete(c.KeyToNodeMap, backNode.Value.(Node).Key)
		if minList.Len() == 0 {
			delete(c.LinkedList, c.MinFreq)
		}
		c.Size--
	}

	// 未超出时，正常插入
	node := Node{Key: key, Value: value, Freq: minFreq}
	newList, ok := c.LinkedList[node.Freq]
	if !ok {
		c.LinkedList[node.Freq] = list.New()
		newList = c.LinkedList[node.Freq]
	}

	newList.PushFront(node)
	c.KeyToNodeMap[key] = newList.Front()
	c.MinFreq = node.Freq
	c.Size++
}

func main() {
	lfu := Constructor(2)
	lfu.Put(1, 1)
	lfu.Put(2, 2)
	println(lfu.Get(1))
	lfu.Put(3, 3)
	println(lfu.Get(2))
	println(lfu.Get(3))
	lfu.Put(4, 4)
	println(lfu.Get(1))
	println(lfu.Get(3))
	println(lfu.Get(4))
	println()
	// 1 -1 3 -1 3 4

	lfu = Constructor(3)
	lfu.Put(2, 2)
	lfu.Put(1, 1)
	println(lfu.Get(2))
	println(lfu.Get(1))
	println(lfu.Get(2))
	lfu.Put(3, 3)
	lfu.Put(4, 4)
	println(lfu.Get(3))
	println(lfu.Get(2))
	println(lfu.Get(1))
	println(lfu.Get(4))
	println()
	// 2 1 2 -1 2 1 4

	lfu = Constructor(2)
	lfu.Put(3, 1)
	lfu.Put(2, 1)
	lfu.Put(2, 2)
	lfu.Put(4, 4)
	println(lfu.Get(2))
	println()
	// 2
}
