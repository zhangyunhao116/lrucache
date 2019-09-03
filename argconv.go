package lrucache

import (
	"unsafe"
)

type mockIface struct {
	_    uintptr
	data unsafe.Pointer
}

func InterfaceToBytes(args ...interface{}) []byte {
	b := make([]byte, 0, len(args))

	for _, v := range args {
		switch v.(type) {
		case uint8, int8, bool:
			iface := *(*mockIface)(unsafe.Pointer(&v))
			value := *(*byte)(iface.data)
			b = append(b, value)
		case uint16, int16:
			iface := *(*mockIface)(unsafe.Pointer(&v))
			value := *(*[2]byte)(iface.data)
			b = append(b, value[0], value[1])
		case uint32, int32:
			iface := *(*mockIface)(unsafe.Pointer(&v))
			value := *(*[4]byte)(iface.data)
			b = append(b, value[0], value[1], value[2], value[3])
		case uint64, int64:
			iface := *(*mockIface)(unsafe.Pointer(&v))
			value := *(*[8]byte)(iface.data)
			b = append(b, value[0], value[1], value[2], value[3], value[4], value[5], value[6], value[7])
		case int, uint:
			iface := *(*mockIface)(unsafe.Pointer(&v))
			if bit == 64 {
				value := *(*[8]byte)(iface.data)
				b = append(b, value[0], value[1], value[2], value[3], value[4], value[5], value[6], value[7])
			} else if bit == 32 {
				value := *(*[4]byte)(iface.data)
				b = append(b, value[0], value[1], value[2], value[3])
			} else {
				panic("bit != (32 or 64)")
			}
		case string, []byte:
			iface := *(*mockIface)(unsafe.Pointer(&v))
			value := *(*string)(iface.data)
			b = append(b, value...)
		case nil: // If the arg is nil, just skip it

		default:
			panic("unknown type")
		}
	}
	return b
}

func InterfaceToBytesWithBuf(b []byte, args ...interface{}) []byte {
	for _, v := range args {
		switch v.(type) {
		case uint8, int8, bool:
			iface := *(*mockIface)(unsafe.Pointer(&v))
			value := *(*byte)(iface.data)
			b = append(b[0:], value)
		case uint16, int16:
			iface := *(*mockIface)(unsafe.Pointer(&v))
			value := *(*[2]byte)(iface.data)
			b = append(b[0:], value[0], value[1])
		case uint32, int32:
			iface := *(*mockIface)(unsafe.Pointer(&v))
			value := *(*[4]byte)(iface.data)
			b = append(b[0:], value[0], value[1], value[2], value[3])
		case uint64, int64:
			iface := *(*mockIface)(unsafe.Pointer(&v))
			value := *(*[8]byte)(iface.data)
			b = append(b[0:], value[0], value[1], value[2], value[3], value[4], value[5], value[6], value[7])
		case int, uint:
			iface := *(*mockIface)(unsafe.Pointer(&v))
			if bit == 64 {
				value := *(*[8]byte)(iface.data)
				b = append(b[0:], value[0], value[1], value[2], value[3], value[4], value[5], value[6], value[7])
			} else if bit == 32 {
				value := *(*[4]byte)(iface.data)
				b = append(b[0:], value[0], value[1], value[2], value[3])
			} else {
				panic("bit != (32 or 64)")
			}
		case string, []byte:
			iface := *(*mockIface)(unsafe.Pointer(&v))
			value := *(*string)(iface.data)
			b = append(b[0:], value...)
		case nil: // If the arg is nil, just skip it

		default:
			panic("unknown type")
		}
	}
	return b
}
