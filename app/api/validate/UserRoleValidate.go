package validate

import "Authority/app/common"

var UserRoleValidate common.Validator

func init() {
	rules := map[string]string {
		"uid"	: "required",
		"rid"	: "required",
		"rids" 	: "required",
	}

	scenes := map[string] []string {
		"add" : {"uid", "rid"},
		"addGroup": {"uid", "rids"},
		"deleteGroup": {"uid", "rids"},
		"delete": {"uid", "rid"},
		"getUserRole": {"uid"},
	}

	UserRoleValidate.Rules = rules
	UserRoleValidate.Scenes = scenes
}
