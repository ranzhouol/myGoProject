package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"onlinePractice/define"
	"onlinePractice/models"
	"strconv"
)

// GetSubmitDetail
// @Tags 公共方法
// @Summary 提交列表
// @Param page query int false "page, 默认为1"
// @Param size query int false "size, 默认为20"
// @Param problem_identity query string flase "problem_identity"
// @Param user_identity query string flase "user_identity"
// @Param status query int flase "status"
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /submit-list [get]
func GetSubmitList(c *gin.Context) {
	// 获取参数,没有取默认值
	size, _ := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	page, _ := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	page = (page - 1) * size //分页查询的起始索引
	var count int64
	problemIdentity := c.Query("problem_identity")
	userIdentity := c.Query("user_identity")
	status, _ := strconv.Atoi(c.Query("status"))

	tx := models.GetSubmitList(problemIdentity, userIdentity, status)
	data := make([]*models.SubmmitBasic, 0)
	// 分页查询
	err := tx.Count(&count).Offset(page).Limit(size).Find(&data).Error
	if err != nil {
		log.Println("GetSubmitList Error: ", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "GetSubmitList Error: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"list":  data,
			"count": count,
		},
	})
}
