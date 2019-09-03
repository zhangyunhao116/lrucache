package lrucache

import "testing"

func TestInterfaceToBytes(t *testing.T) {
	x := InterfaceToBytes(false, true)
	if len(x) != 2 {
		t.Error("bool error")
	}

	x = InterfaceToBytes(uint8(5), int8(6))
	if len(x) != 2 {
		t.Error("uint8 int8 error")
	}

	x = InterfaceToBytes(uint16(5), int16(6))
	if len(x) != 4 {
		t.Error("uint16 int16 error")
	}

	x = InterfaceToBytes(uint32(5), int32(6))
	if len(x) != 8 {
		t.Error("uint32 int32 error")
	}

	x = InterfaceToBytes(uint64(5), int64(6))
	if len(x) != 16 {
		t.Error("uint64 int64 error")
	}

	x = InterfaceToBytes(int(5), int(6))
	if bit == 64 && len(x) != 16 {
		t.Error("uint int error")
	} else if bit == 32 && len(x) != 8 {
		t.Error("uint int error")
	}

	x = InterfaceToBytes([]byte("11111"), "22222")
	if string(x) != "1111122222" {
		t.Error("[]byte string error")
	}

	x = InterfaceToBytes(true, false, uint8(1), int8(1), uint16(1), int16(1), uint32(1), int32(1), uint64(1), int64(1), uint(1), int(1), []byte("111111"), "222222")
	if bit == 64 && len(x) != 60 || bit == 32 && len(x) != 52 {
		t.Error("mixed args error")
	}

	// []byte or string is empty
	var emptyBytes []byte
	emptyStr := ""
	x = InterfaceToBytes(emptyBytes, emptyStr)
	if len(x) != 0 {
		t.Error("empty Bytes or String error")
	}

	x = InterfaceToBytes(nil)
}

func BenchmarkInterfaceToBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InterfaceToBytes(i)
	}
}

func BenchmarkInterfaceToBytesWithBuf(b *testing.B) {
	buf := make([]byte, 0, 32)
	for i := 0; i < b.N; i++ {
		InterfaceToBytesWithBuf(buf, i)
	}
}
