package utils

import "time"

//return time format yyyy-mm-dd hh:ii:ss
func GetCurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

type TIMEFORMAT string

const (
	DATE     TIMEFORMAT = "2006-01-02"
	TIME     TIMEFORMAT = "15:04:05"
	DATETIME TIMEFORMAT = "2006-01-02 15:04:05"
)
