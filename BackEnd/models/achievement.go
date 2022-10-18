package models

import (
	"BackEnd/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type Achievement struct {
	ID                uint `gorm:"PRIMARY_KEY"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
	Identity          string         `gorm:"index;NOT NULL;Type:varchar(36);Column:identity" json:"identity"`
	Name              string         `gorm:"UNIQUE;NOT NULL;Type:varchar(36);Column:name" json:"name"`
	KnowledgeIdentity string         `gorm:"NOT NULL;Type:varchar(36);Column:knowledge_identity" json:"knowledge_identity"`
	Knowledge         *Knowledge     `gorm:"foreignKey:KnowledgeIdentity;references:Identity"`
}

func DeleteAchievement(knowledgeIdentity string) (interface{}, error) {
	data, err := getAchievementByKnowledgeIdentity(knowledgeIdentity)
	if err != nil {
		return nil, err
	}
	err = DB.Delete(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func AddAchievement(name, knowledgeIdentity string) (interface{}, error) {
	data := Achievement{
		Identity:          utils.GetUuid(),
		Name:              name,
		KnowledgeIdentity: knowledgeIdentity,
	}
	err := DB.Create(&data).Error
	if err != nil {
		return nil, err
	}
	return data, err
}

func GetAchievementList(pageStr, pageSizeStr, keyWord string) (interface{}, error) {
	data := make([]*Achievement, 0)
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

func UpdateAchievement(name, knowledgeIdentity string) (interface{}, error) {
	data, err := getAchievementByKnowledgeIdentity(knowledgeIdentity)
	if err != nil {
		return nil, err
	}
	data.Name = name
	err = DB.Save(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func getAchievementByKnowledgeIdentity(knowledgeIdentity string) (*Achievement, error) {
	data := Achievement{}
	err := DB.Where("knowledge_identity = ?", knowledgeIdentity).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}
