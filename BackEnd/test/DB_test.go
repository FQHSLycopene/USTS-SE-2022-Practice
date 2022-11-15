package test

import (
	"BackEnd/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestDB(t *testing.T) {
	dsn := "Lycopene:MiMaJiuShi123321!@tcp(127.0.0.1:3306)/SE2022?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
	}
	problemIdentities := []string{"10bf65da-be8f-4323-a933-f4e8bc101420", "e87a03f5-f915-4ff3-81f0-654db910de03"}
	problems := make([]models.Problem, 0)
	db.Where("identity in ?", problemIdentities).Find(&problems)
	fmt.Println(problems)
}

func TestHasmany(t *testing.T) {
	dsn := "Lycopene:MiMaJiuShi123321!@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
	}
	//db.Create(&User{
	//	Identity: "user_2",
	//	Name:     "cjx",
	//	Cards: []Card{
	//		{Identity: "card_3", Name: "cjx_1"},
	//		{Identity: "card_4", Name: "cjx_1"},
	//	},
	//})
	//db.Create(&Card{
	//	Identity:     "card_1",
	//	Name:         "wjw_1",
	//	UserIdentity: "user_1",
	//})
	//db.Create(&Card{
	//	Identity:     "card_2",
	//	Name:         "wjw_2",
	//	UserIdentity: "user_1",
	//})

	data := make([]User, 0)
	var total int64
	db.Model(data).Preload("Cards").Count(&total).Find(&data)
	fmt.Println(data)
}

type User struct {
	Identity  string     `gorm:"PRIMARY_KEY;index;NOT NULL;Type:varchar(36);Column:identity" json:"identity"`
	Name      string     `gorm:"NOT NULL;Type:varchar(36);Column:name" json:"name"`
	Cards     []Card     `gorm:"foreignKey:UserIdentity;references:Identity"`
	Languages []Language `gorm:"many2many:user_language;foreignKey:Identity;joinForeignKey:UserIdentity;References:Identity;joinReferences:LanguageIdentity"`
}

type Card struct {
	Identity     string `gorm:"PRIMARY_KEY;index;NOT NULL;Type:varchar(36);Column:identity" json:"identity"`
	Name         string `gorm:"NOT NULL;Type:varchar(36);Column:name" json:"name"`
	UserIdentity string `gorm:"NOT NULL;Type:varchar(36);Column:user_identity" json:"user_identity"`
}

type Language struct {
	Identity string `gorm:"PRIMARY_KEY;index;NOT NULL;Type:varchar(36);Column:identity" json:"identity"`
	Name     string `gorm:"NOT NULL;Type:varchar(36);Column:name" json:"name"`
}
