package service

import (
	"BackEnd/define"
	"BackEnd/models"
	"github.com/gin-gonic/gin"
	"time"
)

// DeleteExam
// @Summary 删除考试
// @Tags	老师方法
// @Param	examIdentity query string true "examIdentity"
// @Param	Authorization header string true "Authorization"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/teacher/Exam [delete]
func DeleteExam(c *gin.Context) {
	examIdentity := c.Query("examIdentity")
	_, err := models.DeleteExam(examIdentity)
	if err != nil {
		c.JSON(200, define.Result{
			Code: 401,
			Data: nil,
			Msg:  err.Error(),
		})
		return
	}
	c.JSON(200, define.Result{
		Code: 200,
		Data: nil,
		Msg:  "success",
	})
}

// GetExamPaper
// @Summary	查看试卷
// @Tags	学生方法
// @Param	examIdentity query string true "examIdentity"
// @Param	Authorization header string true "Authorization"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/student/ExamPaper [get]
func GetExamPaper(c *gin.Context) {
	examIdentity := c.Query("examIdentity")
	userIdentity, _ := c.Get("userIdentity")
	data, err := models.GetExamPaper(userIdentity.(string), examIdentity)
	if err != nil {
		c.JSON(200, define.Result{
			Code: 401,
			Data: nil,
			Msg:  err.Error(),
		})
		return
	}
	c.JSON(200, define.Result{
		Code: 200,
		Data: data,
		Msg:  "success",
	})
}

// SaveExamPaper
// @Summary	保存试卷
// @Tags	学生方法
// @Param	json body upExamPaperAccept true "json"
// @Param	Authorization header string true "Authorization"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/student/SaveExamPaper [put]
func SaveExamPaper(c *gin.Context) {
	accept := upExamPaperAccept{}
	userIdentity, _ := c.Get("userIdentity")
	err := c.ShouldBind(&accept)
	if err != nil {
		c.JSON(200, define.Result{
			Code: 401,
			Data: nil,
			Msg:  err.Error(),
		})
		return
	}
	_, err2 := models.SaveExamPaper(userIdentity.(string), accept.ExamIdentity, accept.ExamPaperProblems)
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
		Data: nil,
		Msg:  "success",
	})
}

// UpExamPaper
// @Summary	提交试卷
// @Tags	学生方法
// @Param	json body upExamPaperAccept true "json"
// @Param	Authorization header string true "Authorization"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/student/ExamPaper [put]
func UpExamPaper(c *gin.Context) {
	accept := upExamPaperAccept{}
	userIdentity, _ := c.Get("userIdentity")
	err := c.ShouldBind(&accept)
	if err != nil {
		c.JSON(200, define.Result{
			Code: 401,
			Data: nil,
			Msg:  err.Error(),
		})
		return
	}
	_, err = models.UpExamPaper(userIdentity.(string), accept.ExamIdentity, accept.ExamPaperProblems)
	if err != nil {
		c.JSON(200, define.Result{
			Code: 401,
			Data: nil,
			Msg:  err.Error(),
		})
		return
	}
	c.JSON(200, define.Result{
		Code: 200,
		Data: nil,
		Msg:  "success",
	})
}

type upExamPaperAccept struct {
	ExamIdentity      string                     `binding:"required" json:"exam_identity"`
	ExamPaperProblems []*models.ExamPaperProblem `binding:"required" json:"exam_paper_problems"`
}

// GetStudentExamProblemList
// @Summary	获取学生考试题目列表
// @Tags	学生方法
// @Param	examIdentity query string true "examIdentity"
// @Param	Authorization header string true "Authorization"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/student/ExamProblem [get]
func GetStudentExamProblemList(c *gin.Context) {
	examIdentity := c.Query("examIdentity")
	userIdentity, _ := c.Get("userIdentity")
	data, err := models.GetStudentExamProblemList(userIdentity.(string), examIdentity)
	if err != nil {
		c.JSON(200, define.Result{
			Code: 401,
			Data: nil,
			Msg:  err.Error(),
		})
		return
	}
	c.JSON(200, define.Result{
		Code: 200,
		Data: data,
		Msg:  "success",
	})
}

// PublishExam
// @Summary	发布考试
// @Tags	老师方法
// @Param	Authorization header string true "Authorization"
// @Param   json body publishExamAccept true "json"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/teacher/PublishExam [put]
func PublishExam(c *gin.Context) {
	accept := publishExamAccept{}
	err := c.ShouldBind(&accept)
	if err != nil {
		c.JSON(200, define.Result{
			Code: 401,
			Data: nil,
			Msg:  err.Error(),
		})
		return
	}
	_, err = models.PublishExam(accept.ExamIdentity)
	if err != nil {
		c.JSON(200, define.Result{
			Code: 401,
			Data: nil,
			Msg:  err.Error(),
		})
		return
	}
	c.JSON(200, define.Result{
		Code: 200,
		Data: nil,
		Msg:  "success",
	})
}

type publishExamAccept struct {
	ExamIdentity string `binding:"required" json:"exam_identity"`
}

// GetExamDetail
// @Summary	获取考试详情
// @Tags	老师方法
// @Param	Authorization header string true "Authorization"
// @Param   identity path string true "examIdentity"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/teacher/Exam/{identity} [Get]
func GetExamDetail(c *gin.Context) {
	examIdentity := c.Param("identity")
	data, err := models.GetExamDetail(examIdentity)
	if err != nil {
		c.JSON(200, define.Result{
			Code: 401,
			Data: nil,
			Msg:  err.Error(),
		})
		return
	}
	c.JSON(200, define.Result{
		Code: 200,
		Data: data,
		Msg:  "success",
	})
}

// GetExamProblemList
// @Summary	获取考试题目列表
// @Tags	老师方法
// @Param	examIdentity query string true "examIdentity"
// @Param	page query string false "page"
// @Param	pageSize query string false "pageSize"
// @Param	keyWord query string false "keyWord"
// @Param	Authorization header string true "Authorization"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/teacher/ExamProblem [get]
func GetExamProblemList(c *gin.Context) {
	examIdentity := c.Query("examIdentity")
	page := c.DefaultQuery("page", define.DefaultPage)
	pageSize := c.DefaultQuery("pageSize", define.DefaultPageSize)
	keyWord := c.Query("keyWord")
	list, err := models.GetExamProblemList(examIdentity, page, pageSize, keyWord)
	if err != nil {
		c.JSON(200, define.Result{
			Code: 401,
			Data: nil,
			Msg:  err.Error(),
		})
		return
	}
	c.JSON(200, define.Result{
		Code: 200,
		Data: list,
		Msg:  "success",
	})

}

// AddExamProblem
// @Summary	添加考试题目
// @Tags	老师方法
// @Param	json body addExamProblemAccept true "json"
// @Param	Authorization header string true "Authorization"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/teacher/ExamProblem [post]
func AddExamProblem(c *gin.Context) {
	accept := addExamProblemAccept{}
	err := c.ShouldBind(&accept)
	if err != nil {
		c.JSON(200, define.Result{
			Code: 401,
			Data: nil,
			Msg:  err.Error(),
		})
		return
	}
	data, err2 := models.AddExamProblem(accept.ExamIdentity, accept.ProblemIdentities)
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

type addExamProblemAccept struct {
	ExamIdentity      string   `binding:"required" json:"exam_identity"`
	ProblemIdentities []string `binding:"required" json:"problem_identities"`
}

// DeleteExamProblem
// @Summary	删除考试题目
// @Tags	老师方法
// @Param	examIdentity query string true "examIdentity"
// @Param	problemIdentity query string true "problemIdentity"
// @Param	Authorization header string true "Authorization"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/teacher/ExamProblem [delete]
func DeleteExamProblem(c *gin.Context) {
	examIdentity := c.Query("examIdentity")
	problemIdentity := c.Query("problemIdentity")
	_, err := models.DeleteExamProblem(examIdentity, problemIdentity)
	if err != nil {
		c.JSON(200, define.Result{
			Code: 401,
			Data: nil,
			Msg:  err.Error(),
		})
		return
	}
	c.JSON(200, define.Result{
		Code: 200,
		Data: nil,
		Msg:  "success",
	})

}

// GetExamList
// @Summary	获取老师考试列表
// @Tags	老师方法
// @Param	classIdentity query string true "classIdentity"
// @Param	page query string false "page"
// @Param	pageSize query string false "pageSize"
// @Param	keyWord query string false "keyWord"
// @Param	Authorization header string true "Authorization"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/teacher/Exam [get]
func GetExamList(c *gin.Context) {
	page := c.DefaultQuery("page", define.DefaultPage)
	pageSize := c.DefaultQuery("pageSize", define.DefaultPageSize)
	keyWord := c.Query("keyWord")
	classIdentity := c.Query("classIdentity")

	data, err := models.GetExamList(classIdentity, page, pageSize, keyWord, 2)
	if err != nil {
		c.JSON(200, define.Result{
			Code: 401,
			Data: nil,
			Msg:  err.Error(),
		})
		return
	}
	c.JSON(200, define.Result{
		Code: 200,
		Data: data,
		Msg:  "success",
	})
}

// UpdateExam
// @Summary	更新考试信息
// @Tags	老师方法
// @Param json body updateExamAccept true "StartAt为考试开始时间（例：北京时间2022年12月30日中午12点30分，格式为"2022-12-30T12:30:00.000+08:00"），Duration为考试持续时间单位为秒"
// @Param	Authorization header string true "Authorization"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/teacher/Exam [put]
func UpdateExam(c *gin.Context) {
	accept := updateExamAccept{}
	err := c.ShouldBind(&accept)
	if err != nil {
		c.JSON(200, define.Result{
			Code: 401,
			Data: nil,
			Msg:  err.Error(),
		})
		return
	}
	data, err2 := models.UpdateExam(accept.ExamIdentity, accept.Name, time.Duration(accept.Duration), accept.StartAt)
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

type updateExamAccept struct {
	ExamIdentity string    `binding:"required" json:"exam_identity"`
	Name         string    `binding:"required" json:"name"`
	StartAt      time.Time `binding:"required" json:"StartAt"`
	Duration     int64     `binding:"required" json:"duration"`
}

// GetStudentExamList
// @Summary	获取学生考试列表
// @Tags	学生方法
// @Param	page query string false "page"
// @Param	pageSize query string false "pageSize"
// @Param	keyWord query string false "keyWord"
// @Param	Authorization header string true "Authorization"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/student/Exam [get]
func GetStudentExamList(c *gin.Context) {
	page := c.DefaultQuery("page", define.DefaultPage)
	pageSize := c.DefaultQuery("pageSize", define.DefaultPageSize)
	keyWord := c.Query("keyWord")
	userIdentity, exist := c.Get("userIdentity")
	if !exist {
		c.JSON(200, define.Result{
			Code: 402,
			Data: nil,
			Msg:  "用户没有登陆",
		})
		return
	}
	data, err := models.GetStudentExamList(userIdentity.(string), page, pageSize, keyWord)
	if err != nil {
		c.JSON(200, define.Result{
			Code: 401,
			Data: nil,
			Msg:  err.Error(),
		})
		return
	}
	c.JSON(200, define.Result{
		Code: 200,
		Data: data,
		Msg:  "success",
	})
}

// AddExam
// @Summary	添加考试
// @Tags	老师方法
// @Param json body addExamAccept true "StartAt为考试开始时间（例：北京时间2022年12月30日中午12点30分，格式为"2022-12-30T12:30:00.000+08:00"），Duration为考试持续时间单位为秒"
// @Param	Authorization header string true "Authorization"
// @Success	200  {string}  json{"code":"200","msg":"","data",""}
// @Router	/teacher/Exam [post]
func AddExam(c *gin.Context) {
	accept := addExamAccept{}
	err := c.ShouldBind(&accept)
	if err != nil {
		c.JSON(200, define.Result{
			Code: 401,
			Data: nil,
			Msg:  err.Error(),
		})
		return
	}
	data, err2 := models.AddExam(accept.ClassIdentity, accept.Name, time.Duration(accept.Duration), accept.StartAt, accept.ProblemIdentities)
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

type addExamAccept struct {
	ClassIdentity     string    `binding:"required" json:"class_identity"`
	Name              string    `binding:"required" json:"name"`
	StartAt           time.Time `binding:"required" json:"StartAt"`
	Duration          int64     `binding:"required" json:"duration"`
	ProblemIdentities []string  `binding:"required" json:"problem_identities"`
}
