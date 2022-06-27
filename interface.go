package lru

/*
I_Cache общий интерфейс для lru.Cache и lru.SyncCache
метод Top сюда не включен, т.к. он не может быть потокобезопасным и работать для обеих реализаций интерфейса
*/
type I_Cache[T comparable, X any] interface {
	Size() int                     // текущий размер кэша
	Capacity() int                 // максимальный размер кэша, после достижения которого, последние элементы начнут удаляться
	SetCapacity(capacity int) bool // задать максимальный размер кэша. Лишние элементы будут удалены
	Insert(key T, value X) bool    // добавить элемент
	Get(key T) (X, bool)           // получить элемент
}
