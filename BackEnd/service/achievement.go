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
// @Param	Authorization header string true "Authorization"
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

// GetUserAchievementList
// @Summary	获取成就列表
// @Tags	学生方法
// @Param	Authorization header string true "Authorization"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/student/Achievement [Get]
func GetUserAchievementList(c *gin.Context) {
	userIdentity, _ := c.Get("userIdentity")
	data, err := models.GetUserAchievementList(userIdentity.(string))
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
