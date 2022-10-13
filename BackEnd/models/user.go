package models

import (
	"BackEnd/utils"
	"errors"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint `gorm:"PRIMARY_KEY"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Identity  string         `gorm:"index;NOT NULL;Type:varchar(36);Column:identity" json:"identity"`
	Name      string         `gorm:"NOT NULL;Type:varchar(36);Column:name" json:"name"`
	Password  string         `gorm:"NOT NULL;Type:varchar(255);Column:password" json:"password"`
	Email     string         `gorm:"Type:varchar(255);Column:email" json:"email"`
	Phone     string         `gorm:"Type:varchar(255);Column:phone" json:"phone"`
	Status    int            `gorm:"NOT NULL;Type:int(11);Column:status" json:"status"`
}

func (table *User) TableName() string {
	return "user"
}

func GetUserByIdentity(identity string) (interface{}, error) {
	data := User{}
	err := DB.Where("identity = ?", identity).First(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func AddUser(name, password, email, phone string) (interface{}, error) {
	data := User{
		Identity: utils.GetUuid(),
		Name:     name,
		Password: utils.GetMd5(password),
		Email:    email,
		Phone:    phone,
		Status:   0,
	}
	err := DB.Create(&data).Error
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
