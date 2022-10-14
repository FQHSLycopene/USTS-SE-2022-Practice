package models

import (
	"gorm.io/gorm"
	"time"
)

type Class struct {
	ID            uint `gorm:"PRIMARY_KEY"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	Identity      string         `gorm:"index;NOT NULL;Type:varchar(36);Column:identity" json:"identity"`
	Name          string         `gorm:"NOT NULL;Type:varchar(36);Column:name" json:"name"`
	StudentNumber int            `gorm:"NOT NULL;Type:int;Column:student_number" json:"student_number"`
	Exams         []*Exam        `gorm:"foreignKey:ClassIdentity;references:Identity"`
}
