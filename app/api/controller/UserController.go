package controller

import (
	"Authority/app/api/model"
	"Authority/app/api/validate"
	"Authority/app/common"
	"encoding/json"
	"github.com/gin-contrib/sessions"
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
		return
	}
	c.JSON(http.StatusOK, common.ApiReturn(common.CODE_ERROE, "绑定数据模型失败", false))
	return
}

func Login(c *gin.Context) {
	session := sessions.Default(c)
	if id := session.Get("user_id"); id != nil {
		data := make(map[string]interface{}, 0)
		_ = json.Unmarshal([]byte(session.Get("data").(string)), &data)
		c.JSON(http.StatusOK, common.ApiReturn(common.CODE_SUCCESS, "已登陆", data))
		return
	}

	userValidate := validate.UserValidate
	userModel := model.User{}

	if res, err:= userValidate.Validate(c, "login"); !res {
		c.JSON(http.StatusOK, common.ApiReturn(common.CODE_ERROE, err.Error(), 0))
		return
	}

	var loginUser model.User

	if c.ShouldBind(&loginUser) == nil {
		loginUser.Password = common.GetMd5(loginUser.Password)
		res := userModel.CheckLogin(loginUser)
		if res.Status == common.CODE_SUCCESS {
			userInfo := res.Data.(map[string]interface{})["userInfo"].(model.User)
			returnData := map[string]interface{} {
				"user_id" : userInfo.Uid,
				"nick":		userInfo.Nick,
			}
			jsonData, _ := json.Marshal(returnData)
			session.Set("user_id", returnData["user_id"])
			session.Set("data", string(jsonData))
			session.Save()
			c.JSON(http.StatusOK, common.ApiReturn(res.Status, "登录成功", returnData))
		} else {
			c.JSON(http.StatusOK, common.ApiReturn(common.CODE_ERROE, res.Msg, false))
		}
		return
	}
	c.JSON(http.StatusOK, common.ApiReturn(common.CODE_ERROE, "绑定数据模型失败", false))
	return
}

func Logout(c *gin.Context)  {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, common.ApiReturn(common.CODE_SUCCESS, "注销成功", session.Get("user_id")))
}