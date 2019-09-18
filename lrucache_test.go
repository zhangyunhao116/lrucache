package lrucache

import (
	"fmt"
	"testing"
	"time"
)

func TestLRUCache_Set_Get(t *testing.T) {
	l := New(3)

	// Simple set
	replaceKey := l.Set(1, 15)
	if replaceKey {
		t.Error("replace while Set error")
	}

	// Simple Get
	v, ok := l.Get(1)
	if v != 15 || !ok {
		t.Error("Get exists key error")
	}

	// Get non-existent key
	v, ok = l.Get(999)
	if v != nil || ok {
		t.Error("Get non-existent key error")
	}

	// Get eliminated key
	l.Set(2, 20)
	l.Set(3, 30)
	l.Set(4, 40)

	v, ok = l.Get(1)
	if v != nil || ok {
		t.Error("Get eliminated key error")
	}
	// Get sort
	// Now the key in cache is (4,3,2), left is newer
	v, ok = l.Get(2) // Cache : (2,4,3)
	if v != 20 || !ok {
		t.Error("Get exists key error")
	}

	l.Set(5, 50) // Now should be (5,2,4), root is 4
	v, ok = l.Get(3)
	if v != nil || ok {
		t.Error("Get sort error")
	}

	if l.root.value != 40 || l.root.next.value != 20 || l.root.next.next.value != 50 {
		t.Error("linked list error")
	}

	// Test all types
	testTypes := []interface{}{true, false, uint8(1), int8(1), uint16(1), int16(1), uint32(1), int32(1), uint64(1), int64(1), uint(1), int(1), []byte("111111"), "222222"}
	for i, v := range testTypes {
		l.Set(v, i)
		getValue, ok := l.Get(v)
		if !ok || getValue != i {
			t.Error("Get error")
		}
	}

}

func TestLRUCache_MSet_MGet(t *testing.T) {
	l := New(64)

	replace := l.MSet(1, 2, "666")
	if replace {
		t.Error("MGet error")
	}

	v, ok := l.MGet(1, 2)
	if v != "666" || !ok {
		t.Error("MGet error")
	}

	replace = l.MSet("1", uint8(2), int16(3), int32(4), int64(5), 6, false, float32(7), float64(8), complex64(9), complex128(10), "value")
	if replace {
		t.Error("MSet error")
	}

	v, ok = l.MGet("1", uint8(2), int16(3), int32(4), int64(5), 6, false, float32(7), float64(8), complex64(9), complex128(10))
	if v != "value" || !ok {
		t.Error("MGet error")
	}

}

func TestLRUCache_Rank(t *testing.T) {
	l := New(3)
	l.Set(1, 1)
	l.Set(2, 2)
	if !(l.root.value == nil || l.root.next.value == 1 ||
		l.root.next.next.value == 2 || l.root.next.next.next == l.root) {
		t.Error("rank error")
	}

	l.Set(3, 3)
	if !(l.root.value == 1 || l.root.next.value == 2 ||
		l.root.next.next.value == 3 || l.root.next.next.next == l.root) {
		t.Error("rank error")
	}

	l.Get(2) // Now is 1(root), 3, 2
	if !(l.root.value == 1 || l.root.next.value == 3 ||
		l.root.next.next.value == 2 || l.root.next.next.next == l.root) {
		t.Error("rank error")
	}

	l.Set(4, 4) // Now is 3(root), 2, 4
	if !(l.root.value == 3 || l.root.next.value == 2 ||
		l.root.next.next.value == 4 || l.root.next.next.next == l.root) {
		t.Error("rank error")
	}
}

func TestLRUCache_Corner(t *testing.T) {
	l := New(1)
	l.Set(1, 1)
	if !(l.root.value == 1 || l.root.next == l.root || l.root.prev == l.root) {
		t.Error("maxSize=1 error")
	}

	l.Set(1, 2)
	if !(l.root.value == 2 || l.root.next == l.root || l.root.prev == l.root) {
		t.Error("maxSize=1 error")
	}
}

func TestDataRaces(t *testing.T) {
	l := New(64)
	for i := 0; i < 50; i++ {

		if i%2 == 0 {
			go func() {
				for j := 0; j < 100; j++ {
					l.Set(j, j)
				}
			}()
		} else {
			go func() {
				for j := 0; j < 100; j++ {
					l.Get(j)
				}
			}()
		}
	}
	time.Sleep(time.Millisecond * 100)
}

func TestDataConflict(t *testing.T) {
	l := New(5)

	l.MSet(1, 2, 3, "1")
	v, ok := l.MGet(1, 2, 3)
	if !ok || v != "1" {
		t.Error("data conflict error")
	}

	l.MSet(uint(1), 2, 3, "2")
	v, ok = l.MGet(1, 2, 3)
	if !ok || v != "1" {
		t.Error("data conflict error", fmt.Sprint(v))
	}
}

func TestAppendBuffer(t *testing.T) {
	mockBuf := make([]byte, 0, 5)
	l := New(64)
	l._buf = mockBuf
	l.MSet(1, 2, 3, 4, 5, "value")
	temp := interfaceToBytes(1, 2, 3, 4, 5)
	if cap(l._buf) != len(temp) || &l._buf == &mockBuf {
		t.Error("append buffer error")
	}
}

func TestLruCache_Others(t *testing.T) {
	l := New(64)
	l.Set(1, 1)
	l.Get(1)
	l.Get(2)

	if l.Len() != 1 {
		t.Error("length error")
	}

	if i, j := l.Info(); !(i == j && i == 1) {
		t.Error("info error")
	}

	if l.HitRatio() != float64(0.5) {
		t.Error("hit ratio error")
	}

}
