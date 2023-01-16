/**
@File   : generat_go_file
@Date   : 2023/1/15 5:42 下午
@Author : lyzin
@Desc   :
**/

package actions

import (
	"embed"
	"fmt"
	"gins_cli/utils"
	"os"
	"os/exec"
)

// Assets 声明一个embed fs类型的变量
var Assets embed.FS


func GenerateGoFile(projectPath, projectName string) {
	assetsPath := "assets/tpl"
	
	files, _ := Assets.ReadDir(assetsPath)

	rtfList := ReplaceTplFileList(projectName)
	for _, v := range files {
		for _, rtf := range rtfList {
			if v.Name() == rtf.tplFileName {
				tplFilePath := assetsPath + fmt.Sprintf("/%v", rtf.tplFileName)
				data := rtf.projectName
				newFilePath := projectPath + fmt.Sprintf("%v", rtf.newFilePath)
				utils.ReadTplFileAndCreateFile(Assets, tplFilePath, data, newFilePath)
			}
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
		utils.FatalLog("项目[ %v ]生成go.mod文件错误:%v\n", projectName, err)
	}
	utils.InfoLog("项目[ %v ]生成go.mod文件成功", projectName)
	
	// 同步项目的依赖
	_, err = exec.Command("go",  "mod", "tidy").Output()
	if err != nil {
		utils.FatalLog("项目[ %v ]依赖包同步错误:%v\n", projectName, err)
	}
	utils.InfoLog("项目[ %v ]依赖包同步成功", projectName)
	
	// 生成swagger
	_, err = exec.Command("swag",  "init").Output()
	if err != nil {
		utils.FatalLog("项目[ %v ]生成swagger错误:%v\n", projectName, err)
	}
	utils.InfoLog("项目[ %v ]生成swagger成功", projectName)
	
	// 生成air的启动toml文件
	_, err = exec.Command("air",  "init").Output()
	if err != nil {
		utils.FatalLog("项目[ %v ]生成air的toml配置文件错误:%v\n", projectName, err)
	}
	utils.InfoLog("项目[ %v ]生成air的toml配置文件成功", projectName)
}
