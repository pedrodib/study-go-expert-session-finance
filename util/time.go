package util

import "time"

var timePattern = "2006-01-02T15:04:05"

// StringToTime  Convert String to time.Time
func StringToTime(value string) time.Time {
	convertedTime, _ := time.Parse(timePattern, value)

	return convertedTime
}
