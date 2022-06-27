package lru

type I_Cache[T comparable, X any] interface {
	Size() int
	Capacity() int
	SetCapacity(capacity int) bool
	Insert(key T, value X) bool
	Get(key T) (X, bool)
	Top(count int) []*Item[T, X]
}
