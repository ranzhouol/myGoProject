package service

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"onlinePractice/define"
	"onlinePractice/models"
	"strconv"
)

// GetProblemList
// @Tags 公共方法
// @Summary 问题列表
// @Param page query int false "page, 默认为1"
// @Param size query int false "size, 默认为20"
// @Param keyword query string flase "keyword"
// @Param category_identity query string flase "category_identity"
// @Description 获取问题列表
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /problem-list [get]
func GetProblemList(c *gin.Context) {
	// 获取参数,没有取默认值
	size, _ := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	page, _ := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	page = (page - 1) * size //分页查询的起始索引
	keyword := c.Query("keyword")
	categoryIdentity := c.Query("category_identity")

	// 分页查询
	dataList := make([]*models.ProblemBasic, 0)
	tx := models.GetProblemList(keyword, categoryIdentity)
	// count为数据总量，注意在分页查询之前
	var count int64
	err := tx.Count(&count).Offset(page).Limit(size).Find(&dataList).Error
	if err != nil {
		log.Println("Get problem list error: ", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
		"data": map[string]interface{}{
			"list":  dataList,
			"Count": count,
		},
	})
}

// GetProblemDetail
// @Tags 公共方法
// @Summary 问题详情
// @Param identity query string flase "problem identity"
// @Description 获取问题详情
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /problem-detail [get]
func GetProblemDetail(c *gin.Context) {
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "问题唯一标识不能为空",
		})
		return
	}

	data := new(models.ProblemBasic)
	err := models.DB.Where("identity = ?", identity).
		Preload("ProblemCategories").
		Preload("ProblemCategories.CategoryBasic").
		First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "问题不存在",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "GetProblemDetail Error: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
		"data": data,
	})
}
