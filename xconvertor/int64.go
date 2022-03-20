package xconvertor

import (
	"encoding/binary"
	"fmt"
	"reflect"
	"strconv"
)

// Int64ToString int64 to string
func Int64ToString(e int64) string {
	return strconv.FormatInt(e, 10)
}

// Int64ToBytes int64 to bytes
func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

// IntToInt64 int to int64
func IntToInt64(value int) int64 {
	i, _ := strconv.ParseInt(string(rune(value)), 10, 64)
	return i
}

// ToInt64 convert any numeric value to int64
// any to int64
func ToInt64(value any) (d int64, err error) {
	val := reflect.ValueOf(value)
	switch value.(type) {
	case int, int8, int16, int32, int64:
		d = val.Int()
	case uint, uint8, uint16, uint32, uint64:
		d = int64(val.Uint())
	case string:
		d, err = strconv.ParseInt(val.String(), 10, 64)
	default:
		err = fmt.Errorf("ToInt64 need numeric not `%T`", value)
	}
	return
}
