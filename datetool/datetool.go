// Package datetool 一些封装的时间工具方法
package datetool

import (
	"fmt"
	"strconv"
	"time"
)

// GetNowDate 获取当前日期
func GetNowDate() string {
	return time.Now().Format("2006-01-02")
}

// GetNowTime 获取当前时间(精确到秒)
func GetNowTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// FormatTimeAsDate 将时间格式化为日期
func FormatTimeAsDate(d string) string {
	t, _ := time.Parse("2006-01-02 15:04:05", d)
	return t.Format("2006-01-02")
}

// GetNowTimestamp 获取当前时间戳
func GetNowTimestamp() int64 {
	return time.Now().Unix()
}

// ParseToDateTime 字符串时间转Time
func ParseToDateTime(d string) (time.Time, error) {
	t, err := time.Parse("2006-01-02 15:04:05", d)
	return t, err
}

// ParseToDate 字符串时间转date
func ParseToDate(d string) (time.Time, error) {
	t, err := time.Parse("2006-01-02", d)
	return t, err
}

// GetSubTime 获取两个时间差
func GetSubTime(start string, end string, format string) time.Duration {
	st, _ := time.Parse(format, start)
	en, _ := time.Parse(format, end)
	return en.Sub(st)
}

// GetTimePeriodDesc 获取时间段的中文描述
func GetTimePeriodDesc(duration time.Duration) string {
	spanSeconds := int64(duration.Seconds())
	if spanSeconds/(24*60*60) > 0 {
		return fmt.Sprintf("%d天%d小时%d分%d秒", spanSeconds/(24*60*60), spanSeconds%(24*60*60)/(60*60), spanSeconds%(60*60)/60, spanSeconds%60)
	}
	return fmt.Sprintf("%d小时%d分%d秒", spanSeconds%(24*60*60)/(60*60), spanSeconds%(60*60)/60, spanSeconds%60)
}

// GetZeroTime 获取零点时间
func GetZeroTime(d time.Time) time.Time {
	loc, _ := time.LoadLocation("Local")
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, loc)
}

// GetMaxTime 获取某一天的最大的时间
func GetMaxTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 23, 59, 59, 0, d.Location())
}

// GetZeroTimeString 获取零点时间(返回string)
func GetZeroTimeString(timeStr string) string {
	loc, _ := time.LoadLocation("Local")
	tm, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc)
	tm = time.Date(tm.Year(), tm.Month(), tm.Day(), 0, 0, 0, 0, loc)
	return tm.Format("2006-01-02 15:04:05")
}

// TimeSubDays  返回两个时间相差天数
// Determine how many days difference between the two times
func TimeSubDays(t1, t2 time.Time) int {
	if t1.Location().String() != t2.Location().String() {
		return -1
	}
	hours := t1.Sub(t2).Hours()
	if hours < 0 {
		t1, t2 = t2, t1
		hours = t1.Sub(t2).Hours()
	} else if hours == 0 {
		return 0
	}
	if hours < 24 {
		t1y, t1m, t1d := t1.Date()
		t2y, t2m, t2d := t2.Date()
		isSameDay := (t1y == t2y && t1m == t2m && t1d == t2d)
		if isSameDay {
			return 0
		}
		return 1
	}
	if (hours/24)-float64(int(hours/24)) == 0 {
		return int(hours / 24)
	}
	return int(hours/24) + 1
}

// GetTimeYearMonthDay 分别获取传入时间的年月日
// Get the year, month, and day of the incoming time separately
func GetTimeYearMonthDay(date string) (year, month, day string) {
	loc, _ := time.LoadLocation("Local")
	tm, _ := time.ParseInLocation("2006-01-02 15:04:05", date, loc)
	return strconv.Itoa(tm.Year()), strconv.Itoa(int(tm.Month())), strconv.Itoa(tm.Day())
}

// GetTimeTypeAndFormat 获取前端的时间类型并且格式化为time
// eg:传入 '2019' 返回 year,2019-01-01 2020-01-01
// eg:传入 '2019-02' 返回 month,2019-02-01,2019-03-01
// eg:传入 '2019-02-04' 返回 day,2019-02-04,2019-02-05
func GetTimeTypeAndFormat(date string) (dateType, startDate, endDate string) {
	// 筛选类型判断
	if len(date) == 4 {
		dateType = "year"
		date += "-01-01"
	} else if len(date) == 7 {
		dateType = "month"
		date += "-01"
	} else if len(date) == 10 {
		dateType = "day"
	}
	loc, _ := time.LoadLocation("Local")
	tm, _ := time.ParseInLocation("2006-01-02", date, loc)
	switch dateType {
	case "year":
		startDate = tm.Format("2006-01-02")
		endDate = tm.AddDate(1, 0, 0).Format("2006-01-02")
	case "month":
		startDate = tm.Format("2006-01-02")
		endDate = tm.AddDate(0, 1, 0).Format("2006-01-02")
	case "day":
		startDate = tm.Format("2006-01-02")
		endDate = tm.AddDate(0, 0, 1).Format("2006-01-02")
	}
	return
}

// GetISOYearWeek 获取年周 传入时间获取对应的年和周
func GetISOYearWeek(t time.Time) (y, w int) {
	y, w = t.ISOWeek()
	return
}
