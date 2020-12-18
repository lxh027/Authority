package model

import "Authority/app/common"

type UserRole struct {
	Uid 	int 	`json:"uid" form:"uid"`
	Rid		int 	`json:"rid" form:"rid"`
}


func (model *UserRole) AddUserRole(newUserRole UserRole) common.ReturnType {
	err := db.Create(&newUserRole).Error

	if err != nil {
		return common.ReturnType{Status: common.CodeError, Msg: "创建失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CodeSuccess, Msg: "创建成功", Data: true}
	}
}

func (model *UserRole) DeleteUserRole(newUserRole UserRole) common.ReturnType {
	err := db.Where("uid = ? AND rid = ?", newUserRole.Uid, newUserRole.Rid).Delete(UserRole{}).Error

	if err != nil {
		return common.ReturnType{Status: common.CodeError, Msg: "删除失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CodeSuccess, Msg: "删除成功", Data: true}
	}
}