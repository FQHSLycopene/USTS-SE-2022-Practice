package service

import (
	"BackEnd/define"
	"BackEnd/models"
	"github.com/gin-gonic/gin"
)

// AddAchievement
// @Summary	添加知识点成就
// @Tags	老师方法
// @Param	json body addAchievementAccept true "json"
// @Param	token header string true "token"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/teacher/Achievement [post]
func AddAchievement(c *gin.Context) {
	accept := addAchievementAccept{}
	err2 := c.ShouldBind(&accept)
	if err2 != nil {
		c.JSON(200, define.Result{
			Code: 401,
			Data: nil,
			Msg:  err2.Error(),
		})
		return
	}
	data, err := models.AddAchievement(accept.Name, accept.KnowledgeIdentity)
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

type addAchievementAccept struct {
	Name              string `binding:"required" json:"name"`
	KnowledgeIdentity string `binding:"required" json:"knowledge_identity"`
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
	pageSize := c.DefaultQuery("pageSize", define.DefaultPageSize)
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
