/**
@File   : generat_go_file
@Date   : 2023/1/15 5:42 下午
@Author : lyzin
@Desc   :
**/

package actions

import (
	"embed"
	"gins_cli/utils"
	"os"
	"os/exec"
)

// Assets 声明一个embed fs类型的变量
var Assets embed.FS


// ReplaceTplFileMap 需要被替换的tpl文件结构体
type ReplaceTplFileMap struct {

}

func GenerateGoFile(projectPath, projectName string) {
	assetsPath := "assets/tpl"
	
	files, _ := Assets.ReadDir(assetsPath)
	for _, v := range files {
		if v.Name() == "Makefile.tpl" {
			tplFilePath := assetsPath + "/Makefile.tpl"
			data := projectName
			newFilePath := projectPath + "/Makefile"
			utils.ReadTplFileAndCreateFile(Assets, tplFilePath, data, newFilePath)
		}
		if v.Name() == "main.tpl" {
			tplFilePath := assetsPath + "/main.tpl"
			data := projectName
			newFilePath := projectPath + "/main.go"
			utils.ReadTplFileAndCreateFile(Assets, tplFilePath, data, newFilePath)
		}
		if v.Name() == "api.tpl" {
			tplFilePath := assetsPath + "/api.tpl"
			data := projectName
			newFilePath := projectPath + "/internal/api/users/user.go"
			utils.ReadTplFileAndCreateFile(Assets, tplFilePath, data, newFilePath)
		}
		if v.Name() == "router.tpl" {
			tplFilePath := assetsPath + "/router.tpl"
			data := projectName
			newFilePath := projectPath + "/internal/routers/router.go"
			utils.ReadTplFileAndCreateFile(Assets, tplFilePath, data, newFilePath)
		}
		if v.Name() == "service.tpl" {
			tplFilePath := assetsPath + "/service.tpl"
			data := projectName
			newFilePath := projectPath + "/internal/service/service.go"
			utils.ReadTplFileAndCreateFile(Assets, tplFilePath, data, newFilePath)
		}
		if v.Name() == "global_response.tpl" {
			tplFilePath := assetsPath + "/global_response.tpl"
			data := projectName
			newFilePath := projectPath + "/global/const_response.go"
			utils.ReadTplFileAndCreateFile(Assets, tplFilePath, data, newFilePath)
		}
		if v.Name() == "base_response.tpl" {
			tplFilePath := assetsPath + "/base_response.tpl"
			data := projectName
			newFilePath := projectPath + "/pkg/response/base_response.go"
			utils.ReadTplFileAndCreateFile(Assets, tplFilePath, data, newFilePath)
		}
		if v.Name() == "errcode.tpl" {
			tplFilePath := assetsPath + "/errcode.tpl"
			data := projectName
			newFilePath := projectPath + "/pkg/errcode/errcode.go"
			utils.ReadTplFileAndCreateFile(Assets, tplFilePath, data, newFilePath)
		}
		if v.Name() == "common_errcode.tpl" {
			tplFilePath := assetsPath + "/common_errcode.tpl"
			data := projectName
			newFilePath := projectPath + "/pkg/errcode/common_errcode.go"
			utils.ReadTplFileAndCreateFile(Assets, tplFilePath, data, newFilePath)
		}
		if v.Name() == "README.tpl" {
			tplFilePath := assetsPath + "/README.tpl"
			data := projectName
			newFilePath := projectPath + "/README.md"
			utils.ReadTplFileAndCreateFile(Assets, tplFilePath, data, newFilePath)
		}
	}
}

func GenerateGoModeFile(projectPath, projectName string) {
	err := os.Chdir(projectPath)
	if err != nil {
		utils.FatalLog("切换目录失败:%v\n", projectPath)
		return
	}
	
	// 生成go mod文件
	_, err = exec.Command("go",  "mod", "init", projectName).Output()
	if err != nil {
		utils.FatalLog("项目[ %v ]生成go.mod文件错误:%v\n", err)
	}
	utils.InfoLog("项目[ %v ]生成go.mod文件成功", projectName)
	
	// 同步项目的依赖
	_, err = exec.Command("go",  "mod", "tidy").Output()
	if err != nil {
		utils.FatalLog("项目[ %v ]依赖包同步错误:%v\n", err)
	}
	utils.InfoLog("项目[ %v ]依赖包同步成功", projectName)
	
	// 生成swagger
	_, err = exec.Command("swag",  "init").Output()
	if err != nil {
		utils.FatalLog("项目[ %v ]生成swagger错误:%v\n", err)
	}
	utils.InfoLog("项目[ %v ]生成swagger成功", projectName)
	
	// 生成air的启动toml文件
	_, err = exec.Command("air",  "init").Output()
	if err != nil {
		utils.FatalLog("项目[ %v ]生成air的toml配置文件错误:%v\n", err)
	}
	utils.InfoLog("项目[ %v ]生成air的toml配置文件成功", projectName)
}
