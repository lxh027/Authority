package validate

import "Authority/app/common"

var AuthValidate common.Validator

func init() {
	rules := map[string]string {
		"aid"	: "required",
		"icon"	: "required",
		"title"	: "required",
		"type"	: "required",
		"parent": "required",
	}

	scenes := map[string] []string {
		"add" : {"title", "type", "icon"},
		"delete": {"aid"},
		"find"	: {"rid"},
		"findParent": {"parent"},
	}

	AuthValidate.Rules = rules
	AuthValidate.Scenes = scenes
}
