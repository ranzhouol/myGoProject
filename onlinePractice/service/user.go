package service

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"onlinePractice/helper"
	"onlinePractice/models"
)

// GetUserDetail
// @Tags 公共方法
// @Summary 用户详情
// @Param identity query string true "user identity"
// @Description 获取用户详情
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /user-detail [get]
func GetUserDetail(c *gin.Context) {
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户唯一标识不能为空",
		})
		return
	}

	data := new(models.UserBasic)
	// Omit("password"): 省略密码字段
	err := models.DB.Omit("password").Where("identity = ?", identity).Find(&data).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get User detail By Identity " + identity + " Error: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
		"data": data,
	})
}

// Login
// @Tags 公共方法
// @Summary 用户登录
// @Param username formData string true "username"
// @Param password formData string true "password"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /login [post]
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "必填信息为空",
		})
		return
	}

	password = helper.GetMd5(password)

	data := new(models.UserBasic)
	if err := models.DB.Where("name = ? AND password = ?", username, password).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "用户名或密码错误",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get UserBasic Error:" + err.Error(),
		})
		return
	}

	token, err := helper.GenerateToken(data.Identity, data.Name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "GenerateToken Error:" + err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"date": map[string]interface{}{
			"token": token,
		},
	})
}
