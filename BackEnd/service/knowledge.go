package service

import (
	"BackEnd/define"
	"BackEnd/models"
	"github.com/gin-gonic/gin"
)

// AddKnowledge
// @Summary	添加知识点
// @Tags	老师方法
// @Param	json body addKnowledgeAccept true "json"
// @Param	token header string true "token"
// @Success	200  {string}  define.Result
// @Router	/teacher/Knowledge [post]
func AddKnowledge(c *gin.Context) {
	accept := addKnowledgeAccept{}
	err := c.ShouldBind(&accept)
	if err != nil {
		c.JSON(200, define.Result{
			Code: 401,
			Data: nil,
			Msg:  err.Error(),
		})
		return
	}
	data, err2 := models.AddKnowledge(accept.Name)
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
		Data: data,
		Msg:  "success",
	})
}

type addKnowledgeAccept struct {
	Name string `json:"name" binding:"required"`
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
// @Param	json body updateKnowledgeAccept true "json"
// @Param	token header string true "token"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/teacher/Knowledge [put]
func UpdateKnowledge(c *gin.Context) {
	accept := updateKnowledgeAccept{}
	err := c.ShouldBind(&accept)
	if err != nil {
		c.JSON(200, define.Result{
			Code: 401,
			Data: nil,
			Msg:  err.Error(),
		})
		return
	}
	data, err2 := models.UpdateKnowledge(accept.Name, accept.Identity)
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
		Data: data,
		Msg:  "success",
	})
}

type updateKnowledgeAccept struct {
	Name     string `json:"name" binding:"required"`
	Identity string `json:"identity" binding:"required"`
}

// DeleteKnowledge
// @Summary	删除知识点
// @Tags	老师方法
// @Param	json body deleteKnowledgeAccept true "json"
// @Param	token header string true "token"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/teacher/Knowledge [delete]
func DeleteKnowledge(c *gin.Context) {
	accept := deleteKnowledgeAccept{}
	err := c.ShouldBind(&accept)
	if err != nil {
		c.JSON(200, define.Result{
			Code: 401,
			Data: nil,
			Msg:  err.Error(),
		})
		return
	}
	data, err2 := models.DeleteKnowledge(accept.Identity)
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
		Data: data,
		Msg:  "success",
	})
}

type deleteKnowledgeAccept struct {
	Identity string `json:"identity" binding:"required"`
}
