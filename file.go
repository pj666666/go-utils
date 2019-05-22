package utils

import (
	"path"
	"strings"
)

//获取文件名、扩展名
func GetFilenameAndExt(s string) (fullname, filename, ext string) {
	fullname = path.Base(s)
	ext = path.Ext(fullname)
	filename = strings.TrimSuffix(fullname, ext)
	return
}
