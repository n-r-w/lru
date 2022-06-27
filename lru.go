package lru

import (
	"container/list"
)

type Cache[T comparable, X any] struct {
	capacity int
	order    list.List
	data     map[T]*list.Element
}

type Item[T comparable, X any] struct {
	Key   T
	Value X
}

func New[T comparable, X any](capacity int) *Cache[T, X] {
	if capacity < 1 {
		panic("invalid capacity")
	}

	return &Cache[T, X]{
		capacity: capacity,
		order:    list.List{},
		data:     map[T]*list.Element{},
	}
}

func (c *Cache[T, X]) Size() int {
	return len(c.data)
}

func (c *Cache[T, X]) Capacity() int {
	return c.capacity
}

func (c *Cache[T, X]) SetCapacity(capacity int) bool {
	c.capacity = capacity
	return c.fixCapacity()
}

func (c *Cache[T, X]) Insert(key T, value X) bool {
	if e, ok := c.data[key]; ok {
		e.Value.(*Item[T, X]).Value = value
		c.order.MoveToFront(e)
		return true
	}

	e := c.order.PushFront(
		&Item[T, X]{
			Key:   key,
			Value: value,
		})
	c.data[key] = e

	return c.fixCapacity()
}

func (c *Cache[T, X]) fixCapacity() bool {
	dif := len(c.data) - c.capacity
	if dif <= 0 {
		return true
	}

	e := c.order.Back()
	for i := 0; i < dif; i++ {
		delete(c.data, e.Value.(*Item[T, X]).Key)
		old := e
		e = e.Prev()
		c.order.Remove(old)
	}

	return false
}

func (c *Cache[T, X]) Get(key T) (X, bool) {
	if e, ok := c.data[key]; ok {
		c.order.MoveToFront(e)
		return e.Value.(*Item[T, X]).Value, true
	}

	var res X
	return res, false
}

func (c *Cache[T, X]) Top(count int) []*Item[T, X] {
	res := []*Item[T, X]{}
	e := c.order.Front()
	for i := 0; i < count; i++ {
		res = append(res, e.Value.(*Item[T, X]))
		if e = e.Next(); e == nil {
			break
		}
	}
	return res
}
