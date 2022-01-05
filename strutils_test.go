package gfunc

import (
	"strings"
	"testing"
)

func Test_Bytes2String(t *testing.T) {
	var param = []byte{0x47, 0x4f}
	str := Bytes2String(param)
	t.Log(str)
}

func Test_StringTrim(t *testing.T) {
	s := "   abc\r\n\t"
	s = StringTrim(s)
	if strings.IndexAny(s, "\r\n\t ") >= 0 {
		t.Error(`string trim error`)
	}
}
