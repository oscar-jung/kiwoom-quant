package kiwoom

import "time"

func ParseKiwoomTime(timeStr string) (time.Time, error) {
	// 키움 날짜 포맷: 년(4)월(2)일(2)시(2)분(2)초(2)
	layout := "20060102150405"
	return time.Parse(layout, timeStr)
}
