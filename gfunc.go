package gfunc

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func XorEncrypt(src []byte, key []byte) []byte {
	ssize := len(src)
	ret := make([]byte, ssize)
	copy(ret, src)
	for i := 0; i < len(key); i++ {
		for j := 0; j < ssize; j++ {
			ret[j] = ret[j] ^ key[i]
		}
	}
	return ret
}

func DelphiTrim(v string) (ret string) {
	ret = strings.Trim(v, "\r\n\t ")
	return ret
}

func UnGZIP(v []byte) (ret []byte) {
	var b bytes.Buffer
	b.Write(v)
	r, _ := gzip.NewReader(&b)
	defer r.Close()
	ret, _ = ioutil.ReadAll(r)
	// fmt.Println("ungzip size:", len(ret))
	return ret
}

func HttpDataSizeGet(geturl string, headers map[string]string, params map[string]string) uint64 {
	if headers == nil {
		headers = make(map[string]string)
	}
	headers["Range"] = "bytes=0-0"
	var body io.Reader = nil
	if params != nil {
		urlValues := url.Values{}
		for k, v := range params {
			urlValues.Set(k, v)
		}
		body = ioutil.NopCloser(strings.NewReader(urlValues.Encode())) //把form数据编下码
	}
	httpClient := &http.Client{}
	httpReq, _ := http.NewRequest("GET", geturl, body)
	for k, v := range headers {
		httpReq.Header.Set(k, v)
	}

	httpResp, err := httpClient.Do(httpReq)
	defer httpResp.Body.Close()
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	if httpResp.StatusCode != 206 {
		fmt.Println(`response status code is`, httpResp.StatusCode)
		return 0
	}
	crangemap := httpResp.Header["Content-Range"]
	//	fmt.Println(contentrange)
	if len(crangemap) < 1 {
		fmt.Println("headers key Content-Range is null")
	}
	contentrange := crangemap[0]
	index := strings.Index(contentrange, "/")
	//	fmt.Println(index)
	if index < 1 {
		fmt.Println("headers key Content-Range value is", crangemap)
	}
	length := contentrange[index+1:]
	//	fmt.Println(length)
	ret, _ := strconv.ParseUint(length, 10, 64)
	return ret
}

func HttpGet(geturl string, headers map[string]string, params map[string]string) []byte {
	var body io.Reader = nil
	if params != nil {
		urlValues := url.Values{}
		for k, v := range params {
			urlValues.Set(k, v)
		}
		str := urlValues.Encode()
		if strings.Index(str, "=") == 0 {
			str = str[1:]
		}
		body = ioutil.NopCloser(strings.NewReader(str)) //把form数据编下码
	}
	httpClient := &http.Client{}
	httpReq, _ := http.NewRequest("GET", geturl, body)
	for k, v := range headers {
		httpReq.Header.Set(k, v)
	}

	httpResp, err := httpClient.Do(httpReq)
	defer httpResp.Body.Close()
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	if httpResp.StatusCode != 200 && httpResp.StatusCode != 206 {
		fmt.Println(`response status code is`, httpResp.StatusCode)
		return nil
	}
	//	fmt.Println(httpResp.Header)
	data, _ := ioutil.ReadAll(httpResp.Body)
	return data
}

func HttpPost(posturl string, headers map[string]string, params map[string]string) []byte {
	var body io.Reader = nil
	if params != nil {
		urlValues := url.Values{}
		for k, v := range params {
			urlValues.Set(k, v)
		}
		str := urlValues.Encode()
		if strings.Index(str, "=") == 0 {
			str = str[1:]
		}
		body = ioutil.NopCloser(strings.NewReader(str)) //把form数据编下码
	}
	httpClient := &http.Client{}
	httpReq, _ := http.NewRequest("POST", posturl, body)
	httpReq.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value") //这个一定要加，不加form的值post不过去
	for k, v := range headers {
		httpReq.Header.Set(k, v)
	}
	//	fmt.Printf("%+v\n", httpReq) //看下发送的结构
	httpResp, err := httpClient.Do(httpReq) //发送
	defer httpResp.Body.Close()             //一定要关闭resp.Body
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	if httpResp.StatusCode != 200 && httpResp.StatusCode != 206 {
		fmt.Println(`response status code is`, httpResp.StatusCode)
		return nil
	}
	data, _ := ioutil.ReadAll(httpResp.Body)
	return data
}

func StrToFloat64Def(v string, d float64) float64 {
	var f float64
	var err error
	if f, err = strconv.ParseFloat(v, 64); err != nil {
		return d
	}
	return f
}

func StrToFloat64(v string) float64 {
	var f float64
	var err error
	if f, err = strconv.ParseFloat(v, 64); err != nil {
		panic(err)
	}
	return f
}

func Float64ToStr(v float64) string {
	return strconv.FormatFloat(v, 'f', -1, 64)
}

func UInt64ToStr(v uint64) string {
	return strconv.FormatUint(v, 10)
}

func Int64ToStr(v int64) string {
	return strconv.FormatInt(v, 10)
}

func StrToUInt64Def(v string, d uint64) uint64 {
	var n uint64
	var err error
	if n, err = strconv.ParseUint(v, 10, 64); err != nil {
		return d
	}
	return n
}

func StrToUInt64(v string) uint64 {
	var n uint64
	var err error
	if n, err = strconv.ParseUint(v, 10, 64); err != nil {
		panic(err)
	}
	return n
}

func StrToInt64Def(v string, d int64) int64 {
	var n int64
	var err error
	if n, err = strconv.ParseInt(v, 10, 64); err != nil {
		return d
	}
	return n
}

func StrToInt64(v string) int64 {
	var n int64
	var err error
	if n, err = strconv.ParseInt(v, 10, 64); err != nil {
		panic(err)
	}
	return n
}

func Hex2Bytes(v string) []byte {
	ret, _ := hex.DecodeString(v)
	return ret
}

func Bytes2Hex(v []byte) string {
	return hex.EncodeToString(v)
}

func HMACSHA1(data, key []byte) []byte {
	hmacHash := hmac.New(sha1.New, key)
	hmacHash.Write(data)
	return hmacHash.Sum(nil)
}

func SHA1String(v []byte) string {
	sha1Hash := sha1.New()
	sha1Hash.Write(v)
	return hex.EncodeToString(sha1Hash.Sum(nil))
}
func Md5Bytes(v []byte) string {
	md5Hash := md5.New()
	md5Hash.Write(v)
	sRet := hex.EncodeToString(md5Hash.Sum(nil))
	return sRet
}

func Md5String(v string) string {
	md5Hash := md5.New()
	md5Hash.Write([]byte(v))
	sRet := hex.EncodeToString(md5Hash.Sum(nil))
	return sRet
}

func Base64Encode(v []byte) string {
	encode := base64.StdEncoding
	return encode.EncodeToString(v)
}

func Base64Decode(v string) []byte {
	encode := base64.StdEncoding
	if data, err := encode.DecodeString(v); err != nil {
		fmt.Println(`[Base64Decode]`, err.Error())
		return []byte(``)
	} else {
		return data
	}
}

//先base64解密再zlib解压
func ZBase64Decompress(v string) []byte {
	z := Base64Decode(v)
	return ZlibDecompress(z)
}

//先zlib压缩再base64加密
func ZBase64Compress(v []byte) string {
	z := ZlibCompress(v)
	return Base64Encode(z)
}

func ZlibDecompress(v []byte) []byte {
	b := bytes.NewReader(v)
	var out bytes.Buffer
	if r, e := zlib.NewReader(b); e != nil {
		fmt.Println(`[ZlibDecompress]`, e.Error())
		return []byte(``)
	} else {
		io.Copy(&out, r)
		return out.Bytes()
	}
}

func ZlibCompress(v []byte) []byte {
	var in bytes.Buffer
	w := zlib.NewWriter(&in)
	w.Write(v)
	w.Close()
	return in.Bytes()
}

func AES_CBC_Base64_Decrypt(data string, keysize int, key, iv []byte) ([]byte, error) {
	decdata, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	if decdata, err = AES_CBC_Decrypt(decdata, keysize, key, iv); err != nil {
		return nil, err
	}
	return decdata, nil
}

func AES_CBC_Base64_Encrypt(data []byte, keysize int, key, iv []byte) (string, error) {
	encdata, err := AES_CBC_Encrypt(data, keysize, key, iv)
	if err != nil {
		return ``, err
	}
	ret := base64.StdEncoding.EncodeToString(encdata)
	return ret, nil
}

//aes cbc算法模式 解密
func AES_CBC_Decrypt(data []byte, keysize int, key, iv []byte) (ret []byte, err error) {
	defer func() { //错误处理
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	if err = aes_param_check(data, keysize, key, iv); err != nil {
		return nil, err
	}
	var cipherblock cipher.Block
	if cipherblock, err = aes.NewCipher(key); err != nil {
		return nil, err
	}
	ret = make([]byte, len(data))
	blockmode := cipher.NewCBCDecrypter(cipherblock, iv)
	blockmode.CryptBlocks(ret, data)
	ret = PKCS5UnPadding(ret)
	return ret, nil
}

//aes cbc算法模式 加密
func AES_CBC_Encrypt(data []byte, keysize int, key, iv []byte) (ret []byte, err error) {
	defer func() { //错误处理
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	if err = aes_param_check(data, keysize, key, iv); err != nil {
		return nil, err
	}
	var cipherblock cipher.Block
	if cipherblock, err = aes.NewCipher(key); err != nil {
		return nil, err
	}
	data = PKCS5Padding(data, cipherblock.BlockSize())
	ret = make([]byte, len(data))
	blockmode := cipher.NewCBCEncrypter(cipherblock, iv)
	blockmode.CryptBlocks(ret, data)
	return ret, nil
}

func aes_param_check(data []byte, keysize int, key, iv []byte) error {
	if len(iv) != 16 {
		return errors.New("IV向量长度必须为16个字节")
	}
	switch keysize {
	case 128, 192, 256:
		break
	default:
		return errors.New("密钥长度值只能为 128 192 256")
	}
	if len(key)*8 != keysize {
		return fmt.Errorf("密钥长度必须为 %d位", keysize)
	}
	return nil
}

func PKCS5UnPadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//将当前系统的时间转为字符串 精确到微秒
func CurrTime2Str() string {
	currTime := time.Now()
	hour, min, sec := currTime.Clock()
	microSec := currTime.UTC().Nanosecond() / 1000
	sNow := fmt.Sprintf("%0.2d:%0.2d:%0.2d.%0.6d", hour, min, sec, microSec)
	return sNow
}

//将当前系统的日期转为字符串
func CurrDate2Str() string {
	currTime := time.Now()
	year, month, day := currTime.Date()
	sNow := fmt.Sprintf("%0.4d-%0.2d-%0.2d", year, month, day)
	return sNow
}

//将昨天的日期转化为字符串
func YestodayDate2Str() string {
	currTime := time.Now()
	yestodayTime := currTime.AddDate(0, 0, -1)
	year, month, day := yestodayTime.Date()
	sNow := fmt.Sprintf("%0.4d-%0.2d-%0.2d", year, month, day)
	return sNow
}

//将当前系统的日期时间转为字符串 精确到微秒
func CurrDateTime2Str() string {
	currTime := time.Now()
	year, month, day := currTime.Date()
	hour, min, sec := currTime.Clock()
	microSec := currTime.UTC().Nanosecond() / 1000
	sNow := fmt.Sprintf("%0.4d-%0.2d-%0.2d %0.2d:%0.2d:%0.2d.%0.6d",
		year, month, day, hour, min, sec, microSec)
	return sNow
}

func CurrUnixTime() int64 {
	return time.Now().Unix()
}

func StrToDateTime(v string, location *time.Location) time.Time {
	tm, _ := time.ParseInLocation("2006-01-02 15:04:05", v, location)
	return tm
}

func Str2DateTime(v string) time.Time {
	return StrToDateTime(v, time.Local)
}

func Str2LocalDateTime(v string) time.Time {
	return StrToDateTime(v, time.Local)
}

func Str2UtcDateTime(v string) time.Time {
	loc, _ := time.LoadLocation("UTC")
	return StrToDateTime(v, loc)
}

//创建文件夹 文件夹存在或创建成功返回true 否则返回false
func ForceDirectories(v string) bool {
	if DirExist(v) {
		return true
	}
	if err := os.MkdirAll(v, 777); err != nil {
		return false
	}
	return true
}

//判断文件夹是否存在
func DirExist(v string) bool {
	fileinfo, err := os.Stat(v)
	if err != nil {
		return false
	} else {
		if !fileinfo.IsDir() {
			return false
		}
		return true
	}
}

//判断文件是否存在
func FileExist(v string) bool {
	fileinfo, err := os.Stat(v)
	if err != nil {
		return false
	} else {
		if fileinfo.IsDir() {
			return false
		}
		return true
	}
}

//将byte数据转为string 若byte数据中存在0x0 则只取0x0前的数据
func Bytes2String(v []byte) string {
	nIndex := bytes.IndexByte(v, 0x0)
	if nIndex < 0 {
		return string(v)
	}
	return string(v[:nIndex])
}

//获取当前运行程序所在的路径
func AppFilePath() string {
	var sFileName string
	var err error
	if sFileName, err = exec.LookPath(os.Args[0]); err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		return ``
	}
	if sFileName, err = filepath.Abs(sFileName); err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		return ``
	}
	sRet, _ := filepath.Split(sFileName)
	return sRet
}

//获取文件内容
func GetFileContent(filename string) ([]byte, error) {
	fptr, err := os.Open(filename)
	defer fptr.Close()
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(fptr)
	return data, err
}