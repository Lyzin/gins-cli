package dpath

import (
	"os"
	"path/filepath"
)

// ParentPath 递归返回上一级目录
// path 原始目录
// dth 需要返回上一级目录的深度，1表示往上返回一级，2表示往上返回2级
func ParentPath(path string, dth int) string {
	if dth == 0 {
		return path
	}
	ret := filepath.Dir(path)
	return ParentPath(ret, dth-1)
}

func CheckDirOrFileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}