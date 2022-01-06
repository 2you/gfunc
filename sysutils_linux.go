package gfunc

import "os"

func FileRename(srcname, dstname string) error {
	return os.Rename(srcname, dstname)
}
