package service

import (
	"BackEnd/define"
	"BackEnd/models"
	"github.com/gin-gonic/gin"
)

// DeleteProblemCategory
// @Summary	删除题目类型
// @Tags	管理员方法
// @Param	identity formData string true "identity"
// @Param	token header string true "token"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/teacher/ProblemCategory [delete]
func DeleteProblemCategory(c *gin.Context) {
	identity := c.PostForm("identity")
	data, err := models.DeleteProblemCategory(identity)
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

// GetProblemCategoryList
// @Summary	获取题目类型列表
// @Tags	管理员方法
// @Param	token header string true "token"
// @Param	page query string false "page"
// @Param	pageSize query string false "pageSize"
// @Param	keyWord query string false "keyWord"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/teacher/ProblemCategory [Get]
func GetProblemCategoryList(c *gin.Context) {
	page := c.DefaultQuery("page", define.DefaultPage)
	pageSize := c.DefaultQuery("page", define.DefaultPageSize)
	keyWord := c.Query("keyWord")
	data, err := models.GetProblemCategoryList(page, pageSize, keyWord)
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
