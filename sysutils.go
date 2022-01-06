package gfunc

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

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

//获取当前运行程序所在的路径
func AppFilePath() string {
	var sFileName string
	var err error
	if sFileName, err = exec.LookPath(os.Args[0]); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s", err.Error())
		return ``
	}
	if sFileName, err = filepath.Abs(sFileName); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s", err.Error())
		return ``
	}
	sRet, _ := filepath.Split(sFileName)
	return sRet
}

func AppFileName() string {
	return os.Args[0]
}

//获取文件内容
func GetFileContent(filename string) ([]byte, error) {
	fptr, err := os.Open(filename)
	defer func() {
		_ = fptr.Close()
	}()
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(fptr)
	return data, err
}

func GetFileContent2Str(filename string) (string, error) {
	buf, err := GetFileContent(filename)
	return string(buf), err
}

//获取文件的大小(单位是字节)
func GetFileSize(filename string) int64 {
	fileinfo, err := os.Stat(filename)
	if err != nil && !os.IsExist(err) {
		return 0
	}
	return fileinfo.Size()
}

func AppendStr2File(filename string, content string) error {
	return AppendBytes2File(filename, []byte(content))
}

func AppendBytes2File(filename string, content []byte) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()
	_, err = file.Write(content)
	return err
}

func WriteStr2File(filename string, content string) error {
	return WriteBytes2File(filename, []byte(content))
}

func WriteBytes2File(filename string, content []byte) error {
	return ioutil.WriteFile(filename, content, 0660)
}

func MoveFile(srcname, dstname string) error {
	return FileRename(srcname, dstname)
}

func RemoveFile(filename string) error {
	return os.Remove(filename)
}

//根据参数获取目录下的文件名
func SearchFiles(dirPath string, suffixs []string, containSubDir bool) (filenames []string, err error) {
	var finfos []fs.FileInfo
	if finfos, err = ioutil.ReadDir(dirPath); err != nil {
		return nil, err
	}
	pathSep := string(os.PathSeparator)
	for i := len(dirPath) - 1; i >= 0; i-- {
		if dirPath[i] != os.PathSeparator {
			dirPath = dirPath[0 : i+1]
			break
		}
	}

	for _, finfo := range finfos {
		if finfo.IsDir() {
			if containSubDir {
				tempLst, _ := SearchFiles(dirPath+pathSep+finfo.Name(), suffixs, containSubDir)
				filenames = append(filenames, tempLst...)
			}
		} else {
			fname := strings.ToLower(finfo.Name())
			if suffixs == nil {
				filenames = append(filenames, dirPath+pathSep+finfo.Name())
			} else {
				for _, suffix := range suffixs {
					if strings.HasSuffix(fname, strings.ToLower(suffix)) {
						filenames = append(filenames, dirPath+pathSep+finfo.Name())
						break
					}
				}
			}
		}
	}
	return filenames, nil
}

func ExtractFilePath(filename string) (ret string) {
	ret, _ = filepath.Split(filename)
	return ret
}

func ExtractFileName(filename string) (ret string) {
	_, ret = filepath.Split(filename)
	return ret
}

func ExtractFileExt(filename string) (ret string) {
	ret = filepath.Ext(filename)
	return ret
}

func ExtractFileDir(filename string) (ret string) {
	ret = ExtractFilePath(filename)
	if ret == `` || ret[len(ret)-1] != os.PathSeparator {
		return ret
	}
	return ret[:len(ret)-1]
}
