package models

import (
	"gorm.io/gorm"
	"time"
)

type wrongProblem struct {
	ID              uint `gorm:"PRIMARY_KEY"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	Identity        string         `gorm:"index;NOT NULL;Type:varchar(36);Column:identity" json:"identity"`
	Name            string         `gorm:"NOT NULL;unique;Type:varchar(36);Column:name" json:"name"`
	UserIdentity    string         `gorm:"Type:varchar(36);Column:user_identity" json:"user_identity"`
	ProblemIdentity string         `gorm:"Type:varchar(36);Column:problem_identity" json:"problem_identity"`
	Problem         *Problem       `gorm:"foreignKey:ProblemIdentity;references:Identity"`
	WrongAnswer     string         `gorm:"NOT NULL;Type:text;Column:wrong_answer" json:"wrong_answer"`
}

var WrongProblem *wrongProblem

func (*wrongProblem) GetWrongProblemList(pageStr, pageSizeStr, keyword string) (interface{}, error) {
	data := make([]*wrongProblem, 0)
	DB.Model(data).Preload("Problem").Joins("right join problem p on p.identity = problem_identity")
	return nil, nil
}
