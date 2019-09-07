package lrucache

import (
	"go/types"
	"unsafe"
)

type mockEFace struct {
	_    uintptr
	data unsafe.Pointer
}

// The type constants in go/types not includes byte slice,
// so we just create one for type hints. Since the last
// type in go/types is UntypedNil (25) in Go 1.12.7, we create
// new type from 31.
const (
	TypeByteSlice = uint8(30 + iota)
)

func interfaceToBytes(args ...interface{}) []byte {
	b := make([]byte, 0, len(args)*5)
	var data unsafe.Pointer
	for _, v := range args {
		data = (*(*mockEFace)(unsafe.Pointer(&v))).data
		switch v.(type) {
		case bool:
			b = append(b, uint8(types.Bool), *(*byte)(data))
		case uint8:
			b = append(b, uint8(types.Uint8), *(*byte)(data))
		case int8:
			b = append(b, uint8(types.Int8), *(*byte)(data))
		case uint16:
			b = append(b, uint8(types.Uint16), (*(*[2]byte)(data))[0], (*(*[2]byte)(data))[1])
		case int16:
			b = append(b, uint8(types.Int16), (*(*[2]byte)(data))[0], (*(*[2]byte)(data))[1])
		case uint32:
			b = append(b, uint8(types.Uint32), (*(*[4]byte)(data))[0], (*(*[4]byte)(data))[1],
				(*(*[4]byte)(data))[2], (*(*[4]byte)(data))[3])
		case int32:
			b = append(b, uint8(types.Int32), (*(*[4]byte)(data))[0], (*(*[4]byte)(data))[1],
				(*(*[4]byte)(data))[2], (*(*[4]byte)(data))[3])
		case float32:
			b = append(b, uint8(types.Float32), (*(*[4]byte)(data))[0], (*(*[4]byte)(data))[1],
				(*(*[4]byte)(data))[2], (*(*[4]byte)(data))[3])
		case uint64:
			b = append(b, uint8(types.Uint64), (*(*[8]byte)(data))[0], (*(*[8]byte)(data))[1],
				(*(*[8]byte)(data))[2], (*(*[8]byte)(data))[3], (*(*[8]byte)(data))[4],
				(*(*[8]byte)(data))[5], (*(*[8]byte)(data))[6], (*(*[8]byte)(data))[7])
		case int64:
			b = append(b, uint8(types.Int64), (*(*[8]byte)(data))[0], (*(*[8]byte)(data))[1],
				(*(*[8]byte)(data))[2], (*(*[8]byte)(data))[3], (*(*[8]byte)(data))[4],
				(*(*[8]byte)(data))[5], (*(*[8]byte)(data))[6], (*(*[8]byte)(data))[7])
		case float64:
			b = append(b, uint8(types.Float64), (*(*[8]byte)(data))[0], (*(*[8]byte)(data))[1],
				(*(*[8]byte)(data))[2], (*(*[8]byte)(data))[3], (*(*[8]byte)(data))[4],
				(*(*[8]byte)(data))[5], (*(*[8]byte)(data))[6], (*(*[8]byte)(data))[7])
		case complex64:
			b = append(b, uint8(types.Complex64), (*(*[8]byte)(data))[0], (*(*[8]byte)(data))[1],
				(*(*[8]byte)(data))[2], (*(*[8]byte)(data))[3], (*(*[8]byte)(data))[4],
				(*(*[8]byte)(data))[5], (*(*[8]byte)(data))[6], (*(*[8]byte)(data))[7])
		case complex128:
			b = append(b, uint8(types.Complex128), (*(*[16]byte)(data))[0], (*(*[16]byte)(data))[1],
				(*(*[16]byte)(data))[2], (*(*[16]byte)(data))[3], (*(*[16]byte)(data))[4],
				(*(*[16]byte)(data))[5], (*(*[16]byte)(data))[6], (*(*[16]byte)(data))[7],
				(*(*[16]byte)(data))[8], (*(*[16]byte)(data))[9], (*(*[16]byte)(data))[10],
				(*(*[16]byte)(data))[11], (*(*[16]byte)(data))[12], (*(*[16]byte)(data))[13],
				(*(*[16]byte)(data))[14], (*(*[16]byte)(data))[15])
		case int:
			if bit == 64 {
				b = append(b, uint8(types.Int), (*(*[8]byte)(data))[0], (*(*[8]byte)(data))[1],
					(*(*[8]byte)(data))[2], (*(*[8]byte)(data))[3], (*(*[8]byte)(data))[4],
					(*(*[8]byte)(data))[5], (*(*[8]byte)(data))[6], (*(*[8]byte)(data))[7])
			} else if bit == 32 {
				b = append(b, uint8(types.Int), (*(*[4]byte)(data))[0], (*(*[4]byte)(data))[1],
					(*(*[4]byte)(data))[2], (*(*[4]byte)(data))[3])
			} else {
				panic("bit != (32 or 64)")
			}
		case uint:
			if bit == 64 {
				b = append(b, uint8(types.Uint), (*(*[8]byte)(data))[0], (*(*[8]byte)(data))[1],
					(*(*[8]byte)(data))[2], (*(*[8]byte)(data))[3], (*(*[8]byte)(data))[4],
					(*(*[8]byte)(data))[5], (*(*[8]byte)(data))[6], (*(*[8]byte)(data))[7])
			} else if bit == 32 {
				b = append(b, uint8(types.Uint), (*(*[4]byte)(data))[0], (*(*[4]byte)(data))[1],
					(*(*[4]byte)(data))[2], (*(*[4]byte)(data))[3])
			} else {
				panic("bit != (32 or 64)")
			}
		case string:
			// In this case, we insert a int to indicates how many
			// bytes occupied by this string or []byte to avoid potential
			// data conflict.
			value := *(*string)(data)
			bLen := len(value)
			if bit == 64 {
				b = append(b, uint8(types.String), (*(*[8]byte)(unsafe.Pointer(&bLen)))[0],
					(*(*[8]byte)(unsafe.Pointer(&bLen)))[1], (*(*[8]byte)(unsafe.Pointer(&bLen)))[2],
					(*(*[8]byte)(unsafe.Pointer(&bLen)))[3], (*(*[8]byte)(unsafe.Pointer(&bLen)))[4],
					(*(*[8]byte)(unsafe.Pointer(&bLen)))[5], (*(*[8]byte)(unsafe.Pointer(&bLen)))[6],
					(*(*[8]byte)(unsafe.Pointer(&bLen)))[7])
			} else if bit == 32 {
				b = append(b, uint8(types.String), (*(*[4]byte)(unsafe.Pointer(&bLen)))[0],
					(*(*[4]byte)(unsafe.Pointer(&bLen)))[1], (*(*[4]byte)(unsafe.Pointer(&bLen)))[2],
					(*(*[4]byte)(unsafe.Pointer(&bLen)))[3])
			} else {
				panic("bit != (32 or 64)")
			}

			b = append(b, value...)
		case []byte:
			// In this case, we insert a int to indicates how many
			// bytes occupied by this string or []byte to avoid potential
			// data conflict.
			value := *(*string)(data)
			bLen := len(value)
			if bit == 64 {
				b = append(b, TypeByteSlice, (*(*[8]byte)(unsafe.Pointer(&bLen)))[0],
					(*(*[8]byte)(unsafe.Pointer(&bLen)))[1], (*(*[8]byte)(unsafe.Pointer(&bLen)))[2],
					(*(*[8]byte)(unsafe.Pointer(&bLen)))[3], (*(*[8]byte)(unsafe.Pointer(&bLen)))[4],
					(*(*[8]byte)(unsafe.Pointer(&bLen)))[5], (*(*[8]byte)(unsafe.Pointer(&bLen)))[6],
					(*(*[8]byte)(unsafe.Pointer(&bLen)))[7])
			} else if bit == 32 {
				b = append(b, TypeByteSlice, (*(*[4]byte)(unsafe.Pointer(&bLen)))[0],
					(*(*[4]byte)(unsafe.Pointer(&bLen)))[1], (*(*[4]byte)(unsafe.Pointer(&bLen)))[2],
					(*(*[4]byte)(unsafe.Pointer(&bLen)))[3])
			} else {
				panic("bit != (32 or 64)")
			}

			b = append(b, value...)
		default:
			panic("unknown type")
		}
	}
	return b
}

func interfaceToBytesWithBuf(b []byte, args ...interface{}) []byte {
	var data unsafe.Pointer
	for _, v := range args {
		data = (*(*mockEFace)(unsafe.Pointer(&v))).data
		switch v.(type) {
		case bool:
			b = append(b, uint8(types.Bool), *(*byte)(data))
		case uint8:
			b = append(b, uint8(types.Uint8), *(*byte)(data))
		case int8:
			b = append(b, uint8(types.Int8), *(*byte)(data))
		case uint16:
			b = append(b, uint8(types.Uint16), (*(*[2]byte)(data))[0], (*(*[2]byte)(data))[1])
		case int16:
			b = append(b, uint8(types.Int16), (*(*[2]byte)(data))[0], (*(*[2]byte)(data))[1])
		case uint32:
			b = append(b, uint8(types.Uint32), (*(*[4]byte)(data))[0], (*(*[4]byte)(data))[1],
				(*(*[4]byte)(data))[2], (*(*[4]byte)(data))[3])
		case int32:
			b = append(b, uint8(types.Int32), (*(*[4]byte)(data))[0], (*(*[4]byte)(data))[1],
				(*(*[4]byte)(data))[2], (*(*[4]byte)(data))[3])
		case float32:
			b = append(b, uint8(types.Float32), (*(*[4]byte)(data))[0], (*(*[4]byte)(data))[1],
				(*(*[4]byte)(data))[2], (*(*[4]byte)(data))[3])
		case uint64:
			b = append(b, uint8(types.Uint64), (*(*[8]byte)(data))[0], (*(*[8]byte)(data))[1],
				(*(*[8]byte)(data))[2], (*(*[8]byte)(data))[3], (*(*[8]byte)(data))[4],
				(*(*[8]byte)(data))[5], (*(*[8]byte)(data))[6], (*(*[8]byte)(data))[7])
		case int64:
			b = append(b, uint8(types.Int64), (*(*[8]byte)(data))[0], (*(*[8]byte)(data))[1],
				(*(*[8]byte)(data))[2], (*(*[8]byte)(data))[3], (*(*[8]byte)(data))[4],
				(*(*[8]byte)(data))[5], (*(*[8]byte)(data))[6], (*(*[8]byte)(data))[7])
		case float64:
			b = append(b, uint8(types.Float64), (*(*[8]byte)(data))[0], (*(*[8]byte)(data))[1],
				(*(*[8]byte)(data))[2], (*(*[8]byte)(data))[3], (*(*[8]byte)(data))[4],
				(*(*[8]byte)(data))[5], (*(*[8]byte)(data))[6], (*(*[8]byte)(data))[7])
		case complex64:
			b = append(b, uint8(types.Complex64), (*(*[8]byte)(data))[0], (*(*[8]byte)(data))[1],
				(*(*[8]byte)(data))[2], (*(*[8]byte)(data))[3], (*(*[8]byte)(data))[4],
				(*(*[8]byte)(data))[5], (*(*[8]byte)(data))[6], (*(*[8]byte)(data))[7])
		case complex128:
			b = append(b, uint8(types.Complex128), (*(*[16]byte)(data))[0], (*(*[16]byte)(data))[1],
				(*(*[16]byte)(data))[2], (*(*[16]byte)(data))[3], (*(*[16]byte)(data))[4],
				(*(*[16]byte)(data))[5], (*(*[16]byte)(data))[6], (*(*[16]byte)(data))[7],
				(*(*[16]byte)(data))[8], (*(*[16]byte)(data))[9], (*(*[16]byte)(data))[10],
				(*(*[16]byte)(data))[11], (*(*[16]byte)(data))[12], (*(*[16]byte)(data))[13],
				(*(*[16]byte)(data))[14], (*(*[16]byte)(data))[15])
		case int:
			if bit == 64 {
				b = append(b, uint8(types.Int), (*(*[8]byte)(data))[0], (*(*[8]byte)(data))[1],
					(*(*[8]byte)(data))[2], (*(*[8]byte)(data))[3], (*(*[8]byte)(data))[4],
					(*(*[8]byte)(data))[5], (*(*[8]byte)(data))[6], (*(*[8]byte)(data))[7])
			} else if bit == 32 {
				b = append(b, uint8(types.Int), (*(*[4]byte)(data))[0], (*(*[4]byte)(data))[1],
					(*(*[4]byte)(data))[2], (*(*[4]byte)(data))[3])
			} else {
				panic("bit != (32 or 64)")
			}
		case uint:
			if bit == 64 {
				b = append(b, uint8(types.Uint), (*(*[8]byte)(data))[0], (*(*[8]byte)(data))[1],
					(*(*[8]byte)(data))[2], (*(*[8]byte)(data))[3], (*(*[8]byte)(data))[4],
					(*(*[8]byte)(data))[5], (*(*[8]byte)(data))[6], (*(*[8]byte)(data))[7])
			} else if bit == 32 {
				b = append(b, uint8(types.Uint), (*(*[4]byte)(data))[0], (*(*[4]byte)(data))[1],
					(*(*[4]byte)(data))[2], (*(*[4]byte)(data))[3])
			} else {
				panic("bit != (32 or 64)")
			}
		case string:
			// In this case, we insert a int to indicates how many
			// bytes occupied by this string or []byte to avoid potential
			// data conflict.
			value := *(*string)(data)
			bLen := len(value)
			if bit == 64 {
				b = append(b, uint8(types.String), (*(*[8]byte)(unsafe.Pointer(&bLen)))[0],
					(*(*[8]byte)(unsafe.Pointer(&bLen)))[1], (*(*[8]byte)(unsafe.Pointer(&bLen)))[2],
					(*(*[8]byte)(unsafe.Pointer(&bLen)))[3], (*(*[8]byte)(unsafe.Pointer(&bLen)))[4],
					(*(*[8]byte)(unsafe.Pointer(&bLen)))[5], (*(*[8]byte)(unsafe.Pointer(&bLen)))[6],
					(*(*[8]byte)(unsafe.Pointer(&bLen)))[7])
			} else if bit == 32 {
				b = append(b, uint8(types.String), (*(*[4]byte)(unsafe.Pointer(&bLen)))[0],
					(*(*[4]byte)(unsafe.Pointer(&bLen)))[1], (*(*[4]byte)(unsafe.Pointer(&bLen)))[2],
					(*(*[4]byte)(unsafe.Pointer(&bLen)))[3])
			} else {
				panic("bit != (32 or 64)")
			}

			b = append(b, value...)
		case []byte:
			// In this case, we insert a int to indicates how many
			// bytes occupied by this string or []byte to avoid potential
			// data conflict.
			value := *(*string)(data)
			bLen := len(value)
			if bit == 64 {
				b = append(b, TypeByteSlice, (*(*[8]byte)(unsafe.Pointer(&bLen)))[0],
					(*(*[8]byte)(unsafe.Pointer(&bLen)))[1], (*(*[8]byte)(unsafe.Pointer(&bLen)))[2],
					(*(*[8]byte)(unsafe.Pointer(&bLen)))[3], (*(*[8]byte)(unsafe.Pointer(&bLen)))[4],
					(*(*[8]byte)(unsafe.Pointer(&bLen)))[5], (*(*[8]byte)(unsafe.Pointer(&bLen)))[6],
					(*(*[8]byte)(unsafe.Pointer(&bLen)))[7])
			} else if bit == 32 {
				b = append(b, TypeByteSlice, (*(*[4]byte)(unsafe.Pointer(&bLen)))[0],
					(*(*[4]byte)(unsafe.Pointer(&bLen)))[1], (*(*[4]byte)(unsafe.Pointer(&bLen)))[2],
					(*(*[4]byte)(unsafe.Pointer(&bLen)))[3])
			} else {
				panic("bit != (32 or 64)")
			}

			b = append(b, value...)
		default:
			panic("unknown type")
		}
	}
	return b
}
