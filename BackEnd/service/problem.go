package service

import (
	"BackEnd/define"
	"BackEnd/models"
	"github.com/gin-gonic/gin"
)

// AddProblem
// @Summary	创建题目
// @Tags	老师方法
// @Param	json body addProblemAccept true "json"
// @Param	token header string true "token"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/teacher/Problem [post]
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
	data, err2 := models.AddProblem(accept.Name, accept.Content, accept.Answer, accept.CategoryIdentity, accept.Score, accept.KnowledgeIdentities)
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

type addProblemAccept struct {
	Name                string   `binding:"required" json:"name"`
	Content             string   `binding:"required" json:"content"`
	Answer              string   `binding:"required" json:"answer"`
	CategoryIdentity    string   `binding:"required" json:"categoryIdentity"`
	KnowledgeIdentities []string `binding:"required" json:"knowledgeIdentities"`
	Score               int      `binding:"required" json:"score"`
}
