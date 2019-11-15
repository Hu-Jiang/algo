package medium

// Design and implement a data structure for Least Recently Used (LRU) cache.
// It should support the following operations: get and put.
//
// get(key) - Get the value (will always be positive) of the key if the key
// exists in the cache, otherwise return -1.
//
// put(key, value) - Set or insert the value if the key is not already present.
// When the cache reached its capacity, it should invalidate the least recently
// used item before inserting a new item.
//
// The cache is initialized with a positive capacity.
//
// Follow up:
// Could you do both operations in O(1) time complexity?
//
// Example:
// LRUCache cache = new LRUCache( 2 /* capacity */ );
// cache.put(1, 1);
// cache.put(2, 2);
// cache.get(1);       // returns 1
// cache.put(3, 3);    // evicts key 2
// cache.get(2);       // returns -1 (not found)
// cache.put(4, 4);    // evicts key 1
// cache.get(1);       // returns -1 (not found)
// cache.get(3);       // returns 3
// cache.get(4);       // returns 4

import "container/list"

type LRUCache struct {
	cap       int
	mapCache  map[int]*list.Element
	listCache *list.List
}

type cacheItem struct {
	key, value int
}

var defaultLRUCacheCap = 1 << 10

func ConstructorLRUCache(cap int) LRUCache {
	if cap <= 0 {
		cap = defaultLRUCacheCap
	}
	return LRUCache{
		cap:       cap,
		mapCache:  make(map[int]*list.Element),
		listCache: list.New(),
	}
}

// Get returns the value of the key if the key exists int the
// cache, otherwise returns -1.
func (lru *LRUCache) Get(key int) int {
	v, ok := lru.mapCache[key]
	if !ok {
		return -1
	}

	lru.listCache.MoveToFront(v)
	return v.Value.(*cacheItem).value
}

// Put inserts the value of the key.
// If key exists, then substitute the value of key.
func (lru *LRUCache) Put(key, value int) {
	v, ok := lru.mapCache[key]
	if ok {
		v.Value.(*cacheItem).value = value
		lru.listCache.MoveToFront(v)
		return
	}

	if lru.Len() >= lru.cap {
		lru.RemoveOldest()
	}

	item := &cacheItem{key: key, value: value}
	element := lru.listCache.PushFront(item)
	lru.mapCache[key] = element
}

// RemoveOldest remove the oldest item of LRUCache lru.
func (lru *LRUCache) RemoveOldest() {
	element := lru.listCache.Back()
	if element == nil {
		return
	}
	lru.listCache.Remove(element)
	delete(lru.mapCache, element.Value.(*cacheItem).key)
}

// Len returns the length of LRUCache lru.
func (lru *LRUCache) Len() int {
	return lru.listCache.Len()
}
