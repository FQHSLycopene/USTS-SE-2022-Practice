package models

import (
	"BackEnd/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
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

func GetProblemCategoryByIdentity(identity string) (interface{}, error) {
	data := ProblemCategory{}
	err := DB.Where("identity = ?", identity).First(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetProblemCategoryIdentityByName(name string) (string, error) {
	data := ProblemCategory{}
	err := DB.Where("name = ?", name).First(&data).Error
	if err != nil {
		return "", err
	}
	return data.Identity, nil
}

func GetProblemCategoryList(pageStr, pageSizeStr, keyWord string) (interface{}, error) {
	data := make([]*ProblemCategory, 0)
	var total int64 = 0
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return nil, err
	}
	pageSize, err2 := strconv.Atoi(pageSizeStr)
	if err2 != nil {
		return nil, err2
	}
	err3 := DB.Model(data).Where("name like ?", "%"+keyWord+"%").Offset((page - 1) * pageSize).Limit(pageSize).Count(&total).Find(&data).Error
	if err3 != nil {
		return nil, err3
	}
	return gin.H{
		"total": total,
		"list":  data,
	}, nil
}
