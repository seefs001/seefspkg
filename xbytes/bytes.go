package util

import (
	"bytes"
	"encoding/binary"
	"unsafe"
)

// Bytes2String bytes to string
func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

//字节转换成整形
func Bytes2Int(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}
