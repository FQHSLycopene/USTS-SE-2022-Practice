package service

import (
	"BackEnd/define"
	"BackEnd/models"
	"BackEnd/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

// UpdateAvatar
// @Summary	修改用户头像
// @Tags	公共方法
// @Accept multipart/form-data
// @Param file formData file true "file"
// @Param	Authorization header string true "Authorization"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/Avatar [post]
func UpdateAvatar(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(200, define.Result{
			Code: 401,
			Data: nil,
			Msg:  err.Error(),
		})
		return
	}
	log.Println(file.Filename)
	dst := fmt.Sprintf("./image/avatar/%s", file.Filename)
	// 上传文件到指定的目录
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		c.JSON(200, define.Result{
			Code: 401,
			Data: nil,
			Msg:  err.Error(),
		})
		return
	}
	userIdentity, _ := c.Get("userIdentity")
	_, err2 := models.UpdateAvatar(userIdentity.(string), dst)
	if err2 != nil {
		c.JSON(200, define.Result{
			Code: 401,
			Data: nil,
			Msg:  err.Error(),
		})
		return
	}
	c.JSON(200, define.Result{
		Code: 200,
		Data: nil,
		Msg:  "success",
	})
}

// UpdatePassword
// @Summary	修改用户密码
// @Tags	公共方法
// @Param	json body updatePasswordAccept true "json"
// @Param	Authorization header string true "Authorization"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/Password [post]
func UpdatePassword(c *gin.Context) {
	accept := updatePasswordAccept{}
	err2 := c.ShouldBind(&accept)
	if err2 != nil {
		c.JSON(200, define.Result{
			Code: 401,
			Data: nil,
			Msg:  err2.Error(),
		})
		return
	}
	userIdentity, _ := c.Get("userIdentity")
	data, err := models.UpdatePassword(userIdentity.(string), accept.OldPassword, accept.NewPassword)
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

type updatePasswordAccept struct {
	OldPassword string `binding:"required" json:"old_password"`
	NewPassword string `binding:"required" json:"new_password"`
}

// UpdateUser
// @Summary	修改用户信息
// @Tags	公共方法
// @Param	json body updateUserAccept true "json"
// @Param	Authorization header string true "Authorization"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/User [post]
func UpdateUser(c *gin.Context) {
	accept := updateUserAccept{}
	err2 := c.ShouldBind(&accept)
	if err2 != nil {
		c.JSON(200, define.Result{
			Code: 401,
			Data: nil,
			Msg:  err2.Error(),
		})
		return
	}
	userIdentity, _ := c.Get("userIdentity")
	data, err := models.UpdateUser(userIdentity.(string), accept.Name, accept.Email, accept.Phone)
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

type updateUserAccept struct {
	Name  string `binding:"required" json:"name"`
	Email string `binding:"required" json:"email"`
	Phone string `binding:"required" json:"phone"`
}

// GetUser
// @Summary	获取用户信息
// @Tags	公共方法
// @Param	Authorization header string true "Authorization"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/User [get]
func GetUser(c *gin.Context) {
	userIdentity, _ := c.Get("userIdentity")
	data, err := models.GetUser(userIdentity.(string))
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

// GetAvatar
// @Summary	获取头像
// @Tags	公共方法
// @Param	Authorization header string true "Authorization"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/Avatar [get]
func GetAvatar(c *gin.Context) {
	userIdentity, _ := c.Get("userIdentity")
	data, err := models.GetAvatarByIdentity(userIdentity.(string))
	if err != nil {
		c.JSON(200, define.Result{
			Code: 401,
			Data: nil,
			Msg:  err.Error(),
		})
		return
	}
	c.File("./" + data.(string))
}

// GetClassStudentList
// @Summary	获取班级学生列表
// @Tags	老师方法
// @Param	classIdentity query string true "classIdentity"
// @Param	page query string false "page"
// @Param	pageSize query string false "pageSize"
// @Param	keyWord query string false "keyWord"
// @Param	Authorization header string true "Authorization"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/teacher/ClassStudents [get]
func GetClassStudentList(c *gin.Context) {
	page := c.DefaultQuery("page", define.DefaultPage)
	pageSize := c.DefaultQuery("pageSize", define.DefaultPageSize)
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

// LoginByName
// @Summary	用户名登陆
// @Tags	公共方法
// @Param	json body loginAccept false "登陆信息"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/Login [post]
func LoginByName(c *gin.Context) {
	accept := loginAccept{}
	err := c.ShouldBind(&accept)
	if err != nil {
		c.JSON(200, define.Result{
			Code: 403,
			Msg:  err.Error(),
			Data: nil,
		})
	}
	data, err2 := models.LoginByName(accept.Name, accept.Password)
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
	Name     string `binding:"required" json:"name"`
	Password string `binding:"required" json:"password"`
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
	if !emailIsExist {
		c.JSON(200, gin.H{
			"code": 401,
			"data": nil,
			"msg":  errors.New("邮箱没有注册"),
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
// @Summary	邮箱验证码登陆
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
	data, err3 := models.GetTokenByEmail(accept.Email)
	if err3 != nil {
		c.JSON(200, define.Result{
			401,
			nil,
			err3.Error(),
		})
		return
	}
	c.JSON(200, define.Result{
		200,
		data,
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
// @Param 	json body registerAccept true "status以什么身份注册1为学生2为老师"
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
