package service

import (
	"BackEnd/define"
	"BackEnd/models"
	"BackEnd/utils"
	"errors"
	"github.com/gin-gonic/gin"
)

// Login
// @Summary	Login
// @Tags	公共方法
// @Param	name formData string false "name"
// @Param	password formData string true  "password"
// @Param	email formData string false  "email"
// @Param 	phone formData string false "phone"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/Login [post]
func Login(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	email := c.PostForm("email")
	phone := c.PostForm("phone")

	data, err := models.Login(name, password, email, phone)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 401,
			"data": nil,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(200, define.Result{
		Code: 200,
		Data: data,
		Msg:  "success",
	})
}

// SendCode
// @Summary	SendCode
// @Tags	公共方法
// @Param	email formData string true  "email"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/SendCode [post]
func SendCode(c *gin.Context) {
	email := c.PostForm("email")
	emailIsExist, err := models.EmailIsExist(email)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 401,
			"data": nil,
			"msg":  err.Error(),
		})
		return
	}
	if emailIsExist {
		c.JSON(200, gin.H{
			"code": 401,
			"data": nil,
			"msg":  errors.New("邮箱已存在"),
		})
		return
	}
	code := utils.GetCode()
	err = utils.SendCode(code, email)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 401,
			"data": nil,
			"msg":  err.Error(),
		})
		return
	}
	err = utils.RedisSet(email, code, define.DefaultEmailCodeDuration)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 401,
			"data": nil,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": nil,
		"msg":  "success",
	})
}

// VerifyEmailCode
// @Summary	验证码是否真确
// @Tags	公共方法
// @Param	email formData string true  "此处email由SendCode得来"
// @Param 	code formData string true "code"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/VerifyEmailCode [post]
func VerifyEmailCode(c *gin.Context) {
	email := c.PostForm("email")
	code := c.PostForm("code")
	codeT, err := utils.RedisGet(email)
	if err != nil {
		c.JSON(200, define.Result{
			401,
			nil,
			err.Error(),
		})
		return
	}
	if codeT != code {
		c.JSON(200, define.Result{
			401,
			nil,
			"验证码错误",
		})
		return
	}
	c.JSON(200, define.Result{
		200,
		nil,
		"success",
	})
}

// Register
// @Summary	Register
// @Tags	公共方法
// @Param	name formData string true "name"
// @Param	password formData string true  "password"
// @Param	email formData string true  "此处email由VerifyEmailCode得来"
// @Param 	phone formData string false "phone"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/Register [post]
func Register(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	email := c.PostForm("email")
	phone := c.PostForm("phone")
	data, err := models.AddUser(name, password, email, phone)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 401,
			"data": nil,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": data,
		"msg":  "success",
	})
}
