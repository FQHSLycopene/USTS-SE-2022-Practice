package models

import (
	"BackEnd/utils"
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
	UserIdentity      string         `gorm:"Type:varchar(36);Column:user_identity" json:"user_identity"`
	KnowledgeIdentity string         `gorm:"Type:varchar(36);Column:knowledge_identity" json:"knowledge_identity"`
	Knowledge         *Knowledge     `gorm:"foreignKey:KnowledgeIdentity;references:Identity"`
	Problems          []*Problem     `gorm:"many2many:practise_problems;foreignKey:Identity;joinForeignKey:PractiseIdentity;References:Identity;joinReferences:ProblemIdentity"`
	ProblemNum        int            `gorm:"NOT NULL;Type:int(11);Column:problem_num" json:"problem_num"`
}
type PractiseProblem struct {
	PractiseIdentity string `gorm:"NOT NULL;Type:varchar(36);Column:practise_identity" json:"practise_identity"`
	ProblemIdentity  string `gorm:"NOT NULL;Type:varchar(36);Column:problem_identity" json:"problem_identity"`
	Status           int    `gorm:"NOT NULL;Type:int(11);Column:status;default:0" json:"status"` //0没有做，1做对，2做错
}

//func UpdatePractise(identity, name, knowledgeIdentity string) (interface{}, error) {
//	practise, err := getPractiseByIdentity(identity)
//	if err != nil {
//		return nil, err
//	}
//	if name != "" {
//		practise.Name = name
//	}
//}

func AddPractiseProblem(identity string, problemIdentities []string) (interface{}, error) {
	practise, err := getPractiseByIdentity(identity)
	if err != nil {
		return nil, err
	}
	if len(problemIdentities) != 0 {
		problems := make([]*Problem, 0)
		for _, problemIdentity := range problemIdentities {
			problem, err := getProblemByIdentity(problemIdentity)
			if err != nil {
				return nil, err
			}
			problems = append(problems, problem)
		}
		err := DB.Model(practise).Association("Problems").Append(problems)
		if err != nil {
			return nil, err
		}
		practise.ProblemNum += len(problemIdentities)
	}
	err = DB.Preload("Problems").Save(practise).Error
	if err != nil {
		return nil, err
	}
	return practise, nil
}

func AddPractise(userIdentity, knowledgeIdentity string) (interface{}, error) {
	user, err3 := GetUserByIdentity(userIdentity)
	if err3 != nil {
		return nil, err3
	}
	knowledge, err := getKnowledgeByIdentity(knowledgeIdentity)
	if err != nil {
		return nil, err
	}
	problems, err2 := getProblemByKnowledge(knowledgeIdentity)
	if err2 != nil {
		return nil, err2
	}
	practise := Practise{
		Identity:     utils.GetUuid(),
		Name:         user.Name + ":" + knowledge.Name + ":practise",
		UserIdentity: userIdentity,
		Knowledge:    knowledge,
		Problems:     problems,
		ProblemNum:   len(problems),
	}
	err = DB.Create(&practise).Error
	if err != nil {
		return nil, err
	}
	return practise, nil
}

func getPractiseByIdentity(identity string) (*Practise, error) {
	data := Practise{}
	err := DB.Where("identity = ?", identity).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}
