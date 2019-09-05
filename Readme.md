# lrucache
[![Go Report Card](https://goreportcard.com/badge/github.com/ZYunH/lrucache)](https://goreportcard.com/report/github.com/ZYunH/lrucache)

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



### Tips

- **Multi-keys** : Keep in mind that byte slice or string is better to have only one, this means the key-arguments only actually includes a string or a byte slice, since our strategy is just map interface{} to some bytes, potential data races can be occur if string or byte slice more than one. If you insist on doing so, don't pass binary data as string or byte slice, it can increase the risk of data races. Keep string or byte slice as printable is a good idea to avoid potential data races.

