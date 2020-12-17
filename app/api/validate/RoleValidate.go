package validate

import "Authority/app/common"

var RoleValidate common.Validator

func init() {
	rules := map[string]string {
		"rid"	: "required",
		"desc"	: "required",
	}

	scenes := map[string] []string {
		"add" : {"desc"},
		"delete": {"rid"},
		"update": {"rid", "desc"},
	}

	RoleValidate.Rules = rules
	RoleValidate.Scenes = scenes
}
