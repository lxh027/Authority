package model

type Auth struct {
	Aid 	int 	`json:"aid" form:"aid"`
	Icon 	string 	`json:"icon" form:"icon"`
	Title	string 	`json:"title" form:"title"`
	Href	string 	`json:"href" form:"href"`
	Target 	string 	`json:"target" form:"target"`
	Type 	int 	`json:"type" form:"type"`
	Parent	int 	`json:"parent" form:"parent"`
}


