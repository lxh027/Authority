package controller

import (
	"Authority/app/api/model"
	"Authority/app/api/validate"
	"Authority/app/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRoleAuthsList(c *gin.Context) {
	if res := haveAuth(c, "authAssign"); res != common.Authed {
		c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	roleAuthValidate := validate.RoleAuthValidate//jun
	authModel := model.Auth{}

	if res, err:= roleAuthValidate.Validate(c, "getRoleAuth"); !res {
		c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	allAuths := authModel.GetAuthNoRules()

	authJson := struct {
		Rid 	int 	`json:"rid" form:"rid"`
	}{}

	if c.ShouldBind(&authJson) == nil {
		res := authModel.GetRoleAuth(authJson.Rid)
		if res.Status!=common.CodeSuccess{
			c.JSON(http.StatusOK, common.ApiReturn(res.Status, res.Msg,res.Data))
			return
		}
		auths := res.Data.([]model.Auth)
		var val []int
		for _, auth := range auths {
			val = append(val, auth.Aid)
		}
		c.JSON(http.StatusOK, common.ApiReturn(res.Status, res.Msg, map[string]interface{} {
			"allAuths" 	: allAuths.Data,
			"values"	: val,

		}))
		return
	}
	c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, "绑定数据模型失败", false))
	return
}

