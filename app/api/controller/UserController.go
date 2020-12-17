package controller

import (
	"Authority/app/api/model"
	"Authority/app/api/validate"
	"Authority/app/common"
	"encoding/json"
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)


type menuItem struct {
	Title 	string `json:"title"`
	Icon	string `json:"icon"`
	Href	string `json:"href"`
	Target 	string `json:"target"`
	Child 	[]menuItem `json:"child"`
}

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
				"is_admin": userInfo.IsAdmin,
			}
			if menu, auths, err := getUserAllAuth(userInfo.Uid); err == nil {
				returnData["auths"] = auths
				returnData["menu"] = menu
			} else {
				c.JSON(http.StatusOK, common.ApiReturn(common.CODE_ERROE, "获取权限失败", err.Error()))
				return
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

func GetUserInfo(c *gin.Context)  {
	session := sessions.Default(c)
	if id := session.Get("user_id"); id != nil {
		data := make(map[string]interface{}, 0)
		_ = json.Unmarshal([]byte(session.Get("data").(string)), &data)
		c.JSON(http.StatusOK, common.ApiReturn(common.CODE_SUCCESS, "已登陆", data))
		return
	}
	c.JSON(http.StatusOK, common.ApiReturn(common.CODE_ERROE, "未登陆", false))
}

func getUserAllAuth(userID int) ([]menuItem, []string, error) {
	authModel := model.Auth{}

	if res := authModel.GetUserAllAuth(userID); res.Status == common.CODE_SUCCESS {
		auths := res.Data.([]model.Auth)
		var authsLeft []model.Auth
		var authName []string
		var menu []menuItem

		menuItemCount := 0
		type2Pos := map[int]int{}

		for _, auth := range auths {
			if auth.Type == 2 {
				authName = append(authName, auth.Title)
			} else if auth.Type == 0 {
				item := menuItem{
					Title: auth.Title,
					Target: auth.Target,
					Icon: auth.Icon,
					Href: auth.Href,
				}
				menu = append(menu, item)
				type2Pos[auth.Aid] = menuItemCount
				menuItemCount++
			} else if auth.Type == 1 {
				authsLeft = append(authsLeft, auth)
			}
		}
		if menu != nil {
			for _, auth := range authsLeft {
				item := menuItem{
					Title: auth.Title,
					Target: auth.Target,
					Icon: auth.Icon,
					Href: auth.Href,
				}
				pos := type2Pos[auth.Parent]
				menu[pos].Child = append(menu[pos].Child, item)
			}
		} else {
			menu = make([]menuItem, 0)
		}

		return menu, authName, nil
	} else {
		return nil, nil, errors.New("获取权限列表错误")
	}

}