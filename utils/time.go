package utils

import "time"

// "20190108135039" => "2019-01-08 13:50:39"
func StringToTimestamp(timeStr string) int64 {
	//时间转换的模板，golang里面只能是 "2006-01-02 15:04:05"
	layout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(layout, timeStr[0:4]+"-"+timeStr[4:6]+"-"+timeStr[6:8]+" "+
		timeStr[8:10]+":"+timeStr[10:12]+":"+timeStr[12:14], loc)
	return theTime.UnixNano()
}
