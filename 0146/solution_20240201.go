package _146

type CacheNode struct {
	Key  int
	Val  int
	Prev *CacheNode
	Next *CacheNode
}

type LRUCache struct {
	Head     *CacheNode
	Hash     map[int]*CacheNode
	Capacity int
	Count    int
}

func Constructor(capacity int) LRUCache {

	head := &CacheNode{}
	head.Next = head
	head.Prev = head

	return LRUCache{
		Head:     head,
		Hash:     make(map[int]*CacheNode, 0),
		Capacity: capacity,
		Count:    0,
	}
}

func (cache *LRUCache) Get(key int) int {
	if result, ok := cache.Hash[key]; ok {
		result.Prev.Next = result.Next
		result.Next.Prev = result.Prev
		result.Next = cache.Head.Next
		cache.Head.Next.Prev = result
		cache.Head.Next = result
		result.Prev = cache.Head
		return result.Val
	} else {
		return -1
	}
}

func (cache *LRUCache) Put(key int, value int) {

	if result, ok := cache.Hash[key]; ok {
		result.Val = value
		cache.Get(key)
		return
	}

	if cache.Count == cache.Capacity {
		delete(cache.Hash, cache.Head.Prev.Key)
		cache.Count--
		cache.Head.Prev.Prev.Next = cache.Head
		cache.Head.Prev = cache.Head.Prev.Prev
	}

	node := &CacheNode{Key: key, Val: value}
	cache.Hash[key] = node
	cache.Count++

	node.Next = cache.Head.Next
	cache.Head.Next.Prev = node
	cache.Head.Next = node
	node.Prev = cache.Head
}
