package lruCache

//leetcode submit region begin(Prohibit modification and deletion)
type LRUCache struct {
	size       int
	cap        int
	cache      map[int]*LinkNode
	head, tail *LinkNode
}

type LinkNode struct {
	key, val  int
	pre, next *LinkNode
}

func Constructor(capacity int) LRUCache {
	dummyHead := initLinkNode(0, 0)
	dummyTail := initLinkNode(0, 0)
	dummyHead.next = dummyTail
	return LRUCache{
		size:  0,
		cap:   capacity,
		cache: make(map[int]*LinkNode),
		head:  dummyHead,
		tail:  dummyTail,
	}

}

func initLinkNode(key, val int) *LinkNode {
	return &LinkNode{
		key:  key,
		val:  val,
		pre:  nil,
		next: nil,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.moveToHead(node)
		return node.val
	}
	return -1

}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.val = value
		this.moveToHead(node)
		return
	}

	node := initLinkNode(key, value)
	this.cache[key] = node
	this.addToHead(node)
	this.size++
	if this.size > this.cap {
		tail := this.moveTail()
		delete(this.cache, tail.key)
		this.size--
	}

}

func (thie *LRUCache) removeNode(node *LinkNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (this *LRUCache) moveTail() *LinkNode {
	node := this.tail.pre
	this.removeNode(node)
	return node

}

func (this *LRUCache) addToHead(node *LinkNode) {
	node.next = this.head.next
	node.pre = this.head
	this.head.next.pre = node
	this.head.next = node
}

func (this *LRUCache) moveToHead(node *LinkNode) {
	this.removeNode(node)
	this.addToHead(node)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
//leetcode submit region end(Prohibit modification and deletion)
