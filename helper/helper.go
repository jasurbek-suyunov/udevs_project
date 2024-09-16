package helper

import (
	"strconv"
	"time"
)

func GetCurrentTime() int64 {
	return int64(time.Now().Unix())
}

func CheckIntegers(param string) bool {
	res, err := strconv.Atoi(param)
	if err != nil {
		return false
	} else if param == "" {
		return false
	} else if res < 0 {
		return false
	} else if res == 0 {
	} else {
		return true
	}
	return false
}