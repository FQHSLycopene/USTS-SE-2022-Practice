package service

import (
	"BackEnd/define"
	"BackEnd/models"
	"github.com/gin-gonic/gin"
)

// DeleteProblem
// @Summary	删除题目
// @Tags	老师方法
// @Param	json body deleteProblemAccept true "json"
// @Param	token header string true "token"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/teacher/Problem [delete]
func DeleteProblem(c *gin.Context) {

}

type deleteProblemAccept struct {
	Identity string `binding:"required" json:"identity"`
}

// UpdateProblem
// @Summary	修改题目
// @Tags	老师方法
// @Param	json body updateProblemAccept true "json"
// @Param	token header string true "token"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/teacher/Problem [put]
func UpdateProblem(c *gin.Context) {
	accept := updateProblemAccept{}
	err := c.ShouldBind(&accept)
	if err != nil {
		c.JSON(200, define.Result{
			Code: 401,
			Data: nil,
			Msg:  err.Error(),
		})
		return
	}
	data, err2 := models.UpdateProblem(accept.Identity, accept.Name, accept.Content, accept.Answer, accept.CategoryIdentity, accept.Score, accept.KnowledgeIdentities)
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

type updateProblemAccept struct {
	Identity            string   `binding:"required" json:"identity"`
	Name                string   `binding:"" json:"name"`
	Content             string   `binding:"" json:"content"`
	Answer              string   `binding:"" json:"answer"`
	CategoryIdentity    string   `binding:"" json:"categoryIdentity"`
	KnowledgeIdentities []string `binding:"" json:"knowledgeIdentities"`
	Score               int      `binding:"" json:"score"`
}

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
