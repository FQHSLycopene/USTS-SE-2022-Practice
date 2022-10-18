package router

import (
	"BackEnd/middleware"
	"BackEnd/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	userGroup := r.Group("")
	{
		userGroup.POST("/Login", service.Login)
		userGroup.POST("/SendCode", service.SendCode)
		userGroup.POST("/VerifyEmailCode", service.VerifyEmailCode)
		userGroup.POST("/Register", service.Register)
	}
	teacherGroup := r.Group("/teacher", middleware.AnalyseToken(), middleware.UserIsTeacher())
	{
		teacherGroup.POST("/Knowledge", service.AddKnowledge)
		teacherGroup.POST("/ProblemCategory", service.AddProblemCategory)
		teacherGroup.POST("/Achievement", service.AddAchievement)
		teacherGroup.GET("/ProblemCategory", service.GetProblemCategoryList)
	}
	studentGroup := r.Group("", middleware.AnalyseToken())
	{
		studentGroup.GET("/Knowledge", service.GetKnowledgeList)
		studentGroup.GET("/Achievement", service.GetAchievementList)
	}
	return r
}
