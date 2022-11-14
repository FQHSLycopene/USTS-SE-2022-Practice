package middleware

import (
	"BackEnd/define"
	"github.com/gin-gonic/gin"
)

func IsGetAchievement() gin.HandlerFunc {
	return func(c *gin.Context) {
		userStatus, exists := c.Get("userStatus")
		if !exists {
			c.JSON(200, define.Result{
				Code: 402,
				Msg:  "用户未登录",
			})
			c.Abort()
			return
		}
		if userStatus != define.DefaultTeacherCode {
			c.JSON(200, define.Result{
				Code: 402,
				Msg:  "用户非管理员",
			})
			c.Abort()
			return
		}
	}
}
