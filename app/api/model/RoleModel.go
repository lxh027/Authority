package model

import (
	"Authority/app/common"
)

type Role struct {
	Rid 	int 	`json:"rid" form:"rid"`
	Desc 	string 	`json:"desc" form:"desc"`
	Name 	string 	`json:"name" form:"name"`
}





func (model *Role) GetAllRole(offset int, limit int, name string, desc string) common.ReturnType {
	var roles []Role
	where := "name like ? AND `desc` like ?"
	var count int

	db.Model(&Role{}).Where(where, "%"+name+"%", "%"+desc+"%").Count(&count)


	err := db.Offset(offset).
		Limit(limit).
		Where(where, "%"+name+"%", "%"+desc+"%").
		Find(&roles).
		Error

	if err != nil {
		return common.ReturnType{Status: common.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CodeSuccess, Msg: "查询成功",
			Data: map[string]interface{}{
				"roles": roles,
				"count": count,
			},
		}
	}
}

func (model *Role) GetUserRole(uid int) common.ReturnType {
	var roles []Role

	err := db.
		Joins("JOIN user_role ON role.rid = user_role.rid AND user_role.uid = ? ", uid).
		Find(&roles).
		Error

	if err != nil {
		return common.ReturnType{Status: common.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CodeSuccess, Msg: "查询成功", Data: roles,
		}
	}
}

func (model *Role) GetRoleNoRules() common.ReturnType {
	/*var roles, rolesTotal []Role

	var countTotal, countRole int

	err1 := db.Order("rid").Find(&rolesTotal).
		Count(&countTotal).Error

	err2 := db.Joins("JOIN user_role ON role.rid = user_role.rid AND user_role.uid = ? ", uid).
		Order("rid").
		Find(&roles).
		Count(&countRole).Error

	countLeft := countTotal-countRole
	var rolesLeft []Role
	j := 0
	for i := 0; i < countRole; i++ {
		if roles[i].Rid == rolesTotal[j].Rid {
			j++
			continue
		}
		for roles[i].Rid != rolesTotal[j].Rid {
			rolesLeft = append(rolesLeft, rolesTotal[j])
			j++
		}
	}
	for j < countTotal {
		rolesLeft = append(rolesLeft, rolesTotal[j])
		j++
	}*/
	var rolesTotal []Role

	err := db.Find(&rolesTotal).Error

	if err != nil {
		return common.ReturnType{Status: common.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CodeSuccess, Msg: "查询成功", Data: rolesTotal,
		}
	}
}

func (model *Role) GetRoleByID(rid int) common.ReturnType {//jun
	var getRole Role

	err := db.Select([]string{"name", "`desc`"}).Where("rid = ?", rid).First(&getRole).Error
	if err != nil {
		return common.ReturnType{Status: common.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CodeSuccess, Msg: "查询成功", Data: getRole}
	}
}

func (model *Role) AddRole(newRole Role) common.ReturnType {//jun
	role :=Role{}

	if err := db.Where("name = ? OR `desc` = ?", newRole.Name,newRole.Desc).First(&role).Error; err == nil {
		return common.ReturnType{Status: common.CodeError, Msg: "角色名或描述已存在",  Data: false}
	}

	err := db.Create(&newRole).Error

	if err != nil {
		return common.ReturnType{Status: common.CodeError, Msg: "创建失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CodeSuccess, Msg: "创建成功", Data: true}
	}
}