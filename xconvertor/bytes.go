package xconvertor

import (
	"bytes"
	"encoding/binary"
	"unsafe"
)

// Bytes2String bytes to string
func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// Bytes2Int bytes to int
func Bytes2Int(b []byte) (int, error) {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	err := binary.Read(bytesBuffer, binary.BigEndian, &x)
	if err != nil {
		return 0, err
	}

	return int(x), nil
}

// BytesToInt64 byte to int64
func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}
