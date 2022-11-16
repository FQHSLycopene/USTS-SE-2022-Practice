package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type wrongProblem struct {
	ID        uint `gorm:"PRIMARY_KEY"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Identity  string         `gorm:"index;NOT NULL;Type:varchar(36);Column:identity" json:"identity"`
	Name      string         `gorm:"NOT NULL;unique;Type:varchar(36);Column:name" json:"name"`

	UserIdentity string `gorm:"Type:varchar(36);Column:user_identity" json:"user_identity"`

	ProblemIdentity string   `gorm:"Type:varchar(36);Column:problem_identity" json:"problem_identity"`
	Problem         *Problem `gorm:"foreignKey:ProblemIdentity;references:Identity"`
	WrongAnswer     string   `gorm:"NOT NULL;Type:text;Column:wrong_answer" json:"wrong_answer"`
}

var WrongProblem *wrongProblem

func (*wrongProblem) GetWrongProblemDetail(identity string) (interface{}, error) {
	data, err := getWrongProblemByIdentity(identity)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (*wrongProblem) GetWrongProblemList(pageStr, pageSizeStr, keyWord string) (interface{}, error) {
	var total int64 = 0
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return nil, err
	}
	pageSize, err2 := strconv.Atoi(pageSizeStr)
	if err2 != nil {
		return nil, err2
	}
	data := make([]*wrongProblem, 0)
	DB.Model(&data).
		Where("name like ?", "%"+keyWord+"%").Count(&total).
		Offset((page - 1) * pageSize).Limit(pageSize).
		Find(&data)
	return gin.H{
		"list":  data,
		"total": total,
	}, nil
}

func getWrongProblemByIdentity(identity string) (*wrongProblem, error) {
	data := wrongProblem{}
	err := DB.Model(&data).
		Preload("Problem").Preload("Problem.Knowledge").Preload("Problem.ProblemCategory").
		Where("identity = ?", identity).First(&data).Error

	if err != nil {
		return nil, err
	}
	return &data, nil
}
