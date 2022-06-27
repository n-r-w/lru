package lru

import "sync"

type SyncCache[T comparable, X any] struct {
	cache *Cache[T, X]
	mutex sync.RWMutex
}

func NewSync[T comparable, X any](capacity int) *SyncCache[T, X] {
	return &SyncCache[T, X]{
		cache: New[T, X](capacity),
		mutex: sync.RWMutex{},
	}
}

func (c *SyncCache[T, X]) Size() int {
	c.mutex.RLock()
	defer c.mutex.Unlock()
	return c.cache.Size()
}

func (c *SyncCache[T, X]) Capacity() int {
	c.mutex.RLock()
	defer c.mutex.Unlock()
	return c.cache.Capacity()
}

func (c *SyncCache[T, X]) SetCapacity(capacity int) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.cache.SetCapacity(capacity)
}

func (c *SyncCache[T, X]) Insert(key T, value X) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.cache.Insert(key, value)
}

func (c *SyncCache[T, X]) Get(key T) (X, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.cache.Get(key)
}

func (c *SyncCache[T, X]) Top(count int) []*Item[T, X] {
	c.mutex.RLock()
	defer c.mutex.Unlock()
	return c.cache.Top(count)
}
