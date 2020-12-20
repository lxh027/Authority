package model

import "Authority/app/common"

type Auth struct {
	Aid 	int 	`json:"aid" form:"aid"`
	Icon 	string 	`json:"icon" form:"icon"`
	Title	string 	`json:"title" form:"title"`
	Href	string 	`json:"href" form:"href"`
	Target 	string 	`json:"target" form:"target"`
	Type 	int 	`json:"type" form:"type"`
	Parent	int 	`json:"parent" form:"parent"`
}

func (model *Auth) GetUserAllAuth(userID int) common.ReturnType  {
	var auths []Auth

	db.Joins("JOIN role_auth ON auth.aid = role_auth.aid").
		Joins("JOIN user_role ON role_auth.rid = user_role.rid AND user_role.uid = ?", userID).
		Find(&auths)

	return common.ReturnType{Status: common.CodeSuccess, Msg: "OK", Data: auths}
}

func (model *Auth) GetAllAuth(offset int, limit int, title string) common.ReturnType {
	var auths []Auth
	where := "title like ?"
	var count int

	db.Model(&Auth{}).Where(where, "%"+title+"%").Count(&count)

	err := db.
		Where(where, "%"+title+"%").
		Limit(limit).Offset(offset).
		Find(&auths).
		Error

	if err != nil {
		return common.ReturnType{Status: common.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CodeSuccess, Msg: "查询成功",
			Data: map[string]interface{}{
				"auths": auths,
				"count": count,
			},
		}
	}
}

func (model *Auth) GetParentAuth(parent int) common.ReturnType {
	var auths []Auth
	where := "type = ?"

	err := db.
		Where(where, parent).
		Find(&auths).
		Error

	if err != nil {
		return common.ReturnType{Status: common.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CodeSuccess, Msg: "查询成功", Data: auths,
		}
	}
}

func (model *Auth) AddAuth(newAuth Auth) common.ReturnType {
	auth :=Auth{}

	if err := db.Where("type = ? AND title = ?", newAuth.Type,newAuth.Title).First(&auth).Error; err == nil {
		return common.ReturnType{Status: common.CodeError, Msg: "此类型权限名已存在",  Data: false}
	}

	var err error
	if newAuth.Type == 0 {
		sqlLine := "INSERT INTO auth (icon, title, href, target, type) VALUES (?, ?, ?, ?, ?)"
		err = db.Exec(sqlLine, newAuth.Icon, newAuth.Title, newAuth.Href, newAuth.Target, newAuth.Type).Error
	} else {
		err = db.Create(&newAuth).Error
	}


	if err != nil {
		return common.ReturnType{Status: common.CodeError, Msg: "创建失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CodeSuccess, Msg: "创建成功", Data: true}
	}

}