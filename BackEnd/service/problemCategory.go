package service

import (
	"BackEnd/define"
	"BackEnd/models"
	"github.com/gin-gonic/gin"
)

// DeleteProblemCategory
// @Summary	删除题目类型
// @Tags	老师方法
// @Param	json body deleteProblemCategoryAccept true "json"
// @Param	Authorization header string true "Authorization"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/teacher/ProblemCategory [delete]
func DeleteProblemCategory(c *gin.Context) {
	accept := deleteProblemCategoryAccept{}
	err2 := c.ShouldBind(&accept)
	if err2 != nil {
		c.JSON(200, define.Result{
			Code: 401,
			Data: nil,
			Msg:  err2.Error(),
		})
		return
	}
	data, err := models.DeleteProblemCategory(accept.Identity)
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

type deleteProblemCategoryAccept struct {
	Identity string `binding:"required" json:"identity"`
}

// AddProblemCategory
// @Summary	添加题目类型
// @Tags	老师方法
// @Param	json body addProblemCategoryAccept false "json"
// @Param	Authorization header string true "Authorization"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/teacher/ProblemCategory [post]
func AddProblemCategory(c *gin.Context) {
	accept := addProblemCategoryAccept{}
	err := c.ShouldBind(&accept)
	if err != nil {
		c.JSON(200, define.Result{
			Code: 401,
			Data: nil,
			Msg:  err.Error(),
		})
		return
	}
	data, err2 := models.AddProblemCategory(accept.Name)
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

type addProblemCategoryAccept struct {
	Name string `json:"name" binding:"required"`
}

// GetProblemCategoryList
// @Summary	获取题目类型列表
// @Tags	老师方法
// @Param	Authorization header string true "Authorization"
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
