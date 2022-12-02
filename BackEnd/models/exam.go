package models

import (
	"BackEnd/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

type Exam struct {
	ID            uint `gorm:"PRIMARY_KEY"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	Identity      string         `gorm:"index;NOT NULL;Type:varchar(36);Column:identity" json:"identity"`
	Name          string         `gorm:"NOT NULL;Type:varchar(36);Column:name" json:"name"`
	StartAt       time.Time      `gorm:"NOT NULL;"`
	Duration      time.Duration  `gorm:"NOT NULL;"`
	ClassIdentity string         `gorm:"NOT NULL;Type:varchar(36);Column:class_identity" json:"class_identity"`
	TotalScore    int            `gorm:"NOT NULL;Type:int;Column:total_score" json:"total_score"`
	Publish       int            `gorm:"NOT NULL;Type:int;Column:publish;default:0" json:"publish"` //考试是否发布0为否1为是
	ProblemNumber int            `gorm:"NOT NULL;Type:int;Column:problem_number" json:"problem_number"`
	Problems      []*Problem     `gorm:"many2many:exam_problems;foreignKey:Identity;joinForeignKey:ExamIdentity;References:Identity;joinReferences:ProblemIdentity"`
}

type ExamProblems struct {
	Identity        string `gorm:"index;NOT NULL;Type:varchar(36);Column:identity" json:"identity"`
	ExamIdentity    string `gorm:"NOT NULL;Type:varchar(36);Column:exam_identity" json:"exam_identity"`
	ProblemIdentity string `gorm:"NOT NULL;Type:varchar(36);Column:problem_identity" json:"problem_identity"`
}

type ExamPaperProblem struct {
	ProblemIdentity string `binding:"required" json:"problem_identity"`
	Answer          string `binding:"required" json:"answer"`
}

func SaveExamPaper(userIdentity, examIdentity string, examPaperProblems []*ExamPaperProblem) (interface{}, error) {
	for _, examPaperProblem := range examPaperProblems {
		epp := ExamPaperProblems{}
		err := DB.Model(&epp).
			Where("user_identity = ?", userIdentity).
			Where("exam_identity = ?", examIdentity).
			Where("problem_identity = ?", examPaperProblem.ProblemIdentity).
			First(&epp).Error
		if err != nil {
			return nil, err
		}
		epp.MyAnswer = examPaperProblem.Answer
		err = DB.Where("user_identity = ?", userIdentity).
			Where("exam_identity = ?", examIdentity).
			Where("problem_identity = ?", examPaperProblem.ProblemIdentity).
			Save(&epp).Error
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func GetExamPaper(userIdentity, examIdentity string) (interface{}, error) {
	type data struct {
		Problem
		MyAnswer string
		MyScore  int
	}
	datas := make([]*data, 0)
	var total int64 = 0
	err := DB.Model(&Problem{}).Preload("ProblemCategory").
		Joins("left join exam_paper_problems on problems.identity = exam_paper_problems.problem_identity").
		Where("exam_paper_problems.user_identity = ?", userIdentity).
		Where("exam_paper_problems.exam_identity = ?", examIdentity).
		Count(&total).Find(&datas).Error
	if err != nil {
		return nil, err
	}
	for _, d := range datas {
		DB.Model(&ExamPaperProblems{}).Select("my_score,my_answer").
			Where("user_identity = ?", userIdentity).
			Where("exam_identity = ?", examIdentity).
			Where("problem_identity = ?", d.Problem.Identity).First(&d)
	}
	exam, err2 := getExamByIdentity(examIdentity)
	if err2 != nil {
		return nil, err2
	}
	score := 0
	for _, d := range datas {
		score += d.MyScore
	}
	return gin.H{
		"total":      total,
		"score":      score,
		"totalScore": exam.TotalScore,
		"list":       datas,
	}, nil

}

func UpExamPaper(userIdentity, examIdentity string, examPaperProblems []*ExamPaperProblem) (interface{}, error) {
	for _, examPaperProblem := range examPaperProblems {
		epp := ExamPaperProblems{}
		err := DB.Model(&epp).
			Where("user_identity = ?", userIdentity).
			Where("exam_identity = ?", examIdentity).
			Where("problem_identity = ?", examPaperProblem.ProblemIdentity).
			First(&epp).Error
		if err != nil {
			return nil, err
		}
		problem, err2 := getProblemByIdentity(examPaperProblem.ProblemIdentity)
		if err2 != nil {
			return nil, err2
		}
		epp.MyAnswer = examPaperProblem.Answer
		if problem.ProblemCategory.Name == "选择题" {
			if epp.MyAnswer == problem.Answer {
				epp.Status = 1
				epp.MyScore = problem.Score
			} else {
				epp.MyScore = 0
				epp.Status = 2
			}
		} else {
			eppsp := strings.Split(epp.MyAnswer, ";")
			anssp := strings.Split(problem.Answer, ";")
			for i, s := range eppsp {
				if anssp[i] == s {
					epp.MyScore += 2
				}
			}
			if epp.MyScore == problem.Score {
				epp.Status = 1
			} else {
				epp.Status = 2
			}
		}
		err = DB.Where("user_identity = ?", userIdentity).
			Where("exam_identity = ?", examIdentity).
			Where("problem_identity = ?", examPaperProblem.ProblemIdentity).
			Save(&epp).Error
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func GetStudentExamProblemList(userIdentity, examIdentity string) (interface{}, error) {
	type data struct {
		Problem
		MyAnswer string
		MyScore  int
	}
	datas := make([]*data, 0)
	var total int64 = 0
	err := DB.Model(&Problem{}).Preload("ProblemCategory").
		Joins("left join exam_paper_problems on problems.identity = exam_paper_problems.problem_identity").
		Where("exam_paper_problems.user_identity = ?", userIdentity).
		Where("exam_paper_problems.exam_identity = ?", examIdentity).
		Where("status = 0").Count(&total).
		Find(&datas).Error
	if err != nil {
		return nil, err
	}

	if total == 0 {
		return nil, errors.New("试卷已提交")
	}
	for _, d := range datas {
		DB.Model(&ExamPaperProblems{}).Select("my_score,my_answer").
			Where("user_identity = ?", userIdentity).
			Where("exam_identity = ?", examIdentity).
			Where("problem_identity = ?", d.Problem.Identity).First(&d)
	}
	return gin.H{
		"total": total,
		"list":  datas,
	}, nil
}

func PublishExam(identity string) (interface{}, error) {
	exam, err := getExamByIdentity(identity)
	if err != nil {
		return nil, err
	}
	if exam.Publish == 1 {
		return nil, errors.New("该考试已发布")
	}
	exam.Publish = 1

	problemIdentities := make([]string, 0)
	err = DB.Model(&ExamProblems{}).Where("exam_identity = ?", identity).
		Select("problem_identity").Find(&problemIdentities).Error
	if err != nil {
		return nil, err
	}

	_, err = InitExamPaperProblems(exam.ClassIdentity, exam.Identity, problemIdentities)
	if err != nil {
		return nil, err
	}

	err = DB.Save(exam).Error
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func GetExamDetail(identity string) (interface{}, error) {
	exam := Exam{}
	DB.Model(exam).
		Where("identity = ?", identity).First(&exam)
	return exam, nil
}

func GetExamProblemList(examIdentity, pageStr, pageSizeStr, keyWord string) (interface{}, error) {
	problemIdentities := make([]string, 0)
	err := DB.Model(&ExamProblems{}).Where("exam_identity = ?", examIdentity).
		Select("problem_identity").Find(&problemIdentities).Error
	if err != nil {
		return nil, err
	}
	problems := make([]*Problem, 0)
	var total int64 = 0
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return nil, err
	}
	pageSize, err2 := strconv.Atoi(pageSizeStr)
	if err2 != nil {
		return nil, err2
	}
	err = DB.Model(&ProblemModels).Preload("Knowledge").Preload("ProblemCategory").
		Where("identity in ?", problemIdentities).
		Where("name like ?", "%"+keyWord+"%").Count(&total).
		Offset((page - 1) * pageSize).Limit(pageSize).Find(&problems).Error
	if err != nil {
		return nil, err
	}
	return gin.H{
		"total": total,
		"list":  problems,
	}, nil
}

func AddExamProblem(examIdentity string, problemIdentities []string) (interface{}, error) {
	exam, err := getExamByIdentity(examIdentity)
	if err != nil {
		return nil, err
	}
	if exam.Publish == 1 {
		return nil, errors.New("考试已发布无法修改题目")
	}
	for _, problemIdentity := range problemIdentities {
		problem, err := getProblemByIdentity(problemIdentity)
		if err != nil {
			return nil, err
		}
		ep := ExamProblems{
			Identity:        utils.GetUuid(),
			ExamIdentity:    examIdentity,
			ProblemIdentity: problemIdentity,
		}
		err = DB.Create(&ep).Error
		if err != nil {
			return nil, err
		}
		exam.ProblemNumber += 1
		exam.TotalScore += problem.Score
		err = DB.Model(exam).Preload("Problems").Save(exam).Error
		if err != nil {
			return nil, err
		}
	}
	return exam, nil
}

func DeleteExamProblem(examIdentity, problemIdentity string) (interface{}, error) {
	exam, err := getExamByIdentity(examIdentity)
	if err != nil {
		return nil, err
	}
	if exam.Publish == 1 {
		return nil, errors.New("考试已发布无法修改题目")
	}
	problem, err2 := getProblemByIdentity(problemIdentity)
	if err2 != nil {
		return nil, err2
	}
	ep := ExamProblems{}
	err = DB.Model(&ep).Where("exam_identity = ?", examIdentity).
		Where("problem_identity = ?", problemIdentity).
		Delete(&ep).Error
	if err != nil {
		return nil, err
	}
	exam.ProblemNumber -= 1
	exam.TotalScore -= problem.Score
	err = DB.Save(exam).Error
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func UpdateExam(examIdentity, name string, duration time.Duration, startAt time.Time) (interface{}, error) {
	exam, err := getExamByIdentity(examIdentity)
	if err != nil {
		return nil, err
	}
	if exam.Publish == 1 {
		return nil, errors.New("考试已发布无法修改信息")
	}
	exam.Name = name
	exam.Duration = duration
	exam.StartAt = startAt
	err = DB.Save(exam).Error
	if err != nil {
		return nil, err
	}
	return exam, nil
}

func GetStudentExamList(userIdentity, pageStr, pageSizeStr, keyWord string) (interface{}, error) {
	user, err3 := GetUserByIdentity(userIdentity)
	if err3 != nil {
		return nil, err3
	}

	class := make([]*Class, 0)
	err := DB.Model(user).Association("Classes").Find(&class)
	if err != nil {
		return nil, err
	}
	data, err2 := GetExamList(class[0].Identity, pageStr, pageSizeStr, keyWord, 1)

	if err2 != nil {
		return nil, err2
	}
	return data, nil
}

func GetExamList(classIdentity, pageStr, pageSizeStr, keyWord string, status int) (interface{}, error) {

	data := make([]*Exam, 0)
	var total int64 = 0
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return nil, err
	}
	pageSize, err2 := strconv.Atoi(pageSizeStr)
	if err2 != nil {
		return nil, err2
	}
	if status == 1 {
		err = DB.Model(&data).
			Where("publish = ?", 1).
			Where("class_identity = ?", classIdentity).
			Where("name like ?", "%"+keyWord+"%").Count(&total).
			Offset((page - 1) * pageSize).Limit(pageSize).
			Find(&data).Error
		if err != nil {
			return nil, err
		}
	} else {
		err = DB.Model(&data).
			Where("class_identity = ?", classIdentity).
			Where("name like ?", "%"+keyWord+"%").Count(&total).
			Offset((page - 1) * pageSize).Limit(pageSize).
			Find(&data).Error
		if err != nil {
			return nil, err
		}
	}

	return gin.H{
		"total": total,
		"list":  data,
	}, nil
}

func AddExam(classIdentity, name string, duration time.Duration, startAt time.Time, problemIdentities []string) (interface{}, error) {
	exam := Exam{
		Identity:      utils.GetUuid(),
		Name:          name,
		StartAt:       startAt,
		Duration:      duration,
		ClassIdentity: classIdentity,
		TotalScore:    0,
		ProblemNumber: 0,
	}
	err := DB.Create(&exam).Error
	if err != nil {
		return nil, err
	}
	for _, identity := range problemIdentities {
		ep := &ExamProblems{
			Identity:        identity,
			ExamIdentity:    exam.Identity,
			ProblemIdentity: identity,
		}
		err := DB.Create(&ep).Error
		if err != nil {
			return nil, err
		}
		problem, err2 := getProblemByIdentity(identity)
		if err2 != nil {
			return nil, err2
		}
		exam.ProblemNumber += 1
		exam.TotalScore += problem.Score
	}
	err = DB.Save(&exam).Error
	if err != nil {
		return nil, err
	}

	return nil, err
}

func getExamByIdentity(identity string) (*Exam, error) {
	exam := Exam{}
	err := DB.Model(&exam).Where("identity = ?", identity).First(&exam).Error
	if err != nil {
		return nil, err
	}
	return &exam, nil
}
