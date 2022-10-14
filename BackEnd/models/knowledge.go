package models

import (
	"BackEnd/utils"
	"gorm.io/gorm"
	"time"
)

type Knowledge struct {
	ID        uint `gorm:"PRIMARY_KEY"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Identity  string         `gorm:"index;NOT NULL;Type:varchar(36);Column:identity" json:"identity"`
	Name      string         `gorm:"NOT NULL;Type:varchar(36);Column:name" json:"name"`
}

func AddKnowledge(name string) (interface{}, error) {
	data := Knowledge{
		Identity: utils.GetUuid(),
		Name:     name,
	}
	err := DB.Create(&data).Error
	if err != nil {
		return nil, err
	}
	return data, err
}
