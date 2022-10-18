package service

import (
	"BackEnd/define"
	"BackEnd/models"
	"github.com/gin-gonic/gin"
)

// AddAchievement
// @Summary	添加知识点成就
// @Tags	管理员方法
// @Param	name formData string false "name"
// @Param	knowledgeIdentity formData string false "knowledgeIdentity"
// @Param	token header string true "token"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/teacher/Achievement [post]
func AddAchievement(c *gin.Context) {
	name := c.PostForm("name")
	knowledgeIdentity := c.PostForm("knowledgeIdentity")
	data, err := models.AddAchievement(name, knowledgeIdentity)
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

// GetAchievementList
// @Summary	获取成就列表
// @Tags	公共方法
// @Param	token header string true "token"
// @Param	page query string false "page"
// @Param	pageSize query string false "pageSize"
// @Param	keyWord query string false "keyWord"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/Achievement [Get]
func GetAchievementList(c *gin.Context) {
	page := c.DefaultQuery("page", define.DefaultPage)
	pageSize := c.DefaultQuery("page", define.DefaultPageSize)
	keyWord := c.Query("keyWord")
	data, err := models.GetAchievementList(page, pageSize, keyWord)
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
