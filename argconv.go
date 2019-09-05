package lrucache

import (
	"go/types"
	"unsafe"
)

type mockIface struct {
	_    uintptr
	data unsafe.Pointer
}

func InterfaceToBytes(args ...interface{}) []byte {
	b := make([]byte, 0, len(args)*5)
	for _, v := range args {
		switch v.(type) {
		case bool:
			iface := *(*mockIface)(unsafe.Pointer(&v))
			value := *(*byte)(iface.data)
			b = append(b, uint8(types.Bool), value)
		case uint8:
			iface := *(*mockIface)(unsafe.Pointer(&v))
			value := *(*byte)(iface.data)
			b = append(b, uint8(types.Uint8), value)
		case int8:
			iface := *(*mockIface)(unsafe.Pointer(&v))
			value := *(*byte)(iface.data)
			b = append(b, uint8(types.Int8), value)
		case uint16:
			iface := *(*mockIface)(unsafe.Pointer(&v))
			value := *(*[2]byte)(iface.data)
			b = append(b, uint8(types.Uint16), value[0], value[1])
		case int16:
			iface := *(*mockIface)(unsafe.Pointer(&v))
			value := *(*[2]byte)(iface.data)
			b = append(b, uint8(types.Int16), value[0], value[1])
		case uint32:
			iface := *(*mockIface)(unsafe.Pointer(&v))
			value := *(*[4]byte)(iface.data)
			b = append(b, uint8(types.Uint32), value[0], value[1], value[2], value[3])
		case int32:
			iface := *(*mockIface)(unsafe.Pointer(&v))
			value := *(*[4]byte)(iface.data)
			b = append(b, uint8(types.Int32), value[0], value[1], value[2], value[3])
		case float32:
			iface := *(*mockIface)(unsafe.Pointer(&v))
			value := *(*[4]byte)(iface.data)
			b = append(b, uint8(types.Float32), value[0], value[1], value[2], value[3])
		case uint64:
			iface := *(*mockIface)(unsafe.Pointer(&v))
			value := *(*[8]byte)(iface.data)
			b = append(b, uint8(types.Uint64), value[0], value[1], value[2], value[3], value[4], value[5], value[6], value[7])
		case int64:
			iface := *(*mockIface)(unsafe.Pointer(&v))
			value := *(*[8]byte)(iface.data)
			b = append(b, uint8(types.Int64), value[0], value[1], value[2], value[3], value[4], value[5], value[6], value[7])
		case float64:
			iface := *(*mockIface)(unsafe.Pointer(&v))
			value := *(*[8]byte)(iface.data)
			b = append(b, uint8(types.Float64), value[0], value[1], value[2], value[3], value[4], value[5], value[6], value[7])
		case complex64:
			iface := *(*mockIface)(unsafe.Pointer(&v))
			value := *(*[8]byte)(iface.data)
			b = append(b, uint8(types.Complex64), value[0], value[1], value[2], value[3], value[4], value[5], value[6], value[7])
		case complex128:
			iface := *(*mockIface)(unsafe.Pointer(&v))
			value := *(*[16]byte)(iface.data)
			b = append(b, uint8(types.Complex128), value[0], value[1], value[2], value[3], value[4], value[5], value[6], value[7], value[8], value[9], value[10], value[11], value[12], value[13], value[14], value[15])
		case int:
			iface := *(*mockIface)(unsafe.Pointer(&v))
			if bit == 64 {
				value := *(*[8]byte)(iface.data)
				b = append(b, uint8(types.Int), value[0], value[1], value[2], value[3], value[4], value[5], value[6], value[7])
			} else if bit == 32 {
				value := *(*[4]byte)(iface.data)
				b = append(b, uint8(types.Int), value[0], value[1], value[2], value[3])
			} else {
				panic("bit != (32 or 64)")
			}
		case uint:
			iface := *(*mockIface)(unsafe.Pointer(&v))
			if bit == 64 {
				value := *(*[8]byte)(iface.data)
				b = append(b, uint8(types.Uint), value[0], value[1], value[2], value[3], value[4], value[5], value[6], value[7])
			} else if bit == 32 {
				value := *(*[4]byte)(iface.data)
				b = append(b, uint8(types.Uint), value[0], value[1], value[2], value[3])
			} else {
				panic("bit != (32 or 64)")
			}
		case string:
			// In this case, we insert a int to indicates how many
			// bytes occupied by this string or []byte to avoid potential
			// data conflict.
			iface := *(*mockIface)(unsafe.Pointer(&v))
			value := *(*string)(iface.data)
			if bit == 64 {
				bLen := len(value)
				bLenBytes := *(*[8]byte)(unsafe.Pointer(&bLen))
				b = append(b, uint8(types.String), bLenBytes[0], bLenBytes[1], bLenBytes[2], bLenBytes[3], bLenBytes[4], bLenBytes[5], bLenBytes[6], bLenBytes[7])
			} else if bit == 32 {
				bLen := len(value)
				bLenBytes := *(*[4]byte)(unsafe.Pointer(&bLen))
				b = append(b, uint8(types.String), bLenBytes[0], bLenBytes[1], bLenBytes[2], bLenBytes[3])
			} else {
				panic("bit != (32 or 64)")
			}

			b = append(b, value...)
		case []byte:
			// In this case, we insert a int to indicates how many
			// bytes occupied by this string or []byte to avoid potential
			// data conflict.
			iface := *(*mockIface)(unsafe.Pointer(&v))
			value := *(*string)(iface.data)
			if bit == 64 {
				bLen := len(value)
				bLenBytes := *(*[8]byte)(unsafe.Pointer(&bLen))
				b = append(b, uint8(types.UnsafePointer), bLenBytes[0], bLenBytes[1], bLenBytes[2], bLenBytes[3], bLenBytes[4], bLenBytes[5], bLenBytes[6], bLenBytes[7])
			} else if bit == 32 {
				bLen := len(value)
				bLenBytes := *(*[4]byte)(unsafe.Pointer(&bLen))
				b = append(b, uint8(types.UnsafePointer), bLenBytes[0], bLenBytes[1], bLenBytes[2], bLenBytes[3])
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

func InterfaceToBytesWithBuf(b []byte, args ...interface{}) []byte {
	b = b[0:] // Reset buffer
	for _, v := range args {
		switch v.(type) {
		case bool:
			iface := *(*mockIface)(unsafe.Pointer(&v))
			value := *(*byte)(iface.data)
			b = append(b, uint8(types.Bool), value)
		case uint8:
			iface := *(*mockIface)(unsafe.Pointer(&v))
			value := *(*byte)(iface.data)
			b = append(b, uint8(types.Uint8), value)
		case int8:
			iface := *(*mockIface)(unsafe.Pointer(&v))
			value := *(*byte)(iface.data)
			b = append(b, uint8(types.Int8), value)
		case uint16:
			iface := *(*mockIface)(unsafe.Pointer(&v))
			value := *(*[2]byte)(iface.data)
			b = append(b, uint8(types.Uint16), value[0], value[1])
		case int16:
			iface := *(*mockIface)(unsafe.Pointer(&v))
			value := *(*[2]byte)(iface.data)
			b = append(b, uint8(types.Int16), value[0], value[1])
		case uint32:
			iface := *(*mockIface)(unsafe.Pointer(&v))
			value := *(*[4]byte)(iface.data)
			b = append(b, uint8(types.Uint32), value[0], value[1], value[2], value[3])
		case int32:
			iface := *(*mockIface)(unsafe.Pointer(&v))
			value := *(*[4]byte)(iface.data)
			b = append(b, uint8(types.Int32), value[0], value[1], value[2], value[3])
		case float32:
			iface := *(*mockIface)(unsafe.Pointer(&v))
			value := *(*[4]byte)(iface.data)
			b = append(b, uint8(types.Float32), value[0], value[1], value[2], value[3])
		case uint64:
			iface := *(*mockIface)(unsafe.Pointer(&v))
			value := *(*[8]byte)(iface.data)
			b = append(b, uint8(types.Uint64), value[0], value[1], value[2], value[3], value[4], value[5], value[6], value[7])
		case int64:
			iface := *(*mockIface)(unsafe.Pointer(&v))
			value := *(*[8]byte)(iface.data)
			b = append(b, uint8(types.Int64), value[0], value[1], value[2], value[3], value[4], value[5], value[6], value[7])
		case float64:
			iface := *(*mockIface)(unsafe.Pointer(&v))
			value := *(*[8]byte)(iface.data)
			b = append(b, uint8(types.Float64), value[0], value[1], value[2], value[3], value[4], value[5], value[6], value[7])
		case complex64:
			iface := *(*mockIface)(unsafe.Pointer(&v))
			value := *(*[8]byte)(iface.data)
			b = append(b, uint8(types.Complex64), value[0], value[1], value[2], value[3], value[4], value[5], value[6], value[7])
		case complex128:
			iface := *(*mockIface)(unsafe.Pointer(&v))
			value := *(*[16]byte)(iface.data)
			b = append(b, uint8(types.Complex128), value[0], value[1], value[2], value[3], value[4], value[5], value[6], value[7], value[8], value[9], value[10], value[11], value[12], value[13], value[14], value[15])
		case int:
			iface := *(*mockIface)(unsafe.Pointer(&v))
			if bit == 64 {
				value := *(*[8]byte)(iface.data)
				b = append(b, uint8(types.Int), value[0], value[1], value[2], value[3], value[4], value[5], value[6], value[7])
			} else if bit == 32 {
				value := *(*[4]byte)(iface.data)
				b = append(b, uint8(types.Int), value[0], value[1], value[2], value[3])
			} else {
				panic("bit != (32 or 64)")
			}
		case uint:
			iface := *(*mockIface)(unsafe.Pointer(&v))
			if bit == 64 {
				value := *(*[8]byte)(iface.data)
				b = append(b, uint8(types.Uint), value[0], value[1], value[2], value[3], value[4], value[5], value[6], value[7])
			} else if bit == 32 {
				value := *(*[4]byte)(iface.data)
				b = append(b, uint8(types.Uint), value[0], value[1], value[2], value[3])
			} else {
				panic("bit != (32 or 64)")
			}
		case string:
			// In this case, we insert a int to indicates how many
			// bytes occupied by this string or []byte to avoid potential
			// data conflict.
			iface := *(*mockIface)(unsafe.Pointer(&v))
			value := *(*string)(iface.data)
			if bit == 64 {
				bLen := len(value)
				bLenBytes := *(*[8]byte)(unsafe.Pointer(&bLen))
				b = append(b, uint8(types.String), bLenBytes[0], bLenBytes[1], bLenBytes[2], bLenBytes[3], bLenBytes[4], bLenBytes[5], bLenBytes[6], bLenBytes[7])
			} else if bit == 32 {
				bLen := len(value)
				bLenBytes := *(*[4]byte)(unsafe.Pointer(&bLen))
				b = append(b, uint8(types.String), bLenBytes[0], bLenBytes[1], bLenBytes[2], bLenBytes[3])
			} else {
				panic("bit != (32 or 64)")
			}

			b = append(b, value...)
		case []byte:
			// In this case, we insert a int to indicates how many
			// bytes occupied by this string or []byte to avoid potential
			// data conflict.
			iface := *(*mockIface)(unsafe.Pointer(&v))
			value := *(*string)(iface.data)
			if bit == 64 {
				bLen := len(value)
				bLenBytes := *(*[8]byte)(unsafe.Pointer(&bLen))
				b = append(b, uint8(types.UnsafePointer), bLenBytes[0], bLenBytes[1], bLenBytes[2], bLenBytes[3], bLenBytes[4], bLenBytes[5], bLenBytes[6], bLenBytes[7])
			} else if bit == 32 {
				bLen := len(value)
				bLenBytes := *(*[4]byte)(unsafe.Pointer(&bLen))
				b = append(b, uint8(types.UnsafePointer), bLenBytes[0], bLenBytes[1], bLenBytes[2], bLenBytes[3])
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
