package models

import (
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
	ProblemNumber int            `gorm:"NOT NULL;Type:int;Column:problem_number" json:"problem_number"`
	Problems      []*Problem     `gorm:"many2many:exam_problems;foreignKey:Identity;joinForeignKey:ExamIdentity;References:Identity;joinReferences:ProblemIdentity"`
}

type ExamProblems struct {
	Identity        string `gorm:"index;NOT NULL;Type:varchar(36);Column:identity" json:"identity"`
	ExamIdentity    string `gorm:"NOT NULL;Type:varchar(36);Column:exam_identity" json:"exam_identity"`
	ProblemIdentity string `gorm:"NOT NULL;Type:varchar(36);Column:problem_identity" json:"problem_identity"`
	QuestionNum     int    `gorm:"NOT NULL;Type:int;Column:question_num" json:"question_num"` //题号
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
	data, err2 := GetExamList(class[0].Identity, pageStr, pageSizeStr, keyWord)
	if err2 != nil {
		return nil, err2
	}
	return data, nil
}

func GetExamList(classIdentity, pageStr, pageSizeStr, keyWord string) (interface{}, error) {
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
	err = DB.Model(&data).Where("class_identity = ?", classIdentity).
		Where("name like ?", "%"+keyWord+"%").
		Offset((page - 1) * pageSize).Limit(pageSize).
		Count(&total).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return gin.H{
		"total": total,
		"list":  data,
	}, nil
}
