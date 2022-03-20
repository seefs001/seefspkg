package xconvertor

import (
	"bytes"
	"encoding/binary"
	"strconv"
)

// IntToString int to string
func IntToString(e int) string {
	return strconv.Itoa(e)
}

// Int2Bytes int to bytes
func Int2Bytes(n int) ([]byte, error) {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	err := binary.Write(bytesBuffer, binary.BigEndian, x)
	if err != nil {
		return nil, err
	}
	return bytesBuffer.Bytes(), nil
}
