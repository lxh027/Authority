package model

import "Authority/app/common"

type User struct {
	Uid 	int 	`json:"uid" form:"uid"`
	Nick	string 	`json:"nick" form:"nick"`
	Password string	`json:"password" form:"password"`
	Mail 	string 	`json:"mail" form:"mail"`
	IsAdmin	int 	`json:"is_admin" form:"is_admin"`
}

func (model *User) AddUser(newUser User) common.ReturnType {
	user :=User{}

	if err := db.Where("nick = ? OR mail = ?", newUser.Nick, newUser.Mail).First(&user).Error; err == nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "昵称或邮箱已存在",  Data: false}
	}

	err := db.Create(&newUser).Error

	if err != nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "创建失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CODE_SUCCESS, Msg: "创建成功", Data: true}
	}
}
