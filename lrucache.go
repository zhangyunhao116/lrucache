package lrucache

import (
	"github.com/ZYunH/goutils"
	"sync"
)

type lruCache struct {
	m       map[string]*node
	root    *node
	maxSize int
	hits    int
	misses  int

	lock        sync.RWMutex
	_buf        []byte
	_bufNodePtr *node
}

type node struct {
	key   string
	value interface{}
	prev  *node
	next  *node
}

// Indicates 64-bit or 32-bit system.
const bit = 32 << (^uint(0) >> 63)

// New creates a new LRU cache with max size.
func New(maxSize int) *lruCache {
	if maxSize <= 0 {
		panic("maxSize must be greater than 0")
	}
	root := &node{}
	root.next = root
	root.prev = root
	return &lruCache{m: make(map[string]*node, maxSize), root: root, _buf: make([]byte, 0, 128), maxSize: maxSize}
}

// Set single key and value.
//
// The returned value indicates whether a key is eliminated from cache.
func (c *lruCache) Set(key, value interface{}) (isRemove bool) {
	c.lock.Lock()
	k := goutils.BytesToStringNew(interfaceToBytesWithBuf(c._buf, key))
	// Grow buffer slice to preparing enough space for next conversion.
	if cap(c._buf) < len(k) {
		c._buf = make([]byte, 0, len(k))
	}
	isRemove = c.set(k, value)
	c.lock.Unlock()
	return isRemove
}

func (c *lruCache) set(k string, value interface{}) bool {
	c._bufNodePtr = c.m[k]
	if c._bufNodePtr == nil { // This means the k not in the map
		if len(c.m) < c.maxSize-1 {
			// Cache is not full, insert a new node
			_node := &node{}
			_node.key = k
			_node.value = value
			_node.next = c.root
			_node.prev = c.root.prev
			c.m[k] = _node

			c.root.prev.next = _node
			c.root.prev = _node
		} else {
			// Cache is full, replace the oldest one with the new node,
			// in this case, we just replace the original root with the
			// new root, and make the original root.next become the new root.
			delete(c.m, c.root.key)
			c.root.key = k
			c.root.value = value
			c.m[k] = c.root
			c.root = c.root.next

			return true
		}
	} else {
		// Hits a key, we just update its value.
		c._bufNodePtr.value = value
	}
	return false
}

// Get value via a single key.
func (c *lruCache) Get(key interface{}) (value interface{}, ok bool) {
	c.lock.Lock()
	k := goutils.BytesToString(interfaceToBytesWithBuf(c._buf, key))
	// Grow buffer slice to preparing enough space for next conversion.
	if cap(c._buf) < len(k) {
		c._buf = make([]byte, 0, len(k))
	}
	value, ok = c.get(k)
	c.lock.Unlock()
	return
}

func (c *lruCache) get(k string) (interface{}, bool) {
	c._bufNodePtr = c.m[k]

	if c._bufNodePtr != nil {
		// Hits a key, drop it from the original location, and insert it
		// to the location between root.prev and root (The latest location in cache)
		c.hits++
		c._bufNodePtr.prev.next = c._bufNodePtr.next
		c._bufNodePtr.next.prev = c._bufNodePtr.prev
		c.root = c.root.next
		c._bufNodePtr.prev = c.root.prev
		c._bufNodePtr.next = c.root

		c.root.prev.next = c._bufNodePtr
		c.root.prev = c._bufNodePtr

		return c._bufNodePtr.value, true
	}

	// Here means the k not in the map
	c.misses++
	return nil, false
}

// Set multi-keys and corresponding single value, the last argument in kvs
// is the value, this means that len(kvs) must >= 2, or panic will occur.
//
// Keep in mind that byte slice or string is better to have only one, this
// means the key-arguments only actually includes a string or a byte slice,
// since our strategy is just map interface{} to some bytes, potential data conflict
// can be occur if string or byte slice more than one. If you insist on doing so,
// don't pass binary data as string or byte slice, it can increase the risk of
// data conflict. Keep string or byte slice as printable is a good idea to avoid
// potential data conflict.
func (c *lruCache) MSet(kvs ...interface{}) (isRemove bool) {
	if len(kvs) < 2 {
		panic("at least one key and one value")
	}
	c.lock.Lock()
	key := goutils.BytesToStringNew(interfaceToBytesWithBuf(c._buf, kvs[:len(kvs)-1]...))
	// Grow buffer slice to preparing enough space for next conversion.
	if cap(c._buf) < len(key) {
		c._buf = make([]byte, 0, len(key))
	}
	value := kvs[len(kvs)-1]
	isRemove = c.set(key, value)
	c.lock.Unlock()
	return
}

// Get value via multi-keys.
func (c *lruCache) MGet(keys ...interface{}) (value interface{}, ok bool) {
	c.lock.Lock()
	key := goutils.BytesToString(interfaceToBytesWithBuf(c._buf, keys...))
	// Grow buffer slice to preparing enough space for next conversion.
	if cap(c._buf) < len(key) {
		c._buf = make([]byte, 0, len(key))
	}
	value, ok = c.get(key)
	c.lock.Unlock()
	return
}

func (c *lruCache) Len() int {
	return len(c.m)
}

func (c *lruCache) HitRatio() float64 {
	return float64(c.hits) / float64(c.misses+c.hits)
}

func (c *lruCache) Info() (hits, misses int) {
	return c.hits, c.misses
}
