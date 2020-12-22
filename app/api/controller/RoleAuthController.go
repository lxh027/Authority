package controller

import (
	"Authority/app/api/model"
	"Authority/app/api/validate"
	"Authority/app/common"
	"encoding/json"
	"fmt"
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

func AddRoleAuths(c *gin.Context)  {
	if res := haveAuth(c, "authAssign"); res != common.Authed {
		c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	roleAuthValidate := validate.RoleAuthValidate
	roleAuthModel := model.RoleAuth{}

	if res, err:= roleAuthValidate.Validate(c, "addGroup"); !res {
		c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	roleAuthsJson := struct {
		Rid 	int 	`json:"rid" form:"rid"`
		Aids	string 	`json:"aids" form:"aids"`
	}{}

	if c.ShouldBind(&roleAuthsJson) == nil {
		var aids []int
		_ = json.Unmarshal([]byte((roleAuthsJson.Aids)), &aids)
		fmt.Println(aids)
		for _, aid := range aids {
			res := roleAuthModel.AddRoleAuth(model.RoleAuth{Rid: roleAuthsJson.Rid, Aid: aid})
			if res.Status != common.CodeSuccess {
				c.JSON(http.StatusOK, common.ApiReturn(res.Status, "编号为"+string(rune(aid))+"的权限添加失败", res.Data))
				return
			}
		}
		c.JSON(http.StatusOK, common.ApiReturn(common.CodeSuccess, "添加成功", true))
		return
	}
	c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, "绑定数据模型失败", false))
	return
}

func DeleteRoleAuths(c *gin.Context)  {
	if res := haveAuth(c, "authAssign"); res != common.Authed {//jun
		c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	userRoleValidate := validate.UserRoleValidate
	userRoleModel := model.UserRole{}

	if res, err:= userRoleValidate.Validate(c, "deleteGroup"); !res {
		c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	userRolesJson := struct {
		Uid 	int 	`json:"uid" form:"uid"`
		Rids	string 	`json:"rids" form:"rids"`
	}{}

	if c.ShouldBind(&userRolesJson) == nil {
		var rids []int
		_ = json.Unmarshal([]byte((userRolesJson.Rids)), &rids)
		fmt.Println(rids)
		for _, rid := range rids {
			res := userRoleModel.DeleteUserRole(model.UserRole{Uid: userRolesJson.Uid, Rid: rid})
			if res.Status != common.CodeSuccess {
				c.JSON(http.StatusOK, common.ApiReturn(res.Status, "编号为"+string(rune(rid))+"的权限删除失败", res.Data))
			}
		}
		c.JSON(http.StatusOK, common.ApiReturn(common.CodeSuccess, "删除成功", true))
		return
	}
	c.JSON(http.StatusOK, common.ApiReturn(common.CodeError, "绑定数据模型失败", false))
	return
}

