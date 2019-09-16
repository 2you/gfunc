package gfunc

import (
	"testing"
)

func Test_XorEncrypt(t *testing.T) {
	str1 := "异或数据原文，字符串显示。\nGoGoGo"
	key := "异或密钥"
	buf1 := []byte(str1)
	buf2 := XorEncrypt(buf1, []byte(key))
	t.Logf("原数据\n------------------------------------\n%s\n------------------------------------\n", string(buf1))
	t.Logf("加密后\n------------------------------------\n%v\n------------------------------------\n", string(buf2))
	buf3 := XorEncrypt(buf2, []byte(key))
	t.Logf("解密后\n------------------------------------\n%v\n------------------------------------\n", string(buf3))
}

func Test_AES_CBC_Base64_Decrypt(t *testing.T) {
	src1 := "ev2oYqyOpmrwAR4UnxwvcKZoKv88kP7fgyoEHTKxs2o="
	key1 := "0123456789123456"
	src2 := "fJriagvQS5TZFjuvWMTZhd3ibzMmkuUZHQEVrcEJg3w="
	key2 := "012345678901234567891234"
	src3 := "nRtEw8Wxmb1+nmaUIF+deoL3z80gAa47sHy8lNKSKzM="
	key3 := "01234567890123456789012345678912"
	iv := "1234567812345678"
	dst, err := AES_CBC_Base64_Decrypt(src1, 128, []byte(key1), []byte(iv))
	if err != nil {
		t.Error(err)
	} else {
		t.Log(string(dst))
	}

	dst, err = AES_CBC_Base64_Decrypt(src2, 192, []byte(key2), []byte(iv))
	if err != nil {
		t.Error(err)
	} else {
		t.Log(string(dst))
	}

	dst, err = AES_CBC_Base64_Decrypt(src3, 256, []byte(key3), []byte(iv))
	if err != nil {
		t.Error(err)
	} else {
		t.Log(string(dst))
	}
}

func Test_AES_CBC_Base64_Encrypt(t *testing.T) {
	src1 := "AES CBC 128位密钥加密"
	key1 := "0123456789123456"
	src2 := "AES CBC 192位密钥加密"
	key2 := "012345678901234567891234"
	src3 := "AES CBC 256位密钥加密"
	key3 := "01234567890123456789012345678912"
	iv := "1234567812345678"
	dst, err := AES_CBC_Base64_Encrypt([]byte(src1), 128, []byte(key1), []byte(iv))
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%s %s", src1, dst)
	}

	dst, err = AES_CBC_Base64_Encrypt([]byte(src2), 192, []byte(key2), []byte(iv))
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%s %s", src2, dst)
	}

	dst, err = AES_CBC_Base64_Encrypt([]byte(src3), 256, []byte(key3), []byte(iv))
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%s %s", src3, dst)
	}
}

func Test_CurrTime2Str(t *testing.T) {
	time := CurrTime2Str_Sec()
	t.Log(time)
}

func Test_CurrDate2Str(t *testing.T) {
	date := CurrDate2Str()
	t.Log(date)
}

func Test_YestodayDate2Str(t *testing.T) {
	date := YestodayDate2Str()
	t.Log(date)
}

func Test_CurrDateTime2Str(t *testing.T) {
	dt := CurrDateTime2Str_Sec()
	t.Log(dt)
}

func Test_CurrUnixTime(t *testing.T) {
	now := CurrUnixTime()
	t.Log(now)
}

func Test_Str2DateTime(t *testing.T) {
	td := Str2DateTime("2017-03-10 09:39:23")
	t.Log(td)
}

func Test_Str2LocalDateTime(t *testing.T) {
	td := Str2LocalDateTime("2017-03-10 09:39:23")
	t.Log(td)
}

func Test_ForceDirectories(t *testing.T) {
	b := ForceDirectories("d1/d2/d3/d4/d5")
	t.Logf("%v\n", b)
}

func Test_DirExist(t *testing.T) {
	dir1, dir2 := "mahonia", "mahonia1"
	b := DirExist(dir1)
	t.Logf("%s %v\n", dir1, b)
	b = DirExist(dir2)
	t.Logf("%s %v\n", dir2, b)
}

func Test_FileExist(t *testing.T) {
	file1, file2 := "tcfunc.go", "tcfunc1.go"
	b := FileExist(file1)
	t.Logf("%s %v\n", file1, b)
	b = FileExist(file2)
	t.Logf("%s %v\n", file2, b)
}

func Test_Bytes2String(t *testing.T) {
	var param = []byte{0x47, 0x4f}
	str := Bytes2String(param)
	t.Logf("%s\n", str)
}

func Test_AppFilePath(t *testing.T) {
	str := AppFilePath()
	t.Logf("%s\n", str)
}

func Test_GetFileContent(t *testing.T) {
	if _, err := GetFileContent("gfunc.go"); err != nil {
		t.Errorf("%s\n", err.Error())
	}
}
