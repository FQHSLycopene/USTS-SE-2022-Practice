package service

import (
	"BackEnd/define"
	"BackEnd/models"
	"BackEnd/utils"
	"errors"
	"github.com/gin-gonic/gin"
)

// GetClassStudentList
// @Summary	获取班级学生列表
// @Tags	老师方法
// @Param	classIdentity query string true "classIdentity"
// @Param	page query string false "page"
// @Param	pageSize query string false "pageSize"
// @Param	keyWord query string false "keyWord"
// @Param	token header string true "token"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/teacher/ClassStudents [get]
func GetClassStudentList(c *gin.Context) {
	page := c.DefaultQuery("page", define.DefaultPage)
	pageSize := c.DefaultQuery("page", define.DefaultPageSize)
	keyWord := c.Query("keyWord")
	classIdentity := c.Query("classIdentity")
	data, err := models.GetClassStudentList(classIdentity, page, pageSize, keyWord)
	if err != nil {
		c.JSON(200, define.Result{
			Code: 401,
			Data: nil,
			Msg:  err.Error(),
		})
		return
	}
	c.JSON(200, define.Result{
		Code: 200,
		Data: data,
		Msg:  "success",
	})
}

// Login
// @Summary	Login
// @Tags	公共方法
// @Param	json body loginAccept false "登陆信息"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/Login [post]
func Login(c *gin.Context) {
	accept := loginAccept{}
	err := c.ShouldBind(&accept)
	if err != nil {
		c.JSON(200, define.Result{
			Code: 403,
			Msg:  err.Error(),
			Data: nil,
		})
	}
	data, err2 := models.Login(accept.Name, accept.Password, accept.Email, accept.Phone)
	if err2 != nil {
		c.JSON(200, gin.H{
			"code": 401,
			"data": nil,
			"msg":  err2.Error(),
		})
		return
	}
	c.JSON(200, define.Result{
		Code: 200,
		Data: data,
		Msg:  "success",
	})
}

type loginAccept struct {
	Name     string `binding:"" json:"name"`
	Password string `binding:"required" json:"password"`
	Email    string `binding:"" json:"email"`
	Phone    string `binding:"" json:"phone"`
}

// SendCode
// @Summary	SendCode
// @Tags	公共方法
// @Param	json body sendCodeAccept false  "email"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/SendCode [post]
func SendCode(c *gin.Context) {
	accept := sendCodeAccept{}
	err2 := c.ShouldBind(&accept)
	if err2 != nil {
		c.JSON(200, gin.H{
			"code": 401,
			"data": nil,
			"msg":  err2.Error(),
		})
		return
	}
	emailIsExist, err := models.EmailIsExist(accept.Email)
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
	err = utils.SendCode(code, accept.Email)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 401,
			"data": nil,
			"msg":  err.Error(),
		})
		return
	}
	err = utils.RedisSet(accept.Email, code, define.DefaultEmailCodeDuration)
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

type sendCodeAccept struct {
	Email string `binding:"required" json:"email"`
}

// VerifyEmailCode
// @Summary	验证码是否真确
// @Tags	公共方法
// @Accept  json
// @Param 	code body verifyEmailCodeAccept true "此处email由SendCode得来"
// @Success	200  {string}  define.Result
// @Router	/VerifyEmailCode [post]
func VerifyEmailCode(c *gin.Context) {
	accept := verifyEmailCodeAccept{}
	err := c.ShouldBind(&accept)
	if err != nil {
		c.JSON(200, define.Result{
			401,
			nil,
			err.Error(),
		})
		return
	}
	codeT, err2 := utils.RedisGet(accept.Email)
	if err2 != nil {
		c.JSON(200, define.Result{
			401,
			nil,
			err2.Error(),
		})
		return
	}
	if codeT != accept.Code {
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

type verifyEmailCodeAccept struct {
	Email string `binding:"required" json:"email"`
	Code  string `binding:"required" json:"code"`
}

// Register
// @Summary	Register
// @Tags	公共方法
// @Param 	json body registerAccept true "此处email由VerifyEmailCode得来;status以什么身份注册1为学生2为老师"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/Register [post]
func Register(c *gin.Context) {
	accept := registerAccept{}
	err := c.ShouldBind(&accept)
	if accept.Status == "" {
		accept.Status = define.DefaultUserStatusStr
	}
	if err != nil {
		c.JSON(200, gin.H{
			"code": 401,
			"data": nil,
			"msg":  err.Error(),
		})
		return
	}
	data, err2 := models.AddUser(accept.Name, accept.Password, accept.Email, accept.Phone, accept.Status)
	if err2 != nil {
		c.JSON(200, gin.H{
			"code": 401,
			"data": nil,
			"msg":  err2.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": data,
		"msg":  "success",
	})
}

type registerAccept struct {
	Name     string `binding:"required" json:"name"`
	Password string `binding:"required" json:"password"`
	Email    string `binding:"required" json:"email"`
	Phone    string `binding:"" json:"phone"`
	Status   string `binding:"" json:"status" `
}
