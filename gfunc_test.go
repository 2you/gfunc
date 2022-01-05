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
