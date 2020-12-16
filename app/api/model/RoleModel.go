package model

type Role struct {
	Rid 	int 	`json:"rid" form:"rid"`
	Desc 	string 	`json:"desc" form:"desc"`
}