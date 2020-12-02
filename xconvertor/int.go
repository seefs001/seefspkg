package xconvertor

import (
	"bytes"
	"encoding/binary"
	"strconv"
)

// IntToString int转string
func IntToString(e int) string {
	return strconv.Itoa(e)
}

// Int2Bytes 整形转换成字节
func Int2Bytes(n int) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}
