/**
@File   : rpl_handler
@Date   : 2023/1/15 4:30 下午
@Author : lyzin
@Desc   :
**/

package utils

import (
	"io/fs"
	"os"
	"text/template"
)

type ProjectObj struct {
	ProjectName string
}

// ReadTplFileAndCreateFile 读取 *.tpl 文件，然后依据模板文件生成新的文件
// tplFilePath 模板文件，存在于assets/tpl目录下
// data 需要给模板文件中替换的变量
// newFileName 需要生成的新文件的绝对路径
func ReadTplFileAndCreateFile(assets fs.FS, tplFilePath, data, newFilePath string) {
	tmpl, err := template.ParseFS(assets, tplFilePath)
	if err != nil {
		FatalLog("解析模板文件错误:%v\n", err)
		return
	}
	// 利用给定数据渲染模板，并将结果写入w
	projectData := ProjectObj{
		ProjectName: data,
	}
	
	f, _ := os.OpenFile(newFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	err = tmpl.Execute(f, projectData)
	if err != nil {
		FatalLog("创建go相关文件错误, 错误文件目录:%v", newFilePath)
		FatalLog("创建go相关文件错误:%v\n", err)
		return
	}
	InfoLog("创建go相关文件成功:%v\n", newFilePath)
}