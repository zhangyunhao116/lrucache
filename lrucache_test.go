package lrucache

import (
	lru "github.com/hashicorp/golang-lru"
	"math/rand"
	"testing"
)

func BenchmarkLRU_Rand(b *testing.B) {
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

func BenchmarkLRU_Rand_extra(b *testing.B) {
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

func BenchmarkLRU_Freq(b *testing.B) {
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

func BenchmarkLRU_Freq_extra(b *testing.B) {
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

func TestLRUCache(t *testing.T) {
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

}

func TestLRUCache_MSet(t *testing.T) {
	l := New(64)

	replace := l.MSet(1, 2, 3, "666")
	if replace {
		t.Error("error Mset")
	}

	v, ok := l.MGet(1, 2, 3)
	if v != "666" || !ok {
		t.Error("error mget")
	}

}
