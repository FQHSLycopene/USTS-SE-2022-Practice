package service

import (
	"BackEnd/define"
	"BackEnd/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

// UpdateClass
// @Summary	修改班级
// @Tags	老师方法
// @Param	classIdentity formData string false "classIdentity"
// @Param	name formData string false "name"
// @Param	isChangeCode formData bool false "isChangeCode"
// @Param	studentIdentity formData array false "踢出班级的学生studentIdentity"
// @Param param body string true "上传的JSON"
// @Param	token header string true "token"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/teacher/Class [put]
func UpdateClass(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		c.JSON(200, define.Result{
			Code: 403,
			Data: nil,
			Msg:  err.Error(),
		})
		return
	}
	var m map[string]interface{}
	// 反序列化
	_ = json.Unmarshal(data, &m)
	fmt.Println(m["studentIdentity"])
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
