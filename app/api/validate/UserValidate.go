package validate

import "Authority/app/common"

var UserValidate common.Validator

func init() {
	rules := map[string]string {
		"uid" : "required",
		"nick": "required|maxLen:20",
		"password": "required|minLen:6|maxLen:20",
		"password_check": "required|minLen:6|maxLen:20",
		"mail"	: "required|email",
		"is_admin": "required|bool",
	}

	scenes := map[string] []string {
		"register" 	: {"nick", "password", "password_check", "mail"},
		"login"		: {"nick", "password"},
		"delete"	: {"uid"},
		"update"	: {"uid", "nick", "mail"},
	}

	UserValidate.Rules = rules
	UserValidate.Scenes = scenes
}
