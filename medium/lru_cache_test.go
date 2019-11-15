package medium

import "testing"

func TestLRUCache(t *testing.T) {
	cache := ConstructorLRUCache(2)
	if cache.Len() != 0 {
		t.Fatal("expected the length of lru cache is zero")
	}

	cache.Put(1, 1)
	cache.Put(2, 2)
	if cache.Len() != 2 {
		t.Fatal("expected the length of lru cache is 2")
	}
	head := HeadOfLRUCache(&cache)
	if head != 2 {
		t.Fatalf("expected the front element of list is 2, but got %d", head)
	}

	v := cache.Get(1)
	if v != 1 {
		t.Fatalf("expected value is 1 of key 1, but got %d", v)
	}
	head = HeadOfLRUCache(&cache)
	if head != 1 {
		t.Fatalf("expected the front element of list is 1, but got %d", head)
	}

	cache.Put(3, 3)
	v = cache.Get(2)
	if v != -1 {
		t.Fatal("expected key 2 has been deleted")
	}

	cache.Put(4, 4)
	v = cache.Get(1)
	if v != -1 {
		t.Fatal("expected key 1 has been deleted")
	}

	if v3, v4 := cache.Get(3), cache.Get(4); v3 != 3 || v4 != 4 {
		t.Fatalf("expected v3, v4 = 3, 4. but got v3, v4 = %d, %d", v3, v4)
	}
}

func HeadOfLRUCache(cache *LRUCache) int {
	if cache == nil {
		return -1
	}
	return cache.listCache.Front().Value.(*cacheItem).value
}
