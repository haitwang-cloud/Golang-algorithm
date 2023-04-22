package lru

import "container/list"

// LRUCache is a LRU cache implementation
type LRUCache struct {
    size      int
    capacity  int
    cache     map[interface{}]*list.Element
    eviction *list.List
}

type entry struct {
    key   interface{}
    value interface{}
}

// NewLRUCache creates a new LRUCache with the specified capacity
func NewLRUCache(capacity int) *LRUCache {
    return &LRUCache{
        capacity:  capacity,
        cache:     make(map[interface{}]*list.Element),
        eviction: list.New(),
    }
}

// Put adds an entry to the cache
func (l *LRUCache) Put(key interface{}, value interface{}) {
    if l.cache == nil {
        l.cache = make(map[interface{}]*list.Element)
        l.eviction = list.New()
    }

    if el, ok := l.cache[key]; ok {
        l.eviction.MoveToFront(el)
        el.Value.(*entry).value = value
        return
    }

    newEl := l.eviction.PushFront(&entry{key: key, value: value})
    l.cache[key] = newEl
    l.size++

    if l.size > l.capacity {
        l.evict()
    }
}

// Get retrieves an entry from the cache
func (l *LRUCache) Get(key interface{}) (interface{}, bool) {
    if l.cache == nil {
        return nil, false
    }

    if el, ok := l.cache[key]; ok {
        l.eviction.MoveToFront(el)
        return el.Value.(*entry).value, true
    }

    return nil, false
}

// evict removes the oldest entry from the cache
func (l *LRUCache) evict() {
    if l.cache == nil {
        return
    }

    lastEl := l.eviction.Back()
    if lastEl != nil {
        l.removeElement(lastEl)
    }
}

// removeElement removes an entry from the cache
func (l *LRUCache) removeElement(el *list.Element) {
    l.eviction.Remove(el)
    k := el.Value.(*entry).key
    delete(l.cache, k)
    l.size--
}

// Len returns the number of entries in the cache
func (l *LRUCache) Len() int {
    return l.size
}
