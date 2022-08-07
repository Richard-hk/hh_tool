package util

import (
	"strconv"
	"time"
)

func GetNowDt() (dt int64) {
	dt, _ = strconv.ParseInt(time.Now().Format(TIME_LAYOUT_YMD), 10, 64)
	return dt
}
