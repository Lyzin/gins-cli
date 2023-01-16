/**
@File   : paths
@Date   : 2023/1/15 4:02 下午
@Author : lyzin
@Desc   :
**/

package utils

import "os"

func CurrentPath() string {
	dir, _ := os.Getwd()
	return dir
}

func checkDirExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}

func JoinPathAndCreate(path string, subPath ...string) (string, error){
	for _, v := range subPath {
		path += "/" + v
	}
	
	// check new path exists
	ok := checkDirExists(path)
	if !ok {
		// if exists, creat new path and return
		_ = os.Mkdir(path, 0777)
		return path, nil
	}
	return path, nil
}

func CopyFileToPath() {

}
