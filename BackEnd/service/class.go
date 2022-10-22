package service

import (
	"BackEnd/define"
	"BackEnd/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

type updateClassAccept struct {
	ClassIdentity     string   `json:"class_identity" binding:"required"`
	ClassName         string   `json:"class_name" binding:""`
	IsChangeCode      bool     `json:"is_change_code" binding:"required"`
	StudentIdentities []string `json:"student_identities" binding:""`
}

// UpdateClass
// @Summary	修改班级
// @Tags	老师方法
// @Param   param body updateClassAccept true "上传的JSON"
// @Param	token header string false "token"
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

// CreateClass
// @Summary	创建班级
// @Tags	老师方法
// @Param	name formData string true "name"
// @Param	token header string true "token"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/teacher/Class [post]
func CreateClass(c *gin.Context) {
	name := c.PostForm("name")
	userIdentity, exist := c.Get("userIdentity")
	if !exist {
		c.JSON(200, define.Result{
			Code: 402,
			Data: nil,
			Msg:  "用户没有登陆",
		})
		return
	}
	data, err := models.AddClass(name, userIdentity.(string))
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

// JoinClass
// @Summary	加入班级
// @Tags	学生方法
// @Param	joinCode formData string true "joinCode"
// @Param	token header string true "token"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/student/Class [put]
func JoinClass(c *gin.Context) {
	joinCode := c.PostForm("joinCode")
	userIdentity, exist := c.Get("userIdentity")
	if !exist {
		c.JSON(200, define.Result{
			Code: 402,
			Data: nil,
			Msg:  "用户没有登陆",
		})
		return
	}
	data, err := models.JoinClass(joinCode, userIdentity.(string))
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

// GetClassList
// @Summary	获取班级列表
// @Tags	公共方法
// @Param	page query string false "page"
// @Param	pageSize query string false "pageSize"
// @Param	keyWord query string false "keyWord"
// @Param	token header string true "token"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/Class [get]
func GetClassList(c *gin.Context) {
	page := c.DefaultQuery("page", define.DefaultPage)
	pageSize := c.DefaultQuery("page", define.DefaultPageSize)
	keyWord := c.Query("keyWord")
	userIdentity, exist := c.Get("userIdentity")
	fmt.Println(userIdentity)
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
