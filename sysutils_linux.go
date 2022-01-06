package gfunc

func FileRename(srcname, dstname string) error {
	return os.Rename(srcname, dstname)
}
