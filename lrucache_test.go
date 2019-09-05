package lrucache

import (
	lru "github.com/hashicorp/golang-lru"
	"math/rand"
	"testing"
)

func BenchmarkAll_Rand(b *testing.B) {
	l := New(8192)

	trace := make([]int64, b.N*2)
	for i := 0; i < b.N*2; i++ {
		trace[i] = rand.Int63() % 32768
	}

	b.ResetTimer()

	var hit, miss int
	for i := 0; i < 2*b.N; i++ {
		if i%2 == 0 {
			l.Set(trace[i], trace[i])
		} else {
			_, ok := l.Get(trace[i])
			if ok {
				hit++
			} else {
				miss++
			}
		}
	}
	b.Logf("hit: %d miss: %d ratio: %f", hit, miss, float64(hit)/float64(miss))
}

func BenchmarkAll_Rand_extra(b *testing.B) {
	l, err := lru.New(8192)
	if err != nil {
		b.Fatalf("err: %v", err)
	}

	trace := make([]int64, b.N*2)
	for i := 0; i < b.N*2; i++ {
		trace[i] = rand.Int63() % 32768
	}

	b.ResetTimer()

	var hit, miss int
	for i := 0; i < 2*b.N; i++ {
		if i%2 == 0 {
			l.Add(trace[i], trace[i])
		} else {
			_, ok := l.Get(trace[i])
			if ok {
				hit++
			} else {
				miss++
			}
		}
	}
	b.Logf("hit: %d miss: %d ratio: %f", hit, miss, float64(hit)/float64(miss))
}

func BenchmarkAll_Freq(b *testing.B) {
	l := New(8192)

	trace := make([]int64, b.N*2)
	for i := 0; i < b.N*2; i++ {
		if i%2 == 0 {
			trace[i] = rand.Int63() % 16384
		} else {
			trace[i] = rand.Int63() % 32768
		}
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		l.Set(trace[i], trace[i])
	}
	var hit, miss int
	for i := 0; i < b.N; i++ {
		_, ok := l.Get(trace[i])
		if ok {
			hit++
		} else {
			miss++
		}
	}
	b.Logf("hit: %d miss: %d ratio: %f", hit, miss, float64(hit)/float64(miss))
}

func BenchmarkAll_Freq_extra(b *testing.B) {
	l, err := lru.New(8192)
	if err != nil {
		b.Fatalf("err: %v", err)
	}

	trace := make([]int64, b.N*2)
	for i := 0; i < b.N*2; i++ {
		if i%2 == 0 {
			trace[i] = rand.Int63() % 16384
		} else {
			trace[i] = rand.Int63() % 32768
		}
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		l.Add(trace[i], trace[i])
	}
	var hit, miss int
	for i := 0; i < b.N; i++ {
		_, ok := l.Get(trace[i])
		if ok {
			hit++
		} else {
			miss++
		}
	}
	b.Logf("hit: %d miss: %d ratio: %f", hit, miss, float64(hit)/float64(miss))
}

func TestLRUCache_Set_Get(t *testing.T) {
	l := New(3)

	// Simple set
	replaceKey := l.Set(1, 15)
	if replaceKey {
		t.Error("error replace while set")
	}

	// Simple Get
	v, ok := l.Get(1)
	if v != 15 || !ok {
		t.Error("error get exists key")
	}

	// Get non-existent key
	v, ok = l.Get(999)
	if v != nil || ok {
		t.Error("error get non-existent key")
	}

	// Get eliminated key
	l.Set(2, 20)
	l.Set(3, 30)
	l.Set(4, 40)

	v, ok = l.Get(1)
	if v != nil || ok {
		t.Error("error get eliminated key")
	}
	// Get sort
	// Now the key in cache is (4,3,2), left is newer
	v, ok = l.Get(2) // Cache : (2,4,3)
	if v != 20 || !ok {
		t.Error("error get exists key")
	}

	l.Set(5, 50) // Now should be (5,2,4), root is 4
	v, ok = l.Get(3)
	if v != nil || ok {
		t.Error("error get sort")
	}

	if l.root.value != 40 || l.root.next.value != 20 || l.root.next.next.value != 50 {
		t.Error("error linked list")
	}

	// Test all types
	testTypes := []interface{}{true, false, uint8(1), int8(1), uint16(1), int16(1), uint32(1), int32(1), uint64(1), int64(1), uint(1), int(1), []byte("111111"), "222222"}
	for i, v := range testTypes {
		l.Set(v, i)
		getValue, ok := l.Get(v)
		if !ok || getValue != i {
			t.Error("error Get")
		}
	}

}

func TestLRUCache_MSet_MGet(t *testing.T) {
	l := New(64)

	replace := l.MSet(1, 2, "666")
	if replace {
		t.Error("error Mset")
	}

	v, ok := l.MGet(1, 2)
	if v != "666" || !ok {
		t.Error("error mget")
	}

	replace = l.MSet("1", uint8(2), int16(3), int32(4), int64(5), 6, false, float32(7), float64(8), complex64(9), complex128(10), "value")
	if replace {
		t.Error("error Mset")
	}

	v, ok = l.MGet("1", uint8(2), int16(3), int32(4), int64(5), 6, false, float32(7), float64(8), complex64(9), complex128(10))
	if v != "value" || !ok {
		t.Error("error mget")
	}

}
