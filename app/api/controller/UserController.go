package controller

import (
	"Authority/app/api/model"
	"Authority/app/api/validate"
	"Authority/app/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	userValidate := validate.UserValidate
	userModel := model.User{}

	if res, err:= userValidate.Validate(c, "register"); !res {
		c.JSON(http.StatusOK, common.ApiReturn(common.CODE_ERROE, err.Error(), 0))
		return
	}

	passwordCheck := struct {
		Password string `json:"password" form:"password"`
		PasswordCheck string `json:"password_check" form:"password_check"`
	}{}

	if c.ShouldBind(&passwordCheck) == nil {
		if passwordCheck.Password != passwordCheck.PasswordCheck {
			c.JSON(http.StatusOK, common.ApiReturn(common.CODE_ERROE, "两次密码输入不一致", false))
			return
		}
	}

	var userJson model.User
	if c.ShouldBind(&userJson) == nil {
		userJson.Password = common.GetMd5(userJson.Password)
		res := userModel.AddUser(userJson)
		c.JSON(http.StatusOK, common.ApiReturn(res.Status, res.Msg, res.Data))
	}
	return
}