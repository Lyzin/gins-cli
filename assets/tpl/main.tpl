package main

import (
    "fmt"
    "{{.ProjectName}}/internal/routers"
    "{{.ProjectName}}/pkg/zaplogger"
)

func init() {
	// 设置日志器
	setUpZapLogger()

}

// @title {{.ProjectName}}项目
// @version 1.0
// @description {{.ProjectName}}的接口文档
// @contact.name admin
// @contact.email admin@xx.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 127.0.0.1:8090
// @BasePath /
func main() {
	// 启动服务
	Routers := routers.InitRouter()
	serverErr := Routers.Run(":8090")
	if serverErr != nil {
		panic(fmt.Errorf("start server err:%v\n", serverErr))
	}
}

func setUpZapLogger() {
	// 初始化日志器
	zaplogger.InitLogger()
	defer zaplogger.SugarLogger.Sync()
}