package model

import "Authority/app/common"

type RoleAuth struct {
	Rid 	int 	`json:"rid" form:"rid"`
	Aid 	int 	`json:"aid" form:"aid"`
}

func (model *RoleAuth) AddRoleAuth(newRoleAuth RoleAuth) common.ReturnType {
	err := db.Create(&newRoleAuth).Error

	if err != nil {
		return common.ReturnType{Status: common.CodeError, Msg: "创建失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CodeSuccess, Msg: "创建成功", Data: true}
	}
}
func (model *RoleAuth) DeleteRoleAuth(newRoleAuth RoleAuth) common.ReturnType {
	err := db.Where("rid = ? AND aid = ?", newRoleAuth.Rid, newRoleAuth.Aid).Delete(RoleAuth{}).Error

	if err != nil {
		return common.ReturnType{Status: common.CodeError, Msg: "删除失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CodeSuccess, Msg: "删除成功", Data: true}
	}
}