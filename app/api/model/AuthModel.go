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

	err := db.Offset(offset).
		Limit(limit).
		Where(where, "%"+title+"%").
		Find(&auths).
		Count(&count).
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
