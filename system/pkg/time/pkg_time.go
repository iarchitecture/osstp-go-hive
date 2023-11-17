package pkg_time

import "time"

func GetNowTimeString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
