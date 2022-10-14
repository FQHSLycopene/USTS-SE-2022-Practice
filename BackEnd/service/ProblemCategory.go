package service

import (
	"BackEnd/define"
	"BackEnd/models"
	"github.com/gin-gonic/gin"
)

// AddProblemCategory
// @Summary	添加题目类型
// @Tags	管理员方法
// @Param	name formData string false "name"
// @Param	token header string true "token"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/teacher/ProblemCategory [post]
func AddProblemCategory(c *gin.Context) {
	name := c.PostForm("name")
	data, err := models.AddProblemCategory(name)
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
