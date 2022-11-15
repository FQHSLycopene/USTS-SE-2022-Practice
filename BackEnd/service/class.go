package service

import (
	"BackEnd/define"
	"BackEnd/models"
	"github.com/gin-gonic/gin"
)

// GetClassDetail
// @Summary	获取班级详情
// @Tags	老师方法
// @Param   identity path string true "class_identity"
// @Param	Authorization header string false "Authorization"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/teacher/Class/{identity} [get]
func GetClassDetail(c *gin.Context) {
	classIdentity := c.Param("identity")
	data, err := models.GetClassDetail(classIdentity)
	if err != nil {
		c.JSON(200, define.Result{
			Code: 401,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(200, define.Result{
		Code: 200,
		Msg:  "success",
		Data: data,
	})
}

// UpdateClass
// @Summary	修改班级
// @Tags	老师方法
// @Param   json body updateClassAccept true "上传的JSON"
// @Param	Authorization header string false "Authorization"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/teacher/Class [put]
func UpdateClass(c *gin.Context) {
	data := updateClassAccept{}
	err := c.ShouldBind(&data)
	if err != nil {
		c.JSON(200, define.Result{
			Code: 403,
			Data: nil,
			Msg:  err.Error(),
		})
		return
	}
	result, err2 := models.UpdateClass(data.ClassIdentity, data.ClassName, data.IsChangeCode, data.StudentIdentities)
	if err2 != nil {
		c.JSON(200, define.Result{
			Code: 401,
			Data: nil,
			Msg:  err2.Error(),
		})
		return
	}
	c.JSON(200, define.Result{
		Code: 200,
		Data: result,
		Msg:  "success",
	})
}

type updateClassAccept struct {
	ClassIdentity     string   `json:"class_identity" binding:"required"`
	ClassName         string   `json:"class_name" binding:""`
	IsChangeCode      bool     `json:"is_change_code" binding:"required"`
	StudentIdentities []string `json:"student_identities" binding:""`
}

// CreateClass
// @Summary	创建班级
// @Tags	老师方法
// @Param	json body createClassAccept true "json"
// @Param	Authorization header string true "Authorization"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/teacher/Class [post]
func CreateClass(c *gin.Context) {
	accept := createClassAccept{}
	err2 := c.ShouldBind(&accept)
	if err2 != nil {
		c.JSON(200, define.Result{
			Code: 401,
			Data: nil,
			Msg:  err2.Error(),
		})
		return
	}
	userIdentity, exist := c.Get("userIdentity")
	if !exist {
		c.JSON(200, define.Result{
			Code: 402,
			Data: nil,
			Msg:  "用户没有登陆",
		})
		return
	}
	data, err := models.AddClass(accept.Name, userIdentity.(string))
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

type createClassAccept struct {
	Name string `binding:"required" json:"name"`
}

// JoinClass
// @Summary	加入班级
// @Tags	学生方法
// @Param	json body joinClassAccept true "json"
// @Param	Authorization header string true "Authorization"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/student/Class [put]
func JoinClass(c *gin.Context) {
	accept := joinClassAccept{}
	err2 := c.ShouldBind(&accept)
	if err2 != nil {
		c.JSON(200, define.Result{
			Code: 401,
			Data: nil,
			Msg:  err2.Error(),
		})
		return
	}
	userIdentity, exist := c.Get("userIdentity")
	if !exist {
		c.JSON(200, define.Result{
			Code: 402,
			Data: nil,
			Msg:  "用户没有登陆",
		})
		return
	}
	data, err := models.JoinClass(accept.JoinCode, userIdentity.(string))
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

type joinClassAccept struct {
	JoinCode string `binding:"required" json:"joinCode"`
}

// GetClassList
// @Summary	获取班级列表
// @Tags	公共方法
// @Param	page query string false "page"
// @Param	pageSize query string false "pageSize"
// @Param	keyWord query string false "keyWord"
// @Param	Authorization header string true "Authorization"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/Class [get]
func GetClassList(c *gin.Context) {
	page := c.DefaultQuery("page", define.DefaultPage)
	pageSize := c.DefaultQuery("page", define.DefaultPageSize)
	keyWord := c.Query("keyWord")
	userIdentity, exist := c.Get("userIdentity")
	if !exist {
		c.JSON(200, define.Result{
			Code: 402,
			Data: nil,
			Msg:  "用户没有登陆",
		})
		return
	}
	data, err := models.GetClassList(userIdentity.(string), page, pageSize, keyWord)
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
