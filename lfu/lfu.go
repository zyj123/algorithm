package lfu

type DNode struct {
	key  int
	val  int
	freq int

	prev *DNode
	next *DNode
}

type DList struct {
	head *DNode
	tail *DNode
}

func NewDList() *DList {
	head, tail := &DNode{}, &DNode{}
	head.next = tail
	tail.prev = head

	return &DList{
		head: head,
		tail: tail,
	}
}

func (d *DList) Add(dNode *DNode) {
	headNode, nextNode := d.head, d.head.next

	dNode.prev = headNode
	dNode.next = nextNode
	headNode.next = dNode
	nextNode.prev = dNode
}

func (d *DList) RemoveLast() *DNode {
	last := d.tail.prev
	d.Remove(last)
	return last
}

func (d *DList) Remove(dNode *DNode) {
	prevNode, nextNode := dNode.prev, dNode.next

	prevNode.next = nextNode
	nextNode.prev = prevNode
}

func (d *DList) Evict() {
	last := d.tail.prev
	d.Remove(last)
}

func (d *DList) IsEmpty() bool {
	return d.head.next == d.tail
}

type LFUCache struct {
	cap       int
	minFreq   int
	freqTable map[int]*DList
	keyTable  map[int]*DNode
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		cap:     capacity,
		minFreq: 1,
		freqTable: map[int]*DList{
			1: NewDList(),
		},
		keyTable: map[int]*DNode{},
	}
}

func (lfu *LFUCache) Get(key int) int {
	dNode := lfu.keyTable[key]
	if dNode == nil {
		return -1
	}

	lfu.freqIncr(dNode)
	return dNode.val
}

func (lfu *LFUCache) Put(key int, value int) {
	if dNode, ok := lfu.keyTable[key]; !ok {
		if len(lfu.keyTable) == lfu.cap {
			lfu.evict()
		}
		dNode = &DNode{
			key:  key,
			val:  value,
			freq: 1,
		}
		lfu.keyTable[key] = dNode
		lfu.freqTable[1].Add(dNode)
		lfu.minFreq = 1
	} else {
		dNode.val = value
		lfu.freqIncr(dNode)
	}
}
func (lfu *LFUCache) evict() {
	dNode := lfu.freqTable[lfu.minFreq].RemoveLast()
	delete(lfu.keyTable, dNode.key)
}

func (lfu *LFUCache) freqIncr(dNode *DNode) {
	oldFreq := dNode.freq
	lfu.freqTable[oldFreq].Remove(dNode)
	dNode.freq += 1
	if _, ok := lfu.freqTable[dNode.freq]; !ok {
		lfu.freqTable[dNode.freq] = NewDList()
	}
	lfu.freqTable[dNode.freq].Add(dNode)
	if lfu.minFreq == oldFreq && lfu.freqTable[oldFreq].IsEmpty() {
		lfu.minFreq += 1
	}
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
