package model

import (
	"gin-casbin/global"
	"gin-casbin/utils/response"
)

type User struct {
	Model
	UserName    string `gorm:"type:varchar(20)" json:"username"`
	PassWord    string `gorm:"type:varchar(20)" json:"password"`
	AuthorityId string `gorm:"type:int" json:"authorityId"`
}

func CreateUser(user *User) int {
	if result := global.GVA_DB.Create(&user); result.RowsAffected == 0 {
		return response.ERROR
	}
	return response.SUCCESS
}

func GetUserByUserName(userName string) (User, int) {
	var user User
	if result := global.GVA_DB.Where("user_name=?", userName).Find(&user); result.RowsAffected == 0 {
		return user, response.USER_NOT_EXIST
	}
	return user, response.SUCCESS
}

func GetUserByUid(uid uint) (User, int) {
	var user User
	if result := global.GVA_DB.Where("id=?", uid).Find(&user); result.RowsAffected == 0 {
		return user, response.USER_NOT_EXIST
	}
	return user, response.SUCCESS
}
