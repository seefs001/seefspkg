package xdate

import "time"

// GetCurrntTimeStr  获取当前时间
func GetCurrntTimeStr() string {
	return time.Now().Format("2006/01/02 15:04:05")
}

// GetCurrntTime 获取当前时间
func GetCurrntTime() time.Time {
	return time.Now()
}
