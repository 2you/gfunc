package gfunc

import (
	"testing"
	"time"
)

func Test_CurrTime2Str_Micro(t *testing.T) {
	tm := CurrTime2Str_Micro()
	t.Log(tm)
}

func Test_CurrTime2Str_Mill(t *testing.T) {
	tm := CurrTime2Str_Mill()
	t.Log(tm)
}

func Test_CurrTime2Str_Sec(t *testing.T) {
	tm := CurrTime2Str_Sec()
	t.Log(tm)
}

func Test_CurrDate2Str(t *testing.T) {
	s := CurrDate2Str()
	t.Log(s)
}

func Test_YestodayDate2Str(t *testing.T) {
	s := YestodayDate2Str()
	t.Log(s)
}

func Test_DateTime2Str_Sec(t *testing.T) {
	s := DateTime2Str_Sec(time.Now())
	t.Log(s)
}

func Test_DateTime2Str_Mill(t *testing.T) {
	s := DateTime2Str_Mill(time.Now())
	t.Log(s)
}

func Test_DateTime2Str_Micro(t *testing.T) {
	dt := DateTime2Str_Micro(time.Now())
	t.Log(dt)
}

func Test_CurrDateTime2Str_Sec(t *testing.T) {
	s := CurrDateTime2Str_Sec()
	t.Log(s)
}

func Test_CurrDateTime2Str_Mill(t *testing.T) {
	s := CurrDateTime2Str_Mill()
	t.Log(s)
}

func Test_CurrDateTime2Str_Micro(t *testing.T) {
	s := CurrDateTime2Str_Micro()
	t.Log(s)
}

func Test_CurrUnixTime(t *testing.T) {
	n := CurrUnixTime()
	t.Log(n)
}

func Test_CurrUnixNanoTime(t *testing.T) {
	n := CurrUnixNanoTime()
	t.Log(n)
}

func Test_ParseStrDateTimeInLocation(t *testing.T) {
	if dt, err := ParseStrDateTimeInLocation(`2022-01-01 01:02:03`, time.UTC); err != nil {
		t.Error(err)
	} else {
		t.Log(dt)
	}
}

func Test_ParseStrDateTime(t *testing.T) {
	if dt, err := ParseStrDateTime(`2022-01-01 01:02:03`); err != nil {
		t.Error(err)
	} else {
		t.Log(dt)
	}
}

func Test_StrToDateTime(t *testing.T) {
	dt := StrToDateTime(`2022-01-01 01:02:03`, time.UTC)
	t.Log(dt)
}

func Test_Str2DateTime(t *testing.T) {
	dt := Str2DateTime(`2022-01-01 01:02:03`)
	t.Log(dt)
}

func Test_Str2LocalDateTime(t *testing.T) {
	dt := Str2LocalDateTime(`2022-01-01 01:02:03`)
	t.Log(dt)
}

func Test_Str2UtcDateTime(t *testing.T) {
	dt := Str2UtcDateTime(`2022-01-01 01:02:03`)
	t.Log(dt)
}