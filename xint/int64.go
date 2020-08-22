package xint

import "strconv"

// Int64ToString int64转string
func Int64ToString(e int64) string {
	return strconv.FormatInt(e, 10)
}
