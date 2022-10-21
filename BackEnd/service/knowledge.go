package service

import (
	"BackEnd/define"
	"BackEnd/models"
	"github.com/gin-gonic/gin"
)

// AddKnowledge
// @Summary	添加知识点
// @Tags	老师方法
// @Param	name formData string false "name"
// @Param	token header string true "token"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/teacher/Knowledge [post]
func AddKnowledge(c *gin.Context) {
	name := c.PostForm("name")
	data, err := models.AddKnowledge(name)
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

// GetKnowledgeList
// @Summary	获取知识点列表
// @Tags	公共方法
// @Param	token header string true "token"
// @Param	page query string false "page"
// @Param	pageSize query string false "pageSize"
// @Param	keyWord query string false "keyWord"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/Knowledge [Get]
func GetKnowledgeList(c *gin.Context) {
	page := c.DefaultQuery("page", define.DefaultPage)
	pageSize := c.DefaultQuery("page", define.DefaultPageSize)
	keyWord := c.Query("keyWord")
	data, err := models.GetKnowledgeList(page, pageSize, keyWord)
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

// UpdateKnowledge
// @Summary	修改知识点
// @Tags	老师方法
// @Param	name formData string true "name"
// @Param	identity formData string true "identity"
// @Param	token header string true "token"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/teacher/Knowledge [put]
func UpdateKnowledge(c *gin.Context) {
	name := c.PostForm("name")
	identity := c.PostForm("identity")
	data, err := models.UpdateKnowledge(name, identity)
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

// DeleteKnowledge
// @Summary	删除知识点
// @Tags	老师方法
// @Param	identity formData string true "identity"
// @Param	token header string true "token"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/teacher/Knowledge [delete]
func DeleteKnowledge(c *gin.Context) {
	identity := c.PostForm("identity")
	data, err := models.DeleteKnowledge(identity)
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
