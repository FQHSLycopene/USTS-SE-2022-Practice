package service

import (
	"BackEnd/define"
	"BackEnd/models"
	"github.com/gin-gonic/gin"
)

// GetPractiseList
// @Summary	获取练习列表
// @Tags	学生方法
// @Param	page query string false "page"
// @Param	pageSize query string false "pageSize"
// @Param	keyWord query string false "keyWord"
// @Param	token header string false "token"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/student/Practise [get]
func GetPractiseList(c *gin.Context) {
	page := c.DefaultQuery("page", define.DefaultPage)
	pageSize := c.DefaultQuery("pageSize", define.DefaultPageSize)
	keyWord := c.Query("keyWord")
	userIdentity, _ := c.Get("userIdentity")
	data, err := models.GetPractiseList(userIdentity.(string), page, pageSize, keyWord)
	if err != nil {
		c.JSON(200, define.Result{
			Code: 401,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(200, define.Result{
		Code: 200,
		Msg:  "success",
		Data: data,
	})
}
