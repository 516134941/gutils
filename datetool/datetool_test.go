package datetool

import (
	"testing"
	"time"
)

func TestGetSubTime(t *testing.T) {
	startTime := "2019-09-01 12:00:09"
	endTime := "2019-09-03 13:44:44"
	d := GetSubTime(startTime, endTime, "2006-01-02 15:04:05")
	t.Logf("d:%v", d) // d:49h44m35s
}

func TestGetTimePeriodDesc(t *testing.T) {
	startTime := "2019-09-01 12:00:09"
	endTime := "2019-09-03 13:44:44"
	d := GetSubTime(startTime, endTime, "2006-01-02 15:04:05")
	timeDesc := GetTimePeriodDesc(d)
	t.Logf("timeDesc:%v", timeDesc) // timeDesc:2天1小时44分35秒
}

func TestGetZeroTime(t *testing.T) {
	now := time.Now()
	zeroNow := GetZeroTime(now)
	t.Logf("zeroNow:%v", zeroNow) // timeDesc:2019-09-25 00:00:00 +0800 CST
}

func TestGetMaxTime(t *testing.T) {
	now := time.Now()
	maxNow := GetMaxTime(now)
	t.Logf("maxNow:%v", maxNow) // timeDesc:2019-09-25 00:00:00 +0800 CST
}

func TestGetZeroTimeString(t *testing.T) {
	timeStr := "2019-09-08 09:32:44"
	timeStr = GetZeroTimeString(timeStr)
	t.Logf("timeStr:%v", timeStr) // timeStr:2019-09-08 00:00:00
}

func TestTimeSubDays(t *testing.T) {

	t1 := time.Now()
	t2 := t1.Add(333 * time.Hour)
	days := TimeSubDays(t1, t2)
	t.Logf("days:%v", days) //  days:14
	t3 := t1.Add(-333 * time.Hour)
	days2 := TimeSubDays(t1, t3)
	t.Logf("days:%v", days2) //  days:14
}

func TestGetTimeYearMonthDay(t *testing.T) {
	timeStr := "2019-09-21 12:00:31"
	y, m, d := GetTimeYearMonthDay(timeStr)
	t.Logf("year:%v", y)  //  year:2019
	t.Logf("month:%v", m) //  month:9
	t.Logf("day:%v", d)   //  day:21
}

func TestGetTimeTypeAndFormat(t *testing.T) {
	t1 := "2019"
	t2 := "2019-08"
	t3 := "2019-08-05"
	dateType, startDate, endDate := GetTimeTypeAndFormat(t1)
	t.Logf("start:%v  end:%v  dateType:%v", startDate, endDate, dateType) // start:2019-01-01  end:2020-01-01  dateType:year
	dateType, startDate, endDate = GetTimeTypeAndFormat(t2)
	t.Logf("start:%v  end:%v  dateType:%v", startDate, endDate, dateType) // start:2019-08-01  end:2019-09-01  dateType:month
	dateType, startDate, endDate = GetTimeTypeAndFormat(t3)
	t.Logf("start:%v  end:%v  dateType:%v", startDate, endDate, dateType) // start:2019-08-05  end:2019-08-06  dateType:day
}

func TestGetISOYearWeek(t *testing.T) {
	t1 := time.Now()
	y, w := GetISOYearWeek(t1)
	t.Logf("year:%v, week:%v", y, w) // year:2019, week:45
}
