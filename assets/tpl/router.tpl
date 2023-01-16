package routers

 import (
 	"github.com/gin-gonic/gin"
 	swaggerFiles "github.com/swaggo/files"
 	ginSwagger "github.com/swaggo/gin-swagger"
 	_ "{{.ProjectName}}/docs"
 	"{{.ProjectName}}/internal/api/users"
 )

 // InitRouter 初始化路由
 func InitRouter() *gin.Engine{
 	r := gin.Default()

 	// 示例路由
 	userController := users.UserController{}
 	r.GET("/user-info", userController.GetUserInfo)

 	// 注册swagger
 	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
 	return r
 }