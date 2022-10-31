package models

import (
	"BackEnd/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type User struct {
	ID           uint `gorm:"PRIMARY_KEY"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Identity     string         `gorm:"index;NOT NULL;Type:varchar(36);Column:identity" json:"identity"`
	Name         string         `gorm:"NOT NULL;Type:varchar(36);Column:name" json:"name"`
	Password     string         `gorm:"NOT NULL;Type:varchar(255);Column:password" json:"password"`
	Email        string         `gorm:"Type:varchar(255);Column:email" json:"email"`
	Phone        string         `gorm:"Type:varchar(255);Column:phone" json:"phone"`
	Status       int            `gorm:"NOT NULL;Type:int(11);Column:status" json:"status"`
	Achievements []*Achievement `gorm:"many2many:user_achievements;foreignKey:Identity;joinForeignKey:UserIdentity;References:Identity;joinReferences:AchievementIdentity"`
	Classes      []*Class       `gorm:"many2many:user_classes;foreignKey:Identity;joinForeignKey:UserIdentity;References:Identity;joinReferences:ClassIdentity"`
	Practise     []*Practise    `gorm:"foreignKey:UserIdentity;references:Identity"`
}

func (table *User) TableName() string {
	return "user"
}

func GetClassStudentList(classIdentity, pageStr, pageSizeStr, keyWord string) (interface{}, error) {
	data := make([]*User, 0)
	var total int64 = 0
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return nil, err
	}
	pageSize, err2 := strconv.Atoi(pageSizeStr)
	if err2 != nil {
		return nil, err2
	}
	err = DB.Model(&User{}).Joins("right join user_classes uc on uc.user_identity = identity").
		Where("uc.class_identity = ?", classIdentity).
		Where("status = ?", 1).Offset((page-1)*pageSize).Limit(pageSize).Count(&total).
		Select("Identity", "Name").
		Find(&data).Error
	if err != nil {
		return nil, err
	}
	return gin.H{
		"total": total,
		"list":  data,
	}, nil
}

func GetUserByIdentity(identity string) (*User, error) {
	data := User{}
	err := DB.Where("identity = ?", identity).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func AddUser(name, password, email, phone, statusStr string) (interface{}, error) {
	status, err := strconv.Atoi(statusStr)
	if err != nil {
		return nil, err
	}
	data := User{
		Identity: utils.GetUuid(),
		Name:     name,
		Password: utils.GetMd5(password),
		Email:    email,
		Phone:    phone,
		Status:   status,
	}
	err = DB.Create(&data).Error
	if err != nil {
		return nil, err
	}
	token, err2 := utils.GenerateToken(data.Identity, data.Name, data.Status)
	if err2 != nil {
		return nil, err2
	}
	return token, nil
}

func Login(name, password, email, phone string) (interface{}, error) {
	var data User
	passwordMd5 := utils.GetMd5(password)
	if name != "" {
		err := DB.Where("name = ?", name).Find(&data).Error
		if err != nil {
			return nil, err
		}
	} else if email != "" {
		err := DB.Where("email = ?", email).Find(&data).Error
		if err != nil {
			return nil, err
		}
	} else if phone != "" {
		err := DB.Where("phone = ?", phone).Find(&data).Error
		if err != nil {
			return nil, err
		}
	}
	if data.Password == passwordMd5 {
		token, err := utils.GenerateToken(data.Identity, data.Name, data.Status)
		if err != nil {
			return nil, err
		}
		return token, nil
	} else {
		return nil, errors.New("账号或者密码错误")
	}
}

func EmailIsExist(email string) (bool, error) {
	var total int64 = 0
	err := DB.Model(&User{}).Where("email = ?", email).Count(&total).Error
	if err != nil {
		return true, err
	}
	if total == 0 {
		return false, nil
	} else {
		return true, nil
	}
}
