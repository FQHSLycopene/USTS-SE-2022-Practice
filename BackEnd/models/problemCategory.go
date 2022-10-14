package models

import (
	"BackEnd/utils"
	"gorm.io/gorm"
	"time"
)

type ProblemCategory struct {
	ID        uint `gorm:"PRIMARY_KEY"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Identity  string         `gorm:"index;NOT NULL;Type:varchar(36);Column:identity" json:"identity"`
	Name      string         `gorm:"NOT NULL;Type:varchar(36);Column:name" json:"name"`
}

func AddProblemCategory(name string) (interface{}, error) {
	data := ProblemCategory{
		Identity: utils.GetUuid(),
		Name:     name,
	}
	err := DB.Create(&data).Error
	if err != nil {
		return nil, err
	}
	return data, err
}
