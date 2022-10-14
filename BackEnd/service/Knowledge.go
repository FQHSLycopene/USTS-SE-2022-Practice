package service

import (
	"BackEnd/define"
	"BackEnd/models"
	"github.com/gin-gonic/gin"
)

// AddKnowledge
// @Summary	添加知识点
// @Tags	管理员方法
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
