package middleware

import (
	"BackEnd/define"
	"BackEnd/models"
	"github.com/gin-gonic/gin"
)

func UserIsHasClass() gin.HandlerFunc {
	return func(c *gin.Context) {
		userIdentity, _ := c.Get("userIdentity")
		classISExist, err2 := models.UserIsHasClass(userIdentity.(string))
		if err2 != nil {
			c.JSON(200, define.Result{
				Code: 401,
				Data: nil,
				Msg:  err2.Error(),
			})
			c.Abort()
			return
		}
		if !classISExist {
			c.JSON(200, define.Result{
				Code: 200,
				Data: nil,
				Msg:  "该学生尚未加入班级",
			})
			c.Abort()
			return
		}
	}
}
