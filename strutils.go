package gfunc

import (
	"bytes"
	"strings"
)

//将byte数据转为string 若byte数据中存在0x0 则只取0x0前的数据
func Bytes2String(v []byte) string {
	nIndex := bytes.IndexByte(v, 0x0)
	if nIndex < 0 {
		return string(v)
	}
	return string(v[:nIndex])
}

func StringTrim(v string) (ret string) {
	ret = strings.Trim(v, "\r\n\t ")
	return ret
}
