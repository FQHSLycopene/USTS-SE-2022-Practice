package service

import (
	"BackEnd/define"
	"BackEnd/models"
	"github.com/gin-gonic/gin"
)

// GetWrongProblemList
// @Summary	获取错题列表
// @Tags	学生方法
// @Param	token header string true "token"
// @Param	page query string false "page"
// @Param	pageSize query string false "pageSize"
// @Param	keyWord query string false "keyWord"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/student/WrongProblem [Get]
func GetWrongProblemList(c *gin.Context) {
	page := c.DefaultQuery("page", define.DefaultPage)
	pageSize := c.DefaultQuery("page", define.DefaultPageSize)
	keyWord := c.Query("keyWord")
	data, err := models.WrongProblem.GetWrongProblemList(page, pageSize, keyWord)
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

// GetWrongProblemDetail
// @Summary	获取错题详情
// @Tags	学生方法
// @Param	token header string true "token"
// @Param   identity path string true "wrongProblem_identity"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/student/WrongProblem/{identity} [Get]
func GetWrongProblemDetail(c *gin.Context) {
	wrongProblemIdentity := c.Param("identity")
	data, err := models.WrongProblem.GetWrongProblemDetail(wrongProblemIdentity)
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
