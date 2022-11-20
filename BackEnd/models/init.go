package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB = InitDB()

func InitDB() *gorm.DB {
	dsn := "root:dlq668713@tcp(127.0.0.1:3306)/stus?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
	}
	db.AutoMigrate(&User{}, &Achievement{}, &Class{}, &Exam{}, &Practise{}, &Problem{}, &PractiseProblem{}, &ProblemCategory{}, &Knowledge{}, &ExamProblems{}, &PractiseProblem{}, &ExamPaperProblems{}, &wrongProblem{})

	return db
}
