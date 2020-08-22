package xstring

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

const (
	// PadRight Right padding character
	PadRight int = iota
	// PadLeft Left padding character
	PadLeft
)

// StringToInt64 String转int64
func StringToInt64(e string) (int64, error) {
	return strconv.ParseInt(e, 10, 64)
}

// StringToInt String转int
func StringToInt(e string) (int, error) {
	return strconv.Atoi(e)
}

// GetCurrntTimeStr  获取当前时间
func GetCurrntTimeStr() string {
	return time.Now().Format("2006/01/02 15:04:05")
}

// GetCurrntTime 获取当前时间
func GetCurrntTime() time.Time {
	return time.Now()
}

// StructToJSONStr struct转json str
func StructToJSONStr(e interface{}) (string, error) {
	if b, err := json.Marshal(e); err == nil {
		return string(b), err
	}
	return "", nil

}

// JSONStrToMap json转map
func JSONStrToMap(e string) (map[string]interface{}, error) {
	var dict map[string]interface{}
	if err := json.Unmarshal([]byte(e), &dict); err == nil {
		return dict, err
	}
	return nil, nil
}

// StructToMap struct转map
func StructToMap(data interface{}) (map[string]interface{}, error) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	mapData := make(map[string]interface{})
	err = json.Unmarshal(dataBytes, &mapData)
	if err != nil {
		return nil, err
	}
	return mapData, nil
}

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

// 用掩码进行替换
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func RandStringBytesMaskImpr(n int) string {
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

// RandStringRunes 返回随机字符串
func RandStringRunes(n int) string {
	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// GenValidateCode 生成六位数字验证码
func GenValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}
