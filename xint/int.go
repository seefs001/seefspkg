package xint

import (
	"bytes"
	"encoding/binary"
	"math/rand"
	"strconv"
	"time"
)

// IntToString int转string
func IntToString(e int) string {
	return strconv.Itoa(e)
}

// RandInt random numbers in the specified range
func RandInt(min int, max int) int {
	if max < min {
		max = min
	}
	rand.Seed(rand.Int63n(time.Now().UnixNano()))
	return min + rand.Intn(max+1-min)
}

//整形转换成字节
func Int2Bytes(n int) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}
