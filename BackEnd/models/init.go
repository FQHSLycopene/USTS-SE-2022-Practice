package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB = InitDB()

func InitDB() *gorm.DB {
	dsn := "Lycopene:MiMaJiuShi123321!@tcp(127.0.0.1:3306)/SE2022?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
	}
	db.AutoMigrate(&User{})
	return db
}
