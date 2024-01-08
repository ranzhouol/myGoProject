package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "onlinePractice/docs"
	"onlinePractice/service"
)

func Router() *gin.Engine {
	r := gin.Default()

	// swagger 配置
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// 配置路由规则
	r.GET("/ping", service.Ping)
	// 问题
	r.GET("/problem-list", service.GetProblemList)
	r.GET("/problem-detail", service.GetProblemDetail)

	// 用户
	r.GET("/user-detail", service.GetUserDetail)
	r.POST("/login", service.Login)

	// 提交记录
	r.GET("/submit-list", service.GetSubmitList)

	return r
}
