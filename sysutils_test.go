package gfunc

import (
	"fmt"
	"testing"
)

func Test_ForceDirectories(t *testing.T) {
	b := ForceDirectories("d1/d2/d3/d4/d5")
	t.Log(b)
}

func Test_DirExist(t *testing.T) {
	dir1, dir2 := "mahonia", "mahonia1"
	b := DirExist(dir1)
	t.Logf("%s %v", dir1, b)
	b = DirExist(dir2)
	t.Logf("%s %v", dir2, b)
}

func Test_FileExist(t *testing.T) {
	file1, file2 := "tcfunc.go", "tcfunc1.go"
	b := FileExist(file1)
	t.Logf("%s %v", file1, b)
	b = FileExist(file2)
	t.Logf("%s %v", file2, b)
}

func Test_AppFilePath(t *testing.T) {
	str := AppFilePath()
	t.Log(str)
}

func Test_AppFileName(t *testing.T) {
	s := AppFileName()
	t.Log(s)
}

func Test_GetFileContent(t *testing.T) {
	if _, err := GetFileContent("gfunc.go"); err != nil {
		t.Error(err)
	}
}

func Test_FileRename(t *testing.T) {
	filename := fmt.Sprintf("%d", CurrUnixNanoTime())
	_ = WriteStr2File(filename, ``)
	if err := FileRename(filename, AppFilePath()+filename); err != nil {
		t.Error(err)
	}
}

func Test_RemoveFile(t *testing.T) {
	filename := fmt.Sprintf("%d", CurrUnixNanoTime())
	_ = WriteStr2File(filename, ``)
	if err := RemoveFile(filename); err != nil {
		t.Error(err)
	}
}

func Test_SearchFiles(t *testing.T) {
	if files, err := SearchFiles(AppFilePath(), nil, false); err != nil {
		t.Error(err)
	} else {
		t.Log(files)
	}

	if files, err := SearchFiles(AppFilePath(), nil, true); err != nil {
		t.Error(err)
	} else {
		t.Log(files)
	}

	if files, err := SearchFiles(`./`, []string{`.go`, `.txt`}, true); err != nil {
		t.Error(err)
	} else {
		t.Log(files)
	}
}

func Test_ExtractFileDir(t *testing.T) {
	s := ExtractFileDir(AppFileName())
	t.Log(s)
}

func Test_ExtractFileName(t *testing.T) {
	s := ExtractFileName(AppFileName())
	t.Log(s)
}

func Test_ExtractFileExt(t *testing.T) {
	s := ExtractFileExt(AppFileName())
	t.Log(s)
}
