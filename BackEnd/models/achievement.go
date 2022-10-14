package models

import (
	"gorm.io/gorm"
	"time"
)

type Achievement struct {
	ID                uint `gorm:"PRIMARY_KEY"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
	Identity          string         `gorm:"index;NOT NULL;Type:varchar(36);Column:identity" json:"identity"`
	Name              string         `gorm:"NOT NULL;Type:varchar(36);Column:name" json:"name"`
	KnowledgeIdentity string         `gorm:"NOT NULL;Type:varchar(36);Column:knowledge_identity" json:"knowledge_identity"`
	Knowledge         *Knowledge     `gorm:"foreignKey:KnowledgeIdentity;references:Identity"`
}
