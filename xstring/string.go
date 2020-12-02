package xstring

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

const (
	// PadRight Right padding character
	PadRight int = iota
	// PadLeft Left padding character
	PadLeft
)

// Len string length (utf8)
func Len(str string) int {
	// strings.Count(str,"")-1
	return utf8.RuneCountInString(str)
}

// Substr returns part of a string
func Substr(str string, start int, length ...int) string {
	s := []rune(str)
	sl := len(s)
	if start < 0 {
		start = sl + start
	}

	if len(length) > 0 {
		ll := length[0]
		if ll < 0 {
			sl = sl + ll
		} else {
			sl = ll + start
		}
	}
	return string(s[start:sl])
}

// Pad String padding
func Pad(raw string, length int, padStr string, padType int) string {
	l := length - Len(raw)
	if l <= 0 {
		return raw
	}
	if padType == PadRight {
		raw = fmt.Sprintf("%s%s", raw, strings.Repeat(padStr, l))
	} else if padType == PadLeft {
		raw = fmt.Sprintf("%s%s", strings.Repeat(padStr, l), raw)
	} else {
		left := 0
		right := 0
		if l > 1 {
			left = l / 2
			right = (l / 2) + (l % 2)
		}

		raw = fmt.Sprintf("%s%s%s", strings.Repeat(padStr, left), raw, strings.Repeat(padStr, right))
	}
	return raw
}
