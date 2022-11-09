package models

import (
	"gorm.io/gorm"
	"time"
)

type WrongProblem struct {
	ID              uint `gorm:"PRIMARY_KEY"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	Identity        string         `gorm:"index;NOT NULL;Type:varchar(36);Column:identity" json:"identity"`
	Name            string         `gorm:"NOT NULL;unique;Type:varchar(36);Column:name" json:"name"`
	ProblemIdentity string         `gorm:"Type:varchar(36);Column:problem_identity" json:"problem_identity"`
	Problem         *Problem       `gorm:"foreignKey:ProblemIdentity;references:Identity"`
	WrongAnswer     string         `gorm:"NOT NULL;Type:text;Column:wrong_answer" json:"wrong_answer"`
}
