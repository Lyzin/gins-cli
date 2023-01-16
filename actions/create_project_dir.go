/**
@File   : create_project_dir
@Date   : 2023/1/15 5:08 下午
@Author : lyzin
@Desc   :
**/

package actions

import (
	"gins_cli/utils"
)

// projectDirList 项目目录名
var projectDirList = []string{
	"config",
	"docs",
	"global",
	"internal",
	"internal/api",
	"internal/api/users",
	"internal/dao",
	"internal/middleware",
	"internal/model",
	"internal/routers",
	"internal/service",
	"pkg",
	"pkg/response",
	"pkg/errcode",
	"pkg/path_handler",
	"pkg/zaplogger",
	"storage",
	"third_party",
}

func createGinsDir(projectName string) string {
	currentPath := utils.CurrentPath()
	utils.InfoLog("当前目录：%v", currentPath)
	utils.InfoLog("项目名：%v", projectName)
	
	// 创建项目父目录
	projectPath, _ := utils.JoinPathAndCreate(currentPath, projectName)
	
	for _, v := range projectDirList {
		path, err := utils.JoinPathAndCreate(projectPath, v)
		if err != nil {
			utils.FatalLog("创建目录错误:%v", err)
			continue
		}
		utils.InfoLog("创建目录成功:%v", path)
	}
	return projectPath
}

func CreateProject(args []string) {
	if len(args) == 0 {
		utils.FatalLog("%v", "未传入项目目录名称")
		return
	}
	if len(args) > 1 {
		utils.FatalLog("%v", "不支持传入多个项目目录名称")
		return
	}
	utils.InfoLog("项目目录名称：%v\n", args)
	projectName := args[0]
	
	// 创建项目目录
	projectPath := createGinsDir(projectName)
	
	// 创建gin框架相关的go文件
	GenerateGoFile(projectPath, projectName)
	
	// 创建go mod 管理文件以及自动同步项目依赖
	GenerateGoModeFile(projectPath, projectName)
}