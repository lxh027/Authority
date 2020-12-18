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

func UpdateUser(c *gin.Context)  {
	if res := haveAuth(c, "updateUser"); res != common.Authed {
		c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	userValidate := validate.UserValidate
	userModel := model.User{}

	if res, err:= userValidate.Validate(c, "update"); !res {
		c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	var userJson model.User

	if c.ShouldBind(&userJson) == nil {
		res := userModel.UpdateUser(userJson.Uid, userJson)
		c.JSON(http.StatusOK, common.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, "绑定数据模型失败", false))
	return
}

func DeleteUsers(c *gin.Context) {
	if res := haveAuth(c, "deleteUser"); res != common.Authed {
		c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	userValidate := validate.UserValidate
	userModel := model.User{}

	if res, err:= userValidate.Validate(c, "groupDelete"); !res {
		c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	userArrayJson := struct {
		Users []int `json:"users" form:"users"`
	}{}

	if c.ShouldBind(&userArrayJson) == nil {
		for _, uid := range userArrayJson.Users {
			res := userModel.DeleteUser(uid)
			if res.Status != common.CodeSuccess {
				c.JSON(http.StatusOK, common.ApiReturn(res.Status, "uid为"+string(rune(uid))+"的用户删除失败", res.Data))
				return
			}
		}
		c.JSON(http.StatusOK, common.ApiReturn(common.CodeSuccess, "删除成功", true))
		return
	}
	c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, "绑定数据模型失败", false))
	return
}

func DeleteUser(c *gin.Context)  {
	if res := haveAuth(c, "deleteUser"); res != common.Authed {
		c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	userValidate := validate.UserValidate
	userModel := model.User{}

	if res, err:= userValidate.Validate(c, "delete"); !res {
		c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	userIDJson := struct {
		Uid	int `json:"uid" form:"uid"`
	}{}

	if c.ShouldBind(&userIDJson) == nil {
		res := userModel.DeleteUser(userIDJson.Uid)
		c.JSON(http.StatusOK, common.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, "绑定数据模型失败", false))
	return
}

func GetAllUser(c *gin.Context)  {
	if res := haveAuth(c, "getAllUser"); res != common.Authed {
		c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	userModel := model.User{}

	userJson := struct {
		Offset 	int 	`json:"offset" form:"offset"`
		Limit 	int 	`json:"limit" form:"limit"`
		Where 	struct{
			Nick 	string 	`json:"nick" form:"nick"`
			Mail 	string 	`json:"mail" form:"mail"`
		}
	}{}

	if c.ShouldBind(&userJson) == nil {
		userJson.Offset = (userJson.Offset-1)*userJson.Limit
		res := userModel.GetAllUser(userJson.Offset, userJson.Limit, userJson.Where.Nick, userJson.Where.Mail)
		c.JSON(http.StatusOK, common.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, "绑定数据模型失败", false))
	return
}

func GetUserByID(c *gin.Context) {
	if res := haveAuth(c, "getAllUser"); res != common.Authed {
		c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	userValidate := validate.UserValidate
	userModel := model.User{}

	if res, err:= userValidate.Validate(c, "find"); !res {
		c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	var userJson model.User

	if c.ShouldBind(&userJson) == nil {
		res := userModel.GetUserByID(userJson.Uid)
		c.JSON(http.StatusOK, common.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, "绑定数据模型失败", false))
	return
}

func Register(c *gin.Context) {
	userValidate := validate.UserValidate
	userModel := model.User{}

	if res, err:= userValidate.Validate(c, "register"); !res {
		c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	passwordCheck := struct {
		Password string `json:"password" form:"password"`
		PasswordCheck string `json:"password_check" form:"password_check"`
	}{}

	if c.ShouldBind(&passwordCheck) == nil {
		if passwordCheck.Password != passwordCheck.PasswordCheck {
			c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, "两次密码输入不一致", false))
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
	c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, "绑定数据模型失败", false))
	return
}

func Login(c *gin.Context) {
	session := sessions.Default(c)
	if id := session.Get("user_id"); id != nil {
		data := make(map[string]interface{}, 0)
		_ = json.Unmarshal([]byte(session.Get("data").(string)), &data)
		c.JSON(http.StatusOK, common.ApiReturn(common.CodeSuccess, "已登陆", data))
		return
	}

	userValidate := validate.UserValidate
	userModel := model.User{}

	if res, err:= userValidate.Validate(c, "login"); !res {
		c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	var loginUser model.User

	if c.ShouldBind(&loginUser) == nil {
		loginUser.Password = common.GetMd5(loginUser.Password)
		res := userModel.CheckLogin(loginUser)
		if res.Status == common.CodeSuccess {
			userInfo := res.Data.(map[string]interface{})["userInfo"].(model.User)
			returnData := map[string]interface{} {
				"user_id" : userInfo.Uid,
				"nick":		userInfo.Nick,
				"is_admin": userInfo.IsAdmin,
			}
			if menu, auths, err := getUserAllAuth(userInfo.Uid); err == nil {
				returnData["auths"] = auths
				returnData["menu"] = menu
			} else {
				c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, "获取权限失败", err.Error()))
				return
			}
			jsonData, _ := json.Marshal(returnData)
			session.Set("user_id", returnData["user_id"])
			session.Set("is_admin", returnData["is_admin"])
			session.Set("data", string(jsonData))
			session.Save()
			c.JSON(http.StatusOK, common.ApiReturn(res.Status, "登录成功", returnData))
		} else {
			c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, res.Msg, false))
		}
		return
	}
	c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, "绑定数据模型失败", false))
	return
}

func Logout(c *gin.Context)  {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, common.ApiReturn(common.CodeSuccess, "注销成功", session.Get("user_id")))
}

func GetUserInfo(c *gin.Context)  {
	session := sessions.Default(c)
	if id := session.Get("user_id"); id != nil {
		data := make(map[string]interface{}, 0)
		_ = json.Unmarshal([]byte(session.Get("data").(string)), &data)
		c.JSON(http.StatusOK, common.ApiReturn(common.CodeSuccess, "已登陆", data))
		return
	}
	c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, "未登陆", false))
}
