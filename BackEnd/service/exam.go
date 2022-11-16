package service

import (
	"BackEnd/define"
	"BackEnd/models"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

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
	data, err2 := models.UpdateExam(accept.ExamIdentity, accept.Name, strconv.Itoa(accept.Duration)+"s", accept.StartAt)
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
	Duration     int       `binding:"required" json:"duration"`
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
	pageSize := c.DefaultQuery("page", define.DefaultPageSize)
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
	data, err2 := models.AddExam(accept.ClassIdentity, accept.Name, strconv.Itoa(accept.Duration)+"s", accept.StartAt, accept.ProblemIdentities)
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
	Duration          int       `binding:"required" json:"duration"`
	ProblemIdentities []string  `binding:"required" json:"problem_identities"`
}
