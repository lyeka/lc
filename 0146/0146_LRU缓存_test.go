package lruCache

import (
	"testing"
)

func TestLRU(t *testing.T) {
	cache := Constructor(2)
	if cache.Get(2) != -1 {
		t.Error("expect = -1")
	}
	cache.Put(1, 1)
	cache.Put(2, 2)
	if cache.Get(1) != 1 {
		t.Error("expect = 1")
	}
	cache.Put(3, 3)
	if cache.Get(2) != -1 {
		t.Error("expect = -1")
	}

	cache.Put(4, 4)
	if cache.Get(1) != -1 {
		t.Error("expect = -1")
	}
	if cache.Get(3) != 3 {
		t.Error("expect = -1")
	}

	if cache.Get(4) != 4 {
		t.Error("expect = 4")
	}
}
