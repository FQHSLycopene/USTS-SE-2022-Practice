package middleware

import (
	"BackEnd/define"
	"BackEnd/models"
	"BackEnd/utils"
	"github.com/gin-gonic/gin"
)

func AnalyseToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		analyseToken, err := utils.AnalyseToken(token)
		if err != nil {
			c.JSON(200, define.Result{
				Code: 402,
				Data: nil,
				Msg:  err.Error(),
			})
			c.Abort()
			return
		}
		_, err2 := models.GetUserByIdentity(analyseToken.Identity)
		if err2 != nil {
			c.JSON(200, define.Result{
				Code: 402,
				Data: nil,
				Msg:  err2.Error(),
			})
			c.Abort()
			return
		}
		c.Set("userIdentity", analyseToken.Identity)
		c.Set("userStatus", analyseToken.Status)
	}
}

func UserIsTeacher() gin.HandlerFunc {
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

func UserIsStudent() gin.HandlerFunc {
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
		if userStatus != define.DefaultStudentCode {
			c.JSON(200, define.Result{
				Code: 402,
				Msg:  "用户非学生",
			})
			c.Abort()
			return
		}
	}
}
