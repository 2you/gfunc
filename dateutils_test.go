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

func Test_ParseStrDateInLocation(t *testing.T) {
	if dt, err := ParseStrDateInLocation(`2022-01-01`, time.UTC); err != nil {
		t.Error(err)
	} else {
		t.Log(dt)
	}
}

func Test_ParseStrDate(t *testing.T) {
	if dt, err := ParseStrDate(`2022-01-01`); err != nil {
		t.Error(err)
	} else {
		t.Log(dt)
	}
}

func Test_StrToDate(t *testing.T) {
	dt := StrToDate(`2022-01-01`, time.UTC)
	t.Log(dt)
}

func Test_Str2Date(t *testing.T) {
	dt := Str2Date(`2022-01-01`)
	t.Log(dt)
}

func Test_Str2LocalDate(t *testing.T) {
	dt := Str2LocalDate(`2022-01-01`)
	t.Log(dt)
}

func Test_Str2UtcDate(t *testing.T) {
	dt := Str2UtcDate(`2022-01-01`)
	t.Log(dt)
}

func Test_ParseStrTimeInLocation(t *testing.T) {
	if dt, err := ParseStrTimeInLocation(`01:02:03`, time.UTC); err != nil {
		t.Error(err)
	} else {
		t.Log(dt)
	}
}

func Test_ParseStrTime(t *testing.T) {
	if dt, err := ParseStrTime(`01:02:03`); err != nil {
		t.Error(err)
	} else {
		t.Log(dt)
	}
}

func Test_StrToTime(t *testing.T) {
	dt := StrToTime(`01:02:03`, time.UTC)
	t.Log(dt)
}

func Test_Str2Time(t *testing.T) {
	dt := Str2Time(`01:02:03`)
	t.Log(dt)
}

func Test_Str2LocalTime(t *testing.T) {
	dt := Str2LocalTime(`01:02:03`)
	t.Log(dt)
}

func Test_Str2UtcTime(t *testing.T) {
	dt := Str2UtcTime(`01:02:03`)
	t.Log(dt)
}

func Test_DateSame(t *testing.T) {
	tm1, _ := time.Parse(`2006-01-02 15:04:05`, `2022-08-10 02:03:04`)
	tm2, _ := time.Parse(`2006-01-02 15:04:05`, `2022-08-10 02:03:04`)
	same := DateSame(tm1, tm2)
	t.Log(same)

	tm1, _ = time.Parse(`2006-01-02 15:04:05`, `2022-08-10 03:04:05`)
	tm2, _ = time.Parse(`2006-01-02 15:04:05`, `2022-08-10 02:03:04`)
	same = DateSame(tm1, tm2)
	t.Log(same)

	tm1, _ = time.Parse(`2006-01-02 15:04:05`, `2022-08-10 02:03:04`)
	tm2, _ = time.Parse(`2006-01-02 15:04:05`, `2022-08-11 02:03:04`)
	same = DateSame(tm1, tm2)
	t.Log(same)
}

func Test_ToCnTime(t *testing.T) {
	utc := time.Now().UTC()
	tm := ToCnTime(utc)
	t.Log(utc, `||`, tm)
}
