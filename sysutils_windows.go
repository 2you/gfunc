package gfunc

import "syscall"

func FileRename(srcname, dstname string) error {
	from, err := syscall.UTF16PtrFromString(srcname)
	if err != nil {
		return err
	}
	to, err := syscall.UTF16PtrFromString(dstname)
	if err != nil {
		return err
	}
	return syscall.MoveFile(from, to)
}
