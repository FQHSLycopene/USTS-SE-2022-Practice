package models

import (
	"BackEnd/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type Class struct {
	ID            uint `gorm:"PRIMARY_KEY"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	Identity      string         `gorm:"index;NOT NULL;Type:varchar(36);Column:identity" json:"identity"`
	Name          string         `gorm:"NOT NULL;Type:varchar(36);Column:name" json:"name"`
	JoinCode      string         `gorm:"UNIQUE;NOT NULL;Type:varchar(6);Column:join_code" json:"join_code"`
	StudentNumber int            `gorm:"NOT NULL;Type:int;Column:student_number" json:"student_number"`
	Exams         []*Exam        `gorm:"foreignKey:ClassIdentity;references:Identity"`
}

func DeleteClass(classIdentity string) (interface{}, error) {
	examIdentities, err := getExamIdentityByClassIdentity(classIdentity)
	if err != nil {
		return nil, err
	}
	//删除班级关联的考试
	for _, identity := range examIdentities {
		err := deleteExam(identity)
		if err != nil {
			return nil, err
		}
	}
	//删除班级关联的用户
	userIdentities, err2 := getUserIdentityByClassIdentity(classIdentity)
	if err2 != nil {
		return nil, err2
	}
	for _, identity := range userIdentities {
		err := deleteUserClassAssociation(identity, classIdentity)
		if err != nil {
			return nil, err
		}
	}
	//删除班级
	class, err3 := getClassByIdentity(classIdentity)
	if err3 != nil {
		return nil, err3
	}
	err = DB.Unscoped().Delete(class).Error
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func GetClassDetail(identity string) (interface{}, error) {
	class := Class{}
	err := DB.Preload("Exams").Where("identity = ?", identity).Find(&class).Error
	if err != nil {
		return nil, err
	}
	return class, nil
}

func UpdateClass(identity, name string, isChangCode bool, studentIdentities []string) (interface{}, error) {
	class, err := getClassByIdentity(identity)
	if err != nil {
		return nil, err
	}
	if name != "" {
		class.Name = name
	}
	if isChangCode {
		var code string
		for true {
			code = utils.GetCode()
			exist := joinCodeIsExist(code)
			if !exist {
				break
			}
		}
		class.JoinCode = code
	}
	if len(studentIdentities) != 0 {
		for _, studentIdentity := range studentIdentities {
			//删除学生与班级考试关联
			examIdentities, err2 := getExamIdentityByClassIdentity(identity)
			if err2 != nil {
				return nil, err2
			}
			for _, examIdentity := range examIdentities {
				err := deleteExamPaperProblems(studentIdentity, examIdentity)
				if err != nil {
					return nil, err
				}
			}
			//删除学生与班级关联
			err := deleteUserClassAssociation(studentIdentity, class.Identity)
			if err != nil {
				return nil, err
			}
			class.StudentNumber -= 1
		}
	}
	DB.Save(&class)
	return class, nil
}

func joinCodeIsExist(code string) bool {
	datas := make([]*Class, 0)
	DB.Where("join_code = ?", code).Find(&datas)
	if len(datas) == 0 {
		return false
	} else {
		return true
	}

}

func JoinClass(joinCode, userIdentity string) (interface{}, error) {
	class, err := getClassByJoinCode(joinCode)
	if err != nil {
		return nil, errors.New("班级不存在")
	}
	data := User{}
	users := make([]User, 0)
	user, err2 := GetUserByIdentity(userIdentity)
	if err2 != nil {
		return nil, err2
	}

	err = DB.Preload("Classes").Joins("right join user_classes uc on uc.user_identity = identity").
		Where("uc.class_identity = ?", class.Identity).
		Where("identity = ?", user.Identity).Find(&users).Error
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		err = DB.Model(user).Association("Classes").Append(class)
		if err != nil {
			return nil, err
		}
		class.StudentNumber += 1
	} else {
		return nil, errors.New("已加入班级")
	}

	err = DB.Save(&class).Error
	if err != nil {
		return nil, err
	}
	err = DB.Preload("Classes").
		Where("identity = ?", userIdentity).First(&data).Error
	data.Phone = ""
	data.Email = ""
	data.Password = ""
	return data, nil
}

func AddClass(name, userIdentity string) (interface{}, error) {
	class := Class{
		Identity:      utils.GetUuid(),
		Name:          name,
		JoinCode:      utils.GetCode(),
		StudentNumber: 0,
	}
	data := User{}
	user, err := GetUserByIdentity(userIdentity)
	if err != nil {
		return nil, err
	}
	err = DB.Model(&user).Association("Classes").Append(&class)
	if err != nil {
		return nil, err
	}
	err = DB.Preload("Classes").
		Where("identity = ?", userIdentity).First(&data).Error
	if err != nil {
		return nil, err
	}
	data.Status = -1
	data.Phone = ""
	data.Email = ""
	data.Password = ""
	return data, nil
}

func getClassByJoinCode(joinCode string) (*Class, error) {
	data := Class{}
	err := DB.Where("join_code = ?", joinCode).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func getClassByIdentity(identity string) (*Class, error) {
	data := Class{}
	err := DB.Where("identity = ?", identity).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func GetClassList(userIdentity, pageStr, pageSizeStr, keyWord string) (interface{}, error) {
	data := make([]*Class, 0)
	var total int64 = 0
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return nil, err
	}
	pageSize, err2 := strconv.Atoi(pageSizeStr)
	if err2 != nil {
		return nil, err2
	}
	err = DB.Model(data).
		Joins("right join user_classes uc on uc.class_identity = identity").
		Where("uc.user_identity = ?", userIdentity).Count(&total).
		Where("name like ?", "%"+keyWord+"%").
		Offset((page - 1) * pageSize).Limit(pageSize).Omit("JoinCode").
		Find(&data).Error
	return gin.H{
		"total": total,
		"list":  data,
	}, nil
}
