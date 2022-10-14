package models

import (
	"gorm.io/gorm"
	"time"
)

type Practise struct {
	ID                uint `gorm:"PRIMARY_KEY"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
	Identity          string         `gorm:"index;NOT NULL;Type:varchar(36);Column:identity" json:"identity"`
	Name              string         `gorm:"NOT NULL;Type:varchar(36);Column:name" json:"name"`
	UserIdentity      string         `gorm:"NOT NULL;Type:varchar(36);Column:user_identity" json:"user_identity"`
	KnowledgeIdentity string         `gorm:"NOT NULL;Type:varchar(36);Column:knowledge_identity" json:"knowledge_identity"`
	Knowledge         *Knowledge     `gorm:"foreignKey:KnowledgeIdentity;references:Identity"`
	Problems          []*Problem     `gorm:"many2many:practise_problems;foreignKey:Identity;joinForeignKey:PractiseIdentity;References:Identity;joinReferences:ProblemIdentity"`
}
type PractiseProblem struct {
	PractiseIdentity string `gorm:"NOT NULL;Type:varchar(36);Column:practise_identity" json:"practise_identity"`
	ProblemIdentity  string `gorm:"NOT NULL;Type:varchar(36);Column:problem_identity" json:"problem_identity"`
	Status           int    `gorm:"NOT NULL;Type:int(11);Column:status" json:"status"` //0为做，1做对，2做错
}
