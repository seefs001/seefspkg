package xconvertor

import (
	"encoding/json"
	"strconv"
)

// JSONStrToMap json转map
func JSONStrToMap(e string) (map[string]interface{}, error) {
	var dict map[string]interface{}
	if err := json.Unmarshal([]byte(e), &dict); err == nil {
		return dict, err
	}
	return nil, nil
}

// StringToInt64 String转int64
func StringToInt64(e string) (int64, error) {
	return strconv.ParseInt(e, 10, 64)
}

// StringToInt String转int
func StringToInt(e string) (int, error) {
	return strconv.Atoi(e)
}
