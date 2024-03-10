package service

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"onlinePractice/define"
	"onlinePractice/helper"
	"onlinePractice/models"
	"strconv"
	"time"
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

// SendCode
// @Tags 公共方法
// @Summary 发送验证码
// @Param email formData string true "email"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /send-code [post]
func SendCode(c *gin.Context) {
	email := c.PostForm("email")
	if email == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不正确",
		})

		return
	}

	code := helper.GetRandom()
	// 存放验证码
	models.RedisServer.Set(c, email, code, time.Second*300)
	if err := helper.SendCode(email, code); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Send code err:" + err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "验证码发送成功",
	})
}

// Register
// @Tags 公共方法
// @Summary 用户注册
// @Param email formData string true "email"
// @Param code formData string true "code"
// @Param name formData string true "name"
// @Param password formData string true "password"
// @Param phone formData string false "phone"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /register [post]
func Register(c *gin.Context) {
	email := c.PostForm("email")
	code := c.PostForm("code")
	name := c.PostForm("name")
	password := c.PostForm("password")
	phone := c.PostForm("phone")
	if email == "" || code == "" || name == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不正确",
		})
		return
	}

	// 验证验证码是否正确, 验证码存在redis
	sysCode, err := models.RedisServer.Get(c, email).Result()
	if err != nil {
		log.Printf("get code Error:%v\n", err.Error())

		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "请重新获取验证码: ",
		})
		return
	}

	if sysCode != code {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "验证码不正确",
		})
		return
	}

	// 判断邮箱是否已存在
	var count int64
	err = models.DB.Model(&models.UserBasic{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get User Error: " + err.Error(),
		})
		return
	}
	if count > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "该邮箱已被注册",
		})
		return
	}

	// 数据插入
	data := &models.UserBasic{
		Identity: helper.GetUUid(),
		Name:     name,
		Password: helper.GetMd5(password),
		Phone:    phone,
		Email:    email,
	}

	if err := models.DB.Create(data).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "create user Error: " + err.Error(),
		})
		return
	}

	// 生成token
	token, err := helper.GenerateToken(data.Identity, name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "generate token Error: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"token": token,
		},
	})
}

// GetRankList
// @Tags 公共方法
// @Summary 用户排行榜
// @Param page query int false "page, 默认为1"
// @Param size query int false "size, 默认为20"
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /rank-list [get]
func GetRankList(c *gin.Context) {
	// 获取参数,没有取默认值
	size, _ := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	page, _ := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	page = (page - 1) * size //分页查询的起始索引

	var count int64
	dataList := make([]*models.UserBasic, 0)
	// 按照完成问题个数降序排序，若相同，按照提交次数升序（即完成问题越多、提交次数越少，排名越高）
	err := models.DB.Model(&models.UserBasic{}).Count(&count).Order("finish_problem_num DESC, submit_num ASC").
		Offset(page).Limit(size).First(&dataList).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "get rank list Error:" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"list":  dataList,
			"Count": count,
		},
	})
}
