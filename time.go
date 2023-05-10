package utils

import (
"strconv"
"time"
)

const (
	FormatDate             = "2006-01-02"
	FormatDateTime         = "2006-01-02 15:04:05"
	FormatDateTimeMilliSec = "2006-01-02 15:04:05.999"
	FormatDayString         = "20060102"
)

// 获取毫秒级UTC时间戳
func GetNowUtcUnixMilli() int64 {
	return int64(time.Now().UTC().UnixNano() / 1000000)
}

// 获取秒级UTC时间戳
func GetNowUtcUnix() int64 {
	return time.Now().UTC().Unix()
}

// 纳秒级时间戳转毫秒级时间戳 19位->13位
func UnixNano2UnixMilli(nano int64) int64 {
	return int64(nano / 1000000)
}

// 纳秒级时间戳转秒级时间戳 19位->10位
func UnixNano2Unix(nano int64) int64 {
	return int64(nano / 1000000000)
}

// 毫秒级时间戳转秒级时间戳 13位->10位
func UnixMilli2Unix(milli int64) int64 {
	return int64(milli / 1000)
}

// 格式化毫秒级时间戳 13位->2006-01-02 15:04:05.999
func UnixMilli2TimeStr(milli int64) string {
	return time.Unix(0, int64(milli*1000000)).Format("2006-01-02 15:04:05.999")
}

// 格式化毫秒级时间戳 13位->2006-01-02
func UnixMilli2DateStr(milli int64) string {
	return time.Unix(0, int64(milli*1000000)).Format("2006-01-02")
}

// 毫秒级时间戳转UTC时间  13位->time.Time
func UnixMilli2UtcTime(milli int64) time.Time {
	return time.Unix(int64(milli/1000), 0).UTC()
}

//2014-12-25 18:12:25  --->   1420511108210
func TimeStr2UnixMilli(timeStr string) (int64, error) {
	t, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)
	if err != nil {
		return 0, err
	}
	return int64(t.UnixNano() / 1000000), nil
}

//2014-12-25 18:12:25  --->   1420511108210
func TimeStr2Unix(timeStr string) (int64, error) {
	t, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

func UnixMilli2Hour(milli int64) int {
	t := time.Unix(int64(milli/1000), 0)
	return t.Hour()
}

func RangeDay(start string, end string) ([]string, error) {
	r := []string{}

	s, err := time.Parse(FormatDate, start)
	if err != nil {
		return r, err
	}
	ss := s.Format(FormatDate)

	e, err := time.Parse(FormatDate, end)
	if err != nil {
		return r, err
	}
	ee := e.Format(FormatDate)

	r = append(r, ss)
	if ss == ee {
		return r, nil
	}

	for ee > ss {
		s = s.AddDate(0, 0, 1)
		ss = s.Format(FormatDate)
		r = append(r, ss)
	}

	return r, nil
}

func GetYesterdayDate() string {
	return time.Now().AddDate(0, 0, -1).Format(FormatDate)
}

//返回年月日字符串
func GetYestDayString() string {
	nowTime := time.Now()
	yesterdayString := nowTime.AddDate(0, 0, -1)
	return yesterdayString.Format(FormatDayString)
}

//返回年月日时分秒字符串
func GetFormatDTString() string {
	nowTime := time.Now()
	timeToString := nowTime.Format("20060102150405")
	return timeToString
}

//返回年月日时分秒毫秒字符串
func GetMillSecondDTString() string {
	now := time.Now()
	timeToString := now.Format("20060102150405")
	return timeToString + strconv.Itoa(now.Nanosecond())[0:3]
}

func GetFirstDateOfWeek() (weekMonday string) {
	now := time.Now()

	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}

	weekStartDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	weekMonday = weekStartDate.Format("2006-01-02")
	return
}

/**
获取上周的日期范围
例：05月17日-05月23日
*/
func GetLastWeekRange() (weekMonday, weekSunday string) {
	thisWeekMonday := GetFirstDateOfWeek()
	TimeMonday, _ := time.Parse("2006-01-02", thisWeekMonday)
	lastWeekMonday := TimeMonday.AddDate(0, 0, -7)
	lastWeekSunday := TimeMonday.AddDate(0, 0, -1)
	weekMonday = lastWeekMonday.Format("01月02日")
	weekSunday = lastWeekSunday.Format("01月02日")
	return weekMonday, weekSunday
}
