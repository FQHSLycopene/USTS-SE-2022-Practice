package models

import (
	"BackEnd/utils"
	"errors"
	"fmt"
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
	Key              string           `gorm:"NOT NULL;Type:text;Column:key" json:"Key"` //题解
	Score            int              `gorm:"NOT NULL;Type:int;Column:score" json:"score"`
	Knowledge        []*Knowledge     `gorm:"many2many:problem_knowledge;foreignKey:Identity;joinForeignKey:ProblemIdentity;References:Identity;joinReferences:KnowledgeIdentity"`
	CategoryIdentity string           `gorm:"NOT NULL;Type:varchar(36);Column:category_identity" json:"category_identity"`
	ProblemCategory  *ProblemCategory `gorm:"foreignKey:CategoryIdentity;references:Identity"`
}

var ProblemModels *Problem

func (*Problem) PractiseProblemIsCorrect(userIdentity, problemIdentity, practiseIdentity, answer string) (interface{}, error) {
	problem, err := getProblemByIdentity(problemIdentity)
	if err != nil {
		return nil, err
	}
	practise, err2 := getPractiseByIdentity(practiseIdentity)
	if err2 != nil {
		return nil, err2
	}
	user, err3 := GetUserByIdentity(userIdentity)
	if err3 != nil {
		return nil, err3
	}
	practiseProblem := PractiseProblem{}
	err = DB.Model(practiseProblem).
		Where("practise_identity = ?", practiseIdentity).
		Where("problem_identity = ?", problemIdentity).First(&practiseProblem).Error
	if err != nil {
		return nil, err
	}
	if answer == problem.Answer {
		practise.RightNum++
		practiseProblem.Status = 1
		DB.Save(&practise)
		DB.Where("practise_identity = ?", practiseIdentity).
			Where("problem_identity = ?", problemIdentity).Save(&practiseProblem)
		if practise.ProblemNum == practise.RightNum {
			achievement, err := getAchievementByKnowledgeIdentity(practise.KnowledgeIdentity)
			fmt.Println(achievement)
			if err != nil {
				return nil, err
			}
			err = DB.Model(user).Association("Achievements").Append(achievement)
			if err != nil {
				return nil, err
			}
		}
		return "correct", nil
	} else {
		wrongProblem := wrongProblem{

			Identity:        utils.GetUuid(),
			Name:            problem.Name + time.ANSIC,
			Problem:         problem,
			ProblemIdentity: problemIdentity,
			WrongAnswer:     answer,
		}
		err := DB.Model(&user).Association("WrongProblems").Append(&wrongProblem)
		if err != nil {
			return nil, err
		}
		practiseProblem.Status = 2
		DB.Where("practise_identity = ?", practiseIdentity).
			Where("problem_identity = ?", problemIdentity).Save(&practiseProblem)
		return "wrong", nil
	}

}

func GetRandPractiseProblemDetail(practiseIdentity string) (interface{}, error) {
	//problem := Problem{}
	//problems := make([]*Problem, 0)
	practise, err := getPractiseByIdentity(practiseIdentity)
	if err != nil {
		return nil, err
	}
	if practise.ProblemNum == practise.RightNum {
		return nil, errors.New("该练习已完成")
	}
	practiseProblem := PractiseProblem{}
	DB.Model(practiseProblem).
		Where("practise_identity = ?", practiseIdentity).
		Where("status != ?", 1).Order("RAND()").Limit(1).
		First(&practiseProblem)
	problem, err2 := getProblemByIdentity(practiseProblem.ProblemIdentity)
	if err2 != nil {
		return nil, err2
	}
	return problem, nil
}

func getProblemByKnowledge(knowledgeIdentity string) ([]*Problem, error) {
	problems := make([]*Problem, 0)
	err := DB.Preload("Knowledge").Joins("right join problem_knowledge pk on pk.problem_identity = identity").
		Where("pk.knowledge_identity = ?", knowledgeIdentity).
		Find(&problems).Error
	if err != nil {
		return nil, err
	}
	return problems, nil
}

func AddProblem(name, content, key, answer, categoryIdentity string, score int, knowledgeIdentities []string) (interface{}, error) {
	data := Problem{
		Identity: utils.GetUuid(),
		Name:     name,
		Content:  content,
		Answer:   answer,
		Score:    score,
		Key:      key,
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
	err := DB.Preload("Knowledge").Preload("ProblemCategory").Where("identity = ?", identity).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}
