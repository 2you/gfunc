package gfunc

import (
	"fmt"
	"time"
)

// 将当前系统的时间转为字符串 精确到微秒
func CurrTime2Str_Micro() string {
	currTime := time.Now()
	hour, min, sec := currTime.Clock()
	microSec := currTime.UTC().Nanosecond() / 1000
	sNow := fmt.Sprintf("%0.2d:%0.2d:%0.2d.%0.6d", hour, min, sec, microSec)
	return sNow
}

// 将当前系统的时间转为字符串 精确到毫秒
func CurrTime2Str_Mill() string {
	currTime := time.Now()
	hour, min, sec := currTime.Clock()
	millSec := currTime.UTC().Nanosecond() / 1000 / 1000
	sNow := fmt.Sprintf("%0.2d:%0.2d:%0.2d.%0.3d", hour, min, sec, millSec)
	return sNow
}

// 将当前系统的时间转为字符串 精确到秒
func CurrTime2Str_Sec() string {
	currTime := time.Now()
	hour, min, sec := currTime.Clock()
	sNow := fmt.Sprintf("%0.2d:%0.2d:%0.2d", hour, min, sec)
	return sNow
}

// 将当前系统的日期转为字符串
func CurrDate2Str() string {
	currTime := time.Now()
	year, month, day := currTime.Date()
	sNow := fmt.Sprintf("%0.4d-%0.2d-%0.2d", year, month, day)
	return sNow
}

// 将昨天的日期转化为字符串
func YestodayDate2Str() string {
	currTime := time.Now()
	yestodayTime := currTime.AddDate(0, 0, -1)
	year, month, day := yestodayTime.Date()
	sNow := fmt.Sprintf("%0.4d-%0.2d-%0.2d", year, month, day)
	return sNow
}

func DateTime2Str_Sec(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func DateTime2Str_Mill(t time.Time) string {
	year, month, day := t.Date()
	hour, min, sec := t.Clock()
	millSec := t.UTC().Nanosecond() / 1000 / 1000
	return fmt.Sprintf("%0.4d-%0.2d-%0.2d %0.2d:%0.2d:%0.2d.%0.3d",
		year, month, day, hour, min, sec, millSec)
}

func DateTime2Str_Micro(t time.Time) string {
	year, month, day := t.Date()
	hour, min, sec := t.Clock()
	microSec := t.UTC().Nanosecond() / 1000
	return fmt.Sprintf("%0.4d-%0.2d-%0.2d %0.2d:%0.2d:%0.2d.%0.6d",
		year, month, day, hour, min, sec, microSec)
}

// 将当前系统的日期时间转为字符串 精确到秒
func CurrDateTime2Str_Sec() string {
	return DateTime2Str_Sec(time.Now())
}

// 将当前系统的日期时间转为字符串 精确到毫秒
func CurrDateTime2Str_Mill() string {
	return DateTime2Str_Mill(time.Now())
}

// 将当前系统的日期时间转为字符串 精确到微秒
func CurrDateTime2Str_Micro() string {
	return DateTime2Str_Micro(time.Now())
}

func CurrUnixTime() int64 {
	return time.Now().Unix()
}

func CurrUnixNanoTime() int64 {
	return time.Now().UnixNano()
}

func ParseStrDateTimeInLocation(str string, location *time.Location) (time.Time, error) {
	return time.ParseInLocation(`2006-01-02 15:04:05`, str, location)
}

func ParseStrDateTime(str string) (time.Time, error) {
	return ParseStrDateTimeInLocation(str, time.Local)
}

func StrToDateTime(v string, location *time.Location) time.Time {
	tm, _ := time.ParseInLocation(`2006-01-02 15:04:05`, v, location)
	return tm
}

func Str2DateTime(v string) time.Time {
	return StrToDateTime(v, time.Local)
}

func Str2LocalDateTime(v string) time.Time {
	return StrToDateTime(v, time.Local)
}

func Str2UtcDateTime(v string) time.Time {
	loc, _ := time.LoadLocation(`UTC`)
	return StrToDateTime(v, loc)
}

func ParseStrDateInLocation(str string, location *time.Location) (time.Time, error) {
	return time.ParseInLocation(`2006-01-02`, str, location)
}

func ParseStrDate(str string) (time.Time, error) {
	return ParseStrDateInLocation(str, time.Local)
}

func StrToDate(v string, location *time.Location) time.Time {
	tm, _ := time.ParseInLocation(`2006-01-02`, v, location)
	return tm
}

func Str2Date(v string) time.Time {
	return StrToDate(v, time.Local)
}

func Str2LocalDate(v string) time.Time {
	return StrToDate(v, time.Local)
}

func Str2UtcDate(v string) time.Time {
	loc, _ := time.LoadLocation(`UTC`)
	return StrToDate(v, loc)
}

func ParseStrTimeInLocation(str string, location *time.Location) (time.Time, error) {
	return time.ParseInLocation(`15:04:05`, str, location)
}

func ParseStrTime(str string) (time.Time, error) {
	return ParseStrTimeInLocation(str, time.Local)
}

func StrToTime(v string, location *time.Location) time.Time {
	tm, _ := time.ParseInLocation(`15:04:05`, v, location)
	return tm
}

func Str2Time(v string) time.Time {
	return StrToTime(v, time.Local)
}

func Str2LocalTime(v string) time.Time {
	return StrToTime(v, time.Local)
}

func Str2UtcTime(v string) time.Time {
	loc, _ := time.LoadLocation(`UTC`)
	return StrToTime(v, loc)
}

func DateSame(tm1, tm2 time.Time) bool {
	y1, m1, d1 := tm1.Date()
	y2, m2, d2 := tm2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

func ToCnTime(tm time.Time) time.Time {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	return tm.In(loc)
}
