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
		userGroup.POST("/Login", service.LoginByName)
		userGroup.POST("/SendCode", service.SendCode)
		userGroup.POST("/VerifyEmailCode", service.VerifyEmailCode)
		userGroup.POST("/Register", service.Register)
		userGroup.GET("/image", service.GetImage)
		userGroup.POST("/BodyJSONTest", service.BodyJSONTest)
		userGroup.POST("/PostArrayTest", service.PostArrayTest)
	}
	teacherGroup := r.Group("/teacher", middleware.AnalyseToken(), middleware.UserIsTeacher())
	{
		teacherGroup.GET("/Knowledge", service.GetKnowledgeList)
		teacherGroup.POST("/Knowledge", service.AddKnowledge)
		teacherGroup.PUT("/Knowledge", service.UpdateKnowledge)
		teacherGroup.DELETE("/Knowledge", service.DeleteKnowledge)

		teacherGroup.POST("/ProblemCategory", service.AddProblemCategory)
		teacherGroup.GET("/ProblemCategory", service.GetProblemCategoryList)
		teacherGroup.DELETE("/ProblemCategory", service.DeleteProblemCategory)

		teacherGroup.POST("/Achievement", service.AddAchievement)

		teacherGroup.POST("/Class", service.CreateClass)
		teacherGroup.PUT("/Class", service.UpdateClass)
		teacherGroup.GET("/Class/:identity", service.GetClassDetail)

		teacherGroup.GET("/ClassStudents", service.GetClassStudentList)

		teacherGroup.POST("/Problem", service.AddProblem)
		teacherGroup.PUT("/Problem", service.UpdateProblem)
		teacherGroup.DELETE("/Problem", service.DeleteProblem)

	}

	publicGroup := r.Group("", middleware.AnalyseToken())
	{

		publicGroup.GET("/Achievement", service.GetAchievementList)
		publicGroup.GET("/Class", service.GetClassList)
		publicGroup.GET("/Avatar", service.GetAvatar)
		publicGroup.POST("/Avatar", service.UpdateAvatar)
		publicGroup.GET("/User", service.GetUser)
		publicGroup.POST("/User", service.UpdateUser)
		publicGroup.POST("/Password", service.UpdatePassword)
	}

	studentGroup := r.Group("/student", middleware.AnalyseToken(), middleware.UserIsStudent())
	{
		studentGroup.PUT("/Class", service.JoinClass)
		studentGroup.GET("/Practise", middleware.IsGetAchievement(), service.GetPractiseList)
		studentGroup.GET("/PractiseProblemDetail", service.GetRandPractiseProblemDetail)

		studentGroup.GET("/WrongProblem", service.GetWrongProblemList)
		studentGroup.GET("/WrongProblem/:identity", service.GetWrongProblemDetail)

		studentGroup.PUT("/UpProblemAnswer", service.Problem.UpProblemAnswer)
	}
	return r
}
