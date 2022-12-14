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
	ID                uint `gorm:"PRIMARY_KEY"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt  `gorm:"index"`
	Identity          string          `gorm:"index;NOT NULL;Type:varchar(36);Column:identity" json:"identity"`
	Name              string          `gorm:"unique;NOT NULL;unique;Type:varchar(36);Column:name" json:"name"`
	Password          string          `gorm:"NOT NULL;Type:varchar(255);Column:password" json:"password"`
	Email             string          `gorm:"unique;Type:varchar(255);unique;Column:email" json:"email"`
	Phone             string          `gorm:"Type:varchar(255);Column:phone" json:"phone"`
	Status            int             `gorm:"NOT NULL;Type:int(11);Column:status" json:"status"`
	Avatar            string          `gorm:"default:image/avatar/default.png;Type:varchar(255);Column:avatar" json:"avatar"`
	Achievements      []*Achievement  `gorm:"many2many:user_achievements;foreignKey:Identity;joinForeignKey:UserIdentity;References:Identity;joinReferences:AchievementIdentity"`
	Classes           []*Class        `gorm:"many2many:user_classes;foreignKey:Identity;joinForeignKey:UserIdentity;References:Identity;joinReferences:ClassIdentity"`
	Practise          []*Practise     `gorm:"foreignKey:UserIdentity;references:Identity"`
	WrongProblems     []*wrongProblem `gorm:"foreignKey:UserIdentity;references:Identity"`
	ExamPaperProblems []*Problem      `gorm:"many2many:exam_paper_problems;foreignKey:Identity;joinForeignKey:UserIdentity;References:Identity;joinReferences:ProblemIdentity"`
}

type ExamPaperProblems struct {
	UserIdentity    string `gorm:"index;NOT NULL;Type:varchar(36);Column:user_identity" json:"user_identity"`
	ProblemIdentity string `gorm:"index;NOT NULL;Type:varchar(36);Column:problem_identity" json:"problem_identity"`
	ExamIdentity    string `gorm:"index;NOT NULL;Type:varchar(36);Column:exam_identity" json:"exam_identity"`
	Exam            *Exam  `gorm:"foreignKey:ExamIdentity;references:Identity"`
	MyAnswer        string `gorm:"Type:text;Column:my_answer" json:"my_answer"`
	Status          int    `gorm:"NOT NULL;Type:int(11);Column:status;default:0" json:"status"` //0????????????1?????????2??????
	MyScore         int    `gorm:"NOT NULL;Type:int(11);Column:score;default:0" json:"my-score"`
}

func (table *User) TableName() string {
	return "user"
}

func deleteUserClassAssociation(userIdentity, classIdentity string) error {
	user, err := GetUserByIdentity(userIdentity)
	if err != nil {
		return err
	}
	class, err2 := getClassByIdentity(classIdentity)
	if err2 != nil {
		return err2
	}
	err = DB.Model(user).Association("Classes").Delete(class)
	if err != nil {
		return err
	}
	return nil
}

func getUserIdentityByClassIdentity(classIdentity string) ([]string, error) {
	userIdentities := make([]string, 0)
	err := DB.Model(&User{}).Select("identity").Joins("right join user_classes as uc on uc.user_identity = identity").
		Where("uc.class_identity = ?", classIdentity).Find(&userIdentities).Error
	if err != nil {
		return nil, err
	}
	return userIdentities, nil
}

func deleteExamPaperProblems(userIdentity, examIdentity string) error {
	if userIdentity == "" || examIdentity != "" {
		err := deleteExamPaperProblemsByExamIdentity(examIdentity)
		return err
	} else if userIdentity != "" || examIdentity == "" {
		err := deleteExamPaperProblemsByUserIdentity(examIdentity)
		return err
	}
	epps := make([]*ExamPaperProblems, 0)
	err := DB.Model(&epps).Where("exam_identity = ?", examIdentity).
		Where("user_identity = ?", userIdentity).
		Unscoped().Delete(&epps).Error
	return err
}

func deleteExamPaperProblemsByUserIdentity(userIdentity string) error {
	epps := make([]*ExamPaperProblems, 0)
	err := DB.Model(&epps).Where("user_identity = ?", userIdentity).Unscoped().Delete(&epps).Error
	if err != nil {
		return err
	}
	return nil
}

func deleteExamPaperProblemsByExamIdentity(examIdentity string) error {
	epps := make([]*ExamPaperProblems, 0)
	err := DB.Model(&epps).Where("exam_identity = ?", examIdentity).Unscoped().Delete(&epps).Error
	if err != nil {
		return err
	}
	return nil
}

func InitExamPaperProblems(classIdentity, examIdentity string, problemIdentities []string) (interface{}, error) {
	users := make([]*User, 0)
	err := DB.Model(users).
		Joins("right join user_classes uc on uc.user_identity = identity").
		Where("uc.class_identity = ?", classIdentity).Where("status = 1").
		Find(&users).Error
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		epps := make([]*ExamPaperProblems, 0)
		for _, problemIdentity := range problemIdentities {
			epp := ExamPaperProblems{
				UserIdentity:    user.Identity,
				ProblemIdentity: problemIdentity,
				ExamIdentity:    examIdentity,
				MyAnswer:        "",
				Status:          0,
			}
			epps = append(epps, &epp)
		}
		err := DB.Create(&epps).Error
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func UserIsHasClass(identity string) (bool, error) {
	user, err := GetUserByIdentity(identity)
	if err != nil {
		return false, err
	}
	count := DB.Model(user).Association("Classes").Count()
	if count == 0 {
		return false, err
	}
	return true, nil
}

func UpdateAvatar(identity, dst string) (interface{}, error) {
	user, err := GetUserByIdentity(identity)
	if err != nil {
		return nil, err
	}
	user.Avatar = dst
	err = DB.Save(user).Error
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func UpdatePassword(identity, oldPassword, newPassword string) (interface{}, error) {
	user, err := GetUserByIdentity(identity)
	if err != nil {
		return nil, err
	}
	md5 := utils.GetMd5(oldPassword)
	if user.Password != md5 {
		return nil, errors.New("????????????")
	}
	user.Password = utils.GetMd5(newPassword)
	err = DB.Save(&user).Error
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func UpdateUser(identity, name, email, phone string) (interface{}, error) {
	user, err := GetUserByIdentity(identity)
	if err != nil {
		return nil, err
	}
	if email != user.Email {
		exist, err := EmailIsExist(email)
		if err != nil {
			return nil, err
		}
		if exist {
			return nil, errors.New("??????????????????")
		}
		user.Email = email
	}
	if name != user.Name {
		exist, err := UserNameIsExist(name)
		if err != nil {
			return nil, err
		}
		if exist {
			return nil, errors.New("?????????????????????")
		}
		user.Name = name
	}
	user.Phone = phone
	err = DB.Save(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UserNameIsExist(name string) (bool, error) {
	var total int64 = 0
	err := DB.Model(&User{}).Where("name = ?", name).Count(&total).Error
	if err != nil {
		return true, err
	}
	if total == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func GetUser(identity string) (interface{}, error) {
	user, err := GetUserByIdentity(identity)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetAvatarByIdentity(identity string) (interface{}, error) {
	user, err := GetUserByIdentity(identity)
	if err != nil {
		return nil, err
	}
	return user.Avatar, nil
}

func InitPractise(userIdentity string) (interface{}, error) {
	knowledgeIdentities, err := getKnowledgeIdentities()
	if err != nil {
		return nil, err
	}
	for _, identity := range knowledgeIdentities {
		_, err := AddPractise(userIdentity, identity)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func GetTokenByEmail(email string) (interface{}, error) {
	user := User{}
	err := DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	token, err2 := utils.GenerateToken(user.Identity, user.Status)
	if err2 != nil {
		return nil, err2
	}
	return token, nil
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
		Where("name like ?", "%"+keyWord+"%").
		Where("status = ?", 1).Count(&total).
		Offset((page-1)*pageSize).Limit(pageSize).
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
	exist, err3 := EmailIsExist(email)
	if err3 != nil {
		return nil, err3
	}
	if exist {
		return nil, errors.New("???????????????")
	}
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
	token, err2 := utils.GenerateToken(data.Identity, data.Status)
	if err2 != nil {
		return nil, err2
	}
	if status == 1 {
		_, err := InitPractise(data.Identity)
		if err != nil {
			return nil, err
		}
	}

	return token, nil
}

func LoginByName(name, password string) (interface{}, error) {
	var data User
	passwordMd5 := utils.GetMd5(password)
	if name != "" {
		err := DB.Where("name = ?", name).Find(&data).Error
		if err != nil {
			return nil, err
		}
	}
	if data.Password == passwordMd5 {
		token, err := utils.GenerateToken(data.Identity, data.Status)
		if err != nil {
			return nil, err
		}
		return token, nil
	} else {
		return nil, errors.New("????????????????????????")
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
