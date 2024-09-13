package helper

import "time"

func GetCurrentTime() int64 {
	return int64(time.Now().Unix())
}

