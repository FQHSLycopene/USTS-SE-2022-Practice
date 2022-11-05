package models

import (
	"BackEnd/utils"
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

func AddProblem(name, content, answer, categoryIdentity string, score int, knowledgeIdentities []string) (interface{}, error) {
	data := Problem{
		Identity: utils.GetUuid(),
		Name:     name,
		Content:  content,
		Answer:   answer,
		Score:    score,
	}

	problemCategory, err := getProblemCategoryByIdentity(categoryIdentity)
	if err != nil {
		return nil, err
	}
	data.ProblemCategory = problemCategory
	knowledges := make([]*Knowledge, 0)
	for _, knowledgeIdentity := range knowledgeIdentities {
		knowledge, err := getKnowledgeByIdentity(knowledgeIdentity)
		if err != nil {
			return nil, err
		}
		knowledges = append(knowledges, knowledge)
	}
	data.Knowledge = knowledges
	DB.Preload("Knowledge").Create(&data)
	return data, err
}

func UpdateProblem(identity, name, content, answer, categoryIdentity string, score int, knowledgeIdentities []string) (interface{}, error) {
	data, err := getProblemByIdentity(identity)
	if err != nil {
		return nil, err
	}

	if name != "" {
		data.Name = name
	}
	if content != "" {
		data.Content = content
	}
	if answer != "" {
		data.Answer = answer
	}
	if score != 0 {
		data.Score = score
	}
	err = DB.Model(&data).Preload("Knowledge").Save(&data).Error
	if err != nil {
		return nil, err
	}

	if categoryIdentity != "" {
		problemCategory, err := getProblemCategoryByIdentity(categoryIdentity)
		if err != nil {
			return nil, err
		}
		err = DB.Model(&data).Preload("Knowledge").Association("ProblemCategory").Replace(problemCategory)
		if err != nil {
			return nil, err
		}
	}

	if len(knowledgeIdentities) != 0 {
		knowledges := make([]*Knowledge, 0)
		for _, knowledgeIdentity := range knowledgeIdentities {
			knowledge, err := getKnowledgeByIdentity(knowledgeIdentity)
			if err != nil {
				return nil, err
			}
			knowledges = append(knowledges, knowledge)
		}
		err := DB.Model(&data).Preload("Knowledge").Association("Knowledge").Replace(knowledges)
		if err != nil {
			return nil, err
		}
	}
	return data, nil
}

func getProblemByIdentity(identity string) (*Problem, error) {
	data := Problem{}
	err := DB.Where("identity = ?", identity).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}
