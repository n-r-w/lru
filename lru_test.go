package lru

import (
	"testing"
)

func prepare() *Cache[int, string] {
	c := New[int, string](3)
	c.Insert(1, "111")
	c.Insert(2, "222")
	return c
}

func TestCache_Capacity(t *testing.T) {
	c := prepare()
	if c.Capacity() != 3 {
		t.Error("TestCache_Insert Capacity")
		return
	}
}

func TestCache_Size(t *testing.T) {
	c := prepare()
	if c.Size() != 2 {
		t.Error("TestCache_Insert Size")
		return
	}
}

func TestCache_Insert_Get(t *testing.T) {
	c := prepare()

	if v, ok := c.Get(1); !ok {
		t.Error("TestCache_Insert error 1")
		return
	} else {
		if v != "111" {
			t.Error("TestCache_Insert error 2")
			return
		}
	}

	if v, ok := c.Get(3); ok {
		t.Error("TestCache_Insert error 3")
		return
	} else {
		if v != "" {
			t.Error("TestCache_Insert error 4")
			return
		}
	}

	c.Insert(1, "xxx")
	if c.Size() != 2 {
		t.Error("TestCache_Insert error 5")
		return
	}

	if v, ok := c.Get(1); !ok {
		t.Error("TestCache_Insert error 6")
		return
	} else {
		if v != "xxx" {
			t.Error("TestCache_Insert error 7")
			return
		}
	}

	if !c.Insert(3, "333") {
		t.Error("TestCache_Insert error 8")
		return
	}

	if c.Insert(4, "444") {
		t.Error("TestCache_Insert error 9")
		return
	}

	if c.Size() != 3 {
		t.Error("TestCache_Insert error 10")
		return
	}

	if c.order.Len() != len(c.data) {
		t.Error("TestCache_Insert error 1")
		return
	}

	t1 := c.Top(1)
	if len(t1) != 1 {
		t.Error("TestCache_Top error 12")
		return
	}
	if t1[0].Value != "444" {
		t.Error("TestCache_Top error 13")
		return
	}
}

func TestCache_Top(t *testing.T) {
	c := prepare()

	t1 := c.Top(1)
	if len(t1) != 1 {
		t.Error("TestCache_Top error 1")
		return
	}

	if t1[0].Value != "222" {
		t.Error("TestCache_Top error 2")
		return
	}

	t1 = c.Top(3)
	if len(t1) != 2 {
		t.Error("TestCache_Top error 3")
		return
	}

	if t1[1].Value != "111" {
		t.Error("TestCache_Top error 4")
		return
	}
}
