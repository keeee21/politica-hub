package util

import (
	"fmt"
	"time"
)

func ParsePublishedAt(dateStr string) (time.Time, error) {
	if dateStr == "" {
		return time.Time{}, fmt.Errorf("empty date string")
	}

	// ① RFC1123 形式の解析
	t, err := time.Parse(time.RFC1123, dateStr)
	if err == nil {
		return t, nil
	}

	// ② RFC1123Z 形式の解析
	t, err = time.Parse(time.RFC1123Z, dateStr)
	if err == nil {
		return t, nil
	}

	// ③ RFC3339（ISO8601）形式の解析
	t, err = time.Parse(time.RFC3339, dateStr)
	if err == nil {
		return t, nil
	}

	return time.Time{}, fmt.Errorf("failed to parse date: %s", dateStr)
}