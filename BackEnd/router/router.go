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
		teacherGroup.PUT("/Knowledge", service.UpdateKnowledge)
		teacherGroup.DELETE("/Knowledge", service.DeleteKnowledge)

		teacherGroup.POST("/ProblemCategory", service.AddProblemCategory)
		teacherGroup.GET("/ProblemCategory", service.GetProblemCategoryList)
		teacherGroup.DELETE("/ProblemCategory", service.DeleteProblemCategory)

		teacherGroup.POST("/Achievement", service.AddAchievement)

		teacherGroup.POST("/Class", service.CreateClass)
		//teacherGroup.PUT("/Class", service.UpdateClass)
	}

	publicGroup := r.Group("", middleware.AnalyseToken())
	{
		publicGroup.GET("/Knowledge", service.GetKnowledgeList)
		publicGroup.GET("/Achievement", service.GetAchievementList)
	}

	studentGroup := r.Group("/student", middleware.AnalyseToken(), middleware.UserIsStudent())
	{
		studentGroup.PUT("/Class", service.JoinClass)
	}
	return r
}
