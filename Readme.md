# lrucache

High-performance LRU cache implementation with multi-keys supported for Go



## Features

- Support both single key and multi-keys
- Concurrent-safe API
- Cache statistics



## Usage

##### Single key and single value

```go
l := lrucache.New(64)
l.Set(1, 2)
v, ok := l.Get(1)
if ok {
	print(fmt.Sprint(v)) // 2
}

```

##### Multi-keys and single value

```go
l := lrucache.New(64)
l.MSet(1, 2, 3, "Value")
v, ok := l.MGet(1, 2, 3)
if ok {
	print(fmt.Sprint(v)) // Value
}
```

##### Other informations

```go
l := lrucache.New(64)
l.MSet(1, 2, "Foo", "Value")
l.MGet(1, 2, "Foo")
l.MGet(2, 2, "Foo")

Len := l.Len()
print("len:", Len, "\r\n") // len: 1

hits, misses := l.Info()
print("hits:", hits, " misses:", misses, "\r\n") // hits: 1 misses:1

hitRatio := l.HitRatio()
print("hitRatio:", fmt.Sprint(hitRatio), "\r\n") // hitRatio: 0.5
```

