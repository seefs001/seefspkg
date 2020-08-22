package xfloat

import "strconv"

// Float64ToString float64转string
func Float64ToString(e float64) string {
	return strconv.FormatFloat(e, 'E', -1, 64)
}

