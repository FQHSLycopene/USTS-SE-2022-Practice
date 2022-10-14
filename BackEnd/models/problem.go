package models

import (
	"gorm.io/gorm"
	"time"
)

type Problem struct {
	ID               uint `gorm:"PRIMARY_KEY"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt   `gorm:"index"`
	Identity         string           `gorm:"index;NOT NULL;Type:varchar(36);Column:identity" json:"identity"`
	Name             string           `gorm:"NOT NULL;Type:varchar(36);Column:name" json:"name"`
	Content          string           `gorm:"NOT NULL;Type:text;Column:content" json:"content"`
	Answer           string           `gorm:"NOT NULL;Type:text;Column:answer" json:"answer"`
	Score            int              `gorm:"NOT NULL;Type:int;Column:score" json:"score"`
	Knowledge        []*Knowledge     `gorm:"many2many:problem_knowledge;foreignKey:Identity;joinForeignKey:ProblemIdentity;References:Identity;joinReferences:KnowledgeIdentity"`
	CategoryIdentity string           `gorm:"NOT NULL;Type:varchar(36);Column:category_identity" json:"category_identity"`
	ProblemCategory  *ProblemCategory `gorm:"foreignKey:CategoryIdentity;references:Identity"`
}
