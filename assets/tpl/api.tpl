package users

import (
	"github.com/gin-gonic/gin"
	"{{.ProjectName}}/global"
	"{{.ProjectName}}/internal/service"
    "{{.ProjectName}}/pkg/zaplogger"
)

// UserController 控制器
type UserController struct {}


// GetUserInfo 获取用户信息
// @Summary 获取用户信息
// @Description 获取用户信息接口
// @Tags 用户接口
// @Accept application/json
// @Produce application/json
// @Router /user-info [get]
func (u UserController) GetUserInfo(c *gin.Context) {
    zaplogger.SugarLogger.Infof("c requests:%v", c.Request)
	data, _ := service.GetUserInfoService()
	global.BaseResponse.Success(c, data)
}