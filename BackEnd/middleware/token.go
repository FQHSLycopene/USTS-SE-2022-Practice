package middleware

import (
	"BackEnd/define"
	"BackEnd/models"
	"BackEnd/utils"
	"github.com/gin-gonic/gin"
)

func AnalyseToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		analyseToken, err := utils.AnalyseToken(token)
		if err != nil {
			c.JSON(200, define.Result{
				Code: 401,
				Data: nil,
				Msg:  err.Error(),
			})
			return
		}
		_, err2 := models.GetUserByIdentity(analyseToken.Identity)
		if err2 != nil {
			c.JSON(200, define.Result{
				Code: 401,
				Data: nil,
				Msg:  err2.Error(),
			})
			return
		}
		c.Set("UserIdentity", analyseToken.Identity)
		c.Set("UserStatus", analyseToken.Status)
	}
}
