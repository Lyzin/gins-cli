/**
@File   : replace_tpl_file
@Date   : 2023/1/16 4:59 下午
@Author : lyzin
@Desc   :
**/

package actions

// ReplaceTplFileMap 需要被替换的tpl文件结构体
type ReplaceTplFileMap struct {
	tplFileName string
	projectName string
	newFilePath string
}

func NewReplaceTplFileMap(tplFileName, projectName, newFilePath string) ReplaceTplFileMap {
	return ReplaceTplFileMap{
		tplFileName: tplFileName,
		projectName: projectName,
		newFilePath: newFilePath,
	}
}

// ReplaceTplFileList 生成需要替换的tpl文件列表
func ReplaceTplFileList(projectName string) []ReplaceTplFileMap {
	var rtfList []ReplaceTplFileMap
	rtfList = []ReplaceTplFileMap{
		NewReplaceTplFileMap("api.tpl", projectName, "/internal/api/users/user.go"),
		NewReplaceTplFileMap("base_response.tpl", projectName, "/pkg/response/base_response.go"),
		NewReplaceTplFileMap("common_errcode.tpl", projectName, "/pkg/errcode/common_errcode.go"),
		NewReplaceTplFileMap("errcode.tpl", projectName, "/pkg/errcode/errcode.go"),
		NewReplaceTplFileMap("global_response.tpl", projectName, "/global/const_response.go"),
		NewReplaceTplFileMap("main.tpl", projectName, "/main.go"),
		NewReplaceTplFileMap("Makefile.tpl", projectName,"/Makefile"),
		NewReplaceTplFileMap("path_operate.tpl", projectName,"/pkg/path_handler/path_operate.go"),
		NewReplaceTplFileMap("README.tpl", projectName, "/README.md"),
		NewReplaceTplFileMap("router.tpl", projectName, "/internal/routers/router.go"),
		NewReplaceTplFileMap("service.tpl", projectName, "/internal/service/service.go"),
		NewReplaceTplFileMap("zap_logger.tpl", projectName, "/pkg/zaplogger/zap_logger.go"),
	}
	return rtfList
}
