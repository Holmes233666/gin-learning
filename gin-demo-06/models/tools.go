package models

import "time"

func UnixToTime(unix int64) string {
	t := time.Unix(unix, 0)
	return t.Format("2006-01-02 15:04:05")
}
