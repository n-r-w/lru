package lru

import (
	"container/list"
)

// Cache - LRU кэш, не потокобезопасен
type Cache[T comparable, X any] struct {
	capacity int
	order    list.List
	data     map[T]*list.Element
}

// Item - элемент LRU кэша
type Item[T comparable, X any] struct {
	Key   T
	Value X
}

// New - создать кэш
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

// Size - текущий размер кэша
func (c *Cache[T, X]) Size() int {
	return len(c.data)
}

// Size - максимальный размер кэша, после достижения которого, последние элементы начнут удаляться
func (c *Cache[T, X]) Capacity() int {
	return c.capacity
}

// SetCapacity - задать максимальный размер кэша. Лишние элементы будут удалены
func (c *Cache[T, X]) SetCapacity(capacity int) bool {
	c.capacity = capacity
	return c.fixCapacity()
}

// Insert - добавить элемент
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

// Get - получить элемент
func (c *Cache[T, X]) Get(key T) (X, bool) {
	if e, ok := c.data[key]; ok {
		c.order.MoveToFront(e)
		return e.Value.(*Item[T, X]).Value, true
	}

	var res X
	return res, false
}

// Top - элементы верхнего уровня
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
