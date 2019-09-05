package lrucache

import "testing"

func BenchmarkInterfaceToBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		interfaceToBytes(i)
	}
}

func BenchmarkInterfaceToBytesWithBuf(b *testing.B) {
	buf := make([]byte, 0, 128)
	for i := 0; i < b.N; i++ {
		interfaceToBytesWithBuf(buf, i)
	}
}

func BenchmarkInterfaceToBytesWithBufTypes(b *testing.B) {
	buf := make([]byte, 0, 128)
	for i := 0; i < b.N; i++ {
		interfaceToBytesWithBuf(buf, i, i+1, "testString")
	}
}

func TestInterfaceToBytes(t *testing.T) {
	x := interfaceToBytes(false, true)
	if len(x) != 4 {
		t.Error("bool error")
	}

	x = interfaceToBytes(uint8(5), int8(6))
	if len(x) != 4 {
		t.Error("uint8 int8 error")
	}

	x = interfaceToBytes(uint16(5), int16(6))
	if len(x) != 6 {
		t.Error("uint16 int16 error")
	}

	x = interfaceToBytes(uint32(5), int32(6))
	if len(x) != 10 {
		t.Error("uint32 int32 error")
	}

	x = interfaceToBytes(uint64(5), int64(6))
	if len(x) != 18 {
		t.Error("uint64 int64 error")
	}

	x = interfaceToBytes(int(5), int(6))
	if bit == 64 && len(x) != 18 {
		t.Error("uint int error")
	} else if bit == 32 && len(x) != 10 {
		t.Error("uint int error")
	}

	x = interfaceToBytes([]byte("11111"), "22222")
	if bit == 64 && len(x) != 28 || bit == 32 && len(x) != 20 {
		t.Error("[]byte string error")
	}

	x = interfaceToBytes(true, false, uint8(1), int8(1), uint16(1), int16(1), uint32(1), int32(1), uint64(1), int64(1), uint(1), int(1), []byte("111111"), "222222")
	if bit == 64 && len(x) != 90 || bit == 32 && len(x) != 74 {
		t.Error("mixed args error")
	}

	// []byte or string is empty
	var emptyBytes []byte
	emptyStr := ""
	x = interfaceToBytes(emptyBytes, emptyStr)
	if bit == 64 && len(x) != 18 || bit == 32 && len(x) != 10 {
		t.Error("empty Bytes or String error")
	}

}

func TestInterfaceToBytesWithBuf(t *testing.T) {
	buf := make([]byte, 0, 64)
	mockFunc := func(i ...interface{}) []byte {
		return interfaceToBytesWithBuf(buf, i...)
	}

	// Copied from TestinterfaceToBytes
	x := interfaceToBytes(false, true)
	if len(x) != 4 {
		t.Error("bool error")
	}

	x = mockFunc(uint8(5), int8(6))
	if len(x) != 4 {
		t.Error("uint8 int8 error")
	}

	x = mockFunc(uint16(5), int16(6))
	if len(x) != 6 {
		t.Error("uint16 int16 error")
	}

	x = mockFunc(uint32(5), int32(6))
	if len(x) != 10 {
		t.Error("uint32 int32 error")
	}

	x = mockFunc(uint64(5), int64(6))
	if len(x) != 18 {
		t.Error("uint64 int64 error")
	}

	x = mockFunc(int(5), int(6))
	if bit == 64 && len(x) != 18 {
		t.Error("uint int error")
	} else if bit == 32 && len(x) != 10 {
		t.Error("uint int error")
	}

	x = mockFunc([]byte("11111"), "22222")
	if bit == 64 && len(x) != 28 || bit == 32 && len(x) != 20 {
		t.Error("[]byte string error")
	}

	x = mockFunc(true, false, uint8(1), int8(1), uint16(1), int16(1), uint32(1), int32(1), uint64(1), int64(1), uint(1), int(1), []byte("111111"), "222222")
	if bit == 64 && len(x) != 90 || bit == 32 && len(x) != 74 {
		t.Error("mixed args error")
	}

	// []byte or string is empty
	var emptyBytes []byte
	emptyStr := ""
	x = mockFunc(emptyBytes, emptyStr)
	if bit == 64 && len(x) != 18 || bit == 32 && len(x) != 10 {
		t.Error("empty Bytes or String error")
	}
}
