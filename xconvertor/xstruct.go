package xconvertor

import "encoding/json"

// StructToMap struct to map
func StructToMap(data any) (map[string]any, error) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	mapData := make(map[string]any)
	err = json.Unmarshal(dataBytes, &mapData)
	if err != nil {
		return nil, err
	}
	return mapData, nil
}

// StructToJSONStr struct to json str
func StructToJSONStr(e any) (string, error) {
	if b, err := json.Marshal(e); err == nil {
		return string(b), err
	}
	return "", nil
}
