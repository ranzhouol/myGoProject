package main

import (
	"onlinePractice/models"
	"onlinePractice/router"
)

/*
	项目：基于Gin、Gorm、Vue实现的在线做题系统
	后端：golang、gin、gorm
	前端：vue、elementUI
	数据库：mysql、redis
	整合：swagger API封装
*/

func main() {
	models.Init()
	r := router.Router()

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
