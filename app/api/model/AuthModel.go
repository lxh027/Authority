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
