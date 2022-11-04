package service

import (
	"BackEnd/define"
	"github.com/gin-gonic/gin"
)

// AddProblem
// @Summary	创建题目
// @Tags	老师方法
// @Param	json body addProblemAccept true "json"
// @Param	token header string true "token"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/teacher/Class [post]
func AddProblem(c *gin.Context) {
	accept := addProblemAccept{}
	err := c.ShouldBind(&accept)
	if err != nil {
		c.JSON(200, define.Result{
			Code: 401,
			Data: nil,
			Msg:  err.Error(),
		})
		return
	}

}

type addProblemAccept struct {
}
