package xconvertor

import (
	"encoding/json"
	"strconv"
)

// JSONStrToMap json to map
func JSONStrToMap(e string) (map[string]any, error) {
	var dict map[string]any
	if err := json.Unmarshal([]byte(e), &dict); err == nil {
		return dict, err
	}
	return nil, nil
}

// StringToInt64 String to int64
func StringToInt64(e string) (int64, error) {
	return strconv.ParseInt(e, 10, 64)
}

// StringToInt String to int
func StringToInt(e string) (int, error) {
	return strconv.Atoi(e)
}
