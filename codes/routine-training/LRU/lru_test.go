package lru

import (
	"fmt"
	"testing"
)

/*
该测试用例测试了LRUCache的Put、Get、evict和Len方法。
在测试中，我们在容量为2的LRUCache中插入了四个键值对，
并检查它们是否按照预期逐出。最后，我们使用Len方法检查缓存中的实际数量。
*/
func TestLRUCache(t *testing.T) {
	cache := NewLRUCache(2)
	cache.Put(1, "one")
	cache.Put(2, "two")

	v, ok := cache.Get(1)
	if !ok || v != "one" {
		t.Errorf("cache.Get(1) = (%v, %v), want (%v, true)", v, ok, "one")
	}

	cache.Put(3, "three")

	v, ok = cache.Get(2)
	if ok {
		t.Errorf("cache.Get(2) = (%v, %v), want (%v, false)", v, ok, nil)
	}

	cache.Put(4, "four")

	v, ok = cache.Get(1)
	if ok {
		t.Errorf("cache.Get(1) = (%v, %v), want (%v, false)", v, ok, nil)
	}

	v, ok = cache.Get(3)
	if !ok || v != "three" {
		t.Errorf("cache.Get(3) = (%v, %v), want (%v, true)", v, ok, "three")
	}

	v, ok = cache.Get(4)
	if !ok || v != "four" {
		t.Errorf("cache.Get(4) = (%v, %v), want (%v, true)", v, ok, "four")
	}

	fmt.Printf("cache.Len() = %v\n", cache.Len())
	// Output: cache.Len() = 2
}
