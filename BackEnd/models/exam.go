package models

import (
	"BackEnd/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
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

func DeleteExamProblem(examIdentity string, problemIdentities []string) (interface{}, error) {
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
	}
	return nil, nil
}

func UpdateExam(examIdentity, name, duration string, startAt time.Time) (interface{}, error) {
	exam, err := getExamByIdentity(examIdentity)
	if err != nil {
		return nil, err
	}
	if exam.Publish == 1 {
		return nil, errors.New("考试已发布无法修改信息")
	}
	parseDuration, err2 := time.ParseDuration(duration)
	if err2 != nil {
		return nil, err2
	}
	exam.Name = name
	exam.Duration = parseDuration
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

func AddExam(classIdentity, name, duration string, startAt time.Time, problemIdentities []string) (interface{}, error) {
	parseDuration, err2 := time.ParseDuration(duration)
	if err2 != nil {
		return nil, err2
	}
	exam := Exam{
		Identity:      utils.GetUuid(),
		Name:          name,
		StartAt:       startAt,
		Duration:      parseDuration,
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
