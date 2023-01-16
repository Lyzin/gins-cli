/**
@File   : log_handler
@Date   : 2023/1/15 5:15 下午
@Author : lyzin
@Desc   :
**/

package utils

import (
	"log"
)

func InfoLog(format string, data interface{}) {
	log.Printf("\033[1;32m"+ format +"\033[0m", data)
}

func FatalLog(format string, data ...interface{}) {
	log.Fatalf("\033[1;31m"+ format +"\033[0m", data)
}