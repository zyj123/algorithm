package internal

type cacheNode struct {
	key, value int
	next, prev *cacheNode
}

type LRUCache struct {
	cap        int
	head, tail *cacheNode
	m          map[int]*cacheNode
}

func Constructor(capacity int) LRUCache {
	var (
		head = &cacheNode{}
		tail = &cacheNode{}
	)
	head.next = tail
	tail.prev = head
	return LRUCache{
		cap:  capacity,
		head: head,
		tail: tail,
		m:    map[int]*cacheNode{},
	}
}

func (c *LRUCache) Get(key int) int {
	if node, ok := c.m[key]; ok {
		c.moveToHead(node)
		return node.value
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if node, ok := c.m[key]; ok {
		node.value = value
		c.moveToHead(node)
		return
	}
	if len(c.m) >= c.cap {
		c.removeLast()
	}
	c.putToHead(&cacheNode{
		key:   key,
		value: value,
	})
}

func (c *LRUCache) moveToHead(node *cacheNode) {
	c.remove(node)
	c.putToHead(node)
}

func (c *LRUCache) remove(node *cacheNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
	delete(c.m, node.key)
}

func (c *LRUCache) removeLast() {
	c.remove(c.tail.prev)
}

func (c *LRUCache) putToHead(node *cacheNode) {
	head := c.head
	node.next = head.next
	head.next.prev = node
	node.prev = head
	head.next = node
	c.m[node.key] = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

//实现 LRUCache 类：
//LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
//int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
//void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
//函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
