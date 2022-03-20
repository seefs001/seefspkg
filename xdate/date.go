package xdate

import "time"

// GetCurrentTimeStr get current time string
func GetCurrentTimeStr() string {
	return time.Now().Format("2006/01/02 15:04:05")
}

// GetCurrentTime get current time
func GetCurrentTime() time.Time {
	return time.Now()
}
